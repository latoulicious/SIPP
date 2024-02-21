# Sistem Informasi Perangkat Pembelajaran (Learning Resources Information System)

A comprehensive information system designed to streamline the management of educational resources. Built with Go, Vue.js, and PostgreSQL, and emphasizing clean architecture.

## Demo

Start by including 2-3 eye-catching screenshots of your web app's best features. This immediately shows potential users what the system looks like.

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

* Thorough Documentation and Testing: Ensure reliability and provide clear guidance for developers and users

## Getting Started

This project has two components: a web frontend and a Go backend server. Follow these steps to get up and running.

### Prerequisites

* Node.js and PNPM
* Go 1.21.3 or later
* A PostgreSQL database instance

> Web Frontend Setup

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

```sh
air
```

Running without hot reloading:

```sh
go run cmd/main.go
```

> Additional Notes:

* The server will typically connect to a PostgreSQL database. Setup your database and configure connection settings in the server's environment.

## Contributing

We greatly appreciate contributions to this project! If you have improvements, bug fixes, or new features you'd like to suggest, please:

* Open an issue: Describe the proposed change or issue in detail. This helps us track discussions and understand the context for potential contributions.

* Submit a pull request: Fork the repository, make your changes, and submit a pull request for review. Be sure to follow any established coding standards and conventions.

For more detailed guidelines on our contribution process, please refer to our CONTRIBUTING.md file [link to the CONTRIBUTING.md in your project directory].

## License

This project is licensed under the  MIT license. For the full license text, please see the [LICENSE](https://github.com/latoulicious/SIPP/blob/main/LICENSE).
