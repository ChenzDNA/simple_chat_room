package utils

import (
	"encoding/json"
	"fmt"
)

const SPLIT_ME = "####_split_me_####"

type resData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type chunkedData struct {
	UserId string `json:"userId"`
	Msg    string `json:"msg"`
}

func (d resData) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func OK(msg string, data any) *resData {
	return &resData{Code: 200, Msg: msg, Data: data}
}

func OKChunked(user, msg string) string {
	bytes, _ := json.Marshal(&chunkedData{UserId: user, Msg: msg})
	return fmt.Sprintf("%x\r\n%s%s\r\n", len(bytes)+len(SPLIT_ME), SPLIT_ME, bytes)
}

func EndChunked() string {
	return "0\r\n\r\n"
}

func Err(code int, msg string) *resData {
	return &resData{Code: code, Msg: msg}
}
