# Sistem Informasi Perangkat Pembelajaran 

This project served as my college thesis. It's a comprehensive information system designed to streamline educational resource management. Built with Go, Vue.js, and PostgreSQL, it emphasizes clean architecture principles.

## Demo

![Login](https://i.ibb.co/cTg6BnX/Screenshot-2024-02-21-at-20-34-50-Sistem-Informasi-Perangkat-Pembelajaran.png)

<h4 align="center">Login Page</h4>


![Dashboard](https://i.ibb.co/R0Pvh11/Screenshot-2024-02-21-at-20-35-28-Sistem-Informasi-Perangkat-Pembelajaran.png)

<h4 align="center">Dashboard</h4>

![Database](https://i.ibb.co.com/XtrTxLf/syllabify2-public.png)

<h4 align="center">Database Relation</h4>

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## Key Features

* Powerful CRUD Operations: Easily create, read, update, and delete learning resources.

* Role-Based Access Control: Ensure secure data handling with fine-grained controls over user permissions.

* Clean Architecture: Maintainable and scalable thanks to a well-structured design.

* User Authentication and Authorization: Protect your system with user logins and role-based access.

* User-Friendly Interface: Designed for easy navigation and intuitive use.

* Error Handling and Logging: Robustly handle unexpected issues and track system activity.

## Getting Started

This project has two components: a Vue Frontend and a Go Backend. Follow these steps to get up and running.

### Prerequisites

* Node.js and PNPM

* Go 1.20 or later

* A PostgreSQL database instance

### Setup

> Frontend Setup

Navigate to the web folder:

```sh
cd web
```


Install dependencies:

```sh
pnpm install
```

Development mode (with hot reloading):

```sh
pnpm dev
```

Production build:

```sh
pnpm build
```

Linting:

```sh
pnpm lint
```

Formatting:

```sh
pnpm format
```

> Server Setup

Navigate to the server folder:

```sh
cd server
```

Download dependencies:

```sh
go mod tidy
```

Development mode (with hot reloading):

* Set up environment variables for Air (refer to [Air documentation](https://github.com/cosmtrek/air) )
* or modifying the existing [air.toml](https://github.com/latoulicious/SIPP/blob/main/server/.air.toml)

```sh
air
```

Running without hot reloading:

```sh
go run cmd/main.go
```

> Additional Notes:

* The server will typically connect to a PostgreSQL database. Setup your database and configure connection settings in the server's environment.

## License

This project is licensed under the  [MIT license](https://github.com/latoulicious/SIPP/blob/main/LICENSE).