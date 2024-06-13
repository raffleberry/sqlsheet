EXECUTABLES = go
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

sqlsheet:
	echo "Hello sqlsheet"

build: cmd/app/app.go
	go build -o build/app cmd/app/app.go

test:
	go test -v .

clean:
	rm build/*

run:
	export DEV=1 && go run cmd/app/app.go

