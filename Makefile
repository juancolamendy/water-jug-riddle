compile-service:
	cd lib-service ;\
	go build ./...

test-simulator:
	cd lib-service/wjsimulatorsvc ;\
	go test -v -run TestSimulator
