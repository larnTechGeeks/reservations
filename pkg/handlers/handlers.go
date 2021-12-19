package handlers

import (
	"net/http"

	"github.com/larnTechGeeks/reservations/pkg/config"
	"github.com/larnTechGeeks/reservations/pkg/helpers"
	"github.com/larnTechGeeks/reservations/pkg/models"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewHandler(app *config.AppConfig) *Repository {
	return &Repository{
		app: app,
	}
}

func NewRepo(repo *Repository) {
	Repo = repo
}

func (hr *Repository) HomePage(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	hr.app.Session.Put(r.Context(), "remote-ip", remoteIP)
	StringMap := make(map[string]string)
	StringMap["name"] = "No Name for Home Page"
	helpers.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (hr *Repository) AboutPage(rw http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["name"] = "Vincent Omondi"

	// pull value from session
	remoteIP := hr.app.Session.GetString(r.Context(), "remote-ip")
	// add to string Map
	StringMap["remote-ip"] = remoteIP
	helpers.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
