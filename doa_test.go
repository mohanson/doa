package doa

import (
	"strconv"
	"testing"
)

func TestDoa(t *testing.T) {
	i := Try2(strconv.Atoi("1234")).(int)
	if i != 1234 {
		t.FailNow()
	}
}
