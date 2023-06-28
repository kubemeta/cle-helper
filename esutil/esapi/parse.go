package esapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
)

func ParseResError(res *esapi.Response) error {
	if res.StatusCode == http.StatusNotFound {
		e := &IndexResponse{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return errors.Wrap(err, "Error parsing the response body")
		}
		return parseIndexResponse(res.StatusCode, e)
	}

	var b []byte
	if res.Body != nil {
		b, _ = ioutil.ReadAll(res.Body)
		reader := bytes.NewReader(b)
		res.Body = ioutil.NopCloser(reader)
	}

	e := &ErrorResponse{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		eni := &ErrorResponseNoInfo{}
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&eni); err != nil {
			return errors.Wrap(err, "Error parsing the response body")
		}
		return errors.New(eni.Error)
	}
	return parseErrorInfo(res.StatusCode, e)
}

func parseErrorInfo(statusCode int, res *ErrorResponse) error {
	return errors.New(fmt.Sprintf("%v", res.Info.RootCause))
}

func parseIndexResponse(statusCode int, info *IndexResponse) error {
	return errors.New(info.Result)
}

func ParseResBody(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}
