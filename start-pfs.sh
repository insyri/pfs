#!/bin/bash

PROD_DPSEC="./docker-compose.yml"
DEV_DSPEC="./dev-docker-compose.yml"
CONFIG="./pfs.toml"
ENV="prod"
VERBOSE=false
COLOR="on"
HELP_MAIN="start-pfs

USAGE:
    start-pfs.sh <SUBCOMMAND> [FLAGS] [OPTION...]

SUBCOMMANDS:
    prepare                  Prepares the environment for pfs.
    start                    Prepares and starts the pfs service.

FLAGS:
    -h, --help               Shows help information.
    -e, --env    [dev|prod]  Use dev compose file.        default: prod
    -c, --config string      Path to configuration file.  default: pfs.toml
    -C, --color  [on|off]    Enable color output.         default: true

OPTIONS:
    -v, --verbose            Print verbose output.        default: false
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
            ENV="dev";
        elif [ "$1" != "prod" ]; then
            echo "Invalid environment. Valid environments are only dev and prod."
            exit 1 
        fi
        shift
        ;;
    -C | --color)
        shift
        if [ "$1" = "off" ]; then
            COLOR="off";
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

prepare() {
    verbose_echo "Running prepare function"
}

run() {
    verbose_echo "Running run function"
    if [ "$ENV" = "prod" ]; then
        docker compose -f "$PROD_DPSEC" up -d
    else
        docker compose -f "$DEV_DSPEC" up -d
    fi
}



if [ "$1" = "prepare" ]; then
    prepare
fi

# docker compose -f "$(if [ $ENV = "dev" ]; then echo $DEV_DSPEC; else echo "$PROD_DPSEC"; fi;)" up
