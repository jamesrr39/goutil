package httpextra

import (
	"encoding/json"
	"io"

	"github.com/jamesrr39/goutil/errorsx"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

func DecodeJSONDataResponse(reader io.Reader, dest interface{}) errorsx.Error {
	d := new(DataResponse)

	err := json.NewDecoder(reader).Decode(&d)
	if err != nil {
		return errorsx.Wrap(err)
	}

	dest = d.Data
	return nil
}
