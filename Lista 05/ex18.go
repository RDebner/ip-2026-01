package main

import "fmt"

var vt = [10]int{}
var num int

func main() {
	for i:=0; i < 10; i++ {
		fmt.Scan(&num)
		if len(vt) == 0 {
			vt[0] = num
		} else {
			for j:=0; j < len(vt); j++ {
				menor := false
				if num < vt[j] {
					menor = true
					for k:=len(vt)-1; k > j; k--{
						vt[k] = vt[k-1]
					}
					vt[j] = num
					break
				} 
				if menor == false {
					vt[i] = num
				}
			}
			
		}
	}

	fmt.Print(vt)
}