package tg_a_c_i_su

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Count    int    `json:"count"`
	Messages []Data `json:"messages"`
}
type Data struct {
	Message string `json:"message"`
	PostID  int    `json:"id"`
}

type ACIClient struct{}

var (
	limit = "limit"
	page  = "page"
	uri   = "http://localhost:9504/json/"
)

func CreateACIClient() *ACIClient {
	return &ACIClient{}
}

func (r *ACIClient) Run(chatName string) (*ResponseData, error) {
	uri := fmt.Sprint(uri + chatName)
	resp, err := http.Get(uri)
	if err != nil {
		err := fmt.Errorf("%s", err.Error())
		return nil, err
	}
	if resp.StatusCode > http.StatusOK {
		err := fmt.Errorf("%s", "http code more 200... delay later")
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	data := ResponseData{}
	err = decoder.Decode(&data)
	if err != nil {
		err := fmt.Errorf("%s", "http code more 200... delay later")
		return nil, err
	}

	return &data, nil
}
