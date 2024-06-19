# semaphore-demo-go-gin

###### What is a microservice ?
it is an architectural style that structures an application as a collection of services that are:

- independently deployable
- loosely coupled

services are typically organized around business capabilities. Each service is owned by a single, small team.

---

this application that I'll be building is a simple article manager. It will be able to show articles in HTML, JSON and XML as needed.

Ill be learning 

- Routing
- Custom rendering
- Middleware

###### Routing
it is some sort of technology that allows users to access any web page or API endpoint using a URL.

in our application, we will:
- serve the index page at route `/` (HTTP GET request)
- group article related routes under the `/article` route. Serve the article page at `/article/view/:article_id` (HTTP GET request)

###### Rendering
a web app is able to render a response in many formats such as HTML, text, JSON, XML ...etc.

API endpoints & microservices usually respond with data (JSON), yet it is still possible to respond in other formats.

###### Middleware
It is a piece of code that can be executed at any stage while handling an HTTP request. It is typically used to encapsulate common functionality that you want to apply to multiple routes. We can use middleware before and/or after an HTTP request is handled. Some common uses of middleware include authorization, validation, etc.

###### Creating reusable templates
Since our app has some sort of a UI, we'll definitely be using a number of components multiple times across many pages like the sidebar, menu, sidebar, and footer, and Go allows us to create reusable template snippets that can be imported in other templates.

We'll need to create a `templates` folder which contains all the templates we'll be needing for our application.

here is the template for the menu as `templates/menu.html`:

```
<!--menu.html>
<nav class="navbar navbar-default">
    <div class="container">
        <div class="navbar-header">
            <a class="navbar-brand" href="/">
                Home
            </a>
        </div>
    </div>
</nav>
```

at first the menu will only have the link to the home page. We will add to this as we add more functionality to the application. The template for the header will be placed in the `templates/header.html` file as follows:

```
<!--header.html-->

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
    <script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
</head>
<body class="container">
    <!-- embed the menu.html template at this location -->
    {{ template "menu.html" . }}
```
and the footer:
```
<!-- footer.html -->

    </body>
</html>
```
the things that pop up is `<title> {{ .title }} </title>` which is used to dynamically set the title of the page using the `.title` variable that must be set inside the application.
Another thing that pops up is the following line of code: `{{ template "menu.html" . }}` which is used to import the menu template from the `menu.html` file. This is how Go lets you import one template to another.

The template for the index page makes use of the header and the footer and displays a simple *Hello Gin* message:
```
<!-- index.html -->

<!-- embed the header.html template at this locations -->
{{ template "header.html" . }}

<h1>
    Hello Gin!
</h1>

<!-- embed the footer.html template at this location -->

{{ template "footer.html" . }}
```

###### Completing and validating the setup
once the templates are created, its time to create the entry file for your application. We'll create the `main.go` file for this with the simplest possible web application that will use the index template. We can do this using Gin in 4 steps:

1. Create the router
the default way to create a router in Gin as follows:

`router := gin.Default()`

this creates a router that can used to define the build of the application.

2. Load the templates
once you have created the router, you can load all the templates like this:
`router.LoadHTMLGlob("templates/*")`
this loads all the template files located in the `templates` folder. Once loaded, these dont have to be read again on every request making Gin web applications very fast.

3. Define the router handler
at the heart of Gin is how you divide the application into various routes and define handlers for each route. We will create a route for the index page and an inline router handler.

```
router.GET("/", func(c *gin.Context) {
    // call the HTML method of the context to render a template
    c.HTML(
        // set the HTTP status to 200 (OK)
        http.StatusOK,
        // use the index.html template
        "index.html"
        // pass the data that the pages uses (in this case, 'title')
        gin.H{
            "title": "Home page",
        },
    )
})
```
the `router.GET` method is used to define a router handler for a GET request. It takes in as parameters the route (/) and one or more router handlers which are just functions.

the router handler has a pointer to the context (`gin.Context`) as its parameter.
This context contains all the information about the request that the handler might need to process it. For example, it includes information about the headers, cookies, etc.

the context also has methods to render a response in HTML, text, JSON and XML formats. In this case, we use the `context.html` method to render an HTML template (`index.html`). The call to this method includes additional data in which the value of `title` is set to Home Page. This is a value that the HTML template can make use of. In this case, we use this value in the `<title>` tag in the header's template.

4. Start the application
to start the application, you can use the Run method of the router:
```router.Run()```
this starts the application on `localhost` and serves on the `8080` port by default.

the complete `main.go` file looks like this:
```
// main.go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    // set the router as the default one provided by Gin
    router = gin.Default()

    // process the templates at the start so that they dont have to be loaded 
    // from the disk again. This makes serving HTML pages very fast.
    router.LoadHTMLGlob("templates/*")


    // define the route for the index page and display the index.html template
    // to start with, we'll use an inline route handler. Later on, we'll create
    // standalone functions that will be used as router handlers.

    router.GET("/", func(c *gin.Context) {
        // call the HTML method of the context to render a template
        c.HTML(
            // set the http status to 200 (OK)
            http.StatusOK,
            // use the index.html template
            "index.html",
            // pass the data that the page uses (in this case, 'title')
            gin.H{
                "title": "Home Page",
            },
        )
    })

    // start serving the application
    router.Run()
}
```

to execute the application from the CLI, go to the app dir and run:
```go build -o app```
this will build your app and create an exe app which can be ran using:
```./app```
