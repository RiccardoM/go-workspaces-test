# Go workspaces test
This repository contains a very simple [Go workspaces](https://go.dev/blog/get-familiar-with-workspaces) test. 

The main program is located inside the `main/main.go` file and it depends on a generic `github.com/riccardom/workspaces/chains` package that is not defined inside the `go.mod` file. Instead, this dependency is injected using Go workspaces with the `go.work` file that is generated using the `go work init` command. This works as follow: 

- all packages (`akash`, `default`, `desmos`) have all the same module name within their `go.mod` file: 
  ```
  module github.com/riccardom/workspaces/chain
  ```
- the `go.work` file specifies which module should be used within the `main` package to inject the `github.com/riccardom/workspaces/chain` dependency
- when `main/main.go` is run, the proper dependency will be called and thus the output will change based on the loaded module

Example: 
```
# Init the program to use the default module
$ go work init ./main ./default

# Run the program
$ go run main/main.go
Hello from default!

# Switch module
$ go work edit -dropuse default -use akash

# Run the program again - The output changes!
$ go run main/main.go
Hello from Akash!
```


## Commands
### Init the workspace

```
go work init ./default
```

### Use a different module

```
go work edit -dropuse <used module> -use <new module>

# Eg. Switch from "default" to "akash"
# go work edit -dropuse default -use akash

# Eg. Switch from "akash" to "desmos"
# go work edit -dropuse akash -use desmos
```

### Run the program
```
go run main/main.go
```