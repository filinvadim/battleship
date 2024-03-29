package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getRandom(num int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	randomNumber := random.Intn(num)
	return randomNumber
}

// Sea : selfdocumentated
type Sea struct {
	board [8][8]string
}

// Warship : selfdocumentated, composition with Sea struct
type Warship struct {
	*Sea
	usedShips []int
}

func (s *Sea) paintingSea() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			s.board[i][j] = "~"
		}
	}
	return
}

func (*Sea) randomDirection() string {
	randomNumber := getRandom(2)
	if randomNumber == 0 {
		return "row"
	}
	return "col"
}

func (s *Sea) searchingFreeSlots(direction string, lenShip int) []map[int]int {
	shipCoordinates := make([]map[int]int, 0, lenShip)
	validSlots := 0
	randomSlot := getRandom(8)
	shipStartRandom := getRandom(8 - lenShip)

	for validSlots < lenShip {
		if direction == "row" && s.board[shipStartRandom][randomSlot] == "~" {
			coord := map[int]int{shipStartRandom: randomSlot}
			shipCoordinates = append(shipCoordinates, coord)
			shipStartRandom++
			validSlots++
		} else if direction == "col" && s.board[randomSlot][shipStartRandom] == "~" {
			coord := map[int]int{randomSlot: shipStartRandom}
			shipCoordinates = append(shipCoordinates, coord)
			shipStartRandom++
			validSlots++
		} else {
			shipCoordinates = make([]map[int]int, 0, lenShip)
			randomSlot = getRandom(8)
			shipStartRandom = getRandom(8 - lenShip)
			validSlots = 0
		}
	}
	return shipCoordinates
}

func (w *Warship) unrepeatedShip() int {
	for len(w.usedShips) < 4 {
		ship := len(w.usedShips) + 1
		w.usedShips = append(w.usedShips, ship)
		return ship
	}
	return 0
}

func (w *Warship) placingShipsAndDots(coordinates []map[int]int) {
	lenShip := len(coordinates)
	dotCoord := [][]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}}
	board := &w.Sea.board

	for _, coord := range coordinates {
		for row, col := range coord {
			board[row][col] = strconv.Itoa(lenShip)
			for _, value := range dotCoord {
				if (row+value[0]) < 0 || (row+value[0]) > 7 || col+value[1] < 0 || col+value[1] > 7 {
					continue
				} else if board[row+value[0]][col+value[1]] != "~" {
					continue
				} else {
					board[row+value[0]][col+value[1]] = "."
				}
			}
		}
	}
	return
}

func main() {
	a := Sea{}
	a.paintingSea()
	b := Warship{Sea: &a}

	b.placingShipsAndDots(a.searchingFreeSlots(a.randomDirection(), b.unrepeatedShip()))
	b.placingShipsAndDots(a.searchingFreeSlots(a.randomDirection(), b.unrepeatedShip()))
	b.placingShipsAndDots(a.searchingFreeSlots(a.randomDirection(), b.unrepeatedShip()))
	b.placingShipsAndDots(a.searchingFreeSlots(a.randomDirection(), b.unrepeatedShip()))
	b.placingShipsAndDots(a.searchingFreeSlots(a.randomDirection(), b.unrepeatedShip()))

	for _, value := range a.board {
		fmt.Println(strings.Join(value[:], " "))
	}
}
