#!/bin/bash

PROD_DPSEC="./docker-compose.yml"
DEV_DSPEC="./dev-docker-compose.yml"
CONFIG="./pfs.example.toml"
ENV="prod"
VERBOSE=false
COLOR="on"
HELP_MAIN="start-pfs

USAGE:
    start-pfs.sh [SUBCOMMAND] [FLAGS] [OPTION...]

SUBCOMMANDS:
    prepare               Prepares the docker environment for pfs.
    start                 Starts the dockerized pfs service.
    help                  Outputs the help information.
    clear                 Clears the ./database.env file.

FLAGS:
    -e, --env [dev|prod]  Use dev compose file.        default: prod
    -c, --config string   Path to configuration file.  default: pfs.toml
    -C, --color [on|off]  Enable color output.         default: true

OPTIONS:
    -h, --help            Shows help information.
    -v, --verbose         Print verbose output.        default: false
"

OPTS=$(getopt -o e:c:vhC: --long env:,config:,verbose,help,color: -n 'start-pfs' -- "$@")

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
    case "$1" in
    -e | --env)
        shift
        if [ "$1" = "dev" ]; then
            ENV="dev"
        elif [ "$1" != "prod" ]; then
            echo "Invalid environment. Valid environments are only dev and prod."
            exit 1
        fi
        shift
        ;;
    -C | --color)
        shift
        if [ "$1" = "off" ]; then
            COLOR="off"
        elif [ "$1" != "on" ]; then
            echo "Invalid color. Valid colors are only on and off."
            exit 1
        fi
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

verbose_echo "Recieved options:"
verbose_echo "\$ENV: $ENV"
verbose_echo "\$CONFIG: $CONFIG"
verbose_echo "\$VERBOSE: $VERBOSE"
verbose_echo "\$COLOR: $COLOR"
if [ "$1" != "" ]; then verbose_echo "\$1: $1"; fi

# Note, this will not grab multiline strings. If an issue is filed on this, will fix.
prepare() {
    verbose_echo "Running prepare function"
    pgkeys=("POSTGRES_USER" "POSTGRES_DB" "POSTGRES_PASSWORD" "POSTGRES_INITDB_ARGS" "POSTGRES_INITDB_WALDIR" "POSTGRES_HOST_AUTH_METHOD" "PGDATA")
    for key in "${pgkeys[@]}"; do
        grep "$key" "$CONFIG" >>database.env
    done

    cp database.env frontend
    cp database.env backend

    cp "$CONFIG" backend
}

run() {
    verbose_echo "Running run function"
    if [ "$ENV" = "prod" ]; then
        docker compose -f "$PROD_DPSEC" up
    else
        docker compose -f "$DEV_DSPEC" up
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
    "help") echo "$HELP_MAIN" && exit ;;
    "clear") echo -n "" >"database.env" && echo "Cleared database environment file." && exit ;;
    esac
fi
