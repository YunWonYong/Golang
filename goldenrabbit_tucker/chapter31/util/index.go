package util

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func GenUUID() string {
	return uuid.New().String()
}

func BindBuff[T interface{}](target io.ReadCloser) (*T, []byte, error) {
	instance := new(T)
	buffer, err := GetBuff(target)
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(buffer, instance); err != nil {
		return nil, buffer, err
	}

	return instance, buffer, nil
}

func GetBuff(target io.ReadCloser) ([]byte, error) {
	buffer, err := ioutil.ReadAll(target)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return buffer, nil
}

func Marshal(obj interface{}) (buff []byte, err error) {
	buff, err = json.Marshal(obj)

	if err != nil {
		err = errors.WithStack(err)
	}

	return
}

func Unmarshal[T interface{}](buff []byte) (*T, error) {
	instance := new(T)
	if err := json.Unmarshal(buff, instance); err != nil {
		return nil, errors.WithStack(err)
	}
	return instance, nil
}

func InterfaceToString(obj interface{}) (string, error) {
	str, ok := obj.(string)
	if ok {
		return str, nil
	}

	buff, err := InterfaceToByteArray(obj)

	if err != nil {
		return "", err
	}

	return string(buff), nil
}

func InterfaceToByteArray(obj interface{}) ([]byte, error) {
	buff, ok := obj.([]byte)
	if ok {
		return buff, nil
	}

	return []byte{}, errors.New("interface{} to []byte fail.")
}

func InterfaceToNumber[T int | int8 | int16 | int32 | int64](obj interface{}) (T, error) {
	number, ok := obj.(T)

	if ok {
		return number, nil
	}

	return 0, errors.Errorf("interface{} to number fail. obj: %#v", obj)
}

func RequiredValueValidation(values ...string) bool {
	for _, value := range values {
		if value == "" || len(value) == 0 {
			return false
		}
	}

	return true
}
