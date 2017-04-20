package common

import (
	"fmt"
)

// Comparable interface with single method Compare
type Comparable interface {
	// Compare compares two values of the same type:
	// A.compare(B) returns:
	//  0 if A == B
	//  1 if A > B
	// -1 if A < B
	// error if A and B are not the same type
	Compare(Comparable) (int, error)
}

type ComparableInt int
type ComparableString string
type ComparableFloat64 float64


func (i ComparableInt) Compare(j Comparable) (int, error) {
	compJ, ok := j.(ComparableInt)
	if !ok {
		return 0, fmt.Errorf("j: %v, type: %T is not ComparableInt", j, j)
	}
	if i == compJ {
		return 0, nil
	}
	if i < compJ {
		return -1, nil
	}
	return +1, nil
}

func (i ComparableString) Compare(j Comparable) (int, error) {
	compJ, ok := j.(ComparableString)
	if !ok {
		return 0, fmt.Errorf("j: %v, type: %T is not ComparableString", j, j)
	}
	if i == compJ {
		return 0, nil
	}
	if i < compJ {
		return -1, nil
	}
	return +1, nil
}

func (i ComparableFloat64) Compare(j Comparable) (int, error) {
	compJ, ok := j.(ComparableFloat64)
	if !ok {
		return 0, fmt.Errorf("j: %v, type: %T is not ComparableFloat64", j, j)
	}
	if i == compJ {
		return 0, nil
	}
	if i < compJ {
		return -1, nil
	}
	return +1, nil
}