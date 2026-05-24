package main

import "fmt"

// Implement Higher Order Function (Recive function as parameter)
func procOper(a, b int, oper func(x, y int)) {
	oper(a, b)
}

// Implement Higher Order Function (Return function as result)
func getOper(oper string) func(x, y int) {
	switch oper {
	case "sum":
		return func(x, y int) {
			fmt.Println("Total:", x+y)
		}
	case "sub":
		return func(x, y int) {
			fmt.Println("Sub:", x-y)
		}
	default:
		return func(x, y int) {
			fmt.Println("Invalid operation")
		}
	}
}

// First Order Function
func sum(p, q int) {
	r := p + q
	fmt.Println(r)
}

func main() {
	procOper(10, 20, sum)
	procOper(10, 20, func(x, y int) {
		fmt.Println("Mul:", x*y)
	})

	oper := getOper("sum")
	oper(15, 25)
	oper = getOper("sub")
	oper(15, 25)
}
