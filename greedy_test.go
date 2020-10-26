
package greedy

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestGreedyFanceRepair(t *testing.T) {
  var L []int

  L = []int{}
  Equal(t, 0, GreedyFanceRepair(L))

  L = []int{ 1 }
  Equal(t, 1, GreedyFanceRepair(L))

  L = []int{ 2, 1 }
  Equal(t, 3, GreedyFanceRepair(L))

  L = []int{ 8, 5, 8 }
  Equal(t, 34, GreedyFanceRepair(L))

  L = []int{ 4, 5, 6, 8, 10 }
  Equal(t, 75, GreedyFanceRepair(L))

  L = []int{ 1, 4, 5, 7, 9, 13 }
  Equal(t, 93, GreedyFanceRepair(L))

  L = []int{ 9, 13, 5, 7, 4, 1 }
  Equal(t, 93, GreedyFanceRepair(L))

  L = []int{ 1, 1, 1, 1, 1 }
  Equal(t, 12, GreedyFanceRepair(L))

  L = []int{ 1, 2, 3, 4, 5, 6, 7, 8 }
  Equal(t, 48, GreedyFanceRepair(L))
}
