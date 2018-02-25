package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func main() {
	input, err := get_input()
	if err != nil {
		panic(err)
	}
	max, sorted := get_max_and_sorted(input)
	output(max, sorted, input)
}

func get_input() (map[string]interface{}, error) {
	b, _ := ioutil.ReadAll(os.Stdin)
	input := make(map[string]interface{})
	if err := json.Unmarshal(b, &input); err != nil {
		return nil, err
	}
	return input, nil // e.g. map[string]string{"hello":"world"}
}

func get_max_and_sorted(input map[string]interface{}) (int, []string) {
	max, sorted := 0, make([]string, 0, len(input))
	for key := range input {
		if n := len(key); n > max {
			max = n
		}
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	return max, sorted // e.g. 5, []string{"hello"}
}

func output(max int, sorted []string, input map[string]interface{}) {
	pad := strings.Repeat(" ", max)
	fmt.Println("{")
	for _, key := range sorted {
		fmt.Printf("\t%s: %s%v,\n", key, pad[:len(pad)-len(key)], input[key])
	}
	fmt.Println("}")
}
