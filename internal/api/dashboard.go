package api

import (
	// "bufio"
	// "bytes"
	"net/http"
)

type DashboardService struct {
	//dashboardTemplate *template.Template
	//templates Templates
}


func (s *DashboardService) ShowDashboard(w http.ResponseWriter, r *http.Request) {
	return
	// var buffer = bytes.Buffer{}
	// //buffer.Grow(1024 * 20);
	// bufferWriter := bufio.NewWriter(&buffer)
	// //bufferWriter.Write([]byte("Hello, World!"))
	//
	// err := s.templates.Render(bufferWriter, "dashboard", nil)
	// err = bufferWriter.Flush()
	// //var err error
	//
	// if err != nil {
	// 	logger.Errorf("failed excutetempate: %v", err)
	// 	return
	// }
	//
	// logger.Debugf("Buffer length: %v", buffer.Len())
	//
	// page := Page{Body: buffer.String()}
	//
	// logger.Debugf("page: %v", page)
	//
	// err = s.templates.Render(w, "mainPage", page)
	//
	// if err != nil {
	// 	logger.Errorf("failed excutetempate: %v", err)
	// 	return
	// }
	//
	// logger.Info("dashboard page served")
}
