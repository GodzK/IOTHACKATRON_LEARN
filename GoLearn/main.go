package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//make input
	reader := bufio.NewReader(os.Stdin)

	year, _ := reader.ReadString('\n')
	addnum, errornaja := strconv.ParseFloat(year, 64)
	if errornaja != nil {
		fmt.Println(errornaja)
	} else {
		fmt.Println(addnum + 10)
	}

}
