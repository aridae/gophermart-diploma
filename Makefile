LOCALBIN := ${PWD}/bin

.PHONY: smartimports
smartimports: export SMARTIMPORTS := ${LOCALBIN}/smartimports
smartimports:
	test -f ${SMARTIMPORTS} || GOBIN=${LOCALBIN} go install github.com/pav5000/smartimports/cmd/smartimports@latest
	PATH=${PATH}:${LOCALBIN} ${SMARTIMPORTS} -path . -exclude ./static ./../_mock -local github.com/aridae/go-metrics-store

.PHONY: generate-mocks
generate-mocks: export MOCKGEN := ${LOCALBIN}/mockgen
generate-mocks:
	test -f ${MOCKGEN} || GOBIN=${LOCALBIN} go install go.uber.org/mock/mockgen@latest
	PATH=${PATH}:${LOCALBIN} go generate -run mockgen $(find . -d -name '_mock')

.PHONY: lint
lint: export GOLANGCILINT := ${LOCALBIN}/golangci-lint
lint:
	test -f ${GOLANGCILINT} || GOBIN=${LOCALBIN} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	PATH=${PATH}:${LOCALBIN} ${GOLANGCILINT} run

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/gophermart cmd/gophermart/main.go

.PHONY: autotest
autotest: build
	./bin/gophermarttest \
		 -test.v -test.run=^TestGophermart$$ \
		 -gophermart-binary-path=bin/gophermart \
		 -gophermart-host=localhost \
		 -gophermart-port=8081 \
		 -gophermart-database-uri="***" \
		 -accrual-binary-path=cmd/accrual/accrual_darwin_amd64 \
		 -accrual-host=localhost \
		 -accrual-port=9000 \
		 -accrual-database-uri="***"

perm:
	chmod -R +x bin
