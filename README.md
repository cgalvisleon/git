# Git

Git, creado por Linus Torvalds en 2005, es un sistema de control de versiones distribuido utilizado en desarrollo de software. Surgió para gestionar el kernel de Linux y se popularizó por su eficiencia y versatilidad. Permite a los desarrolladores realizar un seguimiento de cambios, colaborar en proyectos y gestionar ramas de código de forma ágil. Git ha revolucionado la forma en que se construyen y mantienen aplicaciones, siendo fundamental en la industria tecnológica moderna.

url: https://git-scm.com/

## Comandos basicos

```
git init: **Inicia un nuevo repositorio Git en el directorio actual.**
git clone: **Clona un repositorio Git existente a partir de una URL.**
```

## Docker deploy with docker-compose

```
docker network create proxy

docker-compose -p services -f ./images/local.yml up -d
docker-compose -p services -f ./images/local.yml down
```

## Go - Golang

### Initial project

```
go mod init <name>
```

### Run project

```
go run ./main.go
go run ./project/api/main.go
```

### Git recoganice .gitignore

```
git rm -r --cached .
git add .
git commit -m "Actualiza gitignore y limpia caché"

```

## Http server

```
go run ./project/api/cmd/main.go
```

### Dependencis

```
go get -u github.com/go-chi/chi/v5
go get -u github.com/mattn/go-sqlite3
go get -u github.com/rs/cors
```
