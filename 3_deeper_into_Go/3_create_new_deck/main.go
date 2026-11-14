// CHỦ ĐỀ: TẠO BỘ BÀI MỚI (HÀM newDeck, VÒNG LẶP LỒNG NHAU)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

func main() {
	// >>> GỌI newDeck() ĐỂ TẠO BỘ BÀI 16 LÁ. KIỂU CỦA cards ĐƯỢC SUY RA LÀ deck.
	cards := newDeck()

	// Gọi method print của deck để in bộ bài.
	cards.print()
}
