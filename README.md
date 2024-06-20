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

###### Dispalying the list of articles

We'll try to add the functionality to display the list of all articles on the index page.

after refactoring the code, here's the following `main.go` file:
```
package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set the router as the default one provided by Gin
	router = gin.Default()

	// process the templates at the start of so that they dont have to be loaded
	// from the disk again. This makes serving HTML pages very fast.

	router.LoadHTMLGlob("templates/*")

	// handle Index
	router.GET("/", showIndexPage)

	// start serving the application
	router.Run()
}
```

###### Designing the Article model
after getting to know the basic fiels that an article `type` needs, we create a `models.article.go` file to represent our `article`.

since this demo app isnt going to talk about DBs we're not going to be accessing DBs and thus we will be having our data in memory.

we'll be needing a function to return to us the data that's stored.

here's the code that satisfies our needs:
```
package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "article 1", Content: "article 1 body"},
	article{ID: 2, Title: "article 2", Content: "article 2 body"},
}

// return a list of all the articles
func getAllArticles() []article {
	return articleList
}
```

we will also write a test for the above function named `TestGetAllArticles()` and it'll be placed in the `models.article_test.go` file.

here's the code for `models.article_test.go`:
```
package main

import "testing"

// test the fucntion that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check taht the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}

	}
}
```

to run the test just run:
```go test```

###### Creating the view template

since the list of articles will be displayed on the index patge, we dont need to create a new template. However, we need to change the `index.html` template to replace the currect content with the list of articles.

to make this change, we'll assume that the list of articles will be passed to the template in a variable named `payload`. With this assumption, the following snippet should show the list of all articles:

```
{{ range .payload }}
<!-- create the link for the article based on its ID -->
<a href="/article/view/{{.ID}}">
    <!-- display the title of the article -->
    <h2>
        {{.Title}}
    </h2>
</a>
     <!-- display the content of the article -->
<p>
    {{.Content}}
</p>
{{end}}
```

the small code above will loop over all the items in the `payload` variable and will display the title and the content. It will also link to each article.
However, since we havent made (yet) any route handlers for displaying individual articles, these links wont work.

the code above will be placed in the `index.html` file.

###### Specifying the requirement for the route handler with a unit test

first we start with creating the test to define the expected behavior of the handler for the index route.
the test will check for the following conditions:

1. the handler responds with an HTTP status code of 200.
2. the returned HTML contains a title tag containing the text `Home Page`.

the code for the test will be placed in the `TestShowIndexPageUnauthenticated` function in the `handler.article_test.go` file. We will place helper functions used by this function in the `commom_test.go` file.

the content of `handlers.article_test.go` is as follows:

```
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// test that the page title is "Home Page"
		// you can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
```

the contet of the `common_test.go` is as follows:
```
package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tempArticleList []article

// this function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	// set Gin to Test mode
	gin.SetMode(gin.TestMode)

	// run the other tests
	os.Exit(m.Run())
}

// helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}

	return r
}

// helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// this function is used to store the main lists into the temporary one for testing
func saveLists() {
	tempArticleList = articleList
}

// this function is used to restore the main lists from the temporary one
func restoreList() {
	articleList = tempArticleList
}
```

for the implementation of this test, we wrote some helper functions.

the `TestMain` function sets Gin to use the test mode and calls the rest of the test functions. The `getRouter` function creates and returns a router in a manner similar to the main application. The `saveLists()` function saves the original article list in a temporary variable. This temporary variable is used by the `restoreLists()` function to restore the article list to its initial state after a unit test is executed.

the `testHTTPResponse` function executes the function passed in to see if it returns a boolean true value -- indicating a successfull test, or not.
This function helps us avoid duplicating the code needed to test the response of an HTTP request.

to check the HTTP code and the returned HTML, we'll do the following:

1. create a new router.
2. define a route to use the same handler that the main app uses (showIndexPage).
3. create a new request to access this route
4. create a function that processes the response to test the HTTP code and HTML.
5. call the `testHTTPResponse()` with this new function to complete the test.

###### Creating the router handler

we will create all route handlers for article related functionality in the `handler.article.go` file. The handler for the index page, `showIndexPage` performs the following tasks:

1. fetches the list of articles
this can be done using the `getAllArticles()` function defined previously:

```articles := getAllArticles()```

2. renders the index.html template passing it the article list
this can be done using the code below:

```
c.HTML(
	// set the HTTP status to 200 (OK)
	http.StatusOK,
	// use the index.html template
	"index.html",
	// pass the data that the page uses
	gin.H{
		"title": "Home Page",
		"payload": articles,
	},
)
```

the only difference from the version in the previous section is that we're passing the list of articles which will be accessed in the template by the variable named `payload`.

the `handlers.article.go` file should contain the following code:
```