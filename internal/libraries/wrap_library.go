package libraries

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
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"
	fmt.Println("add log", responseData)
	// return &events.APIGatewayProxyResponse{
	// 	Body:            responseData,
	// 	Headers:         responseHeaders,
	// 	StatusCode:      statusCode,
	// 	IsBase64Encoded: false,
	// }, nil
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
		body := strings.NewReader(request.Body)
		req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}

		q := req.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		res := rec.Result()
		responseBody, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return formatAPIResponse(http.StatusInternalServerError, res.Header, err.Error())
		}

		return formatAPIResponse(res.StatusCode, res.Header, string(responseBody))
	}
}
