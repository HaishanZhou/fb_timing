EXEC = $(shell basename $(shell pwd))

$(EXEC):main.go
	@go build -o $(EXEC)

install:
	install $(EXEC) ~/usr/bin
