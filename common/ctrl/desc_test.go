package ctrl

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

func Test_return_error_when_pass_valid_json_without_values(t *testing.T) {

	desc, err := NewDesc("{ \"key\": \"value\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_return_error_when_pass_valid_json_without_id(t *testing.T) {

	desc, err := NewDesc("{ \"endpoint\": \"B\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_return_error_when_pass_valid_json_without_endpoint(t *testing.T) {

	desc, err := NewDesc("{ \"id\": \"A\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_create_desc_instance_with_valid_json(t *testing.T) {

	json := "{ \"id\": \"A\", \"endpoint\": \"B\"}"
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

	if desc.CtrlId != "A" {
		t.Fail()
	}

	if desc.Endpoint != "B" {
		t.Fail()
	}
}
