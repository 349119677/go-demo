package clibrary

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("Random PC UA: %s \n", userAgentPC())
		fmt.Printf("Random Mobile UA: %s \n", userAgentMobile())
	}
}
