package env

import (
	"errors"
	"os"
	"strconv"
)

func GetEnv(key string) (string, error) {
	s := os.Getenv(key)
	if s == "" {
		return s, errors.New("getenv: environment variable empty")
	}
	return s, nil
}

func GetEnvInt(key string) (int, error) {
	s, err := GetEnv(key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(s)
	return v, nil
}

func GetEnvBool(key string) (bool, error) {
	s, err := GetEnv(key)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if nil != err {
		return false, err
	}
	return v, nil
}
