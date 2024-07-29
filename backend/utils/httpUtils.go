package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type ValidatedModel interface {
	Validate() (string, bool)
}

func ReadBody[T ValidatedModel](w http.ResponseWriter, r *http.Request) (T, error) {
	var zeroValue T

	bodyInBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return zeroValue, fmt.Errorf("problem reading body of request: %w", err)
	}

	var body T
	err = json.Unmarshal(bodyInBytes, &body)
	if err != nil {
		return zeroValue, fmt.Errorf("problem parsing body in json: %w", err)
	}
	badField, ok := isValidStruct(body)
	if !ok {
		return zeroValue, fmt.Errorf("invalid data: some required fields are missing or have zero values: %s", badField)
	}
	var badValidation string
	badValidation, ok = body.Validate()
	if !ok {
		return zeroValue, fmt.Errorf("data validation error: %s", badValidation)
	}
	fmt.Printf("Parsed body: %+v\n", body)
	return body, nil
}

func isValidStruct(s interface{}) (string, bool) {
	v := reflect.ValueOf(s)
	for i := 1; i < v.NumField(); i++ {
		if isZeroValue(v.Field(i)) {
			return v.Type().Field(i).Tag.Get("json"), false
		}
	}
	return "", true
}

func isZeroValue(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func HashString(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
