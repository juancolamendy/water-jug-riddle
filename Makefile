CODE_ENTRY=cmd/main.go
SIMULATOR_DIR=lib-service
APISVR_DIR=apisvr
WEBAPP_DIR=webapp
WEBAPP_PORT=3000
WEBAPP_NAME=webapp

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

.PHONY: docker-build-webapp
docker-build-webapp:
	cd $(WEBAPP_DIR) ;\
	docker build . -t $(WEBAPP_NAME)

.PHONY: docker-run-webapp
docker-run-webapp:
	docker run --rm -p $(WEBAPP_PORT):$(WEBAPP_PORT) $(WEBAPP_NAME)

.PHONY: docker-rm-webapp
docker-rm-webapp:
	docker rmi $(WEBAPP_NAME)
