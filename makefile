WIRECMD=wire
GOCMD=go

.PHONY: wire

wire:
	@command -v wire >/dev/null 2>&1 || $(GOCMD) install github.com/google/wire/cmd/wire@latest
	@cd cmd/api && $(WIRECMD)