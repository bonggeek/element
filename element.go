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
}

func (this *Elements) getElementsForInnerWord(word string,charCount int) []Element{
	result := make([]Element, 0)
	s := string(word[0:charCount])
	if val, ok := this.elementList[s]; ok {
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
	return elements
}

func GetElements() *Elements{
	return elements
}

func (this *Elements) getElementsForWordWorker(word string) []Element {
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

func (this *Elements) GetElementsForWord(word string) []Element {
    word = strings.ToLower(word)
    return this.getElementsForWordWorker(word)
}

func (this *Elements) GetAllElements() []Element{
	elements := make([]Element, len(elementList))
	i := 0
	for _, v := range elementList{
		elements[i] = v
		i++
	}

	return elements
}

