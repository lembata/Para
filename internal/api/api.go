package api

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"net/http"
	"path"
	"time"

	"github.com/lembata/para/pkg/logger"
	//"github.com/lembata/para/pkg/database"
	"github.com/lembata/para/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/vearutop/statigz"
)

var logger = log.NewLogger()

type Server struct {
	http.Server
	DashboardService
	LoginService
	AccountService
}

type ApiResponse struct {
	Success   bool   `json:"success"`
	Data      any    `json:"data"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
}

func Init() (*Server, error) {
	logger.Debug("Initializing API...")

	address := "localhost:8080"
	router := chi.NewRouter()

	server := Server{
		Server: http.Server{
			Addr:    address,
			Handler: router,
		},
		DashboardService: DashboardService{},
		LoginService:     LoginService{},
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           10000,
	}))

	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(authenticateHandler())

	httpLogger := log.NewHttpLogger()

	router.Use(httplog.RequestLogger(httpLogger))

	router.Mount("/api/dashboard", server.dashBoardRouter())
	router.Mount("/api/login", server.loginRouter())
	router.Mount("/api/accounts", server.accountRouter())

	staticUI := statigz.FileServer(ui.UIBox.(fs.ReadDirFS))

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		ext := path.Ext(r.URL.Path)

		if ext == ".html" || ext == "" {
			data, err := fs.ReadFile(ui.UIBox, "index.html")
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Cache-Control", "no-cache")

			http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(data))

		} else {
			w.Header().Set("Cache-Control", "no-cache")

			staticUI.ServeHTTP(w, r)
		}
	})

	return &server, nil
}

func (s *Server) Start() error {
	logger.Infof("para is listening on " + s.Addr)
	//logger.Infof("para is running at " + s.displayAddress)

	if s.TLSConfig != nil {
		return s.ListenAndServeTLS("", "")
	} else {
		return s.ListenAndServe()
	}
}

func (s *Server) dashBoardRouter() http.Handler {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/", s.DashboardService.ShowDashboard)
	return r
}

func (s *Server) loginRouter() http.Handler {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/", s.LoginService.ShowLoginPage)
	r.Post("/", s.LoginService.Login)
	return r
}

func (s *Server) accountRouter() http.Handler {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/{id}", s.AccountService.GetAccount)
	r.Post("/add", s.AccountService.CreateAccount)
	r.Post("/all", s.AccountService.All)
	return r
}

func (s *Server) Close() error {
	logger.Info("Shutting down API server...")
	return s.Server.Close()
}

func (s *Server) Shutdown() error {
	logger.Info("Shutting down API server...")
	return s.Server.Shutdown(nil)
}

func (s *Server) addStaticResponses() http.Handler {
	//s.router.Get("/favicon.ico", http.FileServer(http.Dir("web/favicon.ico")))
	//s.router.Get("/static/*", http.FileServer(http.Dir("web/static")))

	return nil
}

func WriteFailure(w http.ResponseWriter, error string, errorCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(Failure(error, errorCode))
}

func WriteSuccess(w http.ResponseWriter) {
	w.Write(Success())
}

func WriteData(w http.ResponseWriter, data any) {
	w.Write(Data(data))
}


func Failure(error string, errorCode int) []byte {
	json, _ := json.Marshal(ApiResponse{
		Success:   false,
		Error:     error,
		ErrorCode: errorCode,
	})

	return json
}

func Success() []byte {
	json, _ := json.Marshal(ApiResponse{
		Success: true,
	})

	return json
}

func Data(data any) []byte {
	json, err := json.Marshal(ApiResponse{
		Success: true,
		Data:    data,
	})

	if err != nil {
		return Failure("failed to serialize data", http.StatusInternalServerError)
	}

	return json
}
