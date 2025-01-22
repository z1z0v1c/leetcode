package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	// Example one
	obj := Constructor(2)

	obj.Put(1, 1); // cache is {1=1}
	obj.Put(2, 2); // cache is {1=1, 2=2}

	param_1 := obj.Get(1);    // return 1
	if param_1 != 1 {
		t.Errorf("Get(1) returned %d instead of 1.", param_1)
	}	

	obj.Put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	
	param_1 = obj.Get(2);    // returns -1 (not found)
	if param_1 != -1 {
		t.Errorf("Get(2) returned %d instead of -1.", param_1)
	}	
	
	obj.Put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}

	param_1 = obj.Get(1);    // return -1 (not found)
	if param_1 != -1 {
		t.Errorf("Get(1) returned %d instead of -1.", param_1)
	}	

	param_1 = obj.Get(3);    // return 3
	if param_1 != 3 {
		t.Errorf("Get(3) returned %d instead of 3.", param_1)
	}	

	param_1 = obj.Get(4);    // return 4
	if param_1 != 4 {
		t.Errorf("Get(4) returned %d instead of 4.", param_1)
	}	
}