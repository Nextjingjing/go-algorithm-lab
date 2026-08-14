# Go Algorithm Learning Roadmap

This roadmap is designed to build long-term problem-solving ability and
job-ready Go skills without turning practice into a race. The goal is to truly
remember patterns, explain the reasoning, write tests, and build software that
other people can use.

> The rough timeline is 10–15 months at three sessions per week. Advance by
> demonstrated ability, not by calendar date. Going faster or slower is fine.

## Read Only This Section First

You are currently in the Foundation stage. The next exercise is:

1. Replay `Valid Anagram` from memory
   ([LC #242](https://leetcode.com/problems/valid-anagram/)).

Do not look at the entire roadmap every time. Choose one exercise at a time.

## Source Labels

Every exercise uses one of these labels:

- `LC #n`: a direct LeetCode problem, linked to the original problem.
- `Lab`: a problem designed for this repository; it is not claimed to be from
  LeetCode.
- `Project`: a small piece of work closer to real software development.
- `Euler #n`: a mathematical exercise from Project Euler for fun days.

Some exercises are selected from
[LeetCode 75](https://leetcode.com/studyplan/leetcode-75/) and
[Top Interview 150](https://leetcode.com/studyplan/top-interview-150/). The goal
is not to complete every item. Use only problems that teach a new pattern or
reinforce a current focus area.

## The Training System

### Weekly rhythm

The normal target is three sessions per week, 35–60 minutes per session:

| Session | Work | Goal |
| --- | --- | --- |
| A | One new exercise | Attempt it independently and write the invariant first |
| B | Continue the exercise or add tests | Clarify the contract and edge cases |
| C | Rebuild an older exercise from memory | Practice retrieval instead of rereading |
| Bonus | A puzzle, benchmark, or project | Keep practice enjoyable |

On a low-energy day, use a 15-minute `minimum session`: trace one input, write
one test case, or explain when a loop stops. That still counts as maintaining
the habit.

### Hint ladder

When starting a new exercise:

1. 0–15 minutes: understand the input/output and trace the examples by hand.
2. 15–30 minutes: write your own brute-force idea or pseudocode.
3. After 30 minutes: ask for hint level 1 — an edge case or guiding question.
4. After 45 minutes: ask for hint level 2 — a pattern or invariant.
5. After 60 minutes: ask for a walkthrough, then close it and rewrite the idea
   yourself.
6. The next day: replay without opening the original code.

Asking for help is not failure. The important part is preserving some time for
your own retrieval attempt. Research on retrieval practice found that actively
recalling knowledge supports long-term learning better than rereading alone
([Karpicke & Blunt, 2011](https://pubmed.ncbi.nlm.nih.gov/21252317/)).

### Spaced replay

After completing an exercise, replay it without looking at the answer on days
`+1`, `+7`, and `+21`. You do not need to retype the entire file every time;
tracing the algorithm, writing the core loop, or explaining the invariant is
enough. Distributed practice is generally more effective than cramming
([distributed-practice review](https://pmc.ncbi.nlm.nih.gov/articles/PMC12189222/)).

### Definition of done

An exercise is `done` when:

- behavior is correct for its contract and relevant edge cases;
- it has a Go doc comment describing behavior, mutation, and preconditions;
- it has table-driven tests whose cases the learner can explain;
- it has a small example in `cmd/<name>/main.go`;
- the learner can explain the loop boundary or base case without guessing;
- the learner can state time and space complexity, at least after a guiding
  question;
- the learner can replay it after seven days, or records that it needs review.

## Current Position

| Exercise | Source | Status | Next evidence needed |
| --- | --- | --- | --- |
| Linear Search | `Lab` | Done | replay without notes |
| Find Max | `Lab` | Done with precondition | explain why `*[]int` is unnecessary |
| Count Occurrences | `Lab` | Done | replay without notes |
| Is Sorted | `Lab` | Done | explain the `len(nums)-1` boundary |
| Reverse Slice | `Lab` | Done | replay and explain `left < right` |
| Contains Duplicate | [LC #217](https://leetcode.com/problems/contains-duplicate/) | Brute-force and hash-map versions done | replay without notes; explain map invariant |
| Two Sum | [LC #1](https://leetcode.com/problems/two-sum/) | Brute-force and hash-map versions done | replay without notes; explain why indexes differ |
| Valid Anagram | [LC #242](https://leetcode.com/problems/valid-anagram/) | Implementation, tests, and cmd example complete | replay without notes; explain character-count invariant |
| Binary Search | [LC #704](https://leetcode.com/problems/binary-search/) | Debugging checkpoint | add empty and exhausted-range tests |
| Merge Sort | `Lab` | First complete version | trace 2- and 3-element inputs by hand |
| Sum Array | `Lab` | Done | replay without notes |
| Move Zeroes | [LC #283](https://leetcode.com/problems/move-zeroes/) | Done with two pointers | replay without notes; explain invariant and `O(n)` |
| Remove Element | [LC #27](https://leetcode.com/problems/remove-element/) | Implementation, tests, and cmd example complete | replay without notes; explain read/write indexes and `nums[:k]` |

The agent helped with several `cmd/` examples, doc comments, and tests. The
next evidence should therefore be learner-owned test design, not immediately
moving to harder algorithms.

## Phase 0 — Foundation Reset

Approximate duration: 2–4 weeks. The goal is to make loops, slices, contracts,
and boundaries feel manageable rather than mentally expensive.

| Order | Exercise | Source | Main skill |
| --- | --- | --- | --- |
| 1 | Sum Array | `Lab` | accumulator and empty input |
| 2 | Min Value | `Lab` | preconditions and initialization |
| 3 | Move Zeroes | [LC #283](https://leetcode.com/problems/move-zeroes/) | in-place two pointers |
| 4 | Remove Element | [LC #27](https://leetcode.com/problems/remove-element/) | read/write indexes |
| 5 | Repair Binary Search | [LC #704](https://leetcode.com/problems/binary-search/) | base case and inclusive range |

Pass this phase when you can write test cases before the code for at least one
exercise and explain `<` versus `<=` from that exercise's index range.

## Phase 1 — Arrays, Maps, And Moving Windows

Approximate duration: 6–10 weeks. Add no more than two new exercises per week.

| Pattern | Exercises |
| --- | --- |
| Hash map | [Contains Duplicate — LC #217](https://leetcode.com/problems/contains-duplicate/), [Two Sum — LC #1](https://leetcode.com/problems/two-sum/), [Valid Anagram — LC #242](https://leetcode.com/problems/valid-anagram/) |
| Two pointers | [Valid Palindrome — LC #125](https://leetcode.com/problems/valid-palindrome/), [Two Sum II — LC #167](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) |
| Prefix/suffix | [Find Pivot Index — LC #724](https://leetcode.com/problems/find-pivot-index/), [Product Except Self — LC #238](https://leetcode.com/problems/product-of-array-except-self/) |
| Sliding window | [Best Time to Buy and Sell Stock — LC #121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/), [Longest Substring Without Repeating Characters — LC #3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) |

Pass this phase when you can propose a brute-force approach for an array
problem, then choose a map, two pointers, or a window and explain the trade-off.

## Phase 2 — Core Data Structures

Approximate duration: 8–12 weeks. Implement small structures yourself before
relying on a library abstraction.

| Topic | Exercises |
| --- | --- |
| Stack | [Valid Parentheses — LC #20](https://leetcode.com/problems/valid-parentheses/), [Min Stack — LC #155](https://leetcode.com/problems/min-stack/) |
| Queue | [Implement Queue Using Stacks — LC #232](https://leetcode.com/problems/implement-queue-using-stacks/), `QueueUsingSlice` (`Lab`) |
| Linked list | [Reverse Linked List — LC #206](https://leetcode.com/problems/reverse-linked-list/), [Merge Two Sorted Lists — LC #21](https://leetcode.com/problems/merge-two-sorted-lists/), [Linked List Cycle — LC #141](https://leetcode.com/problems/linked-list-cycle/) |
| Binary search | [Search Insert Position — LC #35](https://leetcode.com/problems/search-insert-position/), [Find Minimum in Rotated Sorted Array — LC #153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) |

Pass this phase when you can draw pointer/state transitions and write tests for
empty, one-element, cycle, and boundary cases.

## Phase 3 — Recursion, Trees, And Heaps

Approximate duration: 10–14 weeks. Trace the call stack on paper before coding.

| Topic | Exercises |
| --- | --- |
| Recursion warm-up | `Factorial`, `RecursiveSum`, `RecursiveBinarySearch` (`Lab`) |
| Tree DFS | [Maximum Depth — LC #104](https://leetcode.com/problems/maximum-depth-of-binary-tree/), [Same Tree — LC #100](https://leetcode.com/problems/same-tree/), [Invert Binary Tree — LC #226](https://leetcode.com/problems/invert-binary-tree/) |
| Tree BFS/BST | [Level Order Traversal — LC #102](https://leetcode.com/problems/binary-tree-level-order-traversal/), [Validate BST — LC #98](https://leetcode.com/problems/validate-binary-search-tree/) |
| Heap | [Kth Largest Element — LC #215](https://leetcode.com/problems/kth-largest-element-in-an-array/), [Kth Largest in a Stream — LC #703](https://leetcode.com/problems/kth-largest-element-in-a-stream/) |

Pass this phase when you can identify the base case, smaller subproblem, and
combine step without saying that it “just recurses.”

## Phase 4 — Graphs, Backtracking, And Greedy

Approximate duration: 10–14 weeks. Alternate visual/grid problems with
combinatorics problems to keep practice varied.

| Topic | Exercises |
| --- | --- |
| Graph/grid | [Flood Fill — LC #733](https://leetcode.com/problems/flood-fill/), [Number of Islands — LC #200](https://leetcode.com/problems/number-of-islands/), [Rotting Oranges — LC #994](https://leetcode.com/problems/rotting-oranges/) |
| Backtracking | [Subsets — LC #78](https://leetcode.com/problems/subsets/), [Permutations — LC #46](https://leetcode.com/problems/permutations/), [Combination Sum — LC #39](https://leetcode.com/problems/combination-sum/) |
| Greedy | [Maximum Subarray — LC #53](https://leetcode.com/problems/maximum-subarray/), [Jump Game — LC #55](https://leetcode.com/problems/jump-game/) |

Pass this phase when you can explain visited state, branching choices, and
undo/backtrack, and choose BFS or DFS based on the problem's structure.

## Phase 5 — Dynamic Programming Without Fear

Approximate duration: 8–12 weeks. Do not begin with a formula. Begin with the
recursion tree and ask: “What must this state know from the past?”

| Level | Exercises |
| --- | --- |
| Warm-up | [Climbing Stairs — LC #70](https://leetcode.com/problems/climbing-stairs/), [House Robber — LC #198](https://leetcode.com/problems/house-robber/) |
| Core | [Unique Paths — LC #62](https://leetcode.com/problems/unique-paths/), [Coin Change — LC #322](https://leetcode.com/problems/coin-change/) |
| Stretch | [Longest Increasing Subsequence — LC #300](https://leetcode.com/problems/longest-increasing-subsequence/), [Longest Common Subsequence — LC #1143](https://leetcode.com/problems/longest-common-subsequence/) |

Pass this phase when you can derive a recurrence from small examples and turn it
into memoization or a table while explaining the meaning of every state.

## Parallel Track — Go For Real Work

From the end of Phase 0 onward, spend roughly 25–30% of practice time on real
software skills. Algorithms alone are not enough for a Go job.

| Milestone | Work skills | Deliverable |
| --- | --- | --- |
| A | packages, errors, interfaces, table tests, benchmarks | make this repository's contracts consistent and benchmark two algorithms |
| B | `net/http`, JSON, `context`, graceful shutdown | `Project`: Algorithm Runner API |
| C | `database/sql`, migrations, configuration, structured logging | `Project`: Practice Tracker API |
| D | goroutines, channels, mutexes, race detector | `Project`: bounded worker queue with tests |
| E | Docker, CI, profiling, README architecture notes | a capstone that clones and runs with one command |

Every project needs error handling, tests, README usage, and a short design note.
Keep it small, complete, and runnable.

## Math And Fun Lane

Once a month, choose one optional activity. It does not count as unfinished work:

- `Euler #1`: sum of numbers divisible by 3 or 5.
- `Euler #2`: a Fibonacci-sum variant.
- `Euler #6`: difference between the square of a sum and the sum of squares.
- Write a text animation that visualizes merge sort or BFS.
- Benchmark two implementations and predict the result before running them.

Math can be introduced gradually: logarithms for complexity, modular arithmetic,
combinatorics, basic probability, graph terminology, and recurrence relations.

## Checkpoints, Not Grades

| Rank | Observable ability |
| --- | --- |
| Bronze | solve Easy loop/slice problems and write edge-case tests |
| Silver | recognize five patterns and explain time/space trade-offs |
| Gold | solve a familiar-pattern Medium in 45–60 minutes and replay it later |
| Builder | ship a small Go service with tests, errors, docs, and CI |
| Interview-ready | solve a mixed set while communicating reasoning, tests, and complexity |

If you are stuck at a checkpoint, reduce new exercises and increase replay. Do
not respond by adding more hours.

## Agent Review Rules

When new evidence appears, agents must use the `go-algorithm-coach` skill and:

- update this file only when an exercise passes the Definition of done;
- update `docs/learner-profile.md` from evidence, not assumptions;
- distinguish learner-written work from agent-generated docs, tests, and `cmd/`;
- give hints progressively and never reveal the algorithm until explicitly asked;
- suggest one next action that can be completed in a single session.
