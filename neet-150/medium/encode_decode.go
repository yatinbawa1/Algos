package medium

import (
	"strconv"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	strToBeSent := ""
	for _, val := range strs {
		l := strconv.Itoa(len(val))
		strToBeSent += l
		strToBeSent += val
	}

	return strToBeSent
}

func NumberFinder(encoded string, index int, digits *int) (int, error) {
	*digits += 1

	// Check if current Index is Number
	nextLen, err := strconv.Atoi(string(encoded[index]))
	if err != nil {
		return -1, err
	}

	if nextLen > 0 {
		_, err = strconv.Atoi(string(encoded[index+1]))
		if err != nil {
			// If it is not, return currentNumber
			return nextLen, nil
		}
		lastPos, err := NumberFinder(encoded, index+1, digits)
		if err != nil {
			return 0, err
		}

		return (nextLen*10 + lastPos), nil
	}

	return 0, nil
}

func (s *Solution) Decode(encoded string) []string {
	str := []string{}

	for i := 0; i < len(encoded); i++ {
		var digits int
		nextLen, err := NumberFinder(encoded, i, &digits)
		if err == nil {
			if nextLen > 0 {
				str = append(str, encoded[i+digits:i+digits+nextLen])
				i = i + nextLen
			} else {
				str = append(str, "")
			}
		}
	}

	return str
}
