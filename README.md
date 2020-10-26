# Fance Repair (POJ 3253)

You want to cut a board with a length of sum (L) into N planks so that each plank has a length of L1, L2, ..., Ln. When cutting the board to length x and y, it costs x + y. What is the minimum cost?

Constraint
* 1 ≤ N ≤ 20,000
* 1 ≤ Li ≤ 50,000

## Examples

```
Input:

  N = 4
  L = { 8, 5, 8 }

Output:

   34

   34 = 21 + 13 = (13 + 8) + (8 + 5)

```

The length of the original board is 21 (= 8 + 5 + 8). The first cut will cost 21, and should be used to cut the board into pieces measuring 13 and 8. The second cut will cost 13, and should be used to cut the 13 into 8 and 5. This would cost 34 (= 21 + 13). If the 21 was cut into 16 and 5 instead, the second cut would cost 16 for a total of 37 (which is more than 34).

# Approach

From observing the below binary trees that represent how the board is cut when the cutting cost is minimum, we can find that the smaller a plank is, the later it is cut.
In other words, the total minimum cost can be caluculated by merging the cutting cost of the smallest plank one by one in the bottom up order.

```
 L: [ 4, 5, 6, 8, 10 ]
 SUM(L): 33

         33
        /  \
       19  14
      / \  / \
    10  9  8  6
       / \
      5   4
```

|plank length|cost|total|
|------------|----|-----|
|  4 +  5    |   9|    9|
|  6 +  8    |  14|   23|
|  9 + 10    |  19|   42|
| 14 + 19    |  33|   75|


```
 L: [ 1, 2, 3, 4, 5, 6, 7, 8 ]
 SUM(L): 36

              36
             /  \
           21    15
          / \    / \
        12   9  8   7
        / \ / \
       6  6 5  4
      / \
     3   3
    / \
   2   1

```

|plank length|cost|total|
|------------|----|-----|
|  1 +  2    |   3|    3|
|  3 +  3    |   6|    9|
|  4 +  5    |   9|   18|
|  6 +  6    |  12|   30|
|  7 +  8    |  15|   45|
|  9 + 12    |  21|   66|
| 21 + 15    |  36|  102|


# Complexity Analysis

* Time Complexity: O(N*logN)

  Time complexity of min heap insertion (Priority Queue) is O(LogN). For each cut, a new cost is entered to the priority queue. There are total N - 1 cuts.

* Space Complexity: O(N)
