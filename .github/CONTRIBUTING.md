# Contributing

Thanks for showing your interest for contributing to pfs! If you have further
questions on contributing or setup, create an issue.

Before contributing, ensure that an the appropriate issue has reported (unless
it's very small) and an **active** pull request has not yet picked up the
appropriate issue.

## Guidelines

- 🗿 Follow repository structure.
- 🧹 Code is formatted, will be done for via CI.
  - The frontend must use prettier with the given [`.prettierrc`](../frontend/.prettierrc)
  - The backend must use gofmt. <!-- Future ruleset for gofmt -->
- 🐳 Use Docker, pfs was not prepared to be ran on metal.

## Setup

Please create a [fork](https://github.com/insyri/pfs/fork) of the pfs project on GitHub before developing.

[Git](https://git-scm.com/) is required to fetch the project from the internet to your local machine and track project file changes.

```bash
git clone https://github.com/<YOUR_USERNAME>/pfs
cd pfs
```

[Docker](https://docker.com/) is required to run the project under a controlled and containerized process.

If you're using VSCode, extensions will be recommended to you. These extensions can be found in the [`extensions.json`](../.vscode/extensions.json) file.

### \*nix

For Unix and Unix-like based systems, you can simply run the one liner in the project root:

```bash
docker compose -f ./dev-docker-compose.yml up
```

### Windows/WSL

For Windows systems users will have to use [WSL](https://docs.microsoft.com/en-us/windows/wsl/install) for development and live reloading.

1. Open WSL inside the Linux filesystem to a comfortable development directory.
2. Clone the project locally and cd into it.
3. Continue to [\*nix setup](#nix).

No. 1 can be done more simply by using the [VSCode WSL extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl). Directions for opening VSCode into WSL can be found on the extention marketplace page.

#### About Windows live-reloading inside Docker

When mounting volumes to a Docker container, WSL2 fails to propogate file changing events from the Windows file system, to the container's file system. As the files themselves change inside the container, the OS notification is not sent.

Often, these programs create a polling fallback strategy where OS events are not available, however, fsnotify, a file-watching library that air uses, has not integrated a polling fallback method. While it is possible to still have this useful mechanic with certain restraints, it is just easier to resort to using the Linux file system in WSL.

Related issues:

- [`cosmtrek/air#190`](https://github.com/cosmtrek/air/issues/190)
- [`microsoft/WSL#4739`](https://github.com/microsoft/WSL/issues/4739)
- [`fsnotify/fsnotify: ReadDirectoryChangesW`](https://github.com/fsnotify/fsnotify/labels/Windows%20%28ReadDirectoryChangesW%29)
