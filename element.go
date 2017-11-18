package element

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

type Element struct {
	symbol   string
	atNumber uint
	name     string
}

func readElement(elementStr string) (Element, error) {
	splits := strings.Split(elementStr, ",")
	atNum, err := strconv.ParseUint(splits[0], 10, 32)
	if err != nil {
		return Element{}, err
	}

	el := Element{atNumber: uint(atNum), symbol: splits[1], name: splits[2]}

	return el, nil
}

func LoadPeriodicTable(fileName string, elements map[string]Element){
	f, err := os.Open(fileName)
	if err != nil {
		panic("failed to open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		element, err := readElement(line)
		if err != nil {
			fmt.Println("Failed to parse element", err)
			panic("Error!!")
		}

		elements[strings.ToLower(element.symbol)] = element
	}
}

func getElementsForInnerWord(word string, elements map[string]Element, charCount int) []Element{
	result := make([]Element, 0)
	s := string(word[0:charCount])
	if val, ok := elements[s]; ok {
		newWord := string(word[charCount:])
		if len(newWord) > 0 {
			resultTemp := GetElementsForWord(newWord, elements)
			if len(resultTemp) > 0 {
				result = append(result, val)
				result = append(result, resultTemp...)
			}
		} else {
			result = append (result, val)
		}
	}

	return result
}


func GetElementsForWord(word string, elements map[string]Element) []Element {
	result := make([]Element, 0)
	if len(word) > 0 {
		//fmt.Println("Word:", word)
		result = getElementsForInnerWord(word, elements, 1)

		// So using single char as above didn't work lets see
		// is a pair of chars will work
		if len(result) == 0 {
			result = getElementsForInnerWord(word, elements, 2)
		}

	}

	return result
}

