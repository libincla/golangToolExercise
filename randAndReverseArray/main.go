package main

import (
	"fmt"
	"math/rand"
	"time"
)

//步骤一，先生成一个长度为7的100之内的随机数组
//并对数组作首尾置换
func main() {
	var testArray [7]int

	//播种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(testArray); i++ {
		testArray[i] = rand.Intn(100)
	}
	fmt.Println(testArray)

	//首尾置换
	var tmp int = 0

	/*
		假设长度为5的数组，它第一个和最后一个做交换，它第二个和倒数第二个做交换，只需要二次
		长度为9的数组，它的第一个和最后一个做交换，第二个倒数。。。。，只需要四次
		因此，只需要len(array)/2次
	*/
	for i := 0; i < len(testArray)/2; i++ { //交换的次数
		tmp = testArray[i]
		testArray[i] = testArray[len(testArray)-1-i]
		testArray[len(testArray)-1-i] = tmp
	}
	fmt.Println(testArray)

}
