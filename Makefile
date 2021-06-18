CODE_ENTRY=cmd/main.go
APISVR_DIR=apisvr
APISVR_NAME=apisvr
APISVR_PORT=3001
WEBAPP_DIR=webapp
WEBAPP_PORT=3000
WEBAPP_NAME=webapp

.PHONY: test-simulator
test-simulator:
	cd $(APISVR_DIR)/service/wjsimulatorsvc ;\
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

.PHONY: docker-build-apisvr
docker-build-apisvr:
	cd $(APISVR_DIR) ;\
	docker build . -t $(APISVR_NAME)

.PHONY: docker-run-apisvr
docker-run-apisvr:
	docker run --rm -p $(APISVR_PORT):$(APISVR_PORT) $(APISVR_NAME)

.PHONY: docker-rm-apisvr
docker-rm-apisvr:
	docker rmi $(APISVR_NAME)
