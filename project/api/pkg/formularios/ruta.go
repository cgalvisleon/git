package formularios

import (
	"log"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	log.Println("Routes formularios")

	r.Post("/formularios", insertFormulario)
	r.Put("/formularios/{id}", updateFormulario)
	r.Delete("/formularios/{id}", deleteFormulario)
	r.Get("/formularios", getFormularios)
	r.Get("/formularios/{id}", getFormulario)
}
