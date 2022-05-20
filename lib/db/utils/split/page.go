package split

import (
	"strconv"
)

// SizeToLimit ==> get a page start and end row num.
func SizeToLimit(pn, size float64) (start, end string) {
	// 分页参数计算当前页起点值
	pn = (pn - 1) * size
	// 分页参数转类型
	start = strconv.FormatFloat(pn, 'f', -1, 32)
	end = strconv.FormatFloat(size, 'f', -1, 32)
	return
}
