package packageB

//go:generate go run github.com/algorand/oapi-codegen/cmd/oapi-codegen --old-config-style -generate types,skip-prune,spec --package=packageB -o externalref.gen.go spec.yaml