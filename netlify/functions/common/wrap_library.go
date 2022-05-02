package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo/v4"
)

func formatAPIResponse(statusCode int, headers http.Header, responseData string) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("some log 6")
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}
	fmt.Println("some log 7")

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	fmt.Println("some log 8")
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/plain"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              "Hello, World!",
		IsBase64Encoded:   false,
	}, nil
}

// Route wraps echo server into Lambda Handler
func Route(e *echo.Echo) func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		fmt.Println(request.Path)
		body := strings.NewReader(request.Body)
		req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}

		fmt.Println("some log 0")
		q := req.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println("some log 1")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		fmt.Println("some log 2")
		res := rec.Result()
		responseBody, err := ioutil.ReadAll(res.Body)
		fmt.Println("some log 3")
		if err != nil {
			fmt.Println("some log 4")
			return formatAPIResponse(http.StatusInternalServerError, res.Header, err.Error())
		}

		fmt.Println("some log 5")
		return formatAPIResponse(res.StatusCode, res.Header, string(responseBody))
	}
}
