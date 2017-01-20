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

func test_rmjson(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}

func test_before(t *testing.T, filename string) {
	var content = "  jerrylou  "
	err := test_newjson(filename, content)
	if err != nil {
		t.Error("Prepare the data is Interrupted.")
	}
}

func test_after(t *testing.T, filename string) {
	err := test_rmjson(filename)
	if err != nil {
		t.Error("Prepare the data is Interrupted.")
	}
}

func Test_ToString(t *testing.T) {

	test_before(t, ".test.json")

	str, err := ToString(".test.json")
	if err != nil {
		t.Error("ToString exec failed.", err)
	}

	var expect = "  jerrylou  "
	if str != expect {
		t.Error("ToString not expect vaule. expect:", expect, "act:", str)
	}

	test_after(t, ".test.json")
}

func Test_ToString2(t *testing.T) {

	test_before(t, ".test.json")

	_, err := ToString(".test2.json")
	if err == nil {
		t.Error("ToString exec failed.")
	}

	test_after(t, ".test.json")
}

func Test_ToTrimString(t *testing.T) {

	test_before(t, ".test.json")

	str, err := ToTrimString(".test.json")
	if err != nil {
		t.Error("ToTrimString exec failed.", err)
	}

	var expect = "jerrylou"
	if str != expect {
		t.Error("ToTrimString not expect vaule. expect:", expect, "act:", str)
	}

	test_after(t, ".test.json")
}

func Test_ToTrimString2(t *testing.T) {

	test_before(t, ".test.json")

	_, err := ToTrimString(".test2.json")
	if err == nil {
		t.Error("ToTrimString exec failed.")
	}

	test_after(t, ".test.json")
}
