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