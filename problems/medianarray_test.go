package answers

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		exp  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, test := range tests {
		ans := findMedianSortedArrays(test.arr1, test.arr2)
		if ans != test.exp {
			t.Errorf("Error: expected answer %v, got %v", test.exp, ans)
		}
	}
}

func TestFindMedian(t *testing.T) {
	tests := []struct {
		arr []int
		med float64
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{1}, 1},
	}
	for _, test := range tests {
		ans := findMedian(test.arr)
		if ans != test.med {
			t.Errorf("Error: expected median %v, got %v", test.med, ans)
		}
	}
}

func TestJoinSortedArrays(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		exp  []int
	}{
		{[]int{1, 2, 2, 3}, []int{3, 4, 5}, []int{1, 2, 2, 3, 3, 4, 5}},
		{[]int{1}, []int{5, 6, 7}, []int{1, 5, 6, 7}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 4}, []int{1, 2}, []int{1, 1, 2, 3, 4}},
	}
	for _, test := range tests {
		ret := joinSortedArrays(test.arr1, test.arr2)
		if len(ret) != len(test.exp) {
			t.Fatalf("Error: expected length %v, got length %v", len(test.exp), len(ret))
		}
		for i := range ret {
			if ret[i] != test.exp[i] {
				t.Errorf("Error: expected array %v, got array %v", test.exp, ret)
			}
		}
	}
}
