# Go Workspace
This file will discuss Go's (opinionated) project file structure as well as concepts such as GOPATH, GOROOT and modules
 
## File Structure

 - Overall Directory
	 - bin
	 - pkg
	 - src
		 - github.com
			 - JordanInsights (i.e. github.com username)
				 - folder with code for project / repo
				 - folder with code for project / repo
				 - folder with code for project / repo
				 - folder with code for project / repo

## Commands
### Go get
```go get``` is one of the Go commands. Funnily enough, it's used to get go packages

## Environment Variables
### Go Path
Go Path (```GOPATH```) is an environment variable that points to the go workspace that you want to work in. 

This should point at whichever directory you wish to develop from. 

### Go Root
Go Root (```GOROOT```) is another environment variable that points at the binary which contains everything you need to run Go. 

This should point at the filepath where you installed Go. 

## Modules
Go 1.11 introduced Modules. This approach is the default build mode since Go 1.16, therfore the use of ```GOPATH``` is not recommended. 

Modules aim to solve problems related to dependency management, version selection and reproducible build; they also enable users to run Go code outside of ```GOPATH```. 

Using modules is pretty straightforward. Select any directory outside ```GOPATH``` as the root of your project and create a new module with the ```go mod init``` command. 

A ```go.mod``` file will be generated, containing the module path, a Go version, and its dependency requirements, which are the other modules needed for a successfuk build. 

If no ```<modulepath>``` is specified, ```go mode init``` will try to guess the module path from the directory structure. It can also be overridden by supplying an argument


    mkdir my-project
    cd my-project
    go mod init <modulepath>
 A ```go.mod``` file could look like this:
 

    module cmd
    
    go 1.16
The built-in documentation provides an overview of all available ```go mod``` commands

    go help mod
    go help mod init
