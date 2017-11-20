package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		_ = t
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("result") // Write answer to stdout
}
