package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type SortFunction func(list []int) 

func GenerateUniqueRandomNumbers(n, max int) []int {
   rand.Seed(time.Now().UnixNano())
   var result []int
   for len(result) < n {
    value := rand.Intn(max)
   	result = append(result, value)
   }
   return result
}

func Save_json(data interface{}, name string) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(fmt.Sprintf("%v.json",name), file, 0644)
}

func TestSorting(sfc SortFunction, testList []int) int64 {
	scopedTestList := make([]int, len(testList))
	copy(scopedTestList, testList)
	startTimeMerge := time.Now()
	sfc(scopedTestList)
	endTimeMerge := time.Now()
	return endTimeMerge.Sub(startTimeMerge).Nanoseconds()
}