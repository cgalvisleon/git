package imagenes

import (
	"api/pkg/db"
	"api/pkg/res"
	"log"
)

func CreateTable() error {
	log.Println("Creating table imagenes")

	sql := `
	CREATE TABLE IF NOT EXISTS imagenes (
		numIdentificacion INTEGER NOT NULL,
		nombreImagen TEXT,
		fecha TEXT,
		PRIMARY KEY (numIdentificacion, nombreImagen),
		FOREIGN KEY (numIdentificacion) REFERENCES formularios(numIdentificacion)
	);`

	_, err := db.Db.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table imagenes created successfully")

	return nil
}

func InsertImagen(numIdentificacion int, nombreImagen, fecha string) error {
	log.Println("InsertImagen")

	sql := `
	INSERT INTO imagenes (numIdentificacion, nombreImagen, fecha)
	VALUES (?, ?, ?);
	`
	_, err := db.Db.Exec(sql, numIdentificacion, nombreImagen, fecha)
	if err != nil {
		return err
	}

	return nil
}

func UpdateImagen(numIdentificacion int, nombreImagen, fecha string) error {
	log.Println("UpdateImagen")

	sql := `
	UPDATE imagenes SET
	fecha = ?
	WHERE numIdentificacion = ? AND nombreImagen = ?;
	`
	_, err := db.Db.Exec(sql, fecha, numIdentificacion, nombreImagen)
	if err != nil {
		return err
	}

	return nil
}

func DeleteImagen(numIdentificacion int, nombreImagen string) error {
	log.Println("DeleteImagen")

	sql := `
	DELETE FROM imagenes
	WHERE numIdentificacion = ? AND nombreImagen = ?;
	`
	_, err := db.Db.Exec(sql, numIdentificacion, nombreImagen)
	if err != nil {
		return err
	}

	return nil
}

func GetImagenes(numIdentificacion int) ([]res.Json, error) {
	log.Println("GetImagenes")

	sql := `
	SELECT nombreImagen, fecha
	FROM imagenes
	WHERE numIdentificacion = ?;
	`
	rows, err := db.Db.Query(sql, numIdentificacion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imagenes := []res.Json{}
	for rows.Next() {
		var nombreImagen, fecha string
		err := rows.Scan(&nombreImagen, &fecha)
		if err != nil {
			return nil, err
		}

		imagen := res.Json{
			"nombreImagen": nombreImagen,
			"fecha":        fecha,
		}
		imagenes = append(imagenes, imagen)
	}

	return imagenes, nil
}
