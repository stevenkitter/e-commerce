package server

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"io"
)

func ParseJson(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}
	result, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("valid error")
	}
	return nil
}
