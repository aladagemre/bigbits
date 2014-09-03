package bigbits

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMerge(t *testing.T) {
	/*const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}*/

	x := big.NewInt(0)
	p := x.SetBit(x, 3, 1)
	z := x.Lsh(p, 4)
	fmt.Println(z)

	m := big.NewInt(0)
	m.SetString("10011001", 2)
	n := big.NewInt(0)
	n.SetString("11110010", 2)

	fmt.Printf("Result: %v, %v\n", m, n)
	fmt.Println()
	h := AllocateNew(2500)
	h = h.SetBit(h, 2400, 1)
	//fmt.Printf("2500: %v\n", h)

}

func TestBinaryToIndices(t *testing.T) {
	m := big.NewInt(0)
	m.SetString("10011001", 2)
	res := BinaryToIndices(m)
	if res[0] != 0 || res[1] != 3 || res[2] != 4 || res[3] != 7 {
		t.Errorf("Binary mismatch: %v", res)
	}

	n := big.NewInt(0)
	n.SetString("1001000", 2)
	res = BinaryToIndices(n)
	if res[0] != 3 || res[1] != 6 {
		t.Errorf("Binary mismatch: %v", res)
	}
}
