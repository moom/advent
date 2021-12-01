package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	i, p, c := 0, 0, 0

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		c, err = strconv.Atoi(s.Text())
		if c > p && p > 0 {
			i++
		}
		p = c
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(i)
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}
