package zigzagconversion

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
func convert(s string, numRows int) string {
	sArray := [][]byte{}
	row := 0

	for i := 0; i < numRows; i++ {
		sArray = append(sArray, []byte{})
	}

	rowRound := false
	for i := 0; i < len(s); i++ {
		sArray[row] = append(sArray[row], s[i])

		if row == 0 {
			rowRound = true
		}

		if row == numRows-1 {
			rowRound = false
		}

		if rowRound {
			row++
		} else {
			row--
		}
	}

	var ret string

	for i := 0; i < numRows; i++ {
		ret += string(sArray[i])
	}
	return ret
}
