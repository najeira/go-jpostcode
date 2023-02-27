GOTEST=go test -v
TESTPKGS=$(shell go list ./... | grep -v -e example -e statik)

test:
# exclude TestAll_
	$(GOTEST) -race -run='^Test_' $(TESTPKGS)

test/all:
	$(GOTEST) $(TESTPKGS)

test/all/race:
	$(GOTEST) -race $(TESTPKGS)

update:
	cd jpostcode-data && git pull origin master
	go run script/gobdump/main.go

dump:
	rm -f ./internal/data/*.go
	go run script/vardump/main.go
