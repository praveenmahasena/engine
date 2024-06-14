all: test
all: vet
all: package
all: engine_race


test: vet
test: base_test

base_test:
	go test ./... -v

vet:
	go vet ./...

package: engine

package_race: engine_race

engine:
	go build -o ./bin/engine .

engine_race:
	go build --race -o ./bin/engine_race .

run:
	./bin/engine
