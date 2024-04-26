package formularios

import (
	"api/pkg/db"
	"log"
)

func CreateTable() error {
	log.Println("Creating table formularios")

	sql := `
	CREATE TABLE IF NOT EXISTS formularios (
		numdentification INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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
	log.Println("insertFormulario")

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
