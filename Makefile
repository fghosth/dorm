APP = createProject
MAINDIR = ./
default:
	@echo 'Usage of make: [ build | linux_build | windows_build | clean ]'

build:
	@go build -o ./dist/$(APP) $(MAINDIR)

linux_build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/$(APP) $(MAINDIR)

windows_build:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/$(APP).exe $(MAINDIR)

run: build
	@./dist/$(APP) 

install: build
	@mv ./dist/$(APP) $(GOPATH)/bin

clean:
	@rm -f ./dist/*
