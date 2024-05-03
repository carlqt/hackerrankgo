package hackerrankgo

type Number struct {
	Value        int32
	Base         int32
	DecimalValue int32
}

func jokesCount(dates []Number) int64 {
	var count int64

	for i := 0; i < len(dates)-1; i++ {
		currentDate := dates[i]

		for j := i + 1; j <= len(dates)-1; j++ {
			nextDate := dates[j]

			if currentDate.DecimalValue == nextDate.DecimalValue {
				count++
			}
		}
	}

	return count
}

func ValueInDecimal(val int32, base int32) int32 {
	if base == 10 {
		return val
	}

	tens := val / 10 * base
	ones := val % 10

	return tens + ones
}

func isValid(val int32, base int32) bool {
	tens := val / 10
	ones := val % 10

	if tens >= base || ones >= base {
		return false
	}

	return true
}

func NewDates(dates [][]int32) []Number {
	var result []Number

	for _, date := range dates {
		val := date[1]
		base := date[0]

		if isValid(val, base) {
			number := Number{
				Value:        val,
				Base:         base,
				DecimalValue: ValueInDecimal(val, base),
			}

			result = append(result, number)
		}
	}

	return result
}

func jimAndJokes(inputs [][]int32) int64 {
	dates := NewDates(inputs)

	count := jokesCount(dates)

	return count
}
