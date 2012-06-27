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

func (self *NRun) HopelessToWinByN(desired_wins int, desired_length int) bool {
	return (self.Wins() + (desired_length - self.Length())) < desired_wins
}

func (self *NRun) NextRound(is_win bool) NRun {
	var nextrun = make([]bool, self.Length() + 1)
	copy(nextrun, self.Runs)
	nextrun[self.Length()] = is_win
	return NRun{nextrun}
}

func (self *NRun) Winner() bool {
	return self.Wins() > (self.Length() / 2)
}

func main() {
	desired_length := 15
	desired_wins := desired_length / 2
	win_probability := float64(0.)
	narray := []NRun{NRun{[]bool{true}}, NRun{[]bool{false}}}
	for {
		if len(narray) == 0 {
			break
		}
		item := narray[0]
		narray = narray[1:]

		if item.Length() == desired_length && item.Wins() > desired_wins {
			win_probability += item.Probability()
			//fmt.Printf("Here's a winning item: %v -> %v\n", item, item.Probability())
		} else if item.Length() < desired_length {
			if item.HopelessToWinByN(desired_wins, desired_length) {
				//fmt.Printf("Here's a hopeless item: %v\n", item)
			} else {
				narray = append(narray, item.NextRound(true), item.NextRound(false))
			}
		} else {
			//fmt.Printf("Here's a useless item: %v\n", item)
		}
	}
	fmt.Printf("Et voila. %v chance of winning.\n", win_probability * 100.)
}
