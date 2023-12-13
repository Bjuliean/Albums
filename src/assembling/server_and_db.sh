#!bin/bash

clear

echo "\n\n\n"
echo "To run the program, the script will release ports 8080 and 5432."
echo "Allow? (y/n)"

read symbol

if [ $symbol = 'y' ] || [ $symbol = 'Y' ] 
then
    echo "RELEASING PORTS..."
    ./cleanports.sh
    echo "STARTING POSTGRES CONTAINER..."
    sudo docker-compose up -d
    until [ "`docker inspect -f {{.State.Running}} pq_album_container`"=="true" ]; do
        sleep 0.1;
    done;

    echo "STARTING SERVER..."
else
    echo "\nInstallation cancelled\n"
fi