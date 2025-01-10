package kthlargestelementinastream

import "testing"

func TestAdd(t *testing.T) {
	// Example one
	k := 3
	nums := []int{4, 5, 8, 2}

	obj := Constructor(k, nums);
	
	param_1 := obj.Add(3)
	if param_1 != 4 {
		t.Error("Add(3) should return 4.")
	}

	param_1 = obj.Add(5)
	if param_1 != 5 {
		t.Error("Add(5) should return 5.")
	}

	param_1 = obj.Add(10)
	if param_1 != 5 {
		t.Error("Add(10) should return 5.")
	}

	param_1 = obj.Add(9)
	if param_1 != 8 {
		t.Error("Add(9) should return 8.")
	}
	
	param_1 = obj.Add(4)
	if param_1 != 8 {
		t.Error("Add(4) should return 8.")
	}
	
	// Example two
	k = 4
	nums = []int{7, 7, 7, 7, 8, 3}

	obj = Constructor(k, nums);
	
	param_1 = obj.Add(2)
	if param_1 != 7 {
		t.Error("Add(2) should return 7.")
	}

	param_1 = obj.Add(10)
	if param_1 != 7 {
		t.Error("Add(10) should return 7.")
	}

	param_1 = obj.Add(9)
	if param_1 != 7 {
		t.Error("Add(9) should return 7.")
	}

	param_1 = obj.Add(9)
	if param_1 != 8 {
		t.Error("Add(9) should return 8.")
	}

	// Example three
	k = 1
	nums = []int{}

	obj = Constructor(k, nums);
	
	param_1 = obj.Add(-3)
	if param_1 != -3 {
		t.Error("Add(-3) should return -3.")
	}

	param_1 = obj.Add(-2)
	if param_1 != -2 {
		t.Error("Add(-2) should return -2.")
	}

	param_1 = obj.Add(-4)
	if param_1 != -2 {
		t.Error("Add(-2) should return -2.")
	}

	param_1 = obj.Add(0)
	if param_1 != 0 {
		t.Error("Add(0) should return 0.")
	}

	param_1 = obj.Add(4)
	if param_1 != 4 {
		t.Error("Add(4) should return 4.")
	}
	
	// Example four
	k = 2
	nums = []int{0}

	obj = Constructor(k, nums);
	
	param_1 = obj.Add(-1)
	if param_1 != -1 {
		t.Error("Add(-1) should return -1.")
	}

	param_1 = obj.Add(1)
	if param_1 != 0 {
		t.Error("Add(1) should return 0.")
	}

	param_1 = obj.Add(-2)
	if param_1 != 0 {
		t.Error("Add(-2) should return 0.")
	}

	param_1 = obj.Add(-4)
	if param_1 != 0 {
		t.Error("Add(-4) should return 0.")
	}

	param_1 = obj.Add(3)
	if param_1 != 1 {
		t.Error("Add(3) should return 1.")
	}
}