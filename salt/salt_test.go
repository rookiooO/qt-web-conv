package salt

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	sz, _ := Encrypt("admin123")
	fmt.Println(sz)
}
