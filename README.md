# Fetch

simple http client

## Examples


### fetch basic example

```go

fetch := fetch.New()
resp, err := fetch.Get("https://icanhazdadjoke.com/", nil)
if err != nil {
    os.Exit(1)
}
if resp.statusCode == http.StatusBadRequest {
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
    os.Exit(1)
}
if resp.statusCode == http.StatusBadRequest {
    os.Exit(1)
}

```


### fetch with default headers

```go

headers := map[string]string{
    "Authorization": "Bearer boo"
}

fetch := fetch.New(headers)
resp, err := fetch.Get("https://icanhazdadjoke.com/", nil)
if err != nil {
    os.Exit(1)
}
if resp.statusCode == http.StatusBadRequest {
    os.Exit(1)
}

```