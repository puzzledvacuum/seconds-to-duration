package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/puzzledvacuum/seconds-to-duration/duration"
)

func main() {
	if len(os.Args) > 1 {
		secondString := os.Args[1]
		secondInt, err := strconv.Atoi(secondString)
		if err != nil {
			// Output: strconv went wrong: conversion went wrong
			fmt.Printf("ERROR: %q", err)
			return
		}
		fmt.Println(duration.Convert(secondInt))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			secondInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				// Output: strconv went wrong: conversion went wrong
				fmt.Printf("ERROR: %q", err)
				return
			}
			fmt.Println(duration.Convert(secondInt))
		}
		if scanner.Err() != nil {
			// Output: scanner went wrong: First non EOF Error returned
			fmt.Printf("ERROR: %q", scanner.Err())
			return
		}
	}
}
