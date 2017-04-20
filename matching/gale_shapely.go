// Implementation of Galey-Shapely algorithm
// https://en.wikipedia.org/wiki/Stable_marriage_problem
package matching

import (
	"errors"
	"github.com/tornyak/goalg/util/slice"
	"fmt"
)

var (
	ErrDifferentSize  = errors.New("Woman and men set sizes differ")
	ErrNoFreeMan = errors.New("No free man found")
	ErrSliceEmpty = errors.New("Slice is empty")
)

// Engagement status
const (
	Free = 0
	Engaged = 1
)

// Matches
const Unknown = -1

type GaleShapely struct {
	menWishes[][] int
	womenWishes[][] int
	menStatus[] int
	womenStatus[] int
	matches []int // index is a woman and value is a man

}

func InitGaleShapely(menWishes, womenWishes [][]int) (*GaleShapely, error) {
	size := len(menWishes)
	if size == 0 {
		return nil, ErrSliceEmpty
	}

	if size != len(womenWishes) {
		return nil, ErrDifferentSize
	}
	return &GaleShapely{
		menWishes: menWishes,
		womenWishes: womenWishes,
		menStatus: makeAndInitEngagementStatus(size),
		womenStatus: makeAndInitEngagementStatus(size),
		matches: makeAndInitMatches(size),
	}, nil
}

func (gs *GaleShapely) Match() []int  {

	var man, woman int
	var err error

	for  {
		if man, err = getFreeMan(gs.menStatus); err != nil {
			return gs.matches
		}

		if woman, err = popFirstWoman(gs.menWishes[man]); err != nil {
			continue
		}

		if  gs.isWomenFree(woman) {
			gs.engage(man, woman)
		} else {
			gs.engageIfBetterThanFinace(man, woman)
		}
	}
}

func (gs *GaleShapely) engageIfBetterThanFinace(man, woman int) {
	fiance := gs.matches[woman]
	if gs.isBetterThanFiance(man, woman, fiance) {
		gs.breakEngagement(fiance, woman)
		gs.engage(man, woman)
	}
}

// The one found first is higher prio
func (gs *GaleShapely) isBetterThanFiance(man, woman, fiance int) bool {
	for _, choice := range gs.womenWishes[woman] {
		switch choice {
		case man:
			return true
		case fiance:
			return false
		}
	}
	return false
}

func (gs *GaleShapely) engage(man, woman int) {
	fmt.Printf("engage man: %d woman: %d\n", man, woman)
	gs.womenStatus[woman] = Engaged
	gs.menStatus[man] = Engaged
	gs.matches[woman] = man
}

func (gs *GaleShapely) breakEngagement(man, woman int) {
	fmt.Printf("breakEngagement man: %d woman: %d\n", man, woman)
	gs.womenStatus[woman] = Free
	gs.menStatus[man] = Free
	gs.matches[woman] = Unknown
}


func makeAndInitEngagementStatus(size int)[]int {
	status := make([]int, size)
	for i:= 0; i < size; i++ {
		status[i] = Free
	}
	return status
}

func makeAndInitMatches(size int)  []int {
	matches := make([]int, size)
	for i:= 0; i < size; i++ {
		matches[i] = Unknown
	}
	return matches
}

func getFreeMan(men []int) (int, error) {
	for i := range men {
		if men[i] == Free {
			return i, nil
		}
	}
	return 0, ErrNoFreeMan
}

func popFirstWoman(men []int) (int, error) {
	if len(men) > 0 {
		w := men[0]
		slice.DeleteInt(men, 0)
		return w, nil
	}
	return 0, ErrSliceEmpty
}

func (gs *GaleShapely) isWomenFree(woman int) bool {
	return gs.womenStatus[woman] == Free
}