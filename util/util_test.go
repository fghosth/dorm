package util

import (
	"fmt"
	"testing"
)

func TestDstring_UnderToCal(t *testing.T) {
	ds := new(Dstring)
	str := "GUsername"
	res := ds.CalToUnder(str)
	fmt.Println(res)

}
