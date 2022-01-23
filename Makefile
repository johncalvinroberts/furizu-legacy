FE_DIR=client
BE_ENTRYPOINT=main.go
AIR_BIN=./bin/air
BIN=tmp/furizu


build: build-fe build-be

install: install-be install-fe install-air
dev: 
	make -j 2 dev-fe dev-be

install-air: 
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

install-be: 
	go get ./src

build-be:
	go build -o $(BIN) $(BE_ENTRYPOINT)

build-fe:
	cd $(FE_DIR); npm run build;

install-fe:
	cd $(FE_DIR); npm install;

dev-fe:
	cd $(FE_DIR); npm run dev;

dev-be:
	export GIN_MODE=release
	$(AIR_BIN) $(BE_ENTRYPOINT)

run-furizu: 
	$(BIN)

clean:
	rm -rf $(BIN)
	rm -rf tmp/main
	rm -rf $(FE_DIR)/build

gql:
	go run scripts/gqlgen.go generate