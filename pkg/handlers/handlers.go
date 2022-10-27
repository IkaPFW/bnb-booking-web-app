package handlers

import (
	"net/http"

	"github.com/ikapfw/bnb-booking-web-app/pkg/config"
	"github.com/ikapfw/bnb-booking-web-app/pkg/models"
	"github.com/ikapfw/bnb-booking-web-app/pkg/render"
)

// repository used by handler
var Repo *Repository

// repository type
type Repository struct{
	App *config.AppConfig
}

// create new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// set repository for handlers
func NewHandler(r *Repository){
	Repo = r
}

// home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	// fmt.Fprintf(w, "This is the home page")
}

// about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page. The sum of 2 + 2 is %d", sum))
}