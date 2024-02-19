# Creating Rest APIs

- [Introduction to HTTP in Go](#Introduction_to_HTTP_in_Go)
- [Creating a simple HTTP server](#Creating_a_simple_HTTP_server)
- [Handling HTTP requests and responses](#Handling_HTTP_requests_and_responses)
- [Routing in Go (e.g., using `gorilla/mux`)](#Routing_in_Go_)
- [Creating RESTful APIs](#Creating_RESTful_APIs)
- [Middleware for request processing](#Middleware_for_request_processing)
- [JSON handling in Go](#JSON_handling_in_Go)

<div id='Introduction_to_HTTP_in_Go'></div>

## Introduction to HTTP in Go

HTTP (Hypertext Transfer Protocol) is the foundation of data communication on the World Wide Web. It is a protocol that allows clients to communicate with servers, enabling the exchange of text, images, videos, and other media. In Go, the standard library provides robust support for building HTTP servers and clients.

Here's a basic introduction to HTTP in Go:

<div id='Creating_a_simple_HTTP_server'></div>

## Creating a simple HTTP server

### HTTP Server

To create an HTTP server in Go, you typically use the `net/http` package. Below is a simple example of an HTTP server that listens on port 8080 and responds to all requests with "Hello, World!".

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

In this code:

- We import the `net/http` package to utilize HTTP functionality.
- We define a handler function (`handler`) that takes an `http.ResponseWriter` and an `http.Request` as parameters. This function is responsible for generating the response to incoming requests.
- We register our handler function with `http.HandleFunc()`. This function tells the HTTP server to call our handler function whenever a request is made to the root ("/") endpoint.
- We start the HTTP server using `http.ListenAndServe()`, specifying the port to listen on (`:8080` in this case).

<div id='Handling_HTTP_requests_and_responses'></div>

## Handling HTTP requests and responses

### HTTP Client

In Go, you can also create HTTP clients to make requests to HTTP servers. The `net/http` package provides convenient methods for performing HTTP requests. Below is an example of making a GET request to a URL and printing the response body:

```go
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    url := "https://example.com"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response Body:", string(body))
}
```

In this code:

- We import the necessary packages (`fmt`, `net/http`, `io/ioutil`) for making HTTP requests and processing responses.
- We use `http.Get()` to send a GET request to the specified URL (`https://example.com` in this case).
- We check for errors and handle them appropriately.
- We read the response body using `ioutil.ReadAll()` and print it to the console.

These examples provide a basic overview of creating HTTP servers and clients in Go using the standard library. You can extend and customize these examples to build more complex web applications and services.

In Go's `net/http` package, the `http.Request` struct represents an HTTP request received by a server or to be sent by a client. It contains various fields and methods to access information about the request. Some of the important members of the `http.Request` struct include:

1. **Method**: The HTTP method of the request (e.g., GET, POST, PUT, DELETE).

   ```go
   Method string
   ```

2. **URL**: The URL of the request.

   ```go
   URL *url.URL
   ```

3. **Header**: The HTTP headers included in the request.

   ```go
   Header http.Header
   ```

4. **Body**: The request body, which is an `io.ReadCloser`.

   ```go
   Body io.ReadCloser
   ```

5. **ContentLength**: The length of the request body in bytes, as specified in the `Content-Length` header.

   ```go
   ContentLength int64
   ```

6. **Host**: The host name from the `Host` header.

   ```go
   Host string
   ```

7. **RemoteAddr**: The network address of the client that sent the request.

   ```go
   RemoteAddr string
   ```

8. **TLS**: Information about the TLS connection, if present.

   ```go
   TLS *tls.ConnectionState
   ```

9. **Form**: The parsed form data from the request body.

   ```go
   Form url.Values
   ```

10. **PostForm**: The parsed form data from the request body or URL query parameters.

    ```go
    PostForm url.Values
    ```

11. **MultipartForm**: The parsed multipart form data.

    ```go
    MultipartForm *multipart.Form
    ```

12. **Cookies**: The cookies included in the request.

    ```go
    Cookies []*Cookie
    ```

These are some of the important members of the `http.Request` struct in Go. They provide access to various aspects of an incoming HTTP request, allowing you to inspect and process the request in your server-side code.

<div id='Routing_in_Go_'></div>

## Routing in Go (e.g., using `gorilla/mux`)

Routing in Go, particularly in web applications, refers to the process of mapping incoming HTTP requests to the appropriate handler functions based on the request's URL path and HTTP method (GET, POST, PUT, DELETE, etc.). One popular library for routing in Go is `gorilla/mux`, which provides a powerful and flexible router implementation.

### Using `gorilla/mux`

To use `gorilla/mux`, you first need to install it:

```bash
go get -u github.com/gorilla/mux
```

Then, you can import it in your Go code:

```go
import "github.com/gorilla/mux"
```

### Basic Routing

Here's how you can create a basic HTTP server with routing using `gorilla/mux`:

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)

    // Start the server
    http.Handle("/", r)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Products Page")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Articles Page")
}
```

In this example:

- We import `github.com/gorilla/mux`.
- We create a new router using `mux.NewRouter()`.
- We define routes using `HandleFunc()`, specifying the path pattern and the handler function for each route.
- We start the HTTP server using `http.ListenAndServe()` with the router as the handler.

### URL Parameters

`gorilla/mux` allows you to define routes with URL parameters. For example:

```go
r.HandleFunc("/products/{category}/{id}", ProductHandler)
```

You can then access these parameters in your handler function using `mux.Vars(r)`.

### Subrouters

You can also create subrouters to organize your routes:

```go
articles := r.PathPrefix("/articles").Subrouter()
articles.HandleFunc("/", ArticlesHandler)
articles.HandleFunc("/{id}", ArticleHandler)
```

### Middleware

`gorilla/mux` supports middleware, which are functions that are executed before or after a request is handled. You can use middleware for tasks such as logging, authentication, or error handling.

```go
r.Use(LoggerMiddleware)
```

### Handle url parameters in a handler method:

In Go, when using `gorilla/mux` for routing, you can handle URL parameters in a handler method using the `mux.Vars()` function. This function extracts the URL parameters from the request and returns them as a map[string]string.

Here's an example of how you can handle URL parameters in a handler method:

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/books/{id}", BookHandler)
    http.Handle("/", r)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["id"]
    fmt.Fprintf(w, "Book ID: %s", bookID)
}
```

In this example:

- We define a route `/books/{id}` where `{id}` is a placeholder for the book ID.
- We register the `BookHandler` function to handle requests to this route.
- Inside the `BookHandler` function, we use `mux.Vars(r)` to get a map of URL parameters extracted from the request.
- We then access the specific parameter using the key `"id"` to retrieve the book ID.
- Finally, we respond with the book ID in the HTTP response.

When a request is made to `/books/123`, for example, `mux.Vars(r)` will return a map with a single entry `{"id": "123"}`, and the handler will print "Book ID: 123" in the response.

This approach allows you to dynamically handle different URL paths and extract parameters from them, enabling more flexible routing and request handling in your Go web applications.

### Handle query parameters:

In Go, when handling HTTP requests with query parameters, you can access them through the `r.URL.Query()` method, which parses the query string from the URL and returns a map[string][]string. Each query parameter can have multiple values, hence the map value is a slice of strings.

Here's how you can handle query parameters in a handler method:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/search", SearchHandler)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    queryValues := r.URL.Query()

    // Get a single value for a query parameter
    query := queryValues.Get("q")
    fmt.Println("Query:", query)

    // Get all values for a query parameter
    categories := queryValues["category"]
    fmt.Println("Categories:", categories)

    // Loop through all query parameters
    for key, values := range queryValues {
        fmt.Printf("%s: %v\n", key, values)
    }

    // You can now use the query parameters as needed in your application logic
    // Respond to the request accordingly
    fmt.Fprintf(w, "Search Results for query: %s", query)
}
```

In this example:

- We define a route `/search` where clients can send search queries.
- Inside the `SearchHandler` function, we use `r.URL.Query()` to parse the query parameters from the request URL.
- We use `Get()` method to get the value of a specific query parameter (e.g., "q").
- We use array indexing to get all values for a query parameter if it occurs multiple times (e.g., "category").
- We iterate over all query parameters using a loop to handle each parameter accordingly.
- Finally, we respond to the client with the search results or perform any other application-specific logic based on the query parameters.

When a client sends a request to `/search?q=example&category=books&category=movies`, for example, the `SearchHandler` function will print the query parameter values and respond with "Search Results for query: example".

<div id='Creating_RESTful_APIs'></div>

## Creating RESTful APIs

Creating RESTful APIs in Go typically involves defining routes to handle HTTP requests for various resources, such as users, products, or articles. You can use the `net/http` package for basic routing, but libraries like `gorilla/mux` or `gin` offer more features and convenience for building RESTful APIs.

Here's a basic example using `gorilla/mux`:

1. First, install `gorilla/mux`:

```bash
go get -u github.com/gorilla/mux
```

2. Then, create your Go application:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Product struct represents a product
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Data store for products (in-memory for simplicity)
var products = []Product{
	{ID: "1", Name: "Product 1", Price: 29.99},
	{ID: "2", Name: "Product 2", Price: 39.99},
}

// GetProductsHandler returns all products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// GetProductHandler returns a single product by ID
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", GetProductHandler).Methods("GET")

	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
```

In this example:

- We define a `Product` struct to represent product data.
- We define two handler functions: `GetProductsHandler` to return all products, and `GetProductHandler` to return a single product by ID.
- We use `mux.Vars(r)` to extract URL parameters.
- We define routes using `HandleFunc()` and specify the HTTP methods for each route.
- We start the HTTP server using `http.ListenAndServe()` with the router as the handler.

Now, you can test your API using tools like cURL or Postman. For example:

- To get all products: `GET http://localhost:8080/products`
- To get a specific product: `GET http://localhost:8080/products/{id}` (replace `{id}` with the ID of a product)

This example provides a basic foundation for building RESTful APIs in Go. You can extend it by adding more routes, implementing CRUD operations, integrating with a database, adding authentication, and so on, depending on your requirements.

<div id='Middleware_for_request_processing'></div>

## Middleware for request processing

Middleware in Go is a powerful concept often used in web development to preprocess or post-process HTTP requests and responses. It allows you to execute code before or after a request reaches a handler function, enabling tasks such as logging, authentication, authorization, error handling, and more.

In Go, middleware can be implemented using standard functions or third-party packages like `gorilla/mux` or `negroni`. Below, I'll provide examples of implementing middleware using both approaches.

### Standard Middleware

Using standard middleware with the `net/http` package involves defining wrapper functions that take an `http.Handler` as input and return a new `http.Handler`.

```go
package main

import (
	"fmt"
	"net/http"
)

// Middleware function to log requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}

// Handler function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a new router
	r := http.NewServeMux()

	// Register handler with middleware
	r.Handle("/", LoggerMiddleware(http.HandlerFunc(HelloHandler)))

	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
```

### Middleware with `gorilla/mux`

Using `gorilla/mux`, you can define middleware and apply it to specific routes or the entire router.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware function to log requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}

// Handler function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register handler with middleware
	r.HandleFunc("/", HelloHandler).Methods("GET").Name("hello").Handler(LoggerMiddleware)

	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
```

In both examples, the `LoggerMiddleware` function intercepts the incoming request, logs information about it, and then calls the `ServeHTTP` method of the next handler in the chain to continue processing the request.

Middleware allows you to modularize your application logic and apply cross-cutting concerns uniformly across your routes, making your code more maintainable and flexible.

<div id='JSON_handling_in_Go'></div>

## JSON handling in Go

In Go's RESTful services, handling JSON data is a common task, as JSON is a popular format for exchanging data between clients and servers. Go provides built-in support for encoding and decoding JSON data using the `encoding/json` package. Here's how you can handle JSON in Go's RESTful services:

### Encoding JSON Responses

To encode Go data structures into JSON and send them as responses in HTTP handlers, you can use `json.Marshal()` to convert the data to JSON format and then write it to the response writer.

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    // Simulated data
    products := []Product{
        {ID: 1, Name: "Product 1", Price: 29.99},
        {ID: 2, Name: "Product 2", Price: 39.99},
    }

    // Encode products to JSON
    jsonBytes, err := json.Marshal(products)
    if err != nil {
        http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
        return
    }

    // Set Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Write JSON response
    w.Write(jsonBytes)
}

func main() {
    http.HandleFunc("/products", GetProductsHandler)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}
```

### Decoding JSON Requests

To handle JSON data received in HTTP requests, you can use `json.Unmarshal()` to decode the JSON data into Go data structures.

```go
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    // Decode JSON request body into Product struct
    var product Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
        return
    }

    // Do something with the product...
    fmt.Printf("Received product: %+v\n", product)

    // Respond with success message
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Product created successfully")
}
```

In this example, `json.NewDecoder()` is used to create a new JSON decoder that reads from the request body. The `Decode()` method then decodes the JSON data into the provided struct pointer.

### Conclusion

Handling JSON in Go's RESTful services involves encoding JSON responses using `json.Marshal()` and decoding JSON requests using `json.Unmarshal()`. With these tools, you can easily exchange data in JSON format between your Go server and client applications.
