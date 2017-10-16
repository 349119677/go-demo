package clibrary

import (
	"fmt"
	"testing"
)

func TestAuthHeader(t *testing.T) {
	header := MaYiProxy{AppKey: "276536"}.authHeader()
	fmt.Println(header)
}
