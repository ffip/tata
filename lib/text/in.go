package text

import "sort"

// InStringArray ==> 获取sting在[]sting中的位置(不存在返回-1)
func InStringArray(target string, str_array []string) (index int) {
	sort.Strings(str_array)
	index = sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return index
	}
	return -1
}
