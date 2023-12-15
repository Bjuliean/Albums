SRC=main.go
ALL_SRC=main.go handler/*.go storage/*.go logs/*.go
TESTS=./handler/
APPNAME=program
COVFILE=cover.out

all: clean dbdown dbup
	go run $(SRC)

build:
	go build -o $(APPNAME) $(SRC)

test: clean dbdown dbup
	go test $(TESTS)

coverage: clean dbdown dbup
	go test -coverprofile=$(COVFILE) $(TESTS)
	go tool cover -html=$(COVFILE)

clean:
	rm -rf logs/*.txt $(APPNAME) $(COVFILE)

cleanall: dbdown clean
	cd assembling && ./cleanports.sh

dbup:
	cd assembling && sudo sh server_and_db.sh

dbdown:
	cd assembling && sudo docker-compose down

style:
	gofmt -w $(ALL_SRC)


