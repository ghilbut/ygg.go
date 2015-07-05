package target

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

type Desc struct {
	Json     string
	Endpoint string `json:"endpoint"`
}

func NewDesc(text string) (*Desc, error) {

	desc := &Desc{}
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
