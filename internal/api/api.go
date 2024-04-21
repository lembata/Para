package api

import (
	//"html/template"
	"bufio"
	"bytes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/lembata/para/pkg/logger"
)

var logger = log.NewLogger()

type Server struct {
	http.Server
	DashboardService
	LoginService
}

type SidebarButton struct {
	Icon string
	Text string
}

type DashboardService struct {
	//dashboardTemplate *template.Template
	templates Templates
}

type TempDash struct {
	Header    string
	Paragraph string
}



type Page struct {
	Body string
	Sidebar []SidebarButton
}

func Init() (*Server, error) {
	logger.Debug("Initializing API...")

	address := "localhost:8080"
	router := chi.NewRouter()

	templates := Templates{}
	err := templates.LoadTemplates()

	if err != nil {
		logger.Errorf("failed to load templates: %v", err)
		return nil, err
	}

	server := Server{
		Server: http.Server{
			Addr:    address,
			Handler: router,
		},
		DashboardService: DashboardService{
			templates: templates,
		},
		LoginService: LoginService{
			templates: templates,
		},
	}

	server.LoginService.templates.LoadTemplates()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(authenticateHandler())
	//TODO: auth middleware
	//TODO: controllers and routes

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	})
	router.Mount("/dashboard", server.dashBoardRouter())
	router.Mount("/login", server.loginRouter())
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

func (s *Server) loginRouter() http.Handler {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/", s.LoginService.ShowLoginPage)
	r.Post("/", s.LoginService.Login)
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
	var buffer = bytes.Buffer{}
	//buffer.Grow(1024 * 20);
	bufferWriter := bufio.NewWriter(&buffer)
	//bufferWriter.Write([]byte("Hello, World!"))

	err := s.templates.Render(bufferWriter, "dashboard", nil)
	err = bufferWriter.Flush()
	//var err error

	if err != nil {
		logger.Errorf("failed excutetempate: %v", err)
		return
	}

	logger.Debugf("Buffer length: %v", buffer.Len())

	page := Page{Body: buffer.String()}

	logger.Debugf("page: %v", page)

	err = s.templates.Render(w, "mainPage", page)

	if err != nil {
		logger.Errorf("failed excutetempate: %v", err)
		return
	}

	logger.Info("dashboard page served")
}

