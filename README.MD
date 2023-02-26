# Intro
* CRUD service using GIN with request validator for middleware.
* Implement service interface
* Using handler functions as methods of Handler Struct
* Swagger documents included


# API Docs

placeholder


# Some middleware I use
[limit-by-key](https://github.com/yangxikun/gin-limit-by-key)

[cors](https://github.com/gin-contrib/cors)

[xss-mv](https://github.com/dvwright/xss-mw)

[gin-helmet](https://github.com/danielkov/gin-helmet)


# Try on
1. clone this repo
2. create .env file(check out the .env.example file)
3. `CompileDaemon -command="./portfolio"`


# Why this
I usually wouldn't make a simple project this complicated. However, I was trying to learn deeper about abstraction in Go in this project which I found quite appealing and natural.

When I think about interfaces I think about how they should behave. Everything becomes clearer once you designed your interfaces. The only thing left is to implement them. And I certainly enjoy coding in Go.

Before getting started with this project, I checked so many tutorials, blogs, and example projects to figure out the perfect structure for a Go web server project. Only to find out there's no such thing as perfect(should've known). Several things I take into consideration:

1. Is this something that will be reused later?

2. What's it like to be called in another package?

3. If someone wants to find a specific file can they easily guess the directory it's in?
   
4. Will I add more logic/functions in the future?

There is a lot that can be improved in this project. I'm just thrilled to experience new knowledge. After all, learning is always fun.