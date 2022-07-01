## How to use/test

```bash
npm i
npm run build
go run main.go
```

And access http://localhost:8080/index.html

## How it works

Using the basic template engine in Go, it adds variables in the html page to be used by the React frontend.

## How can this be improved

* Better HTTP server. Better templating library, that supports JS files out of the box so that the variables don't have to be in the HTML file
* Using a struct of all the required values to support as many variables as needed
* Better code structure