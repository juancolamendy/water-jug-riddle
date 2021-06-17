CODE_ENTRY=cmd/main.go
SIMULATOR_DIR=lib-service
APISVR_DIR=apisvr
WEBAPP_DIR=webapp

.PHONY: compile-service
compile-service:
	cd $(SIMULATOR_DIR) ;\
	go build ./...

.PHONY: test-simulator
test-simulator:
	cd $(SIMULATOR_DIR)/service/wjsimulatorsvc ;\
	go test -v -run TestSimulator

.PHONY: run-api-dev
run-api-dev:
	cd $(APISVR_DIR) ;\
	go run $(CODE_ENTRY)

.PHONY: install-webapp-dev
install-webapp-dev:
	cd $(WEBAPP_DIR) ;\
	rm -fr node_modules ;\
	npm i

.PHONY: run-webapp-dev
run-webapp-lint:
	cd $(WEBAPP_DIR) ;\
	npm run lint

.PHONY: run-webapp-dev
run-webapp-dev:
	cd $(WEBAPP_DIR) ;\
	npm run dev
