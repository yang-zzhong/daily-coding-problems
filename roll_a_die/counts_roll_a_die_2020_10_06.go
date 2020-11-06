package main

// This problem was asked by Two Sigma.
//
// Alice wants to join her school's Probability Student Club. Membership dues are computed via one of two simple probabilistic games.
//
// The first game: roll a die repeatedly. Stop rolling once you get a five followed by a six. Your number of rolls is the amount you pay, in dollars.
//
// The second game: same, except that the stopping condition is a five followed by a five.
//
// Which of the two games should Alice elect to play? Does it even matter? Write a program to simulate the two games and calculate their expected value.

import (
	"fmt"
	"math/rand"
	"time"
)

func Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func MeetTheSeqSum(seq []int) int {
	sum := 0
	l := len(seq)
	i := 0
	idx := 0
	for {
		if i == l {
			fmt.Println()
			return sum
		}
		rolls := Roll()
		fmt.Printf("%d\t", rolls)
		if idx != 0 && (idx%20 == 0) {
			fmt.Printf("\n")
		}
		idx++
		if rolls == seq[i] {
			i++
			continue
		}
		sum += rolls
		for j := 0; j < i; j++ {
			sum += seq[j]
		}
		i = 0
	}
}

func main() {
	fmt.Printf("sum of meed [5,6]: %d\n", MeetTheSeqSum([]int{5, 6}))
	fmt.Printf("sum of meed [5,5]: %d\n", MeetTheSeqSum([]int{5, 5}))
	fmt.Printf("sum of meed [5,5,5,5]: %d\n", MeetTheSeqSum([]int{5, 5, 5, 5}))
}
