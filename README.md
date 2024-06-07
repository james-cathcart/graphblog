# Graph Blog
This project exists to explore GraphQL via the Go programming language.

## Project Creation
The project directory was created and the go module initialized
```
go mod init graphblog
```
Create a tools.go file in the root directory of the project and include the following dependencies.
```
//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
```
Tidy up
```
go mod tidy
```
Then the GraphQL bootstrapping/generation was done
```
go run github.com/99designs/gqlgen init
```
Some modifications in the directory structure were also made to support the commonly accepted project structure for Go 
projects


## Development Reference
### Regenerate GraphQL Files
```
go run github.com/99designs/gqlgen generate
```
