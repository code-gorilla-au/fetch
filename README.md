# Fetch

simple http client with a basic retry / back off strategy

## Features

- Default retry / back off strategy 1 seconds, 3 seconds, 5 seconds, 10 seconds
- Provide optional custom retry strategy
- Provide optional HTTP client
- Set default headers for every request
- Add additional headers for individual requests
- Response codes > 399 are treated as errors (fetch.APIError)

<br>
<br>

## Examples

- demo [dad jokes](cmd/dad_jokes/dad_jokes.go)

<br>

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

<br>

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

<br>

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

<br>
<br>

## Using functional options

<br>
<br>

```go
fetch := fetch.New(WithOpts(
    /** ... more options */
    WithRetryStrategy(&[]time.Duration{1, 2}),
))

```

### Available options

| option | description |
| ------- | ----------- |
| WithDefaultRetryStrategy | Use default retry strategy            |
| WithHeaders              | Set default headers for every request |
| WithRetryStrategy        | Provide custom retry strategy         | 
| WithHTTPClient           | Provide custom http client            | 


<br>
<br>

## TODO

feature requests being worked on will be added here