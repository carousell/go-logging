module github.com/carousell/go-logging/example

go 1.18

replace (
	github.com/carousell/go-logging => ../
	github.com/carousell/go-logging/gokit => ../gokit
)

require (
	github.com/carousell/go-logging v0.0.0-20230322093349-63592a690170
	github.com/carousell/go-logging/gokit v0.0.0-20230322093349-63592a690170
)

require (
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
)
