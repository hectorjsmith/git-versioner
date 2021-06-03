go/downloadDependencies:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependencies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

format:
	go fmt $(go list ./... | grep -v /vendor/)
	go vet $(go list ./... | grep -v /vendor/)

define prepare_build_vars
    $(eval VERSION_FLAG=-X 'main.version=$(shell git describe --tags)')
endef

build/dev:
	$(call prepare_build_vars)
	go build -a --ldflags "${VERSION_FLAG}" -o dist/git-versioner ./versioner.go

build/snapshot:
	./tools/goreleaser_linux_amd64 --snapshot --rm-dist --skip-publish

build/release:
	git --no-pager diff
	./tools/goreleaser_linux_amd64 --rm-dist --skip-publish

docs/generateChangelog:
	./tools/git-chglog_linux_amd64 --config tools/chglog/config.yml v0.1.0.. > CHANGELOG.md

