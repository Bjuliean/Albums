SRC=main.go
TESTS=./handler/
APPNAME=program
COVFILE=cover.out

all: clean db
	go run $(SRC)

clean:
	rm -rf $(APPNAME) $(COVFILE)
	cd assembling && sudo docker-compose down

push:
	cd .. && git add *
	git commit -m "autopush"
	git push origin develop

test: clean db
	go test $(TESTS)

coverage: clean db
	go test -coverprofile=$(COVFILE) $(TESTS)
	go tool cover -html=$(COVFILE)

db:
	cd assembling && sudo sh server_and_db.sh


