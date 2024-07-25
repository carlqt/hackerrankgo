package jimandjokes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJimAndJokes(t *testing.T) {
	var msg string

	tests := []struct {
		input  [][]int32
		output int64
	}{
		{
			input:  [][]int32{{10, 25}, {8, 31}},
			output: 1,
		},
		{
			input:  [][]int32{{2, 25}, {2, 25}},
			output: 0,
		},
		{
			input:  [][]int32{{11, 10}, {10, 11}},
			output: 1,
		},
	}
	for _, tt := range tests {
		actual := jimAndJokes(tt.input)

		if int64(tt.output) != int64(actual) {
			msg = fmt.Sprintf("There was a problem with the input %v", tt.input)
		}

		assert.Equal(t, tt.output, actual, msg)
	}

	t.Run("from file", func(t *testing.T) {
		// write a program that reads from a file named joke-test.txt and assigns it to a variable with type [][]int32

		// read from the file
		file, err := os.Open("joke-test.txt")
		defer file.Close()

		// loop through the file and append the values to the variable
		var inputs [][]int32
		if err == nil {
			reader := bufio.NewReader(file)
			for {
				line := readLine(reader)
				if line == "" {
					break
				}

				values := strings.Split(line, " ")
				base, _ := strconv.Atoi(values[0])
				value, _ := strconv.Atoi(values[1])

				inputs = append(inputs, []int32{int32(base), int32(value)})
			}
		}

		result := jimAndJokes(inputs)
		assert.Equal(t, int64(65141656), result)
	})

	t.Run("from file", func(t *testing.T) {
		// write a program that reads from a file named joke-test.txt and assigns it to a variable with type [][]int32

		// read from the file
		file, err := os.Open("joke-test2.txt")
		defer file.Close()

		// loop through the file and append the values to the variable
		var inputs [][]int32
		if err == nil {
			reader := bufio.NewReader(file)
			for {
				line := readLine(reader)
				if line == "" {
					break
				}

				values := strings.Split(line, " ")
				base, _ := strconv.Atoi(values[0])
				value, _ := strconv.Atoi(values[1])

				inputs = append(inputs, []int32{int32(base), int32(value)})
			}
		}

		result := jimAndJokes(inputs)
		assert.Equal(t, int64(4999950000), result)
	})
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
