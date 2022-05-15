package json

import (
	"encoding/json"

	"github.com/AntonyIS/url_shortener/shortener"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (*Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serlializer.Redirect.Decode")
	}
	return redirect, nil
}

func (*Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serlializer.Redirect.Decode")
	}
	return rawMsg, nil
}
