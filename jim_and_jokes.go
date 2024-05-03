package hackerrankgo

type Number struct {
	Value int32
	Base  int32
}

func (n Number) IsValid() bool {
	tens := n.Value / 10
	ones := n.Value % 10

	if tens >= n.Base || ones >= n.Base {
		return false
	}

	return true
}

func (n Number) ValueInDecimal() int32 {
	if n.Base == 10 {
		return n.Value
	}

	tens := n.Value / 10 * n.Base
	ones := n.Value % 10

	return tens + ones
}

func jimAndJokes(inputs [][]int32) int32 {
	dates := NewDates(inputs)

	count := jokesCount(dates)

	return count
}

func jokesCount(dates []Number) int32 {
	var count int32

	for i := 0; i < len(dates)-1; i++ {
		currentDate := dates[i]
		currentDateInDecimal := currentDate.ValueInDecimal()

		for j := i + 1; j <= len(dates)-1; j++ {
			nextDate := dates[j]

			if currentDateInDecimal == nextDate.ValueInDecimal() {
				count++
			}
		}
	}

	return count
}

func NewDates(dates [][]int32) []Number {
	var result []Number

	for _, date := range dates {
		number := Number{Value: date[1], Base: date[0]}

		if number.IsValid() {
			result = append(result, Number{Value: date[1], Base: date[0]})
		}
	}

	return result
}
