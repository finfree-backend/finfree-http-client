package finfree_http_client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client interface {
	Get(URI string, responseBody interface{}, queryParams ...pair) (*http.Response, error)
	Post(URI string, requestBody, responseBody interface{}) (*http.Response, error)
	Put(URI string, requestBody, responseBody interface{}) (*http.Response, error)
	Patch(URI string, requestBody, responseBody interface{}) (*http.Response, error)
	Delete(URI string, responseBody interface{}, queryParams ...pair) (*http.Response, error)
	Custom(httpMethod method, URI string, requestBody, responseBody interface{}) (*http.Response, error)
}

type finfreeHttpClient struct {
	client  *http.Client
	baseURL string

	authorizeFn func(r *http.Request)
}

// Get method sends a [GET] request
// response-body can be specified in the request
// query-parameter(s) can be specified in the request
func (cl *finfreeHttpClient) Get(URI string, responseBody interface{}, queryParams ...pair) (*http.Response, error) {
	return cl.withNoRequestBody(http.MethodGet, URI, responseBody, queryParams...)
}

func (cl *finfreeHttpClient) Post(URI string, requestBody, responseBody interface{}) (*http.Response, error) {
	return cl.withRequestBody(http.MethodPost, URI, requestBody, responseBody)
}

func (cl *finfreeHttpClient) Put(URI string, requestBody, responseBody interface{}) (*http.Response, error) {
	return cl.withRequestBody(http.MethodPut, URI, requestBody, responseBody)
}

func (cl *finfreeHttpClient) Patch(URI string, requestBody, responseBody interface{}) (*http.Response, error) {
	return cl.withRequestBody(http.MethodPatch, URI, requestBody, responseBody)
}

func (cl *finfreeHttpClient) Delete(URI string, responseBody interface{}, queryParams ...pair) (*http.Response, error) {
	return cl.withNoRequestBody(http.MethodDelete, URI, responseBody, queryParams...)
}

func (cl *finfreeHttpClient) Custom(httpMethod method, URI string, requestBody, responseBody interface{}) (*http.Response, error) {
	return cl.withRequestBody(string(httpMethod), URI, requestBody, responseBody)
}

func (cl *finfreeHttpClient) withNoRequestBody(method, URI string, responseBody interface{}, queryParams ...pair) (*http.Response, error) {
	// Create the request
	request, err := cl.newRequest(method, URI, nil)
	if err != nil {
		return nil, err
	}

	// If query parameters are passed, set them into request URL
	if queryParams != nil {
		q := request.URL.Query()
		for _, param := range queryParams {
			q.Add(param.getKey(), param.getVal())
		}
		request.URL.RawQuery = q.Encode()
	}

	return cl.do(request, responseBody)
}

func (cl *finfreeHttpClient) withRequestBody(method, URI string, requestBody, responseBody interface{}) (*http.Response, error) {
	// Jsonize (marshal) the request body (if there is)
	var data []byte
	if requestBody != nil {
		var err error
		data, err = json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
	}

	// Create the request
	request, err := cl.newRequest(method, URI, data)
	if err != nil {
		return nil, err
	}

	return cl.do(request, responseBody)
}

// newRequest creates brand-new request
// URI parameter added to base url and this operation creates the full URL of the request
// authorizeFn function is being called in this function and makes the request authorized (if authorizeFn is set)
func (cl *finfreeHttpClient) newRequest(method, URI string, requestBody []byte) (*http.Request, error) {
	// Create the url of the request
	url := cl.baseURL + URI

	// Create the request with given parameters
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Authorize the request
	// If no authorization function created for the client, nothing will change in the request
	cl.authorizeFn(request)

	return request, nil
}

// do send the request
// If request needs to consume a response body, method needs to be called with v variable
// 'v' needs to be a pointer (there is no return statement for unmarshalled data)
func (cl *finfreeHttpClient) do(r *http.Request, v interface{}) (*http.Response, error) {
	// Send the request
	response, err := cl.client.Do(r)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()

	// If length of v is 0, assumed as request consume no response body
	if v == nil {
		return response, nil
	}

	// Read all the data in request body
	// Unmarshall it into 'v'
	if err = json.NewDecoder(response.Body).Decode(v); err != nil {
		return response, err
	}

	return response, nil
}

// Constructor function with no authorization method
func New(baseURL string) Client {
	return new(baseURL, func(r *http.Request) {})
}

// Constructor function with bearer token
func NewWithBearerAuthorization(baseURL string, bearerToken string) Client {
	token := "Bearer " + bearerToken
	return new(baseURL, func(r *http.Request) {
		r.Header.Set("Authorization", token)
	})
}

// Constructor function with key-value pair header (customizable)
func NewWithHeaderAuthorization(baseURL string, pairs ...pair) Client {
	return new(baseURL, func(r *http.Request) {
		for _, p := range pairs {
			r.Header.Set(p.getKey(), p.getVal())
		}
	})
}

func new(baseURL string, authorizeFn func(r *http.Request)) Client {
	return &finfreeHttpClient{
		client:      &http.Client{},
		baseURL:     baseURL,
		authorizeFn: authorizeFn,
	}
}
