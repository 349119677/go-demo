package clibrary

import (
	"fmt"
	"github.com/golibs/uuid"
)

// 生成一个爬取唯一标识，使用 UUID 算法
func CrawlIdGen() string {
	rand := uuid.Rand()
	x := [16]byte(rand)
	return fmt.Sprintf("%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		x[0], x[1], x[2], x[3], x[4],
		x[5], x[6],
		x[7], x[8],
		x[9], x[10], x[11], x[12], x[13], x[14], x[15])
}
