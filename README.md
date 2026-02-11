# 🚀 Go Product Service

![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-blue?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Enabled-blue?logo=docker)
![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow)

API REST desenvolvida em **Go** com integração ao **PostgreSQL**, utilizando **Docker** para ambiente de desenvolvimento.

Este projeto foi criado com o objetivo de estudar arquitetura backend, organização em camadas e integração com banco de dados relacional.

---

## 🛠️ Tecnologias Utilizadas

- Go
- PostgreSQL 16
- Docker
- Docker Compose

🏗️ Arquitetura

O projeto segue uma estrutura baseada em camadas para garantir a separação de responsabilidades:

    Cmd: Ponto de entrada da aplicação.

    Controller: Responsável por lidar com as requisições HTTP e retornar as respostas ao cliente.

    Model: Define a estrutura dos dados (Entidades) e a comunicação direta com o banco de dados.

    Repository: Camada de abstração para persistência de dados.

🚀 Como Executar
Pré-requisitos
```text
    Go (versão 1.22 ou superior)

    Docker e Docker Compose

## 📁 Estrutura do Projeto

GoApi/
├── cmd/
├── controller/
│   └── product_controller.go
├── model/
│   └── product.go
├── main.go
├── docker-compose.yml
├── go.mod
└── go.sum

