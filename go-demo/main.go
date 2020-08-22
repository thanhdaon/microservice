package main

import (
	"fmt"
	"sort"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func solution(n string, m string) string {
	sort.Slice(breads, func(i, j int) bool {
		if breads[i].t < breads[j].t {
			return true
		}

		if breads[i].t == breads[j].t {
			return breads[i].a < breads[j].a
		}

		return false
	})
}

func main() {
	fmt.Println(solution("12003", "00123"))
}
