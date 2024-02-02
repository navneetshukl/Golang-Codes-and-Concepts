package main

import "fmt"

var result []int

func subset(i,n int,arr []int){
	if i>=n{
		for _,value:=range result{
			fmt.Print(value," ")
		}
		fmt.Println()
		return
	}

	result = append(result, arr[i])
	subset(i+1,n,arr)
	result=result[0:len(result)-1]
	subset(i+1,n,arr)

}
func main(){
	arr:=[]int{1,2,3}
	n:=len(arr)

	subset(0,n,arr)
}