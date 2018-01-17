package hello

import (
	"testing"
)

func TestAdd(t *testing.T) {

	sum := Add(1, 2)
	if sum == 3 {

		t.Log("ok")

	} else {

		t.Fatal("wrong")

	}

}
