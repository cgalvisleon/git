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
