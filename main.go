package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpMethod string

type RequestOptions struct {
	URL     string
	Method  HttpMethod
	Headers map[string]string
	Body    string
}

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func parseHeaders(header string) map[string]string {
	headers := make(map[string]string)
	if header != "" {
		headerPairs := bytes.Split([]byte(header), []byte(","))
		for _, pair := range headerPairs {
			header := bytes.SplitN(pair, []byte(":"), 2)
			if len(header) == 2 {
				headers[string(bytes.TrimSpace(header[0]))] = string(bytes.TrimSpace(header[1]))
			}
		}
	}
	return headers
}

func parseFlag() (*RequestOptions, error) {
	var (
		url    = flag.String("url", "", "URL for the HTTP request")
		method = flag.String("method", "GET", "HTTP method (GET, POST, PUT, DELETE)")
		header = flag.String("header", "", "Headers for the request (comma-separated key:value pairs)")
		body   = flag.String("body", "", "Request body")
	)
	flag.Parse()

	if *url == "" {
		return nil, fmt.Errorf("URL is required")
	}

	headers := parseHeaders(*header)

	return &RequestOptions{
		URL:     *url,
		Method:  HttpMethod(*method),
		Headers: headers,
		Body:    *body,
	}, nil
}

func makeRequest(options *RequestOptions) {
	client := &http.Client{}
	req, err := http.NewRequest(string(options.Method), options.URL, bytes.NewBuffer([]byte(options.Body)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body:\n%s\n", body)
}

func main() {
	options, err := parseFlag()
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	makeRequest(options)
}
