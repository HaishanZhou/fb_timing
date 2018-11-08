EXEC = $(shell basename $(shell pwd))

$(EXEC):
	@go build -o $(EXEC)

install:
	install $(EXEC) ~/usr/bin
