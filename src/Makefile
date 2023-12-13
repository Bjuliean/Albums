SRC=main.go
APPNAME=program

all: clean
	cd assembling && sudo docker-compose up -d
	sudo netstat -tulpn | grep :5432
	go run $(SRC)

pqstart:
	cd assembling && sudo sh server_and_db.sh

b:
	./$(APPNAME)

clean:
	rm -rf $(APPNAME)
	cd assembling && sudo docker-compose down