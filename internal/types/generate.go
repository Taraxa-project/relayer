package types

//go:generate rm -f ./hash_types_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path hash_types.go
