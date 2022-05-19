package handlers

import (
	"github.com/seun-otosho/go-course/pkg/config"
	"github.com/seun-otosho/go-course/pkg/models"
	"github.com/seun-otosho/go-course/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

//Repo repos
var Repo *Repository

// NewRepo creates new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// some biz logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hi again"
	// send to page

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
