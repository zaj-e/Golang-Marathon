## Links
### Documentation
[Package Reference](https://golang.org/doc/) \
[Language Specification](https://golang.org/ref/spec)

### Good Reads
[Effective Go](https://golang.org/doc/effective_go.html) \
[Packages](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)\
[Modules](https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16)

### Videos
[Channels](https://www.youtube.com/watch?v=KBZlN0izeiY)\
[Mutex](https://www.youtube.com/watch?v=cjMdUmfzQWs)\
[Install Packages](https://www.youtube.com/watch?v=Kw1ZVXF7s5o)\
[Code Structure](https://www.youtube.com/watch?v=MzTcsI6tn-0)
[Modules](https://www.youtube.com/watch?v=B0EjcYaBm9A)


## Naming Convention
### Variable
- CamelCase, not snake_case
- Single letters to represate indexes (i, j, k)
- Use short and descriptive names for other variables
- Use a variable name that is long enough to understand and not so long that you don't want to type it

### Packages
- Avoid a package-level function name that repeats the package name (log.Info() instead of log.LogInfo())
- Separate your code into logical concerns. If the package deals with multiple types, keep the logic for each type in its own source file:

- If you have a type called an order then you put the orders type and the orders interfaces in orders.go
```
package: postgres
orders.go
suppliers.go
products.go
```

- Go community recommends to use plain and simple names for packages. For example, strutils for string utility functions or http for HTTP requests related functions. A package names with under_scores, hy-phens or mixedCaps should be avoided.

- Like main function, init function is called by Go when a package is initialized. It does not take any arguments and doesn’t return any value. The init function is implicitly declared by Go, hence you can not reference it from anywhere (or call it like init()). You can have multiple init functions in a file or a package. Order of the execution of init function in a file will be according to the order of their appearances. After all the init functions are executed, main function is called. Hence, the main job of init function is to initialize global variables that cannot be initialized in the global context. For example, initialization of an array. Since, for syntax is not valid in package scope, we can initialize the array integers of size 10 using for loop inside init function.

<div align=left ><img src="https://miro.medium.com/max/913/1*7opoqANNzsqLyHViQaK-tg.png"/>
<div align=left ><img src="https://miro.medium.com/max/913/1*UswvbXJpnMDT-kkeAqnzZw.png"/>


### Modules
Go provides a new feature called Modules which provides flexibility to manage your project and dependencies with ease. If you are working on Go packages, then you should consider relocating your project to Go Modules.

### Functions and methods
- Go doesn't have setters and getters

### Interfaces
- If your interface has only one function, append "-er" to the function name. (ex. Stringer)
- If your interface has more than one function, use a name to represent its functionality (ex. CustomerStorage)

### Source Code 
- In the package that defines your domain objects, define the types and interfaces for each object in the same source file 
```
package: inventory
orders.go
--  contains Orders type
    and OrderStorage interface
```

### Comments 
- Make your comments full sentences that end in a period, it improves readability

## Glossary

| Term           | Definition     |
| :------------- | :------------- |
|  Package       | A collection of source files in the same directory that are compiled together    |
|  Module        | A collection of related Go packages that are released together |
|  Repository    | A repository contains one or more modules. A Go repository typically contains only one module, located at the root of the repository.|
|  go.mod        | A file/module that declares the module path: the import path prefix for all packages within the module.|




<div align=left ><img src="https://cdn.filestackcontent.com/E6wA9QGnQUmsLrnJvHLM"/>