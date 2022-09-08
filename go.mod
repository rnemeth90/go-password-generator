module github.com/rnemeth90/go-password-generator

go 1.18

require github.com/spf13/cobra v1.5.0

require github.com/rnemeth90/generator v0.0.0-00010101000000-000000000000

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/rnemeth90/generator => ./generator
