package hackerrankgo

import (
	"fmt"
	"strconv"
)

type Number struct {
	Value int32
	Base  int32
}

func (n Number) IsValid() bool {
	for i := n.Value; i > 0; i = i / 10 {
		r := n.Value % 10

		if r >= n.Base {
			return false
		}
	}

	return true
}

func (n Number) ValueInDecimal() int32 {
	if n.Base == 10 {
		return n.Value
	}

	strValue := strconv.Itoa(int(n.Value))

	result, _ := strconv.ParseInt(strValue, int(n.Base), 32)

	return int32(result)
}

func (n Number) ConvertToBase(base int32) (string, error) {
	if n.IsValid() && (base >= 2 && base <= 36) {
		valueInDecimal64 := int64(n.ValueInDecimal())
		result := strconv.FormatInt(valueInDecimal64, int(base))

		return result, nil
	}

	return "", fmt.Errorf("Number %d is not valid for base %d", n.Value, n.Base)
}

func jimAndJokes(inputs [][]int32) int32 {
	dates := NewDates(inputs)

	count := jokesCount(dates)

	return count
}

func jokesCount(dates []Number) int32 {
	var count int32
	// count = 1

	for i := 0; i < len(dates)-1; i++ {
		currentDate := dates[i]

		if !currentDate.IsValid() {
			continue
		}

		for j := i + 1; j <= len(dates)-1; j++ {
			nextDate := dates[j]

			current, err := currentDate.ConvertToBase(dates[j].Base)
			if err != nil {
				continue
			}

			if current == strconv.Itoa(int(nextDate.Value)) {
				count++
			}
		}
	}

	return count
}

func NewDates(dates [][]int32) []Number {
	var result []Number

	for _, date := range dates {
		result = append(result, Number{Value: date[1], Base: date[0]})
	}

	return result
}
