package formularios

import (
	"api/pkg/db"
	"api/pkg/res"
	"log"
)

func CreateTable() error {
	log.Println("Creating table formularios")

	sql := `
	CREATE TABLE IF NOT EXISTS formularios (
		numIdentificacion INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		nombres TEXT,
		apellidos TEXT,
		tipoIdentificacion TEXT,
		estadoCivil INTEGER,
		fechaNacimiento TEXT,
		numBeneficiarios INTEGER,
		fechaIngreso TEXT
	);
	`
	_, err := db.Db.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table formularios created successfully")

	return nil
}

func InsertFormulario(nombre, apellidos, tipoIdentificacion string, estadoCivil int, fechaNacimiento string, numBeneficiarios int, fechaIngreso string) error {
	log.Println("InsertFormulario")

	sql := `
	INSERT INTO formularios (nombres, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	VALUES (?, ?, ?, ?, ?, ?, ?);
	`
	_, err := db.Db.Exec(sql, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFormulario(numIdentificacion int, nombre, apellidos, tipoIdentificacion string, estadoCivil int, fechaNacimiento string, numBeneficiarios int, fechaIngreso string) error {
	log.Println("UpdateFormulario")

	sql := `
	UPDATE formularios SET
	nombres = ?,
	apellidos = ?,
	tipoIdentificacion = ?,
	estadoCivil = ?,
	fechaNacimiento = ?,
	numBeneficiarios = ?,
	fechaIngreso = ?
	WHERE numIdentificacion = ?;
	`
	_, err := db.Db.Exec(sql, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso, numIdentificacion)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFormulario(numIdentificacion int) error {
	log.Println("DeleteFormulario")

	sql := `
	DELETE FROM formularios
	WHERE numIdentificacion = ?;
	`
	_, err := db.Db.Exec(sql, numIdentificacion)
	if err != nil {
		return err
	}

	return nil
}

func GetFormulario(numIdentificacion int) (res.Json, error) {
	log.Println("GetFormulario")

	sql := `
	SELECT *
	FROM formularios
	WHERE numIdentificacion = ?;
	`
	row := db.Db.QueryRow(sql, numIdentificacion)

	var numIdentificacionResult int
	var nombres, apellidos, tipoIdentificacion, fechaNacimiento, fechaIngreso string
	var estadoCivil, numBeneficiarios int
	err := row.Scan(&numIdentificacionResult, &nombres, &apellidos, &tipoIdentificacion, &estadoCivil, &fechaNacimiento, &numBeneficiarios, &fechaIngreso)
	if err != nil {
		return nil, err
	}

	return res.Json{
		"numIdentificacion":  numIdentificacionResult,
		"nombres":            nombres,
		"apellidos":          apellidos,
		"tipoIdentificacion": tipoIdentificacion,
		"estadoCivil":        estadoCivil,
		"fechaNacimiento":    fechaNacimiento,
		"numBeneficiarios":   numBeneficiarios,
		"fechaIngreso":       fechaIngreso,
	}, nil
}

func AllFormilarios() ([]res.Json, error) {
	log.Println("AllFormularios")

	sql := `
	SELECT *
	FROM formularios;
	`
	rows, err := db.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var formularios []res.Json
	for rows.Next() {
		var numIdentificacionResult int
		var nombres, apellidos, tipoIdentificacion, fechaNacimiento, fechaIngreso string
		var estadoCivil, numBeneficiarios int
		err := rows.Scan(&numIdentificacionResult, &nombres, &apellidos, &tipoIdentificacion, &estadoCivil, &fechaNacimiento, &numBeneficiarios, &fechaIngreso)
		if err != nil {
			return nil, err
		}

		formulario := res.Json{
			"numIdentificacion":  numIdentificacionResult,
			"nombres":            nombres,
			"apellidos":          apellidos,
			"tipoIdentificacion": tipoIdentificacion,
			"estadoCivil":        estadoCivil,
			"fechaNacimiento":    fechaNacimiento,
			"numBeneficiarios":   numBeneficiarios,
			"fechaIngreso":       fechaIngreso,
		}
		formularios = append(formularios, formulario)
	}

	return formularios, nil
}
