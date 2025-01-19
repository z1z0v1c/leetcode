package timebasedkeyvaluestore

import "testing"

func TestTimeMap(t *testing.T) {
	// Example one
	obj := Constructor()
	obj.Set("foo", "bar", 1)

	res := obj.Get("foo", 1)
	if res != "bar" {
		t.Errorf("Get(foo, 1) returned %s instead of bar.", res)
	}

	res = obj.Get("foo", 3)
	if res != "bar" {
		t.Errorf("Get(foo, 3) returned %s instead of bar.", res)
	}

	obj.Set("foo", "bar2", 4)

	res = obj.Get("foo", 4)
	if res != "bar2" {
		t.Errorf("Get(foo, 4) returned %s instead of bar2.", res)
	}
	
	res = obj.Get("foo", 5)
	if res != "bar2" {
		t.Errorf("Get(foo, 5) returned %s instead of bar2.", res)
	}

	// Example two
	obj = Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)

	res = obj.Get("love", 5)
	if res != "" {
		t.Errorf("Get(love, 5) returned %s instead of ''.", res)
	}

	res = obj.Get("love", 10)
	if res != "high" {
		t.Errorf("Get(love, 10) returned %s instead of high.", res)
	}

	res = obj.Get("love", 15)
	if res != "high" {
		t.Errorf("Get(love, 15) returned %s instead of high.", res)
	}

	res = obj.Get("love", 20)
	if res != "low" {
		t.Errorf("Get(love, 20) returned %s instead of low.", res)
	}

	res = obj.Get("love", 25)
	if res != "low" {
		t.Errorf("Get(love, 25) returned %s instead of low.", res)
	}

	// Example three
	obj = Constructor()
	obj.Set("foo", "bar", 1)
	obj.Set("foo", "toilet", 5)
	obj.Set("foo", "bucket", 10)
	obj.Set("foo", "bar2", 20)

	res = obj.Get("foo", 15)
	if res != "bucket" {
		t.Errorf("Get(foo, 15) returned %s instead of bucket.", res)
	}

	// Example four
	obj = Constructor()
	obj.Set("foo", "bar", 1)

	res = obj.Get("foo", 1)
	if res != "bar" {
		t.Errorf("Get(foo, 15) returned %s instead of bar.", res)
	}

	res = obj.Get("foo", 3)
	if res != "bar" {
		t.Errorf("Get(foo, 15) returned %s instead of bar.", res)
	}

	obj.Set("foo", "bar2", 4)

	res = obj.Get("foo", 4)
	if res != "bar2" {
		t.Errorf("Get(foo, 15) returned %s instead of bar2.", res)
	}

	res = obj.Get("foo", 5)
	if res != "bar2" {
		t.Errorf("Get(foo, 15) returned %s instead of bar2.", res)
	}

	obj.Set("foo", "zigzag", 7)
	obj.Set("foo", "conundrum", 8)
	obj.Set("foo", "hyperbole", 9)
	obj.Set("foo", "silhouette", 10)
	obj.Set("foo", "blasphemy", 11)

	res = obj.Get("foo", 9)
	if res != "hyperbole" {
		t.Errorf("Get(foo, 15) returned %s instead of hyperbole.", res)
	}
}