package imagenes

import (
	"log"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	log.Println("Routes Imagenes")

	r.Post("/imagenes", insertImagen)
	r.Put("/imagenes/{id}/{nombre}", updateImagen)
	r.Delete("/imagenes/{id}/{nombre}", deleteImagen)
	r.Get("/imagenes/{id}", getImagenes)
}
