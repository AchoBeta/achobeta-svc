lint:
	@echo "Checking protobuf..."
	@buf lint ./proto
	@echo "Check done"

build:
	@rm -rf ./out
	@ buf generate --template ./buf.gen.yaml
	@echo "Build done"
