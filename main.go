package main

import (
	"bufio"
	"fmt"
	"os"
)

func findFirstDigit(line string) (int, error) {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			return int(char - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found in line: %s", line)
}
func reverseString(line string) string {
	runes := []rune(line)
	for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findLastDigit(line string) (int, error) {

	//revLine := reverseString(line)
	/*	for _, char := range revLine {
		fmt.Print(string(char))
		if char >= '0' && char <= '9' {
			return int(char - '0'), nil
		}
	}*/
	for i := len(line) - 1; i >= 0; i-- {
		fmt.Println(line[i:len(line)])
		if line[i] >= '0' && line[i] <= '9' {
			return int(line[i] - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found in line: %s", line)
}

func process(gioename string) (int, error) {

	answer := 0
	file, err := os.Open(gioename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rev, err := findLastDigit(line)
		if err != nil {
			fmt.Println(err)
		}
		result, err := findFirstDigit(line)
		if err != nil {
			fmt.Println(err)
		}

		lineAns := result * 10
		fmt.Printf("\nfirst: %d Second: %d\n", lineAns, rev)
		answer += result*10 + rev
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return answer, nil
}

func main() {

	filename := "aocp1"
	fmt.Println(process(filename))
}
