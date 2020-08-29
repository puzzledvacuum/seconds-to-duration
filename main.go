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
			panic(err) // Output: strconv went wrong: conversion went wrong
		}
		fmt.Println(duration.Convert(secondInt))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			secondInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err) // Output: strconv went wrong: conversion went wrong
			}
			fmt.Println(duration.Convert(secondInt))
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		// err := errors.New("No input supplied")
		// panic(err)
	}
}
