# Votio

<div align="center">TODO - insert a logo and screenshot here</div>

<h2 align="center">Turn customer feedback into actionable insights with Votio</h2>

Votio is an open source Customer feedback management tool

## ⭐️ Why Novu?

TODO

## ✨ Features

TODO

## 👨‍💻 Development

### Prerequisites

Votio is powered by [PocketBase](https://pocketbase.io/) and [SolidJS](https://www.solidjs.com/) so you should have:

- [Go v1.20+](https://gist.github.com/MichaelCurrin/ca6b3b955172ff993184d39807dd68d4)
- [NodeJS v18+](https://nodejs.org/en/download/package-manager/)
- [Task](https://taskfile.dev/installation/) - optional, useful to manage the whole project

### How to run locally

#### Using Task

This is the easiest way to get started:

```
# install
task frontend-install
# This command will start the development servers for both backend and frontend
task run-app
```

#### Manually

1. Clone this repo
2. Install and build backend dependencies:

```
cd pocketbase
go mod download
go build
```

3. Run server: `./pocketbase serve`
4. Open a new terminal at the root of the repo
5. Install frontend dependencies:

```
cd app
pnpm install
pnpm dev
```

6. Open a new terminal and run: `pnpm dev`

#### Using Docker

if you don't have Go installed, you can use docker to work only in the frontend by:

```
cd pocketbase
docker-compose up -d

# then run the frontend server
cd ../app
pnpm install
pnpm dev
```

## 🚀 Deployment

TODO

## 🛡️ License

Novu is licensed under the MIT License - see the [LICENSE](https://github.com/gengue/votio/blob/main/LICENSE) file for details.
