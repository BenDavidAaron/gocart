module gocart

go 1.15

require (
	github.com/bendavidaaron/gocart v0.0.0-20211030165947-8e8850e509c7
	github.com/google/subcommands v1.2.0
	internal/gocart v0.0.1
)

replace internal/gocart => ./internal/gocart
