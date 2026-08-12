# Algorithm Practice with Go

Repository: <https://github.com/Nextjingjing/go-algorithm-lab>

## Progress Dashboard

The progress dashboard is available at `docs/index.html`.

GitHub Pages URL:
<https://nextjingjing.github.io/go-algorithm-lab/>

To show it on GitHub, enable GitHub Pages:

1. Open repository `Settings`.
2. Go to `Pages`.
3. Set source to `Deploy from a branch`.
4. Select the `main` branch and `/docs` folder.

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
go run ./cmd/find-max
go run ./cmd/linear-search
go run ./cmd/merge-sort
go run ./cmd/two-sum
```

## Structure

- `algorithms/`: Main algorithm implementations, grouped by category.
- `cmd/`: Small runnable examples with `main`.
- `docs/`: Static pages for GitHub Pages, including the progress dashboard.
- `tests/`: Separated test packages, grouped by algorithm category.

Current layout:

```text
go-algorithm-lab/
├── algorithms/
│   ├── brute-force/
│   │   ├── find_max.go
│   │   ├── linear_search.go
│   │   └── two_sum.go
│   └── divide-conquer/
│       ├── binary_search.go
│       └── merge_sort.go
├── cmd/
│   ├── binary-search/
│   │   └── main.go
│   ├── find-max/
│   │   └── main.go
│   ├── linear-search/
│   │   └── main.go
│   ├── merge-sort/
│   │   └── main.go
│   └── two-sum/
│       └── main.go
├── docs/
│   └── index.html
├── tests/
│   ├── brute-force/
│   │   ├── find_max_test.go
│   │   ├── linear_search_test.go
│   │   └── two_sum_test.go
│   └── divide-conquer/
│       ├── binary_search_test.go
│       └── merge_sort_test.go
├── go.mod
└── README.md
```
