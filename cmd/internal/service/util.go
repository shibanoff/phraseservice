package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

func decodeRequest(r *http.Request, v interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(content, v)
	if err != nil {
		return errors.WithStack(err)
	}

	if validatable, ok := v.(validation.Validatable); ok {
		return validatable.Validate()
	}

	return nil
}

func sendJSONResponse(w http.ResponseWriter, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errors.WithStack(err)
	}

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(b)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func sendError(w http.ResponseWriter, err error) {
	const internalError = "internal service error"

	var userError = struct {
		Error string
	}{internalError}

	var vErr validation.Errors
	if errors.As(err, &vErr) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		userError.Error = vErr.Error()
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := sendJSONResponse(w, userError); err != nil {
		panic(err)
	}
}
