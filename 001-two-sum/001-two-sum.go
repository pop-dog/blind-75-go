package main

func twoSum(nums []int, target int) []int {
	// Create a map to store the diffs from the target
	diffs := make(map[int]int)
	for i, num := range nums {
		diffs[target-num] = i
	}
	for i, num := range nums {
		// Second pass: see if what we need is in diffs
		idx, found := diffs[num]
		if found && i != idx {
			return []int{i, idx}
		}
	}
	return []int{}
}

func main() {
	x := twoSum([]int{2, 7, 11, 15}, 9)
	println("x:", x)
}
