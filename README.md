# bigint

Go's `big.Int` is mutable to enable flexibility in performance tuning but sometimes immutability is more desired. To write immutable code with `big.Int` we need to be really verbose like `new(big.Int).Add(a, b)`. This package provides an immutable wrapper around `big.Int`. Other than the immutability, it tries to be as similar to `big.Int` as possible. Examples:

| `int`                 | `big.Int`                                     | `bigint`                       |
| --------------------- | --------------------------------------------- | ------------------------------ |
| `42`                  | `big.NewInt(42)`                              | `bigint.New(42)`               |
| `strconv.Atoi("123")` | `new(big.Int).SetString("123", 10)`           | `bigint.FromString("123", 10)` |
| `a := b`              | `a := new(big.Int).Set(b)`                    | `a := b`                       |
| `a := b + c`          | `a := new(big.Int).Add(b, c)`                 | `a := b.Add(c)`                |
| `a += b`              | `a.Add(a, b)`                                 | `a = a.Add(b)`                 |
| `a + b * c`           | `new(big.Int).Add(a, new(big.Int).Mul(b, c))` | `a.Add(b.Mul(c))`              |
| `(a + b) * c`         | `new(big.Int).Mul(new(big.Int).Add(a, b), c)` | `a.Add(b).Mul(c)`              |
| `a == b`              | `a.Cmp(b) == 0`                               | `a.EQ(b)`                      |
| `a < b`               | `a.Cmp(b) < 0`                                | `a.LT(b)`                      |

Check documentation for all utility functions: to be updated...

## Install

```
go get github.com/minswap/bigint
```

## Roadmap

- Support all functions from `big.Int`
- 100% test coverage
