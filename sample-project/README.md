src:https://www.youtube.com/watch?v=XCsL89YtqCs

To tell the go tool about this workspace, set the go path environment variable
~$ export GOPATH=$HOME/{pwd}/sample-project
~$ go env GOPATH (visualize current)


Since we are going to install and run a bunch of fo programs in the foreseeable future, it makes sense to add the bin
directory to our system path

export PATH=$HOME/sample-project/bin:#PATH

So now i can just type hello to run the hello world program



src
Source code resides in the src/ directory of your workspace
mkdir -p src/unique_namespace

namespace
A unique base path inside which all of your go code will reside

tests
The test for a package live in that packages directory, in files that end in *_test


Commands
go install (execute next to hello.go)
(Executable) The go tool has installed it to the bin directory inside the workspace 
(Library) To make the string package available to other code, we need to install the package object to our workspace with go install. Now we can find the package object  inside the pkg 

go build
check that the file compiles. When used on go packages, go build just builds the code and throws away the output. 

go get