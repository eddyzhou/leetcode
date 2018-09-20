package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range(nums) {
		v := target - n
		if idx, ok := m[v]; ok {
			return []int{idx, i}
		}

		m[n] = i
	}
	panic("no result")
}

//func main() {
//	nums := []int{2, 7, 11, 15}
//	taget := 9
//	fmt.Println(twoSum(nums, taget))
//}
