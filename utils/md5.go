package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

const saltBit = uint(8)
const saltShift = uint(8)
const increaShift = saltBit + saltShift

type Mist struct {
	sync.Mutex
	increase int64
	saltA    int64
	saltB    int64
}

// 构造函数
func NewMist() *Mist {
	mist := Mist{increase: 1}
	return &mist
}

//生成唯一编号
func (c *Mist) Generate() int64 {
	c.Lock()
	c.increase++
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = randB.Int64()

	mist := int64((c.increase << increaShift) | (c.saltA << saltShift) | c.saltB)
	c.Unlock()
	return mist
}

func main() {
	mist := NewMist()
	for i := 0; i < 3; i++ {
		fmt.Println(mist.Generate())
	}
}
