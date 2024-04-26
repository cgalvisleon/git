package formularios

import (
	"log"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	log.Println("Routes formularios")

	r.Post("/formularios", insertFormulario)
}
