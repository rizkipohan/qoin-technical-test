package main

import (
	"fmt"
	"math/rand"
)

type attribute struct {
	Score         int
	GetDiceNumber []int
	DiceToRoll    int
	// ContinuePlay  bool
}

func main() {

	var jPemain int
	var jDadu int

	fmt.Print("masukkan jumlah pemain: ")
	fmt.Scan(&jPemain)
	fmt.Print("masukkan jumlah dadu: ")
	fmt.Scan(&jDadu)
	fmt.Println("==============================")
	fmt.Println("pemain:", jPemain, ", dadu:", jDadu)

	a := 0
	isNext := true
	ddMin := 1
	ddMax := 7
	sisaPlayer := jPemain

	playersData := make(map[int]attribute)

	for isNext {

		fmt.Println("==============================")
		fmt.Println("Giliran ke", a+1, ":")

		for p := 0; p < jPemain; p++ {
			noPemain := p + 1
			score_total := playersData[noPemain].Score
			dice_to_roll := playersData[noPemain].DiceToRoll

			if a+1 == 1 {
				dice_to_roll = jDadu
			}

			var getDD = []int{}

			for d := 0; d < dice_to_roll; d++ {
				ddNumb := rand.Intn(ddMax-ddMin) + ddMin
				getDD = append(getDD, ddNumb)
			}

			playersData[noPemain] = attribute{
				Score:         score_total,
				GetDiceNumber: getDD,
				DiceToRoll:    dice_to_roll,
			}

			fmt.Println("\t Pemain #", noPemain, "(", score_total, "):", getDD)
		}

		fmt.Println("----------")

		fmt.Println("Setelah Evaluasi:")

		for p := 0; p < jPemain; p++ {
			noPemain := p + 1
			var getDD = []int{}
			for _, diceNumb := range playersData[noPemain].GetDiceNumber {
				if diceNumb == 1 {
					nextPemain := noPemain + 1
					if nextPemain <= jPemain {
						playersData[nextPemain] = attribute{
							Score:         playersData[nextPemain].Score,
							DiceToRoll:    playersData[nextPemain].DiceToRoll + 1,
							GetDiceNumber: append(playersData[nextPemain].GetDiceNumber, diceNumb),
						}
					}
				}
				if diceNumb == 6 {
					playersData[noPemain] = attribute{
						Score:         playersData[noPemain].Score + 1,
						DiceToRoll:    playersData[noPemain].DiceToRoll - 1,
						GetDiceNumber: playersData[noPemain].GetDiceNumber,
					}
				}
				if diceNumb != 6 {
					getDD = append(getDD, diceNumb)
				}
			}

			fmt.Println("\t Pemain #", noPemain, "(", playersData[noPemain].Score, "):", getDD)
		}

		for p := 0; p < jPemain; p++ {
			if playersData[p+1].DiceToRoll == 0 {
				sisaPlayer = sisaPlayer - 1
			}
		}

		if sisaPlayer == 1 {
			isNext = false
			playerNo := 0
			maxScore := 0
			for p := 0; p < jPemain; p++ {
				if playersData[p+1].DiceToRoll > 0 {
					fmt.Println("permainan berakhir hanya pemain #", p+1, " yang masih memiliki sisa dadu sebanyak:", playersData[p+1].DiceToRoll)
				}

				if playersData[p+1].Score > playersData[p+2].Score {
					maxScore = playersData[p+1].Score
					playerNo = p + 1
				} else if playersData[p+1].Score < playersData[p+2].Score {
					maxScore = playersData[p+2].Score
					playerNo = p + 2
				}
			}
			fmt.Println("pemenangnya adalah pemain #", playerNo, " dengan skor:", maxScore)
		}
		a++
	}
}
