package api

import (
	"net/http"
)

type LoginService struct {
	//templates Templates
}


func (s *LoginService) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	return
	// err := s.templates.Render(w, "login", nil)
	//
	// if err != nil {
	// 	logger.Errorf("failed excutetempate: %v", err)
	// 	return
	// }
	//
	// logger.Info("Login page served")
}

func (s *LoginService) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		logger.Errorf("failed to parse form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !r.Form.Has("username") || !r.Form.Has("password") {
		logger.Errorf("username or password not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Debugf("form: %v", r.Form)
	err = Authenticate(w, r, r.Form.Get("usetname"), r.Form.Get("password"))

	if err != nil {
		logger.Errorf("failed to authenticate: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var returnUrl string

	if r.URL.Query().Has("returnURL") {
		returnUrl = r.URL.Query().Get("returnURL")
	} else {
		returnUrl = "/dashboard"
	}

	w.Header().Set("HX-Redirect", returnUrl)
}
