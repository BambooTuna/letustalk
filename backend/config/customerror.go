package config

import (
	"errors"
	"fmt"
	"strings"
)

type CustomError string

func ValidateError(fieldName string, errType string) CustomError {
	errorMessage := "を正しく入力してください"
	switch errType {
	case "email":
		errorMessage = "がメールアドレスの形式になっていません"
	case "min":
		errorMessage = "は１以上で入力してください"
	case "max":
		errorMessage = "が大きすぎます"
	case "gte":
		errorMessage = "が短すぎます"
	case "lte":
		errorMessage = "が長すぎます"
	case "required":
		errorMessage = "を入力してください"
	}
	return CustomError(fmt.Sprintf("%s%s", fieldName, errorMessage))
}

func Error(message CustomError) error {
	return errors.New(string(message))
}

func Errors(messages []CustomError) error {
	r := make([]string, len(messages))
	for i, e := range messages {
		r[i] = string(e)
	}
	return errors.New(strings.Join(r, ","))
}
