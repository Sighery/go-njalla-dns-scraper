package updater

import "testing"

func TestUdpdater(t *testing.T) {
	updater := New()
	if updater == nil {
		t.Fail()
	}
}
