package tool

import "fmt"

/**
 * @Author: yucky
 * @Date: 2024/1/2 21:17
 * @Desc:
 */

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func FuckNumberOne(num int64) {
	fmt.Println("test-FuckNumberOne:", num)
}
