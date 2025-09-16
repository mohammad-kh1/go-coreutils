BUILD_DIR := ./bin
CMD_DIR := ./cmd

COMMANDS := $(notdir $(wildcard $(CMD_DIR)/*))

GO := go
.DEFAULT_GOAT := build

build: ## Build all commands
	@for cmd in $(COMMANDS); do \
		echo "Building $$cmd..."; \
		mkdir -p $(BUILD_DIR); \
		$(GO) build -o $(BUILD_DIR)/$$cmd $(CMD_DIR)/$$cmd; \
	done
	@echo "All commands built successfully!"


# make run CMD=command
run:
ifndef CMD
		@echo "Please specify a command to run. Example: make run CMD=command"
		@exit 1
endif
	@$(BUILD_DIR)/$(CMD)

install:
	@echo "building"
	@for cmd in $(COMMANDS) ; do \
		$(GO) install $(CMD_DIR)/$$cmd; \
	done

# Run tests
test:
	@echo "TEST"

# Clean build dir
clean:
	@rm -rf $(BUILD_DIR)

help:
	@echo "  make build        - Build all commands"
	@echo "  make run CMD=name - Run a specific command (e.g., make run CMD=rev)"
	@echo "  make install      - Install all binaries to GOPATH/bin"
	@echo "  make test         - Run all tests"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make help         - Show this help message"
