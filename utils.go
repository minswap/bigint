package bigint

func Min(a ...BigInt) BigInt {
	if len(a) == 0 {
		panic("bigint.Min: input is an empty array")
	}
	ret := a[0]
	for i := 1; i < len(a); i++ {
		if a[i].LT(ret) {
			ret = a[i]
		}
	}
	return ret
}

func Max(a ...BigInt) BigInt {
	if len(a) == 0 {
		panic("bigint.Max: input is an empty array")
	}
	ret := a[0]
	for i := 1; i < len(a); i++ {
		if a[i].GT(ret) {
			ret = a[i]
		}
	}
	return ret
}

// Sum returns 0 if the input array is empty
func Sum(a ...BigInt) BigInt {
	ret := New(0)
	for _, x := range a {
		ret.Add(x)
	}
	return ret
}

// Product returns 1 if the input array is empty
func Product(a ...BigInt) BigInt {
	ret := New(1)
	for _, x := range a {
		ret.Mul(x)
	}
	return ret
}
