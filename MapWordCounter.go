package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	
	temp:=strings.Fields(s)
	count :=make(map[string]int)
	
	for i:=0;i<len(temp);i++{
			count[temp[i]]+=1	
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
