package main

import ("fmt")

func sums(nums ...int)int{
	total:=0
	for _,num:=range nums{
		total=total+num
	}
	return total
}

func main(){
	nums:=[]int{1,2,3,4,5}
	sum:=sums(1,2,3,4,5)
	fmt.Println(sum)
	sum2:=sums(nums...)
	fmt.Println(sum2)


}

