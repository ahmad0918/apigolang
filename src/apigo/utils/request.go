package utils

import (
	"github.com/go-resty/resty/v2"
	logger "github.com/sirupsen/logrus"
	"time"
)

var (
	client = resty.New()
)

var ContentT = ContentType{
	JSON:     "application/json",
	FormData: "multipart/form-data; boundary=<calculated when request is sent>",
}

type ContentType struct {
	JSON     string
	FormData string
}

// HTTPRequest create a simple request object with resty
func HTTPRequest(headers, params map[string]string, body, result interface{}, retryCount int) *resty.Request {
	if headers == nil {
		headers = map[string]string{}
	}

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = ContentT.JSON
	}

	client.SetRetryCount(retryCount)
	return client.RemoveProxy().
		SetDebug(false).
		SetTimeout(10 * time.Second).
		R().SetHeaders(headers).
		SetPathParams(params).
		SetBody(body).
		SetResult(&result)
}

// PostRequest get response for PostRequest
func PostRequest(required, path string, headers, params map[string]string, body []byte, result interface{}, retryCount int) (*resty.Response, error) {
	response, err := HTTPRequest(headers, params, body, &result, retryCount).Post(path)
	Log3rdParty(required, response.Request.Method, response.Request.URL, string(body), string(response.Body()))
	return response, err
}

// GetRequest get response for GetRequest
func GetRequest(required, path string, headers, params map[string]string, result interface{}, retryCount int) (*resty.Response, error) {
	response, err := HTTPRequest(headers, params, nil, &result, retryCount).Get(path)
	Log3rdParty(required, response.Request.Method, response.Request.URL, "-", string(response.Body()))
	return response, err
}

func Log3rdParty(reqID string, method, url, req, resp interface{}) {
	// Assign detail with parameter above and generate console
	// Request
	logger.WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": req,
		"reqID":   reqID,
	}).Info("Resty Request")
	Loge.Info("Resty Request => method =", method,
		"||url =", url,
		"||request =", req,
		"||reqID =", reqID)

	// Response
	logger.WithFields(logger.Fields{
		"response": resp,
		"reqID":    reqID,
	}).Info("Resty Response")
	Loge.Info("Resty Response ==> response =", resp,
		"||reqID =", reqID)
}
