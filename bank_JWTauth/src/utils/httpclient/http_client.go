package httpclient

import (
	"fmt"

	"bank_JWTauth/src/utils/errors"

	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New()
)

func HTTPClientPost(url string, body map[string]interface{}) (*resty.Response, *errors.RestErr) {
	resp, err := client.R().SetBody(body).EnableTrace().Post(url)
	if err != nil {
		fmt.Println(resp.RawResponse)
		return nil, errors.NewInternalServerError("Error in response")
	}
	if resp.StatusCode() > 299 {
		fmt.Println(resp.StatusCode())
		return nil, errors.NewInternalServerError("unsuccessful response")
	}

	return resp, nil

}
