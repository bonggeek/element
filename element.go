package element

import (
	"strings"
)

type Element struct {
	Symbol   string
	AtNumber uint
	Name     string
}

type Elements struct {
    elementList map[string]Element
    fakeElementList map[string]Element
}

func (this *Elements) getElementsForInnerWord(word string,charCount int,useFake bool) []Element{
	result := make([]Element, 0)
	s := string(word[0:charCount])
	var elementList map[string]Element
	if useFake {
		elementList = this.fakeElementList
	} else {
		elementList = this.elementList
	}

	if val, ok := elementList[s]; ok {
		newWord := string(word[charCount:])
		if len(newWord) > 0 {
			resultTemp := this.getElementsForWordWorker(newWord)
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

var elements *Elements = newElements()
func newElements() *Elements{
	var elements *Elements = new(Elements)
	elements.elementList = elementList
	elements.fakeElementList = fakeElementList
	return elements
}

func GetElements() *Elements{
	return elements
}

func (this *Elements) getElementsForWordWorker(word string) []Element {
	result := make([]Element, 0)
	if len(word) > 0 {
		//fmt.Println("Word:", word)
		result = this.getElementsForInnerWord(word, 1, false)

		// So using single char as above didn't work lets see
		// is a pair of chars will work
		if len(result) == 0 && len(word) > 1 {
			result = this.getElementsForInnerWord(word, 2, false)
		}

		// Single and pair didn't work, try with fake elements now
		if len(result) == 0 && len(word) > 0 {
			result = this.getElementsForInnerWord(word, 1, true)
		}
	}

	return result
}

func (this *Elements) GetElementsForWord(word string) []Element {
    word = strings.ToLower(word)
    return this.getElementsForWordWorker(word)
}

func (this *Elements) GetAllElements(useFakeElements bool) []Element{
	nElements := len(elementList)
	if useFakeElements {
		nElements = nElements + len(fakeElementList)
	}

	elements := make([]Element, nElements)

	i := 0
	for _, v := range elementList{
		elements[i] = v
		i++
	}

	if useFakeElements {
		for _,v := range fakeElementList {
			elements[i] = v
			i++
		}
	}

	return elements
}

