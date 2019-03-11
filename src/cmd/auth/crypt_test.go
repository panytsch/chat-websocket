package auth

import (
	"log"
	"testing"
	"unicode/utf8"
)

func TestEncrypt(t *testing.T) {
	cases := []string{
		"asad",
		"",
		"0",
		"dasdsaioyeua asdka2e waeuiyh2313 -",
	}
	for _, v := range cases {
		re := Encrypt(v, DefaultSecret)
		rd := Decrypt(re, DefaultSecret)
		log.Println(re, "string valid : ", utf8.ValidString(string(re)))
		if rd != v {
			t.Error(
				"before enc ", v,
				"\nafter enc", re,
				"\ndecrypt encrypted ", rd,
			)
		}
	}
}
