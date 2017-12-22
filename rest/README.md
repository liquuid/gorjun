General Info to build swagger generated server

* All Go source code files, except `restapi/configure_api_specification_for_gorjun.go` will be overwriten by: `swagger generate server -f swagger.yml` command

* To run server:

  `go run cmd/api-specification-for-gorjun-server/main.go`

    you need provide a tls-certificate and key for https connections


* The Api documentations provided by swagger-ui will available in: `http[s]://localhost/[port]/docs`

* To generate and install a separated binary to serve rest api you will need build it with: `go install ./cmd/api-specification-for-gorjun-server`

