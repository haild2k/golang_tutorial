// CHỦ ĐỀ: HÀM TRẢ VỀ NHIỀU GIÁ TRỊ (MULTIPLE RETURN VALUES)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

func main() {
	// Tạo bộ bài 16 lá.
	cards := newDeck()

	// >>> NHẬN NHIỀU GIÁ TRỊ TRẢ VỀ: hand, remainingCards := deal(...)
	// >>> GIÁ TRỊ THỨ NHẤT GÁN CHO hand (5 LÁ ĐẦU), GIÁ TRỊ THỨ HAI GÁN CHO remainingCards (11 LÁ CÒN LẠI).
	// >>> SỐ BIẾN BÊN TRÁI PHẢI BẰNG SỐ GIÁ TRỊ HÀM TRẢ VỀ; KHÔNG CẦN GIÁ TRỊ NÀO THÌ DÙNG _ .
	hand, remainingCards := deal(cards, 5)

	// In bài trên tay, một dòng phân cách, rồi in phần bài còn lại.
	hand.print()
	fmt.Println("======================")
	remainingCards.print()
}
