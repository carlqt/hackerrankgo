package hackerrankgo

import (
	"fmt"
	"strconv"
)

func jimAndJokes(inputs [][]int32) int32 {
	dates := NewDates(inputs)

	if !dates.CanJoke() {
		return 0
	}

	return dates.JokesCount()
}

func reverseSlice(s []int32) []int32 {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}

	return s
}

type Dates []Number

func NewDates(dates [][]int32) Dates {
	var result Dates

	for _, date := range dates {
		result = append(result, Number{Value: date[1], Base: date[0]})
	}

	return result
}

func (d Dates) CanJoke() bool {
	for _, date := range d {
		if !date.IsValid() {
			return false
		}
	}

	return true
}

func (d Dates) JokesCount() int32 {
	var count int32

	for i := 0; i < len(d)-1; i++ {
		next, err := d[i+1].ConvertToBase(d[i].Base)
		if err != nil {
			return 0
		}

		current, err := d[i].ConvertToBase(d[i+1].Base)
		if err != nil {
			return 0
		}

		if current == next {
			count++
		}
	}

	return count
}

type Number struct {
	Value int32
	Base  int32
}

func (n Number) IsValid() bool {
	numSlice := n.toSlice()

	for _, num := range numSlice {
		if num >= n.Base {
			return false
		}
	}

	return true
}

func (n Number) BaseValue() int32 {
	baseN := strconv.FormatInt(int64(n.Value), int(n.Base))

	// convert string to int32
	result, err := strconv.ParseInt(baseN, int(n.Base), 32)
	if err != nil {
		return 0
	}

	return int32(result)
}

func (n Number) ConvertToBase(base int32) (int32, error) {
	if !n.IsValid() {
		return 0, fmt.Errorf("Number %d is not valid for base %d", n.Value, n.Base)
	}

	baseN := strconv.FormatInt(int64(n.BaseValue()), int(base))

	// convert string to int32
	result, err := strconv.Atoi(baseN)
	if err != nil {
		return 0, err
	}

	return int32(result), nil
}

func (n Number) toSlice() []int32 {
	var result []int32

	for i := n.Value; i > 0; i = i / 10 {
		result = append(result, n.Value%10)
	}

	return reverseSlice(result)
}
