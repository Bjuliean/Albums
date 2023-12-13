SRC=main.go
APPNAME=program

all: clean
	cd assembling && sudo sh server_and_db.sh
	go run $(SRC)

clean:
	rm -rf $(APPNAME)
	cd assembling && sudo docker-compose stop

push:
	cd .. && git add *
	git commit -m "autopush"
	git push origin develop