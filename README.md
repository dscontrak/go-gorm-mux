# Project in Go with GORM and MUX 

## Steps

Base on: https://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html

### 1. Create init with module

```sh
# Create module (Optional)
go mod init github.com/dscontrak/go-gorm-mux
# Download the dependencies automatically 
go mod tidy
# if the command above doen't work so you can use the following commands:
# Download the dependencies for this project (These dependencies is going to donwload to GOPATH)
go get -u github.com/gorilla/mux
go get -u gorm.io/driver/sqlite
go get -u github.com/jinzhu/gorm

```

### 2. Run and compile

```sh
# Run
go run main.go
# Build
go build
```

## Notes

By default the GOPATH is following folder: `echo $HOME/go`, you could modify with below commands:

```sh
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```

So, in short:

GOROOT is for compiler/tools that comes from go installation.
GOPATH is for your own go projects / 3rd party libraries (downloaded with "go get").
