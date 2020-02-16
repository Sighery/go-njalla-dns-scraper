package records

import "testing"

func TestConstructor(t *testing.T) {
	record := New(map[string]interface{}{
		"test": 1,
		"key2": "Privileged",
	})
	if record == nil {
		t.Fail()
	}
}

func TestGetValue(t *testing.T) {
	record := New(map[string]interface{}{
		"test": 1,
		"key2": "Privileged",
	})
	if record.getValue("test") != 1 {
		t.Fail()
	}
}

func TestGetInfo(t *testing.T) {
	record := New(map[string]interface{}{
		"test": 1,
		"key2": "Privileged",
	})
	info := record.getInfo()
	if info["key2"] == nil {
		t.Fail()
	}
}
