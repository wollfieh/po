
BINARY=pushover

${BINARY}:
  go mod tidy
	go build -o ${BINARY} .

dist: ${BINARY}
	install -D pushover dist/pushover 
clean:
	rm -rf dist ${BINARY}

