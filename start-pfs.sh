#!/bin/bash

PROD_DPSEC="./docker-compose.yml"
DEV_DSPEC="./dev-docker-compose.yml"
CONFIG="./pfs.toml"
DBENV="./database.env"
ENV="prod"
VERBOSE=false
COLOR="on"
HELP_MAIN="start-pfs

USAGE:
    start-pfs.sh [SUBCOMMAND] [FLAGS] [OPTION...]

SUBCOMMANDS:
    prepare  Prepares the docker environment for pfs.
    start    Starts the dockerized pfs service.
    build    Builds the pfs Docker image.
    help     Shows help information.
    clear    Clears the database environment file.

FLAGS:
    -e, --env [dev|prod|path/to/file.yml]
        Use environment specific Docker compose file.  default: prod
    -c, --config string  Path to configuration file.   default: pfs.toml
    -d, --dbenv string                                 default: database.env
        Path to database environment file, used in the Docker compose file.

OPTIONS:
    -h, --help      Shows help information.
    -n, --no-color  Disables color output.
    -v, --verbose   Print verbose output.
"

OPTS=$(getopt -o d:e:c:vhn --long dbenv:env:,config:,verbose,help,no-color: -n 'start-pfs' -- "$@")

verbose_echo() {
    if [ "$VERBOSE" = true ]; then
        echo -e "$(if [ "$COLOR" = "on" ]; then echo -e "\e[1;33m[VERBOSE]\e[0m" "$@"; else echo "[VERBOSE]" "$@"; fi)"
    fi
}

if [ ! "$OPTS" != 0 ]; then
    echo "Failed parsing options."
    exit 1
fi

eval set -- "$OPTS"

while true; do
    case $1 in
    -e | --env)
        shift
        if [ "$1" = "dev" ]; then
            ENV="dev"
        elif [ "$1" = "prod" ]; then
            :
        elif [ "$1" != "prod" ]; then
            if [ ! -e "$1" ]; then
                echo "Docker compose file does not exist."
                exit 1
            fi
            ENV="$1"
        fi
        shift
        ;;
    -C | --color)
        shift
        COLOR="off"
        shift
        ;;
    -d | --dbenv)
        shift
        if [ ! -e "$1" ]; then
            echo "Database environment file does not exist."
            exit
        fi
        DBENV="$1"
        shift
        ;;
    -c | --config)
        shift
        if [ ! -e "$1" ]; then
            echo "Config file does not exist."
            exit
        fi
        CONFIG="$1"
        shift
        ;;
    -h | --help)
        echo "$HELP_MAIN" && exit
        shift
        ;;
    -v | --verbose)
        VERBOSE=true
        shift
        ;;
    --)
        shift
        break
        ;;
    *) break ;;
    esac
done

verbose_echo "Parsed as"
verbose_echo "$OPTS"
verbose_echo "Recieved options:"
verbose_echo "\$ENV: $ENV"
verbose_echo "\$CONFIG: $CONFIG"
verbose_echo "\$VERBOSE: $VERBOSE"
verbose_echo "\$COLOR: $COLOR"
if [ "$1" != "" ]; then verbose_echo "\$1: $1"; fi

# Note, this will not grab multiline strings. If an issue is filed on this, will fix.
prepare() {
    verbose_echo "Running prepare function"
    clear
    pgkeys=("POSTGRES_USER" "POSTGRES_DB" "POSTGRES_PASSWORD" "POSTGRES_INITDB_ARGS" "POSTGRES_INITDB_WALDIR" "POSTGRES_HOST_AUTH_METHOD" "PGDATA")
    for key in "${pgkeys[@]}"; do
        grep "$key" "$CONFIG" >>"$DBENV"
    done

    cp "$DBENV" frontend
    cp "$DBENV" backend

    cp "$CONFIG" backend

    if [ ! -d "frontend/node_modules" ]; then
        cd frontend && npm install && cd ..
    fi
}

run() {
    verbose_echo "Running run function"
    if [ "$ENV" = "prod" ]; then
        docker compose -f "$PROD_DPSEC" up
    elif [ "$ENV" = "dev" ]; then
        docker compose -f "$DEV_DSPEC" up
    else
        docker compose -f "$ENV" up
    fi
}

clear_dbe() {
    verbose_echo "Clearing $DBENV"
    echo -n "" >"$DBENV"
}

build() {
    prepare
    verbose_echo "Running build function"
    if [ "$ENV" = "prod" ]; then
        docker compose -f "$PROD_DPSEC" build
    elif [ "$ENV" = "dev" ]; then
        docker compose -f "$DEV_DSPEC" build
    else
        docker compose -f "$ENV" build
    fi
}

if [ "$#" = 0 ]; then
    verbose_echo "\$# is 0"
    prepare
    run
else
    case "$1" in
    "prepare") prepare ;;
    "start") run ;;
    "build") build ;;
    "help") echo "$HELP_MAIN" && exit ;;
    "clear") clear_dbe && echo "Cleared database environment file." && exit ;;
    esac
fi
