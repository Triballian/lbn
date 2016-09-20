package main

type FlashCard struct {
	Simplified string
	English string
	Dictionary *dictionary.Dictionary
}

type FlashCards struct {
	Name string
	CardOrder string
	ShowHalf string
	Cards []*FlashCards
}
func main() {

}
