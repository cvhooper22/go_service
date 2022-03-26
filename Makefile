BINARY_NAME=go-service

build:
	go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}