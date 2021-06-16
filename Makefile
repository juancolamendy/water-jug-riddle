compile-service:
	cd lib-service ;\
	go build ./...

test-simulator:
	cd lib-service/wjsimulatorsvc ;\
	go test -v -run TestSimulator

run-api-dev:
	cd apisvr ;\
	go run cmd/main.go
