package idgen

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	id := generate()
	fmt.Println(id)
	if id == 0 {
		t.Error("idgen.generate() returned 0")
	}
}