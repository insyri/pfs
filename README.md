<img align="right" width="150" src="./frontend/src/assets/directory_computer.svg"/>

# pfs

📂 Server for pasting, uploading, and downloading files. Built on an API that serves data to the frontend, allowing both developers and casual i-net users to browse file contents.

[![](https://shields.io/badge/Svelte-05122A?logo=svelte&style=for-the-badge)](https://svelte.dev/)
[![](https://shields.io/badge/Go-05122A?logo=go&style=for-the-badge)](https://go.dev/)
[![](https://shields.io/badge/Docker-05122A?logo=docker&style=for-the-badge)](https://docker.com/)
[![](https://shields.io/badge/TypeScript-05122A?logo=typescript&style=for-the-badge)](https://www.typescriptlang.org/)
[![](https://shields.io/badge/Tailwind%20CSS-05122A?logo=tailwindcss&style=for-the-badge)](https://tailwindcss.com/)
[![](https://shields.io/badge/PostgreSQL-05122A?logo=postgresql&style=for-the-badge)](https://postgresql.org/)
[![](https://shields.io/badge/Nginx-05122A?logo=nginx&logoColor=009639&style=for-the-badge)](https://nginx.org/)

<!-- ## Features

...

## Public Instances

... -->

## To start using pfs

pfs uses [Docker](https://www.docker.com/) to create and manage the programs used to make pfs work.

<!-- TODO: make configurable settings in pfs and add them here -->

```bash
git clone https://github.com/insyri/pfs
cd pfs
docker compose up
```

## Development

pfs has a [contributing guide](./.github/CONTRIBUTING.md), it contains contributing guidelines and more in-depth development setup specifications.

If you want to build and run pfs right away, you can run this through

On \*nix[^1] systems, users can run this one liner in the project directory:

```bash
docker compose -f ./dev-docker-compose.yml
```

For Windows systems, users will have to use WSL and open the project inside the Linux file system.

[^1]: Unix, unix-like, and unix based. E.g.: MacOS, Linux, and Windows Subsystem for Linux.
