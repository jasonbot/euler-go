/*
Problem 121

A bag contains one red disc and one blue disc. In a game of chance a player
takes a disc at random and its colour is noted. After each turn the disc is
returned to the bag, an extra red disc is added, and another disc is taken at
random.

The player pays £1 to play and wins if they have taken more blue discs than red
discs at the end of the game.

If the game is played for four turns, the probability of a player winning is
exactly 11/120, and so the maximum prize fund the banker should allocate for
winning in this game would be £10 before they would expect to incur a loss.
Note that any payout will be a whole number of pounds and also includes the
original £1 paid to play the game, so in the example given the player actually
wins £9.

Find the maximum prize fund that should be allocated to a single game in which
fifteen turns are played.
*/

package main

import (
	"fmt"
)

func ProbabilityAtN(iteration int) float64 {
	return float64(1.) / (float64(iteration) + float64(1.))
}

type NRun struct {
	Runs []bool
}

func (self *NRun) Length() int {
	return int(len(self.Runs))
}

func (self *NRun) Wins() int {
	wins := int(0)
	for index := int(0); index < len(self.Runs); index++ {
		if self.Runs[index] == true {
			wins += 1
		}
	}
	return wins
}

func (self *NRun) Probability() float64 {
	probability := float64(ProbabilityAtN(1))
	if self.Runs[0] == false {
		probability = float64(1.) - probability
	}
	for index := int(1); index < len(self.Runs); index++ {
		if self.Runs[index] == true {
			probability *= ProbabilityAtN(index + 1)
		} else {
			probability *= (float64(1.) - ProbabilityAtN(index+1))
		}
	}
	return probability
}

func (self *NRun) NextRound(is_win bool) NRun {
	return NRun{append(self.Runs, is_win)}
}

func main() {
	desired_length := 6
	narray := []NRun{NRun{[]bool{true}}, NRun{[]bool{false}}}
	fmt.Printf("This is %v\n", narray)
	for {
		if len(narray) == 0 {
			break
		}
		item := narray[0]
		narray = narray[1:]
		if item.Length() == desired_length && item.Wins() >= (desired_length/2) {
			fmt.Printf("Here's a winning item: %v -> %v\n", item, item.Probability())
		} else {
			fmt.Printf("Here's a losing item: %v -> %v\n", item, item.Probability())
		}

		if item.Length() < desired_length {
			narray = append(narray, item.NextRound(true), item.NextRound(false))
		}
	}
	fmt.Printf("Et voila\n")
}
