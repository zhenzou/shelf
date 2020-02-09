package shelf

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	NamePlaceholder = "${name}"
	PagePlaceholder = "${page}"
)

type Request struct {
	Method string
	URL    string
	Args   Args
}

func (req Request) BuildRequest() (*http.Request, error) {
	url := req.URL
	url = strings.ReplaceAll(url, NamePlaceholder, req.Args.Name)
	url = strings.ReplaceAll(url, PagePlaceholder, strconv.FormatUint(req.Args.Page, 10))
	return http.NewRequest(req.Method, url, nil)
}

type Response struct {
	Request     Request
	RawResponse *http.Response
	Data        []byte
}
