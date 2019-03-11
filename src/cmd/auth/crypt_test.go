package auth

import "testing"

const test_secret = "sec"

func TestEncrypt(t *testing.T) {
	cases := []string{
		"asad",
		"",
		"0",
		"dasdsaioyeua asdka2e waeuiyh2313 -",
	}
	for _, v := range cases {
		re := Encrypt([]byte(v), test_secret)
		rd := Decrypt(re, test_secret)
		if string(rd) != v {
			t.Error(
				"before enc ", v,
				"\nafter enc", re,
				"\ndecrypt encrypted ", rd,
			)
		}
	}
}
