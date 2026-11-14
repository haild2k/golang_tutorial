// CHỦ ĐỀ: ĐỌC/GHI FILE (FILE I/O) VÀ TEST CHO FILE I/O
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> PHẦN CHÍNH CỦA BÀI NẰM Ở saveToFile / newDeckFromFile (deck.go) VÀ TEST TRONG deck_test.go.

package main

func main() {
	// Tạo bộ bài 16 lá.
	cards := newDeck()

	// Xáo trộn bộ bài.
	cards.shuffle()

	// In bộ bài ra màn hình.
	cards.print()
}
