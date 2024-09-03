package sortthejumblednumbers

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		originalNum int
		mappedNum   int
		index       int
	}

	mapNumber := func(num int) int {
		if num == 0 {
			return mapping[0]
		}
		var mapped int
		multiplier := 1
		for num > 0 {
			digit := num % 10
			mapped += mapping[digit] * multiplier
			multiplier *= 10
			num /= 10
		}
		return mapped
	}

	mappedValues := make(map[int][]pair)

	// 数値をマッピングしてマップに追加
	for i, num := range nums {
		mapped := mapNumber(num)
		mappedValues[mapped] = append(mappedValues[mapped], pair{num, mapped, i})
	}

	// マッピングされた値をソート
	var sortedMapped []int
	for mapped := range mappedValues {
		sortedMapped = append(sortedMapped, mapped)
	}
	sort.Ints(sortedMapped)

	// 結果を構築
	result := make([]int, 0, len(nums))
	for _, mapped := range sortedMapped {
		pairs := mappedValues[mapped]
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].index < pairs[j].index
		})
		for _, p := range pairs {
			result = append(result, p.originalNum)
		}
	}

	return result
}
