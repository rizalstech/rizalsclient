package rizalsclient

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -package rizalsclient -o ./types.gen.go ../api/api/dist/api.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -package rizalsclient -o ./client.gen.go ../api/api/dist/api.json
