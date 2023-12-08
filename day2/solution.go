package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/churmd/adventofcode2023/common"
	"github.com/churmd/higherorder"
)

func Solution() {
	games := parseInput(input)
	isGameValidPart1 := func(game Game) bool {
		return isGameValid(game, 12, 13, 14)
	}
	validgames := higherorder.Filter(isGameValidPart1, games)
	validGameIds := higherorder.Map(func(game Game) int { return game.ID }, validgames)
	sumOfgameIds := higherorder.Foldl(func(x, y int) int { return x + y }, 0, validGameIds)

	fmt.Println("Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes.")
	fmt.Println("What is the sum of the IDs of those games?")
	fmt.Printf("%d\n", sumOfgameIds)
}

type Round struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

type Game struct {
	ID     int
	Rounds []Round
}

func parseInput(input string) []Game {
	lines := common.SplitNewLines(input)
	games := make([]Game, 0)
	for _, line := range lines {
		game := parseLine(line)
		games = append(games, game)
	}

	return games
}

func parseLine(line string) Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	gameSplit := strings.Split(line, ":")
	gameSection := gameSplit[0]
	gameIdSplit := strings.Split(gameSection, " ")
	gameId, err := strconv.Atoi(gameIdSplit[1])
	if err != nil {
		panic("game id " + gameIdSplit[1] + " is not a number: " + err.Error())
	}

	game := Game{
		ID: gameId,
	}

	roundsSection := gameSplit[1]
	rounds := strings.Split(roundsSection, ";")
	for _, round := range rounds {
		currentRound := Round{}
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cubeDetail := strings.Split(cube, " ")
			count, err := strconv.Atoi(cubeDetail[0])
			if err != nil {
				panic("cube count " + cubeDetail[0] + " is not a number: " + err.Error())
			}

			switch cubeDetail[1] {
			case "red":
				currentRound.Red = count
			case "green":
				currentRound.Green = count
			case "blue":
				currentRound.Blue = count
			default:
				panic(cubeDetail[1] + " is not red, green or blue")
			}
		}
		game.Rounds = append(game.Rounds, currentRound)
	}

	return game
}

func isGameValid(game Game, red, green, blue int) bool {
	for _, round := range game.Rounds {
		if round.Red > red {
			return false
		}
		if round.Green > green {
			return false
		}
		if round.Blue > blue {
			return false
		}
	}

	return true
}
