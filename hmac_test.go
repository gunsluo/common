package common

import (
	"bytes"
	"testing"
)

func Test_HmacSha1ToHex(t *testing.T) {
	var message = []byte("test_message")
	var secret = "test_secret"
	var expect = "2dec96a5811c3bab60cdf636a07250e9d0f7e926"

	sha := HmacSha1ToHex(message, secret)

	if sha != expect {
		t.Error("HmacSha1ToHex not expect vaule. expect:", expect, "act:", sha)
	}
}

func Test_HmacSha1(t *testing.T) {
	var message = []byte("test_message")
	var secret = "test_secret"
	var expect = []byte{45, 236, 150, 165, 129, 28, 59, 171, 96, 205, 246, 54, 160, 114, 80, 233, 208, 247, 233, 38}

	sha := HmacSha1(message, secret)
	if !bytes.Equal(sha, expect) {
		t.Error("HmacSha1 not expect vaule. expect:", expect, "act:", sha)
	}
}

func Test_VerifyHmacSha1(t *testing.T) {
	var message = []byte("test_message")
	var secret = "test_secret"
	var signature = "2dec96a5811c3bab60cdf636a07250e9d0f7e926"

	flag := VerifyHmacSha1(message, secret, signature)
	if !flag {
		t.Error("VerifyHmacSha1 not expect vaule. expect:", !flag, "act:", flag)
	}
}

func Test_VerifyHmacSha12(t *testing.T) {
	var message = []byte("test_message")
	var secret = "test_secret"
	var signature = ""

	flag := VerifyHmacSha1(message, secret, signature)
	if flag {
		t.Error("VerifyHmacSha1 not expect vaule. expect:", !flag, "act:", flag)
	}
}
