package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type Stencil struct {
	host       string
	httpclient *resty.Client
}

func NewStencilClient(host string) *Stencil {
	return &Stencil{
		host:       host,
		httpclient: resty.New(),
	}
}

func (s Stencil) CreateNamespace(namespace, description string) (responseCode int, err error) {
	url := fmt.Sprintf("%s/v1beta1/namespaces", s.host)
	payload := fmt.Sprintf(`{"id": "%s", "format": "FORMAT_PROTOBUF", "compatibility": "COMPATIBILITY_BACKWARD", "description": "%s"}`, namespace, description)

	resp, err := s.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(url)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode() >= http.StatusBadRequest && resp.StatusCode() < http.StatusInternalServerError {
		var errResponse StencilErrorResponse

		err = json.Unmarshal(resp.Body(), &errResponse)
		if err != nil {
			return 0, err
		}

		return errResponse.Code, errors.New(errResponse.Message)
	}

	fmt.Printf("[Create stencil namespace] [%d] success: %s\n", resp.StatusCode(), string(resp.Body()))
	return resp.StatusCode(), nil
}

func (s Stencil) UploadSchema(namespace, schemaName, schemaDescriptor string) error {
	reader, err := os.Open(schemaDescriptor)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1beta1/namespaces/%s/schemas/%s", s.host, namespace, schemaName)
	resp, err := s.httpclient.R().
		SetBody(reader).
		Post(url)

	if err != nil {
		return err
	}

	fmt.Printf("[Create schema] [%d] success: %s\n", resp.StatusCode(), string(resp.Body()))
	return nil
}
