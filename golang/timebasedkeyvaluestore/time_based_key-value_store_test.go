package timebasedkeyvaluestore

import "testing"

func TestTimeMap(t *testing.T) {
	// Example one
	obj := Constructor()
	obj.Set("foo", "bar", 1)

	res := obj.Get("foo", 1)
	if res != "bar" {
		t.Errorf("Get('foo', 1) should returned %s instead 'bar'.", res)
	}

	res = obj.Get("foo", 3)
	if res != "bar" {
		t.Errorf("Get('foo', 3) should returned %s instead 'bar'.", res)
	}

	obj.Set("foo", "bar2", 4)

	res = obj.Get("foo", 4)
	if res != "bar2" {
		t.Errorf("Get('foo', 4) should returned %s instead 'bar2'.", res)
	}
	
	res = obj.Get("foo", 5)
	if res != "bar2" {
		t.Errorf("Get('foo', 5) should returned %s instead 'bar2'.", res)
	}
}