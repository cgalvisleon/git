package formularios

import (
	"api/pkg/res"
	"log"
	"net/http"
	"strconv"
)

func insertFormulario(w http.ResponseWriter, r *http.Request) {
	log.Panicln("insertFormulario")
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
