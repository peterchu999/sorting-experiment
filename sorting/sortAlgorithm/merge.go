package sortAlgorithm


func mergeV1(left []int, right []int) []int { 
    newList := []int{}
    p1 := 0
    p2 := 0
    for p1 < len(left) || p2 < len(right) {
        newInt := 0
        if (p1 >= len(left)) {
            newInt = right[p2]
            p2 += 1
        } else if ( p2 >= len(right)) {
            newInt = left[p1]
            p1 += 1
        } else {
            if (left[p1] < right[p2]) {
                newInt = left[p1]
                p1 += 1
            } else {
                newInt = right[p2]
                p2 += 1
			}
        }
        newList = append(newList, newInt)
    }
    return newList
}

func merge(left []int, right []int) []int { 
	lLeft := len(left)
	lRight := len(right)
    newList := make([]int, lLeft + lRight)
    p1 := 0
    p2 := 0
	
    for p1 < lLeft || p2 < lRight {
        newInt := 0
        if (p1 >= len(left)) {
            newInt = right[p2]
            p2 += 1
        } else if ( p2 >= lRight) {
            newInt = left[p1]
            p1 += 1
        } else {
            if (lLeft < lRight) {
                newInt = left[p1]
                p1 += 1
            } else {
                newInt = right[p2]
                p2 += 1
			}
        }
        newList[p1+p2-1] = newInt
    }
    return newList
}

func divide(arr []int) ([]int, []int) {
	length := len(arr)
	p1 := length / 2
	return arr[:p1], arr[p1:]	
}

func MergeSortV1(list []int)[]int {
	a, b := divide(list)
	if len(a) < 2 && len(b) < 2 {
		return mergeV1(a,b)
	} 
	return mergeV1(
		MergeSort(a),
		MergeSort(b),
	)
}

func MergeSortV1Void(list []int) {
	MergeSortV1(list)
}

func MergeSort(list []int)[]int {
	a, b := divide(list)
	if len(a) < 2 && len(b) < 2 {
		return merge(a,b)
	} 
	return merge(
		MergeSort(a),
		MergeSort(b),
	)
}

func MergeSortVoid(list []int) {
	MergeSort(list)
}