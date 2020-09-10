# Client library generator

# Prerequisite
```
alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
export API_VERSION=v0.2
export PUBLIC_API_DOCS=<somewhere>
export API_KEY=<something>
export ORGANIZATION=<something>
export PROJECT=<something>
```

## Steps
```
go run transformer/transformer.go rename-operation -i $PUBLIC_API_DOCS -o temp/renamed-op.json
go run transformer/transformer.go fix-content-type -i temp/renamed-op.json -o temp/fix-content-type.json
go run transformer/transformer.go basic-response -i temp/fix-content-type.json -o temp/basic-resp.json
swagger generate client -f temp/basic-resp.json --target temp -A lightstep-api --template-dir templates/
go run generator/generator.go json-response --apikey $APIKEY -o temp/json --org $ORGANIZATION -p $PROJECT
go run generator/generator.go struct-response -i temp/json -o ../pkg/v0.2/response
go run transformer/transformer.go full-response -i temp/fix-content-type.json -o temp/full-resp.json  -p ../pkg/v0.2/response/
swagger generate client -f temp/full-resp.json --target ../pkg/v0.2/ -A lightstep-api --template-dir templates/
```