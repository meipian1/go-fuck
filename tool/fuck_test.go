package tool

import "testing"

/**
 * @Author: yucky
 * @Date: 2024/1/2 21:17
 * @Desc:
 */

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
