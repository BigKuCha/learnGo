package funcs

import "testing"

func TestIsContains(t *testing.T) {
	tests := []struct {
		str  string
		dst  string
		want bool
	}{
		{"abcdefg", "cde", true},
		{"中华人民共和国万岁", "万岁", true},
		{"学而思教育科技有限公司", "乐外教", false},
		{"红尘中，你的无上清凉", "红尘", false}, //报错
	}

	for _, test := range tests {
		if got := IsContains(test.str, test.dst); got != test.want {
			t.Errorf("IsContains('%s', '%s');want %t, got %t ", test.str, test.dst, test.want, got)
		}
	}
}
