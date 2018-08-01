package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	err := mainInner()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(-1)
	}
}

func mainInner() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("require 1 arg")
	}
	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		return fmt.Errorf("compiling filter as regex: %v", err)
	}
	var latch *string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if latch == nil {
			groups := re.FindStringSubmatch(line)
			if groups != nil {
				if len(groups) != 2 {
					return fmt.Errorf("filter must have exactly one capture group but had %v", len(groups)-1)
				}
				latch = &groups[1]
				fmt.Printf("latched onto '%v'\n", *latch)
			}
		}
		if latch != nil {
			if strings.Contains(line, *latch) {
				fmt.Println(scanner.Text())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
