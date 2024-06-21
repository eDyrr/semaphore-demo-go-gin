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
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	c.HTML(
		http.StatusOk,
		"index.html",
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}
```

###### Displaying a single article

in the last section, while we displayed a list of articles, the links to the articles didnt work. In this section, we'll add handlers and templates to display an article when it is selected.

###### Setting up the route

we can set up a new route to handle requests for a single article in the same manner as in the previous route. However, each article has its own route, and with Gin we can pass parameters to routes, as follows:

```router.GET("article/view/:article_id", geArticle)```

this route will match all requests matching the above path and will store the value of the last part of the route in the route parameter named `article_id` which we can access in the route handler. For this route, we will define the handler in a function named `getArticle`.

the updated `main.go` file should contain the following code:

```
func main() {
	router := gin.Default()
	
	// handle Index
	router.GET("/", showIndexPage)

	// handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}
```

###### Creating the view templates

since we're aiming to display an article, we should create an interface for it, thus a template as follows:
```
<!-- article.html -->

<!-- embed the header.html template at this location -->

{{ template "header.html" . }}

<!-- display the title of the article -->

<h1>
    {{ .payload.Title }}
</h1>

<!-- display the content of the article -->
<p>
    {{ .payload.Content }}
</p>

<!-- embed the footer.html template at this location -->
{{ template "footer.html" . }}
```

###### Specifying the requirement for the Go microservice router

the test for the handler of this route will check fort the following conditions:

- the handler responds with an HTTP status code of 200,
- the returned HTML contains a title tag containing the title of the article that was fetched.

the code for the test will be placed in the `TestArticleUnauthenticated` function in the `handlers.article_test.go` file. We will place helper functions used by this function in the `common_test.go` file.

###### Creating the router handler

the handler for the article page, `getArticle` performs the following tasks:
1. extracts the ID of the article to display
to fetch and display the right article, we first need to extract its ID from the context. This can be extracted as follows:

```c.Param("article_id")```

where `c` is the Gin Context which is parameter to any route handler when using Gin.

2. fetches the article
this can be done using the `getArticleByID()` function defined in the `models.article.go` file:

```article, err := getArticleByID(articleID)```

after adding `getArticleByID`, the `models.article.go` file should look like this:

```
package main

import "errors"

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

func getArticleByID(id int) (*article, error) {
	for _, v := range articleList {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("Article not found")
}
```

3. renders the `article.html` template passing it the article

this can be done using the code below:
```
c.HTML(
	http.StatusOK,
	"article.html",
	gin.H{
		"title": article.Title,
		"payload": article,
	},
)
```

the updated `handlers.article.go` file should contain the following code:

```
package main

import(
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	// check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// call the HTML method of the context to render a template
			c.HTML(
				// set the HTTP status to 200 (OK)
				http.StatusOK,
				// use the 
				"article.html",
				// pass the data that the page uses
				gin.H{
					"title": article.Title,
					"payload": article,
				},
			)
		} else {
			// if the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
```
now we build and run the application and visit `http://localhost:8080/article/view/1` in a browser

###### Responding with JSON/XML
we'll do some refactoring of the app so that, depending on the request headers, our application can respond in HTML, JSON or XML format.

###### Creating a reusable function
so far, we've been using the HTML method of Gin's context to render directly from route handlers. This is fine when we always want to render HTML.
however, if we want to change the format of the response based on the request, we should refactor this part out into a single function that takes care of the rendering. By doing this, we can let the route handler focus on validation and data fetching.

a route handler has to do the same kind of validation, data fetching and data processing irrespective of the desired response format. Once this part is done, this data can be used to generate a response in the desired format. If we need an HTML response, we can pass this data to the HTML template and generate the page. If we need a JSON response, we can convert this data to JSOn and send it back. Likewise for XML.

we'll create a `render` function in `main.go` that will be used by all the route handlers. This function will take care of rendering in the right format based on the request's `Accept` header.

in Gin, the `Context` passed to a route handler contains a field named `Request`. This field contains the `Header` field which contains all the request headers. We can use the `Get` method on `Header` to extract the `Accept` header as follows:
```
// c is the Gin Context
c.Request.Header.Get("Accept")
```

- if this is set to `application/json`, the function will render JSON,
- if this is set to `application/xml`, the function will render XML, and
- if this is set to anything else or is empty, the function will render HTML.

the `render` function is as follows, add it in the `handler.article.go` file:

```
// render one of HTML, JSON or CSV based on the 'Accept' header of the request
// if the header doesnt specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
		case "application/json":
			// respond with JSON
			c.JSON(http.StatusOK, data["payload"])
		case "application/xml":
			// respond with XML
			c.XML(http.StatusOK, data["payload"])
		default:
			// respond with HTML
			c.HTML(http.StatusOK, templateName, data)
	}
}
```

###### Modifying the requirement for the route handlers with a unit test

since we're now expecting JSON and XML responses if the respective headers are set, we should add tests to the `handlers.article_test.go` file to test these conditions. We will add tests to:

1. test that the application returns a JSON list of articles when the `Accept` header is set to `application/json`

2. test the application returns an article in XML format when the `Accept` header is set to `application/xml`

these will be added as functions named `TestArticleListJSON` and `TestArticleListXML`.

###### Updating the route handlers

the route handlers dont really need to change much as the logic for rendering in any format is pretty much the same. All that needs to be done is use the `render` function instead of rendering using the `c.HTML` methods.

for example, the `showIndexPage` route handler in `handlers.article.go` will change from:

```
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}
```

to:

```
func showIndexPage(c *gin.Context) {
	article := getAllArticles()

	// call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Home Page",
		"payload": articles}, "index.html")
}
```

**retrieving the list of articles in JSON format**

to see our latest updates in action, build and run your application. Then execute the following command:

```
curl -X GET -H "Accept: application/json" http://localhost:8080/"
```

this should return a response as follows:

```
[{"id":1,"title":"Article 1","content":"Article 1 body"},{"id":2,"title":"Article 2","content":"Article 2 body"}]
```

as you can see, our request got a response in the JSOn format because we set the `Accept` header to `application/json`.

**retrieving an article in XML format**

lets now get our application to respond with the details of a particular article in the XML format. To do this, first, start your application as mentioned above. Now execute the following command:

```
curl -X GET -H "Accept: application/XML" http://localhost:8080/
```

this should return a response as follows:

```
<article><ID>1</ID><Title>Article 1</Title><Content>Article 1 body</Content></article>
```

as you can see, our request got a response in the XML format because we set the `Accept` header to `application/xml`.

###### Testing the application

since we've been using tests to create specifications for our route handlers and models, we should constantly be running them to ensure that the functions work as expected. Lets now run the tests that we have written and see the results. In your project directory, execute the following command:

```
go test -v
```
executing this commnad should look like this:

```

```