package chief

import (
	"math/big"
	"testing"
)

func TestA(t *testing.T) {
	var (
		fn func(_vrfn *big.Int) int
		nl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		//vl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		vl = []int{2, 4, 6, 8}
	)

	for i := 10; i < 20; i++ {
		vrfn := big.NewInt(int64(i))
		fn = func(_vrfn *big.Int) int {
			m := big.NewInt(int64(len(nl)))
			x := new(big.Int).Sub(_vrfn, vrfn)
			if x.Cmp(m) >= 0 {
				return 0
			}
			idx := new(big.Int).Mod(_vrfn, m)
			// skip if `n` in volunteer list
			v := nl[idx.Int64()]
			for _, vol := range vl {
				if vol == v {
					return fn(new(big.Int).Add(_vrfn, big.NewInt(1)))
				}
			}
			return v
		}
		r := fn(vrfn)
		t.Log(vrfn, "-->", r)
	}
}
