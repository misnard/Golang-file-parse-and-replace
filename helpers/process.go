package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	file, err := os.Open(src)
	if err != nil {
		fmt.Printf("An error occured : %s", err)
		return
	}

	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("An error occured : %s", err)
		return
	}

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)

	var output string
    var lineNumber int

	for scanner.Scan() {
		lineNumber++
		t := scanner.Text()
		found, res := ProcessLine(t, old, new)
		if found {
			occ += strings.Count(t, old)
			lines = append(lines, lineNumber)
		}

		_, err = fmt.Fprintln(writer, res)
		if err != nil {
			fmt.Printf("An error occured : %s", err)
			return
		}
		output += res + "\n"
	}

	fmt.Printf("Final Data result :\n %s\n", output)

	return occ, lines, nil
}

func ProcessLine(line, old, new string) (found bool, res string) {
	found = strings.Contains(line, old)

	lineSlice := strings.Split(line, " ")
	replacedLineSlice := make([]string, 0, len(lineSlice))

	for i := range lineSlice {
		replacedLineSlice = append(replacedLineSlice, ReplaceWordWithTypography(lineSlice[i], old, new))
	}



	res = strings.Join(replacedLineSlice, " ")

	return found, res
}

//Replace word with sensitive case but we can upgrade this function by checking all letters
func ReplaceWordWithTypography(word, old, new string) (res string) {
	if word == strings.ToUpper(word) && word == strings.ToUpper(old) {
		res = strings.Replace(word, strings.ToUpper(old), strings.ToUpper(new), -1)
	} else if word == strings.ToLower(word) && word == strings.ToLower(old) {
		res = strings.Replace(word, strings.ToLower(old), strings.ToLower(new), -1)
	} else if word == strings.Title(word) && word == strings.Title(old) {
		res = strings.Replace(word, strings.Title(old), strings.Title(new), -1)
	} else {
		res = word
	}

	return res
}