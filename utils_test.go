package common

import (
	"bufio"
	"os"
	"testing"
)

func test_newjson(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	/*
		_, err := f.Write([]byte(content))
		if err != nil {
			return err
		}
	*/

	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func test_rmjson(filename string) {
}

func test_before(t *testing.T) {
	var content = "  jerrylou  "
	err := test_newjson(".test.json", content)
	if err != nil {
		t.Error("Prepare the data is Interrupted.")
	}
}

func test_after(t *testing.T) {
	err := os.Remove(".test.json")
	if err != nil {
		t.Error("Prepare the data is Interrupted.")
	}
}

func Test_ToString(t *testing.T) {

	test_before(t)

	str, err := ToString(".test.json")
	if err != nil {
		t.Error("ToString exec failed.", err)
	}

	var expect = "  jerrylou  "
	if str != expect {
		t.Error("ToString not expect vaule. expect:", expect, "act:", str)
	}

	test_after(t)
}

func Test_ToTrimString(t *testing.T) {

	test_before(t)

	str, err := ToTrimString(".test.json")
	if err != nil {
		t.Error("ToString exec failed.", err)
	}

	var expect = "jerrylou"
	if str != expect {
		t.Error("ToString not expect vaule. expect:", expect, "act:", str)
	}

	test_after(t)
}
