package api

import (
	//"html/template"
	"bytes"
	"io/fs"
	"net/http"
	"path"
	"time"

	"github.com/lembata/para/pkg/logger"
	"github.com/lembata/para/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/vearutop/statigz"
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

type TempDash struct {
	Header    string
	Paragraph string
}

type Page struct {
	Body    string
	Sidebar []SidebarButton
}

func Init() (*Server, error) {
	logger.Debug("Initializing API...")

	address := "localhost:8080"
	router := chi.NewRouter()

	//var templates Templates
	//err := templates.LoadTemplates()

	// if err != nil {
	// 	logger.Errorf("failed to load templates: %v", err)
	// 	return nil, err
	// }

	server := Server{
		Server: http.Server{
			Addr:    address,
			Handler: router,
		},
		DashboardService: DashboardService{},
		LoginService:     LoginService{},
	}

	//server.LoginService.templates.LoadTemplates()

	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(authenticateHandler())

	httpLogger := httplog.NewLogger("Stash", httplog.Options{
		Concise: true,
	})
	router.Use(httplog.RequestLogger(httpLogger))

	//TODO: auth middleware
	//TODO: controllers and routes

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	})
	router.Mount("/api/dashboard", server.dashBoardRouter())
	router.Mount("/api/login", server.loginRouter())

	//customUILocation := ""
	staticUI := statigz.FileServer(ui.UIBox.(fs.ReadDirFS))

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		ext := path.Ext(r.URL.Path)

		if ext == ".html" || ext == "" {
			//themeColor := cfg.GetThemeColor()
			data, err := fs.ReadFile(ui.UIBox, "index.html")
			if err != nil {
				panic(err)
			}
			//indexHtml := string(data)
			//prefix := getProxyPrefix(r)
			//indexHtml = strings.ReplaceAll(indexHtml, "%COLOR%", themeColor)
			//indexHtml = strings.Replace(indexHtml, `<base href="/"`, fmt.Sprintf(`<base href="%s/"`, prefix), 1)

			w.Header().Set("Content-Type", "text/html")
			//setPageSecurityHeaders(w, r, pluginCache.ListPlugins())

			// if r.URL.Query().Has("t") {
			// 	w.Header().Set("Cache-Control", "private, max-age=31536000, immutable")
			// } else {
			// 	w.Header().Set("Cache-Control", "no-cache")
			// }
				
			w.Header().Set("Cache-Control", "no-cache")

			//w.Header().Set("ETag", GenerateETag(data))

			http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(data))

		} else {
			isStatic, _ := path.Match("/assets/*", r.URL.Path)
			if isStatic {
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			} else {
				w.Header().Set("Cache-Control", "no-cache")
			}

			staticUI.ServeHTTP(w, r)
		}
	})
	//fs := http.FileServer(http.Dir("ui/dist"))
	//router.Handle("/", fs.ServeHTTP(r, w))// http.StripPrefix("/", fs))

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
