package main

import "fmt"

func makeOddGenerator() func() uint {
	odd_num := uint(1)
	return func() (num uint) {
		num = odd_num
		odd_num += 2
		return
	}
}
func main() {
	bam := makeOddGenerator()
	for i := 1; i <= 5; i++ {
		fmt.Println(bam())
	}
}
