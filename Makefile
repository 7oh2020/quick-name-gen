MAIN_FILE="./src/cmd/main/main.go"
TEST_FILE="./src/cmd/test/main.go"
DIST_DIR="./dist"

.Phony: fmt
fmt:
	@gofmt -l -s -w .
	@echo "Done Format complete."

.PHONY: runtest
runtest:
	@go run $(TEST_FILE) "Hello World"

.Phony: run
run:
	@go run $(MAIN_FILE)

.Phony: build
build:
	@rm -rf $(DIST_DIR)
	@mkdir $(DIST_DIR)
	@go build -o $(DIST_DIR)/main $(MAIN_FILE)
	@echo "Done Build complete."
