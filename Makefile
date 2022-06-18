BINARY_NAME=prisonersdilemmago

all: build run

build:
	go build -o ${BINARY_NAME} ./examine.go ./game.go ./prisoner.go ./strategy*.go

run:
	go build -o ${BINARY_NAME} ./examine.go ./game.go ./prisoner.go ./strategy*.go
	./${BINARY_NAME}

clean:
	go clean
	rm ./${BINARY_NAME}
