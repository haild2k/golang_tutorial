// Trong Go, package là đơn vị phạm vi (scope), không phải file:

// Mọi file .go trong cùng một thư mục và cùng tên package được compiler gộp lại thành một package duy nhất.
// Những gì khai báo ở top-level (type, func, var, const) trong một file sẽ tự động dùng được ở mọi file khác cùng package, không cần import.

package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
