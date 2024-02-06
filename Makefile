
BINARY=pushover

${BINARY}: *.go
	go mod tidy
	go mod vendor
	go build -o ${BINARY} .

dist: ${BINARY}
	install -D pushover dist/pushover 
clean:
	rm -rf dist ${BINARY}

