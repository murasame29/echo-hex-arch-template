package helpers

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// 最小値と最大値を渡したときランダムな整数をかえす
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// 整数を与えたとき、その文字列の長さのランダム文字列をかえす
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// 長さ6文字の名前をかえす
func RandomOwner() string {
	return RandomString(6)
}

// 0 ~ 1000　のランダムな金額をかえす
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10))
}
