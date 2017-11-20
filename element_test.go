
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
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res := e.GetElementsForWord(line)
		if len(res) > 0 {
			count++
			fmt.Println()
			fmt.Print(line, " - ")
			for _, k := range res {
				fmt.Print(k.symbol, " ")
			}
		}
	}

	fmt.Println()
	fmt.Println("Found words: ", count)
}