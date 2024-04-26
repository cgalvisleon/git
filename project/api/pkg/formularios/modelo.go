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
		fechaIngreso TEXT,
	);
	`
	_, err := db.Db.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table formularios created successfully")

	return nil
}
