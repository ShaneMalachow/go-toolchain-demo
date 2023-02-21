/*
Copyright © 2023 Shane Malachow <shane.malachow@sciencelogic.com>
*/
package cmd

import (
	"strconv"
	"testing"
)

func Test_isPerfect(t *testing.T) {
	perfectNumbers := []int{6, 28, 496}
	for x := 1; x < 1000; x++ {
		perf := false
		for _, n := range perfectNumbers {
			if x == n {
				perf = true
				break
			}
		}
		if perf != isPerfect(x) {
			t.Errorf("%d want: %s, got: %s", x, strconv.FormatBool(perf), strconv.FormatBool(!perf))
		}
	}
}
