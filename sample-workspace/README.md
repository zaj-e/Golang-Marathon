## Folders

| Item        | Purpose           |
| :------------- |:------------- |
| src            | Source code resides in the src/ directory of your workspace |
| namespace      | A unique path i nside which all of your go code will reside      |
| tests | The test for a packagelive in that packages directory, in all files that end in *_test      |
___
## Workspace
___
To tell the go tool about this workspace, set the go path environment variable
```
$ export GOPATH=$HOME/{pwd}/sample-project
```
#### Imports
Go first searches for package directory inside GOROOT/src directory and if it doesn’t find the package, then it looks for GOPATH/src. Since, fmt package is part of Go’s standard library which is located in GOROOT/src, it is imported from there. But since Go cannot find "foo" package inside GOROOT, it will lookup inside GOPATH/src and we have it there.

___
Visualize current
```
$ go env GOPATH 
```
___
If you are going to install and run a bunch of fo programs in the foreseeable future, it makes sense to add the bin directory to our system path. After this you canjust type hello and run the hello world program.

```
export PATH=$HOME/sample-project/bin:#PATH
```
___
To be honest i'm not sure what the -p does
```
mkdir -p src/unique_namespace
```
___


## Commands
The ./... pattern matches all the packages in a module. So when you run go build ./... or go test ./..., you are building or testing all the packages inside a module (including module itself if it is a package). These commands also download and install dependencies inside go.mod file.
___
```
go build
```
Check that the file compiles. When used on go packages, go build just builds the code and throws away the output. 
___
```
go install 
```
<ins>Executable:</ins> The go tool has installed it to the bin directory inside the workspace \
<ins>Library</ins>: To make the string package available to other code, we need to install the package object to our workspace with go install. Now we can find the package object  inside the pkg 
___
```
go mod graph
```
Shows you the dependencies of your module
___
```
go mod vendor
```
Sometimes, when you are running automated tests for your project (executable module like main in our case), there is a chance that your test machine may face some network issues while downloading the dependencies or when a test machine is running in an isolated environment. In such cases, you need to provide dependencies beforehand. This is called vendoring. When you use go build, it looks for the dependencies inside module cache directory but you can force Go to use dependencies from vendor directory of the module using the command
```
go build -mod vendor
```
___
```
 GO111MODULE=on go install <package>@<version>
 ```
 To install a CLI application (module) from anywhere in the terminal, use the following command. This will install a module inside the module cache directory and generate a binary executable file inside $GOPATH/bin directory.Even though you could install multiple versions of a module, you can only have a single binary executable file. So the module cache directory will have source code of multiple versions, however, the bin directory will have only one instance of the module binary.

## Reference
- https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16
