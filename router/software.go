package router

import (
	"github.com/chinx/maitre/module/svrmux"
	"github.com/chinx/maitre/models"
)

func SoftwareRouter(m svrmux.Router) {
	m.Group("/abc/:bcd/:def", func() {
		m.Get("/a","10110", &models.Software{})
		m.Post("/b","10110", &models.Software{})
		m.Put("/c","10110", &models.Software{})
	})
}
