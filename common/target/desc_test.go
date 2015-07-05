package target

import "testing"

func Test_return_error_when_pass_empty_json_string(t *testing.T) {

	desc, err := NewDesc("")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_return_error_when_pass_invlid_json(t *testing.T) {

	desc, err := NewDesc("{ qwerty")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_return_error_when_pass_valid_json_without_endpoint(t *testing.T) {

	desc, err := NewDesc("{ \"key\": \"value\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_create_desc_instance_with_valid_json(t *testing.T) {

	json := "{ \"endpoint\": \"A\"}"
	desc, err := NewDesc(json)

	if desc == nil {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	if desc.Json != json {
		t.Fail()
	}

	if desc.Endpoint != "A" {
		t.Fail()
	}
}
