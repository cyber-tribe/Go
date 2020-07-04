package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var exp string
	// 空白文字を許すためにこの手順が必要
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	exp = scanner.Text()
	exp = strings.TrimSpace(exp)

	fmt.Println("exp:", exp)
	ans := rpn(&exp)
	fmt.Println(ans)
}

func pop(slice []float64) (float64, []float64) {
	// []float64にpop関数を定義
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	fmt.Println("pop: ", ans)
	return ans, slice
}
func rpn(exp *string) float64 {
	stack := make([]float64, 1)
	ans := 0.0
	for _, token := range strings.Split(*exp, " ") {
		fmt.Println("stack:", stack)
		f, err := strconv.ParseFloat(token, 64)
		if err == nil { //tokenが数値のとき
			fmt.Println("push:", token)
			stack = append(stack, f)
		} else {
			switch token { // 演算子を独自に定義できる. 例えばsinやcosなど
			case "+":
				stack = apply2(stack, "+")
			case "-":
				stack = apply2(stack, "-")
			case "*":
				stack = apply2(stack, "*")
			case "/":
				stack = apply2(stack, "/")
			default:
				fmt.Println("Error: Unknown operator token->" + token)
				os.Exit(1)
			}
		}
	}
	ans, stack = pop(stack)
	return ans
}
func apply2(stack []float64, operator string) []float64 {
	x, y := 0.0, 0.0
	y, stack = pop(stack)
	x, stack = pop(stack)
	fmt.Println("operator:", operator)
	switch operator {
	case "+":
		stack = append(stack, (x + y))
	case "-":
		stack = append(stack, (x - y))
	case "*":
		stack = append(stack, (x * y))
	case "/":
		if y == 0 {
			fmt.Println("error: zero divide")
			os.Exit(1)
		} else {
			stack = append(stack, x/y)
		}
	}
	return stack
}
