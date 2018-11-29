package array_test

import (
	"testing"

	"microprogram/array"
)

func TestCommonSearch(t *testing.T) {
	var (
		source = [][]int{
			{3, 5, 7, 8},
			{5, 6, 9, 14},
			{11, 18, 20, 32},
		}
		target = 9
	)
	t.Log(array.CommonSearch(source, target))
}

func TestEfficiencySearch(t *testing.T) {
	var (
		source = [][]int{
			{3, 5, 7, 8},
			{5, 6, 9, 14},
			{11, 18, 20, 32},
		}
		target = 9
	)
	t.Log(array.EfficiencySearch(source, target))
}