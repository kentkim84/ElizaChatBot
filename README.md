GMIT Software Development 3rd year.

Module: Data Representation.

### Data Representation Project
##### This repository contains some codes written in the programming language Go.
##### The author is Yongjin Kim, a Student at GMIT.
---
This project is to create a chatbot web application in Go. Users will be able to visit the web application through their browser, type in sentences such as “Hi, my name is John.” and the web application will reply with sentences such as “Hello John, how are you?”.

## How to clone
1. Open **Git Bash**.
2. Change the current working directory to the location where you want the cloned directory to be made.
3. Type 'git clone', and then paste the URL `https://github.com/kentkim84/WebChatBot.git` or `git@github.com:kentkim84/WebChatBot.git`
```
git clone https://github.com/kentkim84/WebChatBot.git
```
## Coding Standards
// Version 0.2 using C standards

### 1. Naming Conventions and Style
1.1. Use Pascal casing for class and structs
    
    type Rectangle struct {
        Name string
        Width, Height float64
    }

    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

1.2. Use camel casing for local variable names and function parameters
    
    func SomeMethod(someParameter const int)
    {
        someNumber int
    }

1.3. Use verb-object pairs for method names
a.	Use pascal casing for public methods
        
    public:
    func DoSomething()

b.	Use camel casing for other methods
        
    private:
    func doSomething()

1.4. Put constant variables after import

    import (
        "fmt"
    )

    const CONSTANT_NUMBER int = 100

### References
This project is from
* [Data Representation and Querying - Ian McLoughlin](https://data-representation.github.io/problems/project.html)

This coding standard is followed by
* [popeKim's c/c++ coding standards](https://docs.google.com/document/d/1cT8EPgMXe0eopeHvwuFmbHG4TJr5kUmcovkr5irQZmo/edit#heading=h.r2n9mhxbh2gg)