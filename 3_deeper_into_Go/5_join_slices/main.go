package main

import "fmt"

// main là điểm bắt đầu của chương trình.
func main() {
	// Tạo một bộ bài mới (16 lá) bằng hàm newDeck trong deck.go
	cards := newDeck()

	// Chuyển bộ bài thành một chuỗi, các lá cách nhau bằng dấu phẩy, rồi in ra
	fmt.Println(cards.toString())
}
