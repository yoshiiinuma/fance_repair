
package greedy

import (
  "container/heap"
  "fmt"
)

const DEBUG = true

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(val interface{}) {
  *pq = append(*pq, val.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
  last := len(*pq) - 1
  val := (*pq)[last]
  *pq = (*pq)[:last]
  return val
}

func createPriorityQueue(L []int) *PriorityQueue {
  N := len(L)
  pq := make(PriorityQueue, 0)
  heap.Init(&pq)
  for i := 0; i < N; i++ {
    heap.Push(&pq, L[i])
  }
  return &pq
}

/**
 * Generates a PriorityQueue from the given array
 * Returns the following accessors to the queue
 *
 *  - push(int)
 *  - pop() int
 *  - print()
 *
 */
func getPriorityQueueAccessors(L []int) (func(int), func() int, func() int, func()) {
  pq := createPriorityQueue(L)
  return func(val int) {           /* push */
      heap.Push(pq, val)
    },
    func() int {                   /* pop */
      return heap.Pop(pq).(int)
    },
    func() int {                   /* len */
      return len(*pq)
    },
    func() {                       /* print */
      fmt.Printf("%v", *pq)
    }
}

func GreedyFanceRepair(L []int) int {
  if DEBUG {
    fmt.Println("=========================================")
    fmt.Printf("L: %v\n", L)
    fmt.Println("-----------------------------------------")
  }
  N := len(L)
  if N == 0 {
    return 0
  } else if N == 1 {
    return L[0]
  }
  total := 0
  pushToPQ, popFromPQ, lenPQ, printPQ := getPriorityQueueAccessors(L)
  if DEBUG {
    fmt.Printf("total %d: ", total)
    printPQ()
    fmt.Println("\n-----------------------------------------")
  }

  for lenPQ() > 1 {
    first := popFromPQ()
    second := popFromPQ()
    cost := first + second
    total += cost
    pushToPQ(cost)
    if DEBUG {
      fmt.Printf("total %d cost %d (= %d + %d): ", total, cost, first, second)
      printPQ()
      fmt.Println()
    }
  }
  return total
}

