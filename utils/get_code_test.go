package utils

import "testing"

func TestGet(t *testing.T) {
	code := GenValidateCode(6)
	t.Log(code)
}
