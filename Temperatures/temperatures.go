package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var target int64 = 0
	var result int64 = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	fmt.Fprintln(os.Stderr, "DEBUG: Number of temperatures is: ", n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		fmt.Fprintln(os.Stderr, "DEBUG: Index i =", t)
		if t == 0 {
			result = t
			break
		}
		if i == 0 {
			result = t
			fmt.Fprintln(os.Stderr, "DEBUG: Setting result to:", result)
			continue
		}
		fmt.Fprintln(os.Stderr, "DEBUG: t | target | result :", t, target, result)
		if math.Abs(float64(t)) > float64(target) && math.Abs(float64(t)) < math.Abs(float64(result)) {
			result = t
		}
		if math.Abs(float64(result)) == float64(t) && t > result {
			result = t
		}
		_ = t
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(result) // Write answer to stdout
}
