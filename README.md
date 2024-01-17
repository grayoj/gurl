# gURL

gURL is a command-line interface (CLI) tool written in Go for making HTTP requests. It provides a convenient way to send HTTP requests and view the responses from the command line.

> I plan to make a Homebrew formula, hopefully.

## Features

- Supports common HTTP methods: GET, POST, PUT, DELETE
- Customizable HTTP headers
- Ability to include a request body
- Easy-to-use command-line interface with flags

## Installation

Make sure you have Go installed on your machine. Then, you can install gURL using the following command:

```bash
go install github.com/grayoj/gurl
```

## Using gURL

> gURL features some flags to help you make requests better.

- url: The URL for the HTTP request (required).
- method: HTTP method (default: GET).
- header: Headers for the request (comma-separated key:value pairs).
- body: Request body.

Here's a sample request: `gurl -url "https://api.gurl.com/data" -method POST -header "Content-Type: application/json" -body '{"key": "value"}'`

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it according to the terms of the license.
