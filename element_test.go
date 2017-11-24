
package element

import (
	"testing"
	"os"
	"bufio"
	"fmt"
)

func TestLoadElement(t *testing.T){
	e := GetElements()

	if(e == nil) {
		t.Error("Expected not nil elements")
	}
}


func TestAllWords(t *testing.T){
	e := GetElements()
	f, err := os.Open("./english.txt")
	if err != nil {
		t.Error("Failed to open dictionary file")
	}

	defer f.Close()
	t.Log("Opened dict")
	foundCount := 0
	notFoundCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res := e.GetElementsForWord(line)

		if len(res) > 0 {
			foundCount++
			fmt.Print(line, " - ")
			for _, k := range res {
				fmt.Print(k.Symbol, " ")
			}
			fmt.Println()
		} else {
			fmt.Println("!===================", line)
			notFoundCount++
		}
	}

	fmt.Println()
	fmt.Println("Found words: ", foundCount)
	fmt.Println("Not Found words: ", notFoundCount)
}


func TestAllElements(t *testing.T){
	e := GetElements()
	res := e.GetAllElements(false)
	if len(res) != 118 {
		t.Error("Count of elements ")
	}

	res = e.GetAllElements(true)
	if len(res) <= 118 {
		t.Error("Count of fake elements ")
	}
}