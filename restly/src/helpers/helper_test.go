package helpers

import (
    "testing"
)

func TestGenUUID(t *testing.T) {
	_, err := GenUUID()
	if err != nil{
		t.FailNow()
	}
}

