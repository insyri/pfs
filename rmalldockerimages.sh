# chmod +x ./rmalldockerimages.sh
docker rmi -f $(docker images -aq)