# pfs

📂 Server for pasting, uploading, and downloading files.

[![](https://shields.io/badge/Svelte-05122A?logo=svelte&style=for-the-badge)](https://svelte.dev/)
[![](https://shields.io/badge/TypeScript-05122A?logo=typescript&style=for-the-badge)](https://www.typescriptlang.org/)
[![](https://shields.io/badge/Tailwind%20CSS-05122A?logo=tailwindcss&style=for-the-badge)](https://www.typescriptlang.org/)
[![](https://shields.io/badge/Vite-05122A?logo=vite&style=for-the-badge)](https://www.typescriptlang.org/)
[![](https://shields.io/badge/Go-05122A?logo=go&style=for-the-badge)](https://go.dev/)
[![](https://shields.io/badge/PostgreSQL-05122A?logo=postgresql&style=for-the-badge)](https://postgresql.org/)
[![](https://shields.io/badge/Nginx-05122A?logo=nginx&logoColor=009639&style=for-the-badge)](https://nginx.org/)
[![](https://shields.io/badge/Docker-05122A?logo=docker&style=for-the-badge)](https://docker.com/)

## Development

To install this project on your system, you will need Git and the other appropriate dependencies for application specific development.

- [Git](https://git-scm.com/)

```bash
git clone https://github.com/insyri/pfs.git
cd pfs
```

### Frontend

- [Node.js](https://nodejs.org/)

```bash
cd frontend
npm install
npm run dev
```

### Backend

- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Go](https://go.dev/)

```bash
cd backend
# Direct install
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# Or, if you have Go already installed
go install github.com/cosmtrek/air # hot reload tool
air
```

## Deployment/Production

- [Docker](https://www.docker.com/)

```bash
docker compose up
```
