.PHONY: build clean

GO := go
SRC := main.go
BIN := stats

build: $(BIN)

$(BIN): $(SRC) metrics/metrics.go metrics/input.go
	$(GO) build -o $@

run: build
	./$(BIN)

test:
	@cd tests && go test -v

clean:
	rm -f $(BIN)
