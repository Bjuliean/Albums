SRC=main.go
APPNAME=program

all: clean
	cd assembling && sudo sh server_and_db.sh
	go run $(SRC)

clean:
	rm -rf $(APPNAME)
	cd assembling && sudo docker-compose stop