export PATH=$(go env GOPATH)/bin:$PATH
swag init -g apidoc.go -o "./api" -d  "./api","./internal/transport/rest/","./internal/models"