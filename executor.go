package shelf

import (
	"compress/gzip"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewExecutor(client *http.Client) Executor {
	return &executor{client: client}
}

type executor struct {
	client *http.Client
}

func (e executor) Exec(ctx context.Context, req Request) (Response, error) {
	request, err := req.BuildRequest()
	if err != nil {
		return Response{}, NewExecutorError(err, req)
	}
	request = request.WithContext(ctx)
	resp, err := e.client.Do(request)
	if err != nil {
		return Response{}, NewExecutorError(err, req)
	}
	body := resp.Body
	defer body.Close()
	if strings.EqualFold(resp.Header.Get("Content-Encoding"), "gzip") && resp.ContentLength != 0 {
		body, err = gzip.NewReader(body)
		if err != nil {
			return Response{}, NewExecutorError(err, req)
		}
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return Response{}, NewExecutorError(err, req)
	}
	return Response{req, resp, data}, nil
}
