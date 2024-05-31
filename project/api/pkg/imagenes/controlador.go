package imagenes

import (
	"api/pkg/res"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func insertImagen(w http.ResponseWriter, r *http.Request) {
	log.Println("insertImagen")
	body, _ := res.GetBody(r)

	numIdentificacionStr := body["numIdentificacion"].(string)
	nombreImagen := body["nombreImagen"].(string)
	fecha := body["fecha"].(string)

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "El campo numIdentificacion debe ser un número",
		})
		return
	}

	err = InsertImagen(numIdentificacion, nombreImagen, fecha)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Imagen insertada exitosamente",
	})
}

func updateImagen(w http.ResponseWriter, r *http.Request) {
	log.Println("updateImagen")

	numIdentificacionStr := chi.URLParam(r, "id")
	nombreImagen := chi.URLParam(r, "nombre")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "El campo numIdentificacion debe ser un número",
		})
		return
	}

	body, _ := res.GetBody(r)
	fecha := body["fecha"].(string)

	err = UpdateImagen(numIdentificacion, nombreImagen, fecha)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Imagen actualizada exitosamente",
	})
}

func deleteImagen(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteImagen")

	numIdentificacionStr := chi.URLParam(r, "id")
	nombreImagen := chi.URLParam(r, "nombre")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "El campo numIdentificacion debe ser un número",
		})
		return
	}

	err = DeleteImagen(numIdentificacion, nombreImagen)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Imagen eliminada exitosamente",
	})
}

func getImagenes(w http.ResponseWriter, r *http.Request) {
	log.Println("getImagenes")

	numIdentificacionStr := chi.URLParam(r, "id")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "El campo numIdentificacion debe ser un número",
		})
		return
	}

	imagenes, err := GetImagenes(numIdentificacion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, imagenes)
}
