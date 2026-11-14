// CHỦ ĐỀ: XÁO TRỘN / HOÁN ĐỔI PHẦN TỬ TRONG SLICE VÀ KIỂM TRA (ASSERT) BẰNG TEST
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

func main() {
	// Tạo bộ bài 16 lá theo thứ tự cố định.
	cards := newDeck()

	// >>> XÁO TRỘN BỘ BÀI: shuffle SỬA TRỰC TIẾP cards, KHÔNG CẦN GÁN LẠI cards = ...
	cards.shuffle()

	// In bộ bài sau khi xáo; mỗi lần chạy sẽ ra thứ tự khác.
	cards.print()
}
