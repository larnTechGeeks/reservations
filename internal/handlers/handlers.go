package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/larnTechGeeks/reservations/internal/config"
	"github.com/larnTechGeeks/reservations/internal/helpers"
	"github.com/larnTechGeeks/reservations/internal/models"
	"net/http"
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
	helpers.RenderTemplate(rw, r, "home.page.tmpl", &models.TemplateData{
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
	helpers.RenderTemplate(rw, r, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (hr *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(rw, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (hr *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(rw, r, "generals.page.tmpl", &models.TemplateData{})
}

func (hr *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(rw, r, "majors.page.tmpl", &models.TemplateData{})
}

func (hr *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("Contact Page Hit")
	helpers.RenderTemplate(rw, r, "contact.page.tmpl", &models.TemplateData{})
}

func (hr *Repository) Availability(rw http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(rw, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (hr *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	// get the post data and Write Back
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	rw.Write([]byte(fmt.Sprintf("Arrival Date: %s Departure Date: %s", start, end)))

	//helpers.RenderTemplate(rw, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// define the JSON type that will be returned to the user
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:message`
}

func (hr *Repository) AvailabilityJSON(rw http.ResponseWriter, r *http.Request) {
	// Initialize the struct with data
	res := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	//Construct Marshal the JSON  struct to json bytes using json package
	// data variable is marshalled to bytes
	data, err := json.MarshalIndent(res, "", "    ")

	//check for errors
	if err != nil {
		fmt.Println(err)
	}
	//Set Content-Type header to rw to hold application/json
	rw.Header().Set("Content-Type", "application/json")
	// write the json data to response writer
	rw.Write(data)

}
