# Fetch

simple http client with a basic retry / back off strategy

- 1 seconds
- 3 seconds
- 5 seconds
- 10 seconds

## Examples


### fetch basic example

```go

fetch := fetch.New()
var apiErr *fetch.APIError
resp, err := fetch.Get("https://icanhazdadjoke.com/", nil)
if err != nil {
    if errors.As(err, &apiErr) {
        // non 2xx,3xx response
        // StatusCode: 400
        // StatusText: Bad Request
        // Message: ""
        fmt.PrintLn(apiErr)
    }
    // standard errors
    os.Exit(1)
}

```

### fetch with headers

```go

headers := map[string]string{
    "Authorization": "Bearer boo"
}

fetch := fetch.New()
resp, err := fetch.Get("https://icanhazdadjoke.com/", headers)
if err != nil {
  if errors.As(err, &apiErr) {
        // non 2xx,3xx response
        // StatusCode: 400
        // StatusText: Bad Request
        // Message: ""
        fmt.PrintLn(apiErr)
    }
    // standard errors
    os.Exit(1)
}

```


### fetch with default headers

```go

headers := map[string]string{
    "Authorization": "Bearer boo"
}

options := fetch.Options{
    DefaultHeaders = headers,
}

fetch := fetch.New(options)
resp, err := fetch.Get("https://icanhazdadjoke.com/", nil)
if err != nil {
  if errors.As(err, &apiErr) {
        // non 2xx,3xx response
        // StatusCode: 400
        // StatusText: Bad Request
        // Message: ""
        fmt.PrintLn(apiErr)
    }
    // standard errors
    os.Exit(1)
}

```