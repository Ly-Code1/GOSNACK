package main

import (
	"fmt"
	"strings"
)

type gameTennis struct {
	p1Point   int
	p2Point   int
	scoreCall string
	p1Name    string
	p2Name    string
}

func (game *gameTennis) setPoint1(player string) {

	if !strings.HasPrefix(game.scoreCall, "Win") {
		if player == game.p1Name {
			game.p1Point++
		} else {
			game.p2Point++
		}
		game.CurrentScore()
	} else {
		fmt.Println("*****Game already end*****")
	}
	printPoint(game.p1Point, game.p2Point, game.scoreCall)
}

func (game *gameTennis) setPoint2(point1, point2 int) {

	game.p1Point = point1
	game.p2Point = point2
	game.CurrentScore()
	printPoint(game.p1Point, game.p2Point, game.scoreCall)
}

func (game *gameTennis) CurrentScore() {

	score := ""

	if game.p1Point < 4 && game.p2Point < 4 {
		scoreStr := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
		if game.p1Point == game.p2Point {
			score = scoreStr[game.p1Point] + "-All"
		} else {
			score = scoreStr[game.p1Point] + "-" + scoreStr[game.p2Point]
		}
	}

	if game.p1Point == game.p2Point && game.p1Point >= 3 {
		score = "Deuce"
	}

	if game.p1Point > game.p2Point && game.p2Point >= 3 {
		score = "Advantage " + game.p1Name
	}

	if game.p2Point > game.p1Point && game.p1Point >= 3 {
		score = "Advantage " + game.p2Name
	}

	if game.p1Point >= 4 && game.p2Point >= 0 && (game.p1Point-game.p2Point) >= 2 {
		score = "Win for " + game.p1Name
	}

	if game.p2Point >= 4 && game.p1Point >= 0 && (game.p2Point-game.p1Point) >= 2 {
		score = "Win for " + game.p2Name
	}

	game.scoreCall = score
}

func printGame(gNo int, nam1, nam2 string) {
	fmt.Println("=====================================================================")
	fmt.Printf("TennisGame%-2d: %-15s vs %-15s\n", gNo, nam1, nam2)
	fmt.Println("=====================================================================")
}

func printPoint(p1, p2 int, score string) {
	fmt.Printf("Point       : %8d        vs %8d       | ScoreCall: %s\n", p1, p2, score)
}

func main() {
	//use setPoint1(playerName) add 1 score by player name
	game1 := gameTennis{p1Name: "Roger Federer", p2Name: "Rafael Nadal"}
	printGame(1, game1.p1Name, game1.p2Name)
	game1.setPoint1(game1.p1Name)
	game1.setPoint1(game1.p2Name)
	game1.setPoint1(game1.p1Name)
	game1.setPoint1(game1.p1Name)
	game1.setPoint1(game1.p1Name)
	game1.setPoint1(game1.p1Name)

	/*
		//use setPoint2(pointPlayer1,pointPlayer2)  set score following input
		game2 := gameTennis{p1Name: "Serena Williams", p2Name: "Naomi Osaka"}
		printGame(2, game2.p1Name, game2.p2Name)
		game2.setPoint2(0, 0)
		game2.setPoint2(1, 1)
		game2.setPoint2(2, 2)
		game2.setPoint2(3, 3)
		game2.setPoint2(4, 4)

		game2.setPoint2(1, 0)
		game2.setPoint2(0, 1)
		game2.setPoint2(2, 0)
		game2.setPoint2(0, 2)
		game2.setPoint2(3, 0)
		game2.setPoint2(0, 3)
		game2.setPoint2(4, 0)
		game2.setPoint2(0, 4)

		game2.setPoint2(2, 1)
		game2.setPoint2(1, 2)
		game2.setPoint2(3, 1)
		game2.setPoint2(1, 3)
		game2.setPoint2(4, 1)
		game2.setPoint2(1, 4)

		game2.setPoint2(3, 2)
		game2.setPoint2(2, 3)
		game2.setPoint2(4, 2)
		game2.setPoint2(2, 4)

		game2.setPoint2(4, 3)
		game2.setPoint2(3, 4)
		game2.setPoint2(5, 4)
		game2.setPoint2(4, 5)
		game2.setPoint2(15, 14)
		game2.setPoint2(14, 15)

		game2.setPoint2(6, 4)
		game2.setPoint2(4, 6)
		game2.setPoint2(16, 14)
		game2.setPoint2(14, 16)
	*/

}
