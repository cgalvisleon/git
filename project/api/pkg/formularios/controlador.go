package formularios

import (
	"api/pkg/res"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func insertFormulario(w http.ResponseWriter, r *http.Request) {
	log.Println("insertFormulario")
	body, _ := res.GetBody(r)

	nombre := body["nombre"].(string)
	apellidos := body["apellidos"].(string)
	tipoIdentificacion := body["tipoIdentificacion"].(string)
	estadoCivilStr := body["estadoCivil"].(string)
	fechaNacimiento := body["fechaNacimiento"].(string)
	numBeneficiariosStr := body["numBeneficiarios"].(string)
	fechaIngreso := body["fechaIngreso"].(string)

	estadoCivil, err := strconv.Atoi(estadoCivilStr)
	if err != nil {
		estadoCivil = 0
	}

	numBeneficiarios, err := strconv.Atoi(numBeneficiariosStr)
	if err != nil {
		numBeneficiarios = 0
	}

	err = InsertFormulario(nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, nil)
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Formulario creado exitosamente",
	})
}

func updateFormulario(w http.ResponseWriter, r *http.Request) {
	log.Println("updateFormulario")
	body, _ := res.GetBody(r)

	numIdentificacionStr := chi.URLParam(r, "id")
	nombre := body["nombre"].(string)
	apellidos := body["apellidos"].(string)
	tipoIdentificacion := body["tipoIdentificacion"].(string)
	estadoCivilStr := body["estadoCivil"].(string)
	fechaNacimiento := body["fechaNacimiento"].(string)
	numBeneficiariosStr := body["numBeneficiarios"].(string)
	fechaIngreso := body["fechaIngreso"].(string)

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "numIdentificacion debe ser un número",
		})
		return
	}

	estadoCivil, err := strconv.Atoi(estadoCivilStr)
	if err != nil {
		estadoCivil = 0
	}

	numBeneficiarios, err := strconv.Atoi(numBeneficiariosStr)
	if err != nil {
		numBeneficiarios = 0
	}

	err = UpdateFormulario(numIdentificacion, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, nil)
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Formulario actualizado exitosamente",
	})
}

func deleteFormulario(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteFormulario")

	numIdentificacionStr := chi.URLParam(r, "id")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "numIdentificacion debe ser un número",
		})
		return
	}

	err = DeleteFormulario(numIdentificacion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, nil)
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{
		"message": "Formulario eliminado exitosamente",
	})
}

func getFormulario(w http.ResponseWriter, r *http.Request) {
	log.Println("getFormulario")

	numIdentificacionStr := chi.URLParam(r, "id")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Json{
			"message": "numIdentificacion debe ser un número",
		})
		return
	}

	log.Println("numIdentificacion", numIdentificacion)

	result, err := GetFormulario(numIdentificacion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}

func getFormularios(w http.ResponseWriter, r *http.Request) {
	log.Println("getFormularios")

	result, err := AllFormilarios()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, nil)
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}
