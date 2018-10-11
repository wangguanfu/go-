package main

import "fmt"

func main1001() {
	//百钱百鸡
	//公鸡5钱1只  母鸡3钱1只 小鸡1钱3只  花费100钱买100只鸡
	//公鸡 20  母鸡 33 小鸡 100
	count := 0 //计数器
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken += 3 {
				count++
				if cock+hen+chicken == 100 && cock*5+hen*3+chicken/3 == 100 {
					fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n", cock, hen, chicken)
				}
			}
		}
	}

	fmt.Println(count)
}

func main() {
	count := 0 //计数器
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			//百鸡
			chicken := 100 - cock - hen
			//百钱
			if chicken%3 == 0 && cock*5+hen*3+chicken/3 == 100 {
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n", cock, hen, chicken)

			}
		}
	}

	fmt.Println(count)
}
