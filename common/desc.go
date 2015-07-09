package common

import (
	"encoding/json"
	"log"
)

type descError struct {
	msg string
}

func (e *descError) Error() string {
	return e.msg
}

type CtrlDesc struct {
	Json     string
	CtrlId   string `json:"id"`
	Endpoint string `json:"endpoint"`
}

func NewCtrlDesc(text string) (*CtrlDesc, error) {

	desc := &CtrlDesc{}
	if err := json.Unmarshal([]byte(text), &desc); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(desc.CtrlId) < 1 {
		msg := "there is no id value in json string."
		log.Println(msg)
		return nil, &descError{msg}
	}

	if len(desc.Endpoint) < 1 {
		msg := "there is no endpoint value in json string."
		log.Println(msg)
		return nil, &descError{msg}
	}

	desc.Json = text
	return desc, nil
}

type TargetDesc struct {
	Json     string
	Endpoint string `json:"endpoint"`
}

func NewTargetDesc(text string) (*TargetDesc, error) {

	desc := &TargetDesc{}
	if err := json.Unmarshal([]byte(text), &desc); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(desc.Endpoint) < 1 {
		msg := "there is no endpoint value in json string."
		log.Println(msg)
		return nil, &descError{msg}
	}

	desc.Json = text
	return desc, nil
}
