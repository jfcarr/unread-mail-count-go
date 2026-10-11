OUTPUT_EXE=unread-mail-count

default:
	@echo "Targets:"
	@echo "  run"
	@echo "  build"
	@echo "  clean"

deploy: build
	cp ./$(OUTPUT_EXE) $(HOME)/bin
	cp ./$(OUTPUT_EXE)-live.json $(HOME)/bin/$(OUTPUT_EXE).json

run: build
	./$(OUTPUT_EXE)

build:
	go build

clean:
	go clean