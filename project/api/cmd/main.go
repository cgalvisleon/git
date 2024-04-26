package main

import (
	"api/pkg/db"
	"api/pkg/formularios"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	err := db.ConectDB()
	if err != nil {
		log.Panicln(err)
	}

	err = CreateTables()
	if err != nil {
		log.Panicln(err)
	}

	err = HttpServer()
	if err != nil {
		log.Panicln(err)
	}
}

func CreateTables() error {
	log.Println("Creating tables")

	err := formularios.CreateTable()
	if err != nil {
		log.Panicln(err)
		return err
	}

	return nil
}

func HttpServer() error {
	log.Println("Starting http server")

	// Create a new router
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
		<html lang="es">
		<head>
				<meta charset="UTF-8">
				<title>Página HTML Básica</title>
				<meta name="description" content="Esta es una página HTML básica.">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<!-- Puedes agregar más metadatos, como iconos, autores, etc. -->
				<link rel="stylesheet" href="styles.css"> <!-- Archivo CSS externo -->
		</head>
		<body>
				<header>
						<h1>Bienvenido a Mi Página</h1>
				</header>
				<nav>
						<ul>
								<li><a href="#home">Inicio</a></li>
								<li><a href="#about">Acerca de</a></li>
								<li><a href="#contact">Contacto</a></li>
						</ul>
				</nav>
				<main>
						<section id="home">
								<h2>Inicio</h2>
								<p>Este es el contenido de la sección de inicio.</p>
						</section>
						<section id="about">
								<h2>Acerca de</h2>
								<p>Esta es la sección sobre nosotros.</p>
						</section>
				</main>
				<footer>
						<p>© 2024 Mi Sitio Web</p>
				</footer>
		</body>
		</html>`))
	})

	handler := cors.AllowAll().Handler(r)
	addr := ":3000"
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	fmt.Println("Run http server on http://localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
