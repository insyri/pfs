# Default files
COMPOSE_FILE="docker-compose.yml"
CONFIG="pfs.example.toml"

while [ True ]; do
if [ "$1" = "--dev" -o "$1" = "-d" ]; then
    COMPOSE_FILE="dev-docker-compose.yml"
    shift 1
elif [ "$1" = "--config" -o "$1" = "-c" ]; then
    CONFIG=$2
    shift 2
    if [ ! -e $CONFIG  ]; then
        echo "Config file does not exist."
        exit
    fi
else
    break
fi
done

cp $CONFIG backend
cp $CONFIG nginx/conf
cp $CONFIG frontend

docker compose -f $COMPOSE_FILE up