package lib

import "testing"

func TestRestore(t *testing.T) {
	bin, err := Restore("/tmp/smc-devne", true, true)
	t.Log(err, bin)
}
