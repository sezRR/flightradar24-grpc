# Variables
PROTOC = protoc
SCRIPT_DIR = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PROTO_DIR = $(SCRIPT_DIR)../protos
OUT_DIR = $(SCRIPT_DIR)../
PROTO_FILES = $(shell find $(PROTO_DIR) -name "*.proto")

.DEFAULT_GOAL := generate

# Generate .pb.go files
generate:
	@echo "Found proto files: $(PROTO_FILES)"
	@$(foreach file, $(PROTO_FILES), $(PROTOC) -I=$(PROTO_DIR) --go_out=$(OUT_DIR) $(file);)

# Clean generated files (optional)
clean:
	find $(PROTO_DIR) -name "*.pb.go" -exec rm -f {} +

# Add phony targets
.PHONY: generate clean
