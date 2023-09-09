module github.com/NikSays/go-maib-ecomm

go 1.20

require (
	github.com/google/go-querystring v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/stretchr/testify v1.8.4
	software.sslmate.com/src/go-pkcs12 v0.2.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.3.0 // Parsing was broken
)
