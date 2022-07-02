#!/bin/bash

PROD_DPSEC="./docker-compose.yml"
DEV_DSPEC="./dev-docker-compose.yml"
CONFIG="./pfs.toml"
ENV="prod"
VERBOSE=false
HELP="start-pfs

USAGE:
    start-pfs.sh [FLAGS] [OPTION...]

FLAGS:
    -h, --help                Shows help information.
    -e, --env    [dev|prod]   Use dev compose file.          default: prod
    -c, --config FILE         Path to configuration file.    default: pfs.toml

OPTIONS:
    -v, --verbose             Print verbose output.          default: false

EXAMPLES:
    start-pfs.sh -c example.toml
    start-pfs.sh -c=example.toml
    start-pfs.sh --env=prod
    start-pfs.sh --config=spec.toml --verbose
"
OPTS=$(getopt -o e:c:vh --long env:,config:,verbose,help -n 'startpfs' -- "$@")

verbose_echo() {
    if [ "$VERBOSE" = true ]; then
        echo "[VERBOSE]" "$@"
    fi
}

spec_to_file() {
    if [ "$1" = "prod" ]; then
        return $PROD_DPSEC
    elif [ "$1" = "dev" ]; then
        return $DEV_DSPEC
    else
        echo "Invalid environment provided. Exiting."
        exit 1
    fi
}
# used to be $?
if [ ! "$OPTS" != 0 ]; then echo "Failed parsing options." >&2 ; exit 1 ; fi

eval set -- "$OPTS"

while true; do
    case "$1" in
        -e | --env ) shift; if [ "$1" = "dev" ]; then ENV="dev"; fi; shift ;;
        -c | --config ) shift; if [ ! -e "$1"  ]; then echo "Config file does not exist."; exit; fi; CONFIG="$1"; shift ;;
        -h | --help ) echo "$HELP" && exit; shift ;;
        -v | --verbose ) VERBOSE=true; shift ;;
        -- ) shift; break ;;
        * ) break ;;
    esac
done

FILE=$(spec_to_file $ENV)
docker compose -f "$FILE" up