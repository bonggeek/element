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

func ReadElement(elementStr string) (Element, error) {
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
		element, err := ReadElement(line)
		if err != nil {
			fmt.Println("Failed to parse element", err)
			panic("Error!!")
		}

		elements[strings.ToLower(element.symbol)] = element
	}
}
