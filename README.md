# Factorial Calculator

## Overview

This repository contains a RESTful service that calculates factorials of given integers using goroutines. 
The service utilizes the `httprouter` package for creating the server and provides an endpoint for factorial calculation.
The service will start and listen on port 8989.

## Usage

* To run this project use `go run cmd/main.go`
* To run all tests use `go test ./...`
* To run linter use `golangci-lint run ./...`



## API Endpoint

### `POST /calculate`

* Request JSON Structure: {"a": int, "b": int}
* Response JSON Structure: {"factorial A": uint64, "factorial B": uint64}
The endpoint calculates the factorials of a and b using goroutines and returns a JSON response with the calculated values.

### Request

To test the API, you can use the curl command-line tool to make POST requests, like this:

`curl localhost:8989/calculate -d {\"a\":6,\"b\":12}`

This command sends a JSON request with the parameter a & b set to 6 & 12 to the server's /calculate endpoint, and it will return a pair of factorials.
