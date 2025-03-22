# Go Workspace and Multi-Module Example

This project is designed to demonstrate how Go's workspace and multi-module setup works. It includes multiple modules, each with its own functionality, and a workspace configuration that ties everything together.

## Project Structure

The project consists of the following directories:

```python
.
├── go.work               # The Go workspace configuration
├── hello                 # The 'hello' module
│   ├── go.mod            # The Go module file for the 'hello' module
│   ├── lib               # The library containing functions used in 'hello'
│   │   └── hello.go      # The 'hello' function
│   └── main.go           # The entry point for the 'hello' module
├── hello-world           # The 'hello-world' module that uses 'hello' and 'world'
│   ├── go.mod            # The Go module file for the 'hello-world' module
│   └── hello-world.go    # The entry point for the 'hello-world' module
└── world                 # The 'world' module
    ├── go.mod            # The Go module file for the 'world' module
    ├── lib               # The library containing functions used in 'world'
    │   └── world.go      # The 'world' function
    └── main.go           # The entry point for the 'world' module
```


### Modules Overview

- **`hello`**: This module contains a library (`lib/hello.go`) and a `main.go` file. The `main.go` is the entry point for the module, which prints a message to the console. The `lib/hello.go` contains a function (`SayHello`) that can be used by other modules.

- **`world`**: This module is similar to `hello` but prints a message related to the world. It contains a library (`lib/world.go`) and a `main.go` file. The `lib/world.go` contains a function (`SayWorld`) that can be used by other modules.

- **`hello-world`**: This module uses both the `hello` and `world` modules to print both messages. It imports the functions from the `hello` and `world` modules and calls them in its `hello-world.go` file.

### Workspace Configuration

The `go.work` file is used to define a Go workspace, allowing multiple modules to work together. It references the `hello`, `hello-world`, and `world` modules, enabling you to work with them as a single project.

```go
go 1.23.1

use (
	./hello
	./hello-world
	./world
)
```

### How to Set Up the Project

To set up the project locally, follow these steps:
Clone the repository to your local machine:

```sh
git clone https://github.com/amanshaw4511/go-multi-module.git
cd go-multi-module
```

Make sure you have Go 1.18 or later installed. You can check your Go version by running:

```sh
go version
```
Run the Go modules tidy command to download dependencies:

```sh
go mod tidy
```

Run the hello-world module to see the combined output from both the hello and world modules:

```sh
go run ./hello-world
```

This should output:

    Hello from the hello module!
    Hello from the world module!

#### Running Individual Modules


To run the hello module:
```sh
go run ./hello
```

To run the world module:
```sh
go run ./world
```

### Go Workspace and Multi-Module Workflow

In Go, a workspace is a collection of Go modules that can work together. The go.work file helps manage this workspace by defining which modules should be included. This setup allows you to:
- Develop multiple modules concurrently without the need to publish them to a repository.
- Share code between modules by importing packages from other modules within the workspace.
- Use go run or go build to work with the entire workspace or individual modules.

This project demonstrates how to structure a Go workspace with multiple modules, helping you better understand how Go handles dependencies and project organization in more complex scenarios.

### Understanding repositoy, modules, package
A module is a collection of related Go packages that are versioned together as a single unit.

Modules record precise dependency requirements and create reproducible builds.

Most often, a version control repository contains exactly one module defined in the repository root. (Multiple modules are supported in a single repository, but typically that would result in more work on an on-going basis than a single module per repository).

Summarizing the relationship between repositories, modules, and packages:

    A repository contains one or more Go modules.
    Each module contains one or more Go packages.
    Each package consists of one or more Go source files in a single directory.

source of the Quote: https://go.dev/wiki/Modules#modules
