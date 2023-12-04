package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFileAndSeperate(filePath string) []string{
	res, err:= os.ReadFile(filePath)
	if(err!=nil) {
		panic("can not able to read file")
	}
	input:= string(res[:])
	inputArr := strings.Split(input, "\n")[:]
	return inputArr
}

var strNumbers []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	}

func day1(){
	arr:=readFileAndSeperate("./test.txt")
	arr = arr[:len(arr)-1]
	/* var numsArr []string */
	sum:=0
	for _, ele := range arr {
		newEle:=strings.Clone(ele)
		for _, i:= range strNumbers{
			abc:= strings.Contains(newEle, i)
			if (abc){
				index:=fmt.Sprint(slices.Index(strNumbers, i)+1) 
				fmt.Println(index)
				newStr:=strings.ReplaceAll(ele, i, index)
				newEle=newStr
			}
		}
		nums := ""
		for _, i := range newEle{
			inti, err := strconv.Atoi(string(i))
			if err != nil {
				continue	
			}
			nums = fmt.Sprint(nums,inti)
		}
		first, _:=strconv.Atoi(string(nums[0]))
		last, _:=strconv.Atoi(string(nums[len(nums)-1]))
		final:=fmt.Sprintf("%v%v",first,last)
		f, _:=strconv.Atoi(final)
		sum+=f
	}
	fmt.Println(sum)
}

func main(){
	day1()
}
