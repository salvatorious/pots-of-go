package memo

import (
	"fmt"
	"strings"

	types "github.com/salvatorious/pots-of-go/types"
)

// Func ...
type Func func(key []int) types.Solution

// Memo ...
type Memo struct {
	f     Func
	cache map[string]result
}

type result struct {
	value types.Solution
}

// New ...
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get ...
func (memo *Memo) Get(key []int) types.Solution {
	keyString := arrayToString(key, ",")
	res, ok := memo.cache[keyString]

	if !ok {
		res.value = memo.f(key)
		memo.cache[keyString] = res
	}

	return res.value
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
