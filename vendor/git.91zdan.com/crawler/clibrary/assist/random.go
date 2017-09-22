package assist

import (
	"time"
	"math/rand"
)

// 获取设置了种子的 *rand.Rand
func RandSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
