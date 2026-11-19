# Minimalist Echo Boilerplate
This project is used to create a minimalistic boilerplate for golang echo.
Project structure:

```
|-- cmd: used for main entrypoint to main.go
|-- deployment: used for deployment script like Dockerfile or docker-compose.yaml
|-- internal: used for config, database, domain, response constant, and routing
    |-- config: used for config echo or project in general
    |-- database: used for database config, currently only support for postgresql
    |-- domain: this area will be used for handler, service, repository, models based on your business process (example like below)
        |-- user:
            |-- handler.go
            |-- models.go
            |-- repository.go
            |-- service.go
        |-- role:
            |-- handler.go
            |-- models.go
            |-- repository.go
            |-- service.go
    |-- http_response: used for create constant http_response success or failed
    |-- server: used for routing and register routes each domain
```
# Depedency:
1. [Golang Echo Framework](https://echo.labstack.com): Main framework
2. [Gorm](https://gorm.io)
3. Postgres driver for gorm
3. [joho/godotenv](https://github.com/joho/godotenv)