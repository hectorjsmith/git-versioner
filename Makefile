install-deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependecies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

define prepare_build_vars
    $(eval VERSION_FLAG=-X 'main.appVersion=$(shell git describe --tags)')
endef

build/local:
	$(call prepare_build_vars)
	go build -a --ldflags "${VERSION_FLAG}" -o build/git-versioner.bin ./versioner.go
