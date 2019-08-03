package main

import (
	"fmt"
	"math"
	"time"
)

/* *************************************************************
 * Copyright  2018 Bridge-ruijiezhi@163.com. All rights reserved.
 *
 * FileName: main
 *
 * @Author: Bridge 2019/3/6 11:04
 *
 * @Version: 1.0
 * *************************************************************/

var aa map[int]int
var Ac []int

func main() {
	fmt.Println("go please waite a moment. ")

	n := 1

	count := 0

	pos := 0

	t := time.Now()

	for n = 1; n < 200000; n++ {
		aa = make(map[int]int)
		Ac = make([]int, 10)

		pos = Fn(n)

		if n == pos {
			count++
			fmt.Println("n is :", n, " Fn(n) is : ", pos)
			if count > 1 {
				break
			}
		}
		fmt.Println("time is : ", time.Now().Sub(t).String())
		fmt.Println(n, " stop")
		fmt.Println("test 1998 is ", Fn_test(1998))
	}

}

func Fn(n int) int {
	i := 0
	y := 0

	for x := n; x >= 1; x /= 10 {
		y = x % 10
		aa[i] = y
		Ac[i] = y
		i++
	}

	Ac = Ac[0:i]
	m := len(Ac) - 1
	h := aa[m]

	if n%gPow(m) == 0 {
		return fnA(h, m)
	} else {
		return fnA(h, m) + fnB(m-1)
	}
}

func fnA(h int, m int) int {
	if h == 0 {
		return 0
	} else if h == 1 {
		if m == 0 {
			return 1
		} else {
			return m*gPow(m-1) + 1
		}
	} else {
		if m == 0 {
			return 1
		} else {
			return gPow(m) + h*m*gPow(m-1)
		}
	}
}

func fnB(m int) int {
	return 9
}

func gPow(m int) int {
	return int(math.Pow10(m))
}

func Fn_test(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total += Count(i)
	}

	return total
}

func Count(n int) int {
	num := 0
	for t := n; t >= 1; t /= 10 {
		if t%10 == 1 {
			num++
		}
	}

	return num
}
