
package greedy

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestGreedyFanceRepair(t *testing.T) {
  var L []int

  L = []int{ 8, 5, 8 }
  Equal(t, 34, GreedyFanceRepair(L))
}
