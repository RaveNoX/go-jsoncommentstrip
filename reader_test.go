package jsoncommentstrip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func process(input, expected string) (err error) {
	jsonReader := NewReader(bytes.NewReader([]byte(input)))
	result, err := ioutil.ReadAll(jsonReader)

	if err != nil {
		return
	}

	resultStr := string(result)
	if expected != resultStr {
		err = fmt.Errorf("Result is not equals expected\nExpected:\n%#v\n\nResult:\n%#v\n\n", expected, resultStr)
	}

	return
}

func TestWithoutComments(t *testing.T) {
	input := "{\n\"one\": 1,\n\"two\": 2,\n\"string\": \"value\"\n}"

	err := process(input, input)
	if err != nil {
		t.Error(err)
	}

	input = "{\"one\": 1,\"two\": 2,\"string\": \"value\"}"
	err = process(input, input)
	if err != nil {
		t.Error(err)
	}
}

func TestSlComment(t *testing.T) {
	input := "{\n\"one\": 1, // test //\n\"two\": 2, //test //\r\n\"string\": \"value\"\n//test\n}"
	expected := "{\n\"one\": 1, \n\"two\": 2, \r\n\"string\": \"value\"\n\n}"

	err := process(input, expected)
	if err != nil {
		t.Error(err)
	}

	input = "{// woot\n\"one\": 1, // test //\n\"two\": 2, //test //\r\n\"string\": \"value\"\n//test\n}"

	err = process(input, expected)
	if err != nil {
		t.Error(err)
	}
}

func TestMlComment(t *testing.T) {
	expected := "{\"one\":1}"
	input := "{/* multi\nline\r\ncomment */\"one\":1}"

	err := process(input, expected)
	if err != nil {
		t.Error(err)
	}
}

func TestQuotationEscape(t *testing.T) {
	expected := "{\"one\": \"a value \\\" // /*woot\"\r\n}"
	input := "{/* multi\nline\r\ncomment */\"one\": \"a value \\\" // /*woot\"/* m\nl *///woot\r\n}"

	err := process(input, expected)
	if err != nil {
		t.Error(err)
	}
}
