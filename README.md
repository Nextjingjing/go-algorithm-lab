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
```

## Structure

- `algorithms/`: โค้ดอัลกอริทึมหลัก แยกตามหมวด
- `cmd/`: โปรแกรมตัวอย่างที่มี `main`
- `tests/`: แบบทดสอบที่แยกตามหมวดของอัลกอริทึม

Current layout:

```text
my-algorithm/
├── algorithms/
│   └── divide-conquer/
│       ├── binary_search.go
│       └── merge_sort.go
├── cmd/
│   ├── binary-search/
│   │   └── main.go
│   └── merge-sort/
│       └── main.go
├── tests/
│   └── divide-conquer/
│       ├── binary_search_test.go
│       └── merge_sort_test.go
├── go.mod
└── README.md
```
