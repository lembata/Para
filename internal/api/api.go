package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lembata/para/pkg/logger"
	"net/http"
	//"github.com/go-chi/chi/v5/render"
	"html/template"
)

var logger = log.NewLogger()

type Server struct {
	http.Server
	DashboardService
}

type DashboardService struct {
	dashboardTemplate *template.Template
}

type LoginService struct {
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
	}

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(authenticateHandler())
	//TODO: auth middleware
	//TODO: controllers and routes
	router.Mount("/dashboard", server.dashBoardRouter())
	fs := http.FileServer(http.Dir("web/public"))
	router.Handle("/*", http.StripPrefix("/", fs))

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

func (s *DashboardService) ShowDashboard(w http.ResponseWriter, r *http.Request) {
	if s.dashboardTemplate == nil {
		t, err := template.New("index").ParseFiles("web/index.html")
		if err != nil {
			logger.Errorf("failed to parse template: %v", err)
			return
		}

		s.dashboardTemplate = t
	}

	err := s.dashboardTemplate.Execute(w, nil)

	if err != nil {
		logger.Errorf("failed excutetempate: %v", err)
		return
	}
}

func (s *LoginService) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	return
	// if s.dashboardTemplate == nil {
	// 	t, err := template.New("index").ParseFiles("web/index.html")
	// 	if err != nil {
	// 		logger.Errorf("failed to parse template: %v", err)
	// 		return
	// 	}
	//
	// 	s.dashboardTemplate = t
	// }
	//
	// err := s.dashboardTemplate.Execute(w, nil)
	//
	// if err != nil {
	// 	logger.Errorf("failed excutetempate: %v", err)
	// 	return
	// }
}
