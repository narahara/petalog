package hash

import (
	"fmt"
	"hash/fnv"
)

// Fowler-Noll-Voハッシュ関数を使用して文字列のハッシュ値を生成する
func Create32Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	sum := h.Sum32()
	return fmt.Sprintf("%08x", sum)
}
