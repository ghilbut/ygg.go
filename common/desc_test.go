package common_test

import (
	. "github.com/ghilbut/ygg.go/common"
	"testing"
)

func Test_CtrlDesc_return_error_when_pass_empty_json_string(t *testing.T) {

	desc, err := NewCtrlDesc("")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_CtrlDesc_return_error_when_pass_invlid_json(t *testing.T) {

	desc, err := NewCtrlDesc("{ qwerty")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_CtrlDesc_return_error_when_pass_valid_json_without_values(t *testing.T) {

	desc, err := NewCtrlDesc("{ \"key\": \"value\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_CtrlDesc_return_error_when_pass_valid_json_without_id(t *testing.T) {

	desc, err := NewCtrlDesc("{ \"endpoint\": \"B\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_CtrlDesc_return_error_when_pass_valid_json_without_endpoint(t *testing.T) {

	desc, err := NewCtrlDesc("{ \"id\": \"A\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_CtrlDesc_create_desc_instance_with_valid_json(t *testing.T) {

	json := "{ \"id\": \"A\", \"endpoint\": \"B\"}"
	desc, err := NewCtrlDesc(json)

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

func Test_TargetDesc_return_error_when_pass_empty_json_string(t *testing.T) {

	desc, err := NewTargetDesc("")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_TargetDesc_return_error_when_pass_invlid_json(t *testing.T) {

	desc, err := NewTargetDesc("{ qwerty")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_TargetDesc_return_error_when_pass_valid_json_without_endpoint(t *testing.T) {

	desc, err := NewTargetDesc("{ \"key\": \"value\" }")

	if desc != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func Test_TargetDesc_create_desc_instance_with_valid_json(t *testing.T) {

	json := "{ \"endpoint\": \"A\"}"
	desc, err := NewTargetDesc(json)

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
