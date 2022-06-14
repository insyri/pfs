# Contributing

Thanks for showing your interest for contributing to pfs! If you have further
questions on contributing or setup, create an issue.

Before contributing, ensure that an the appropriate issue has reported (unless
it's very small) and an **active** pull request has not yet picked up the
appropriate issue.

## Setup

### Programs
- [Docker](https://www.docker.com/)
- [Node.js](https://nodejs.org/)
- [Go](https://golang.org/)

If you're using VSCode, extensions will be recommended to you. These extensions can be found in the [`extensions.json`](../.vscode/extensions.json) file.

### Running

For Unix and Unix-like based systems, you can simply run the one liner in the project root:
```bash
docker compose -f ./dev-docker-compose.yml up
```

For Windows systems, there is a more complex process.

1. Run the database in a container.
```powershell
# In the project root
.\Start-PostgresInDocker.ps1
# Will output some hash, this is the container ID.
```
2. Run the specific project uncontained.
```powershell
# For backend development
cd backend
# Verify .air.toml is using the Windows config lines.
# Also, verify that main.go line 27 is using "localhost" instead of "database".
go mod download
go install github.com/cosmtrek/air@latest
air

# For frontend development
cd frontend
npm run dev
```

<!-- <details>
  <summary>About Windows live-reloading inside Docker</summary>
  some info later
  https://github.com/cosmtrek/air/issues/190
  https://github.com/microsoft/WSL/issues/4739
</details> -->
