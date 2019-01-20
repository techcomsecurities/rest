package rest

import (
	"net/http"
)

const (
	defaultRetryTimes = 3
	defaultTimeout = 30 // seconds
)

// Request represents an abstracted request sent to server
// MUST use NewRequest for creating new Request object
type Request struct {
	maxRetryTimes int
	timeout int
	header http.Header
	client *http.Client
}

func NewRequest() *Request {
	var r = Request{
		maxRetryTimes: defaultRetryTimes,
		timeout: defaultTimeout,
		client: &http.Client{},
	}

	return &r
}

// AddHeader adds the key, value pair to the header. 
// It appends to any existing values associated with key.
func (r *Request) AddHeader(key, value string) *Request {
	r.header.Add(key, value)
	return r
}

// SetHeader sets the header entries associated with key to the single element value. 
// It replaces any existing values associated with key.
func (r *Request) SetHeader(key, value string) *Request {
	r.header.Set(key, value)
	return r
}

// Retry sets maximum number of retry times, default is 3 times
// n <= 0 will not retry
func (r *Request) Retry(n int) *Request {
	if n < 0 {
		n = 0
	}
	r.maxRetryTimes = n
	return r 
}

// TODO: implement later
func (r *Request) Body(data interface) *Request {
	return r
}

// TODO: implement later
func (r *Request) Timeout(n int) *Request {
	return r
}

// TODO: implement later
func (r *Request) Get(url string) (resp *http.Response, err error) {
	return &http.Response{}, nil
}

// TODO: implement later
func (r *Request) Post(url string) (resp *http.Response, err error) {
	return &http.Response{}, nil
}

func (r *Request) Put(url string) (resp *http.Response, err error) {
	// TODO: implement later
	return &http.Response{}, nil
}

// TODO: implement later
func (r *Request) Delete(url string) (resp *http.Response, err error) {
	return &http.Response{}, nil
}

func (r *Request) request(method string, endpoint string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrapf(err, "NewRequest failed: %s %s\n", method, endpoint)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(REQUEST_TIMEOUT) * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "HTTP Do failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "ReadAll failed: body=%+v\n", resp.Body)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d\n%s", resp.StatusCode, string(body))
	}

	return body, nil
}