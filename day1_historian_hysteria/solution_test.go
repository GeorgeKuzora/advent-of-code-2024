package day1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_lomutoPartition(t *testing.T) {
	testSlice := []int{5,3,1,7,9,2,4}
	expectedSlice := []int{3,1,2,4,9,5,7}

	pivotId := lomutoPartition(&testSlice, 0, len(testSlice)-1)

	if !reflect.DeepEqual(testSlice, expectedSlice) {
		t.Error("Slices are not equal", testSlice, expectedSlice)
	}
	if pivotId != 3 {
		t.Error("pivotId not right", pivotId, 3)
	}
}

func TestQuickSort(t *testing.T) {
	testSlice := []int{5,3,1,7,9,2,4}
	expectedSlice := []int{1,2,3,4,5,7,9}
	fmt.Println(testSlice)
	quickSort(&testSlice)
	fmt.Println(testSlice)

	if !reflect.DeepEqual(testSlice, expectedSlice) {
		t.Error("Slices are not equal", testSlice, expectedSlice)
	}
}

func TestFindDistance(t *testing.T) {
	testSlice1 := []int{5,3,1,7,9,2,4}
	testSlice2 := []int{5,3,1,6,9,2,4}
	
	differ := FindDistance(testSlice1, testSlice2)
	
	if differ != 1 {
		t.Error("Expected difference 1, got", differ)
	}
}