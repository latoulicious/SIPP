# Sistem Informasi Perangkat Pembelajaran

"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

## Server

> Project Setup

```sh
go mod tidy
```

> Compile and Hot-Reload for Development

```sh
air
```
* Need to install air first

* [Air - Live reload for Go apps](https://github.com/cosmtrek/air)

> Run without Hot-Reload

```
go run cmd/main.go
```

## Web

> Project Setup

```sh
pnpm
```

> Compile and Hot-Reload for Development

```sh
pnpm dev
```

> Type-Check, Compile and Minify for Production

```sh
pnpm build
```

> Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
pnpm test:unit
```

> Lint with [ESLint](https://eslint.org/)

```sh
pnpm lint
```

> Format with [Prettier](https://prettier.io/)

```sh
pnpm format
```
