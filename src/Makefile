SRC=main.go
APPNAME=program

all: clean
	go run $(SRC)

b: clean
	go build -o $(APPNAME) $(SRC)
	./$(APPNAME)

clean:
	rm -rf $(APPNAME)