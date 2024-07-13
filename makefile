WIRECMD=wire
GOCMD=go

.PHONY: wire

wire:
	@command -v wire >/dev/null 2>&1 || $(GOCMD) install github.com/google/wire/cmd/wire@latest
	@cd cmd/api && $(WIRECMD)

run:
	cd cmd/api/ && $(GOCMD) run main.go wire_gen.go

mocks-download:
	$(GOCMD) mod download
	$(GOCMD) install -mod=mod go.uber.org/mock/mockgen@latest

mocks-gen: mocks-download 
	@~/go/bin/mockgen -source=internal/infrastructure/web/repositories/weather-repository.go -destination=internal/usecase/mocks/weather-repository.go -typed=true -package=mock
	@~/go/bin/mockgen -source=internal/infrastructure/web/repositories/zipcode-repository.go -destination=internal/usecase/mocks/zipcode-repository.go -typed=true -package=mock

fmt:
	go fmt ./...

test-clean: fmt
	$(GOCMD) clean -testcache

tests: fmt test-clean
	$(GOCMD) test -cover -p=1 ./...
