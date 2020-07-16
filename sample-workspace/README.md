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
