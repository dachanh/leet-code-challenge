# leet-code-challenge

A curated set of LeetCode problems grouped by **DSA pattern** for FAANG-style
interview practice. Each problem lives in its own folder with a Go solution and
table-driven tests.

## Problems by Pattern

### Two Pointers / Opposite-End

| # | Problem | Difficulty |
|---|---------|------------|
| 167 | [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Easy |
| 11  | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | Medium |
| 42  | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | Hard |
| 15  | [3Sum](https://leetcode.com/problems/3sum/) | Medium |
| 16  | [3Sum Closest](https://leetcode.com/problems/3sum-closest/) | Medium |

### Sliding Window

| # | Problem | Difficulty |
|---|---------|------------|
| 209 | [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | Medium |
| 904 | [Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/) | Medium |
| 424 | [Longest Repeating Char Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) | Medium |
| 567 | [Permutation in String](https://leetcode.com/problems/permutation-in-string/) | Medium |
| 713 | [Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/) | Medium |
| 76  | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | Hard |

### Prefix Sum

| # | Problem | Difficulty |
|---|---------|------------|
| 560 | [Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/) | Medium |

### Hashmap

| # | Problem | Difficulty |
|---|---------|------------|
| 1   | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy |
| 49  | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium |
| 128 | [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) | Medium |

### Stack / Monotonic Stack

| # | Problem | Difficulty |
|---|---------|------------|
| 20  | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | Easy |
| 739 | [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) | Medium |

### Binary Search

| # | Problem | Difficulty |
|---|---------|------------|
| 704 | [Binary Search](https://leetcode.com/problems/binary-search/) | Easy |
| 33  | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Medium |

### Linked List (incl. fast-slow pointers)

| # | Problem | Difficulty |
|---|---------|------------|
| 21  | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | Easy |
| 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | Easy |
| 206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | Easy |

### Heap / Priority Queue

| # | Problem | Difficulty |
|---|---------|------------|
| 703 | [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | Easy |
| 215 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium |

### Sorting

| # | Problem | Difficulty |
|---|---------|------------|
| 75  | [Sort Colors (Dutch National Flag)](https://leetcode.com/problems/sort-colors/) | Medium |
| 56  | [Merge Intervals](https://leetcode.com/problems/merge-intervals/) | Medium |

### Tree (DFS / BFS)

| # | Problem | Difficulty |
|---|---------|------------|
| 104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | Easy |
| 226 | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | Easy |
| 102 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Medium |

### Graph (Matrix DFS/BFS)

| # | Problem | Difficulty |
|---|---------|------------|
| 200 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) | Medium |

### Dynamic Programming

| # | Problem | Difficulty |
|---|---------|------------|
| 70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | Easy |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) | Medium |

### Greedy / Kadane

| # | Problem | Difficulty |
|---|---------|------------|
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy |
| 53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | Medium |

### Backtracking

| # | Problem | Difficulty |
|---|---------|------------|
| 78  | [Subsets](https://leetcode.com/problems/subsets/) | Medium |

### Bit Manipulation

| # | Problem | Difficulty |
|---|---------|------------|
| 136 | [Single Number](https://leetcode.com/problems/single-number/) | Easy |

## Topics Roadmap (FAANG)

A recommended order for interview prep, from foundational to advanced:

1. **Arrays & Hashing** — 1, 242, 49, 128, 560
2. **Two Pointers** — 167, 11, 15, 16, 42
3. **Sliding Window** — 209, 904, 567, 424, 713, 76
4. **Stack / Monotonic Stack** — 20, 739
5. **Binary Search** — 704, 33
6. **Linked List** — 21, 141, 206
7. **Trees (DFS/BFS)** — 104, 226, 102
8. **Heap / Top-K** — 703, 215, 347
9. **Sorting** — 75, 56
10. **Backtracking** — 78
11. **Graphs** — 200
12. **Dynamic Programming** — 70, 198, 121, 53
13. **Bit Manipulation** — 136

## Structure

```
{problem_number}/golang/
├── go.mod              # module problem{N}
├── solution.go         # implementation
└── solution_test.go    # table-driven tests
```

Each `solution.go` has a header block:

```go
// {N}. {Title}
// Difficulty: {Easy|Medium|Hard}
// Pattern:    {pattern-tag}
// Link:       https://leetcode.com/problems/...
```

## Run Tests

Run a single problem:

```bash
cd 1/golang && go test -v -run .
```

Run all problems:

```bash
for d in */golang; do (cd "$d" && go test ./...); done
```
