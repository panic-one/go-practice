package main

import (
	"fmt"
	"strconv"
)

type GameScore struct {
	Name   string
	Score  int
	Number int
}

func main() {
	p1, err := NewGameScore("yukiru", 0, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2, err := NewGameScore("yukiru", 100, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	p3, err := NewGameScore("yukiru", 60, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	p4, err := NewGameScore("yukiru", 80, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	p5, err := NewGameScore("yukiru", 50, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p1, p2, p3, p4, p5)
	scoreList := []GameScore{*p1, *p2, *p3, *p4, *p5}
	syukei(scoreList)
}

func NewGameScore(name string, score int, number int) (*GameScore, error) {
	if score >= 0 && score <= 100 {
		return &GameScore{
			Name:   name,
			Score:  score,
			Number: number,
		}, nil
	} else {
		return nil, fmt.Errorf("Invalid score: %s", strconv.Itoa(score))
	}
}

func syukei(scoreList []GameScore) int{
	sum := 0
	for _, score := range scoreList{
		sum += score.Score
	}
	return sum
}
