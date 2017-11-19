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

type Elements struct {
    elementList map[string]Element
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

func (this *Elements) LoadPeriodicTable(fileName string){

    if this.elementList == nil {
        this.elementList = make(map[string]Element)
    }

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

		this.elementList[strings.ToLower(element.symbol)] = element
	}

}

func (this *Elements) getElementsForInnerWord(word string,charCount int) []Element{
	result := make([]Element, 0)
	s := string(word[0:charCount])
	if val, ok := this.elementList[s]; ok {
		newWord := string(word[charCount:])
		if len(newWord) > 0 {
			resultTemp := this.GetElementsForWord(newWord)
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


func (this *Elements) GetElementsForWord(word string) []Element {
	result := make([]Element, 0)
	if len(word) > 0 {
		//fmt.Println("Word:", word)
		result = this.getElementsForInnerWord(word, 1)

		// So using single char as above didn't work lets see
		// is a pair of chars will work
		if len(result) == 0 && len(word) > 1 {
			result = this.getElementsForInnerWord(word, 2)
		}

	}

	return result
}

