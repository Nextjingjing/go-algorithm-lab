# Algorithm Practice with Go

Repository: <https://github.com/Nextjingjing/go-algorithm-lab>

## Learning Plan

The practice roadmap is available at [`docs/learning-plan.md`](docs/learning-plan.md).

The evidence-based progress review is available at
[`docs/learner-profile.md`](docs/learner-profile.md).

## AI Coaching

Codex and Claude share the `go-algorithm-coach` project skill. It reviews the
learner's own implementation, maintains Go doc comments, tests and runnable
examples when requested, and updates progress without revealing a full
algorithm solution unless explicitly asked.

- Canonical skill: `.agents/skills/go-algorithm-coach/SKILL.md`
- Claude project path: `.claude/skills/go-algorithm-coach/SKILL.md`

## Run tests

Run only separated tests:

```bash
go test ./tests/...
```

Show each test case:

```bash
go test -v ./tests/...
```

Run tests without using cache:

```bash
go test -count=1 ./tests/...
```

Show each test case and skip cache:

```bash
go test -v -count=1 ./tests/...
```

Run every package in the project:

```bash
go test ./...
```

When using `go test ./...`, Go may show `[no test files]` for packages such as
`cmd/binary-search` or `algorithms/divide-conquer`. That is normal because the
actual tests are stored under `tests/`.

When Go prints `(cached)`, it means the tests passed using a previous result
because the related files have not changed.

## Run an example

```bash
go run ./cmd/binary-search
go run ./cmd/contains-duplicate
go run ./cmd/hash-map
go run ./cmd/count-occurrences
go run ./cmd/find-max
go run ./cmd/is-sorted
go run ./cmd/linear-search
go run ./cmd/merge-sort
go run ./cmd/reverse-slice
go run ./cmd/two-sum
go run ./cmd/two-sum-map
go run ./cmd/valid-anagram
go run ./cmd/sum-array
go run ./cmd/move-zeroes
go run ./cmd/remove-element
go run ./cmd/valid-palindrome
go run ./cmd/two-sum-sorted
go run ./cmd/pivot-index
go run ./cmd/product-except-self
go run ./cmd/highest-altitude
go run ./cmd/left-right-difference
go run ./cmd/best-time-to-buy-and-sell-stock
go run ./cmd/best-time-to-buy-and-sell-stock-sliding
go run ./cmd/max-sum-subarray
go run ./cmd/max-consecutive-ones
go run ./cmd/longest-substring-without-repeating-characters
go run ./cmd/valid-parentheses
go run ./cmd/min-stack
go run ./cmd/queue-using-stacks
go run ./cmd/queue-using-slice
go run ./cmd/stack-using-queues
```

## Structure

- `algorithms/`: Main algorithm implementations, grouped by category.
- `algorithms/hash-map/`: Hash-map based algorithm implementations.
- `algorithms/greedy/`: Greedy algorithms that make locally optimal choices.
- `algorithms/stack/`: Stack-based algorithms for nested or ordered state.
- `algorithms/queue/`: Queue-based algorithms and FIFO data structures.
- `cmd/`: Small runnable examples with `main`.
- `docs/`: Learning plan and practice notes.
- `tests/`: Separated test packages, grouped by algorithm category.

Current layout:

```text
go-algorithm-lab/
├── algorithms/
│   ├── brute-force/
│   │   ├── contains_duplicate.go
│   │   ├── count_occurrences.go
│   │   ├── find_max.go
│   │   ├── is_sorted.go
│   │   ├── linear_search.go
│   │   ├── move_zeroes.go
│   │   ├── reverse_slice.go
│   │   ├── sum_array.go
│   │   └── two_sum.go
│   ├── divide-conquer/
│   │   ├── binary_search.go
│   │   └── merge_sort.go
│   ├── greedy/
│   │   └── best_time_to_buy_and_sell_stock.go
│   ├── prefix-suffix/
│   │   ├── highest_altitude.go
│   │   ├── left_right_difference.go
│   │   ├── pivot_index.go
│   │   └── product_except_self.go
│   ├── sliding-window/
│   │   ├── best_time_to_buy_and_sell_stock.go
│   │   ├── longest_substring_without_repeating_characters.go
│   │   ├── max_consecutive_ones.go
│   │   └── max_sum_subarray.go
│   ├── stack/
│   │   ├── implement_stack_using_queues.go
│   │   ├── min_stack.go
│   │   └── valid_parentheses.go
│   ├── queue/
│   │   ├── implement_queue_using_stacks.go
│   │   └── queue_using_slice.go
│   └── two-pointers/
│       ├── move_zeroes.go
│       ├── remove_element.go
│       ├── valid_palindrome.go
│       └── two_sum_sorted.go
├── cmd/
│   ├── best-time-to-buy-and-sell-stock/
│   │   └── main.go
│   ├── best-time-to-buy-and-sell-stock-sliding/
│   │   └── main.go
│   ├── max-sum-subarray/
│   │   └── main.go
│   ├── max-consecutive-ones/
│   │   └── main.go
│   ├── longest-substring-without-repeating-characters/
│   │   └── main.go
│   ├── binary-search/
│   │   └── main.go
│   ├── contains-duplicate/
│   │   └── main.go
│   ├── count-occurrences/
│   │   └── main.go
│   ├── find-max/
│   │   └── main.go
│   ├── is-sorted/
│   │   └── main.go
│   ├── linear-search/
│   │   └── main.go
│   ├── move-zeroes/
│   │   └── main.go
│   ├── remove-element/
│   │   └── main.go
│   ├── pivot-index/
│   │   └── main.go
│   ├── product-except-self/
│   │   └── main.go
│   ├── highest-altitude/
│   │   └── main.go
│   ├── left-right-difference/
│   │   └── main.go
│   ├── two-sum-sorted/
│   │   └── main.go
│   ├── valid-palindrome/
│   │   └── main.go
│   ├── merge-sort/
│   │   └── main.go
│   ├── reverse-slice/
│   │   └── main.go
│   ├── sum-array/
│   │   └── main.go
│   ├── two-sum/
│   │   └── main.go
│   ├── two-sum-map/
│   │   └── main.go
│   └── valid-anagram/
│       └── main.go
├── docs/
│   ├── learner-profile.md
│   └── learning-plan.md
├── tests/
│   ├── brute-force/
│   │   ├── contains_duplicate_test.go
│   │   ├── count_occurrences_test.go
│   │   ├── find_max_test.go
│   │   ├── is_sorted_test.go
│   │   ├── linear_search_test.go
│   │   ├── reverse_slice_test.go
│   │   └── two_sum_test.go
│   ├── divide-conquer/
│   │   ├── binary_search_test.go
│   │   └── merge_sort_test.go
│   ├── greedy/
│   │   └── best_time_to_buy_and_sell_stock_test.go
│   ├── prefix-suffix/
│   │   ├── highest_altitude_test.go
│   │   ├── left_right_difference_test.go
│   │   ├── pivot_index_test.go
│   │   └── product_except_self_test.go
│   ├── sliding-window/
│   │   ├── best_time_to_buy_and_sell_stock_test.go
│   │   ├── longest_substring_without_repeating_characters_test.go
│   │   ├── max_consecutive_ones_test.go
│   │   └── max_sum_subarray_test.go
│   ├── stack/
│   │   ├── implement_stack_using_queues_test.go
│   │   ├── min_stack_test.go
│   │   └── valid_parentheses_test.go
│   ├── queue/
│   │   ├── implement_queue_using_stacks_test.go
│   │   └── queue_using_slice_test.go
│   └── two-pointers/
│       ├── move_zeroes_test.go
│       ├── remove_element_test.go
│       ├── two_sum_sorted_test.go
│       └── valid_palindrome_test.go
├── go.mod
└── README.md
```
