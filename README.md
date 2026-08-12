# Algorithm Practice with Go

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
go run ./cmd/merge-sort
go run ./cmd/two-sum
```

## Structure

- `algorithms/`: Main algorithm implementations, grouped by category.
- `cmd/`: Small runnable examples with `main`.
- `tests/`: Separated test packages, grouped by algorithm category.

Current layout:

```text
my-algorithm/
├── algorithms/
│   ├── brute-force/
│   │   └── two_sum.go
│   └── divide-conquer/
│       ├── binary_search.go
│       └── merge_sort.go
├── cmd/
│   ├── binary-search/
│   │   └── main.go
│   ├── merge-sort/
│   │   └── main.go
│   └── two-sum/
│       └── main.go
├── tests/
│   ├── brute-force/
│   │   └── two_sum_test.go
│   └── divide-conquer/
│       ├── binary_search_test.go
│       └── merge_sort_test.go
├── go.mod
└── README.md
```
