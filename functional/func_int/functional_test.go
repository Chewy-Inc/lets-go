package func_int_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var empty []int

func TestAll(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	passing := []int{12, 25, 36, 56, 90, 99}
	failing := []int{3, 5, 8, 12, 20, 32}

	assert.True(t, All(passing, condition))
	assert.False(t, All(failing, condition))
	assert.True(t, All(empty, condition))
}

func TestAny(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	passing := []int{12, 25, 36, 56, 90, 99}
	failing := []int{3, 5, 8, 9, 0, 2}

	assert.True(t, Any(passing, condition))
	assert.False(t, Any(failing, condition))
	assert.False(t, Any(empty, condition))
}

func TestFilter(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	initial := []int{12, 8, 2, 0, 22, 87}

	filtered := Filter(initial, condition)

	assert.Equal(t, 3, len(filtered))
	assert.EqualValues(t, filtered, []int{12, 22, 87})

	assert.Equal(t, len(Filter(empty, condition)), 0)
}

func TestFilterNot(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	initial := []int{12, 8, 2, 0, 22, 87}

	filtered := FilterNot(initial, condition)

	assert.Equal(t, 3, len(filtered))
	assert.EqualValues(t, filtered, []int{8, 2, 0})

	assert.Equal(t, len(FilterNot(empty, condition)), 0)
}

func TestFlatten(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9 ,10}
	slice3 := []int{11, 12, 13, 14, 15}

	matrix := [][]int{slice1, slice2, slice3}

	assert.EqualValues(t, []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}, Flatten(matrix))
	assert.Equal(t, len(empty), len(Flatten([][]int{empty})))
}

func TestFold(t *testing.T) {
	foldFn := func(x int, y int) int {
		return x + y
	}

	sliceToFold := []int{10, 20, 30, 40, 50}

	assert.Equal(t, Fold(sliceToFold, 50, foldFn), 200)
	assert.Equal(t, Fold(empty, 0, foldFn), 0)
}

func TestFoldR(t *testing.T) {
	foldFn := func(x int, y int) int {
		return x + y
	}

	sliceToFold := []int{10, 20, 30, 40, 50}

	assert.Equal(t, FoldR(sliceToFold, 50, foldFn), 200)
	assert.Equal(t, FoldR(empty, 0, foldFn), 0)
}

func TestForEach(t *testing.T) {
	sliceToPrint := []int{10, 20, 30, 40, 50}

	ForEach(sliceToPrint, func(x int){fmt.Println(x)})
}

func TestIndexOf(t *testing.T) {
	sliceToPrint := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	assert.Equal(t, IndexOf(sliceToPrint, 60), 5)
	assert.Equal(t, IndexOf(empty, 10), -1)
}

func TestMap(t *testing.T) {
	mapFn := func(x int) int {
		return x + 1
	}

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	assert.EqualValues(t, []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}, Map(slice, mapFn))
	assert.EqualValues(t, []int{}, Map(empty, mapFn))
}

func TestMax(t *testing.T) {
	slice := []int{18, 75, 79, 29, 51, 94, 71}

	assert.Equal(t, 94, Max(slice))
	assert.Equal(t, 0, Max(empty))
}

func TestMin(t *testing.T) {
	slice := []int{18, 75, 79, 29, 51, 94, 71}

	assert.Equal(t, 18, Min(slice))
	assert.Equal(t, 0, Min(empty))
}

func TestReduce(t *testing.T) {
	foldFn := func(x int, y int) int {
		return x + y
	}

	sliceToFold := []int{10, 20, 30, 40, 50}

	assert.Equal(t, 150, Reduce(sliceToFold, foldFn))
	assert.Equal(t, 0, Reduce(empty, foldFn))
}

func TestReduceR(t *testing.T) {
	foldFn := func(x int, y int) int {
		return x + y
	}

	sliceToFold := []int{10, 20, 30, 40, 50}

	assert.Equal(t, ReduceR(sliceToFold, foldFn), 150)
	assert.Equal(t, ReduceR(empty, foldFn), 0)
}