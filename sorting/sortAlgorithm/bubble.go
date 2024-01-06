package sortAlgorithm

func BubbleSort(list []int)[]int {
	for i:=0; i < len(list); i++ {
		for j:=1; j < len(list) - i; j++ {
			if list[j] < list[j-1] {
				temp := list[j-1]
				list[j-1] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}

func BubbleSortVoid(list []int) {
	BubbleSort(list)
}