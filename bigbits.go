package bigbits

import (
	"container/list"
	"math/big"
	"strings"
)

func Difference(x, y *big.Int) *big.Int {
	return x.AndNot(x, y)
}

func Merge(x, y *big.Int) *big.Int {
	return x.Or(x, y)
}

func SetBit(x *big.Int, bit int, value uint) *big.Int {
	/*if x.Bit(bit) != 0 {
		fmt.Println("%v was already 1\n", bit)
	}*/
	return x.SetBit(x, bit, value)
}

func IndicesToBinary(array []int) *big.Int {
	result := big.NewInt(0)
	for _, v := range array {
		result = result.SetBit(result, v, 1)
	}
	return result
}

func BinaryToIndices(x *big.Int) []int {
	L := list.New()
	count := 0
	for i := 0; i < x.BitLen(); i++ {
		if x.Bit(i) == 1 {
			L.PushBack(i)
			count++
		}
	}
	/*
		i := 0

		for x.Sign() != 0 {

			if x.Bit(0) == 1 {
				L.PushBack(i)
				count++
			}
			x.Rsh(x, 1)
			i++
		}
	*/
	result := make([]int, count)
	j := 0
	for item := L.Front(); item != nil; item = item.Next() {
		result[j] = item.Value.(int)
		j++
	}
	return result
}

func AllocateNew(size int) *big.Int {
	m := big.NewInt(0)
	m.SetString(strings.Repeat("0", size), 2)
	return m
}
