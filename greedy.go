
package greedy

import "fmt"

const DEBUG = true

func swap(L []int, i, j int) {
  temp := L[j]
  L[j] = L[i]
  L[i] = temp
}

func GreedyFanceRepair(L []int) int {
  if DEBUG {
    fmt.Println("=========================================")
    fmt.Printf("L: %v\n", L)
    fmt.Println("-----------------------------------------")
  }
  N := len(L)
  var cost int
  total := 0
  first := 0
  second := 1
  if L[first] > L[second] { swap(L, first, second) }
  for i := 1; i < N; i++ {
    for j := i; j < N; j++ {
      if L[j] < L[first] {
        second = first
        first = j
      } else  if L[j] < L[second] {
        second = j
      }
    }
    cost = L[first] + L[second]
    total += cost
    swap(L, i - 1, first)
    swap(L, i, second)
    if DEBUG {
      fmt.Printf("i %d first %d second %d cost %d: %v\n", i, L[i-1], L[i], cost, L)
    }
    L[i] = cost
    first = i
    second = i + 1
  }
  return total
}
