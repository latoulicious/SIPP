# Sistem Informasi Perangkat Pembelajaran

"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

## Server

> 🚧 Work in Progress

## Web

> Project Setup

```sh
yarn
```

> Compile and Hot-Reload for Development

```sh
yarn dev
```

> Type-Check, Compile and Minify for Production

```sh
yarn build
```

> Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
yarn test:unit
```

> Run End-to-End Tests with [Cypress](https://www.cypress.io/)

```sh
yarn test:e2e:dev
```

### This runs the end-to-end tests against the Vite development server. It is much faster than the production build. But it's still recommended to test the production build with `test:e2e` before deploying (e.g. in CI environments)

```sh
yarn build
yarn test:e2e
```

> Lint with [ESLint](https://eslint.org/)

```sh
yarn lint
```

> Format with [Prettier](https://prettier.io/)

```sh
yarn format
```
