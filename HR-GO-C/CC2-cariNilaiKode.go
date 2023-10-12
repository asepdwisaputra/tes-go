package main

import "strconv"

func getQueryAnswers(cache_entries [][]string, queries [][]string) []int32 {
	// Write your code here
	var cache = make(map[string]map[string]string)

	for _, entry := range cache_entries {
		var timeStamp string = entry[0]
		var key string = entry[1]
		var value string = entry[2]

		if _, ok := cache[key]; !ok {
			cache[key] = make(map[string]string)
		}
		cache[key][timeStamp] = value
	}

	var result []int32

	for _, query := range queries {
		var key string = query[0]
		var timeStamp string = query[1]

		value, ok := cache[key][timeStamp]
		if !ok {

			result = append(result, -1)
		} else {

			intvalue, _ := strconv.Atoi(value)
			result = append(result, int32(intvalue))
		}
	}

	return result
}

func main() {

}
