package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/peterchu999/sort/sortAlgorithm"
	"github.com/peterchu999/sort/utils"
)



func run_test(arrLen int) {
	mergeTimeTracker := []int64{}	
	bubbleTimeTracker := []int64{}
	builtInTimeTracker := []int64{}
	quickTimeTracker := []int64{}

	for i:=1; i < arrLen+1; i++ {
		testCase := utils.GenerateUniqueRandomNumbers(i, 100)
		mergeT := utils.TestSorting(sortAlgorithm.MergeSortVoid, testCase)
		time.Sleep(time.Millisecond * 5) // incase computation degenerate machine and make later case worst
		bubbleT := utils.TestSorting(sortAlgorithm.BubbleSortVoid, testCase)
		time.Sleep(time.Millisecond * 5) // incase computation degenerate machine and make later case worst
		quickT := utils.TestSorting(sortAlgorithm.QuickSortVoid, testCase)
		time.Sleep(time.Millisecond * 5) // incase computation degenerate machine and make later case worst
		biT :=  utils.TestSorting(sort.Ints, testCase)

		// fmt.Printf("%v. merge: %v , bubble: %v , sort: %v\n", i, mergeT,bubbleT,biT)
		mergeTimeTracker = append(mergeTimeTracker,mergeT)
		bubbleTimeTracker = append(bubbleTimeTracker,bubbleT)
		builtInTimeTracker = append(builtInTimeTracker,biT)
		quickTimeTracker = append(quickTimeTracker, quickT)
	}
	maps := map[string][]int64 {}
	maps["merge"] = mergeTimeTracker
	maps["bubble"] = bubbleTimeTracker
	maps["sort"] = builtInTimeTracker
	maps["quick"] = quickTimeTracker

	utils.Save_json(maps, "sort-test")
	fmt.Println("Done...")
}

func getArrLen() int {
	args := os.Args[1:]
	if len(args) > 0 {
		a := args[0]
		arrLenFromArg, err := strconv.Atoi(a) 
		if err == nil {
			return arrLenFromArg
		} 
	}
	return 300
}


func main() {
	arrLen := getArrLen()
	run_test(arrLen)
}