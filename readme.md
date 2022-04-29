# GOLAND SANDBOX README

## THIS REPO IDEA

Idea for this repository is to learn golang together in monorepo environment, which allow us to code review our code and
let to see potentially solutions. Golang has easy and dev-friendly monorepo building process. Just create own project
root folder and have fun.

## GOLANG INSTALLATION

https://go.dev/doc/install

PREFERABLE IDE: GOLAND

## CREATE OWN PROJECT

Goland has predefined options to that, but if you don't have goland use this steps

1. Create folder for your project
2. Create main.go
3. Fill main.go with

```
package main

func main() {

}
```

4. Create go.mod file and fill with

```
module example

go 1.17 (or another version)
```

5. Write your code

## Helpful commands

``go ...``

```
        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

```