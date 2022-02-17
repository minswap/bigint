# bigint

Go's `big.Int` is mutable to enable flexibility in performance tuning but sometimes immutability is more desired. To write immutable code with `big.Int` we need to be really verbose like `new(big.Int).Add(a, b)`. This package provides an immutable wrapper around `big.Int`. Other than the immutability, it tries to be as similar to `big.Int` as possible. Examples:

| int                 | big.Int                                     | bigint                       |
| ------------------- | ------------------------------------------- | ---------------------------- |
| 42                  | big.NewInt(42)                              | bigint.New(42)               |
| strconv.Atoi("123") | new(big.Int).SetString("123", 10)           | bigint.FromString("123", 10) |
| a+b                 | new(big.Int).Add(a, b)                      | a.Add(b)                     |
| a\*(b+c)            | new(big.Int).Mul(a, new(big.Int).Add(b, c)) | a.Mul(b.Add(c))              |
| a == b              | a.Cmp(b) == 0                               | a.EQ(b)                      |
| a < b               | a.Cmp(b) < 0                                | a.LT(b)                      |

Check documentation for all utility functions: to be updated...

## Install

```
go get github.com/minswap/bigint
```

## Roadmap

- Support all functions from `big.Int`
- 100% test coverage
