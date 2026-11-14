// CHỦ ĐỀ: LƯU DỮ LIỆU XUỐNG Ổ CỨNG (GHI FILE)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

func main() {
	// Tạo bộ bài 16 lá.
	cards := newDeck()

	// >>> GHI BỘ BÀI RA FILE "my_cards" TRONG THƯ MỤC HIỆN TẠI (NƠI CHẠY LỆNH go run).
	// >>> GIÁ TRỊ error TRẢ VỀ ĐANG BỊ BỎ QUA; BÀI SAU (7_error_handle) SẼ HỌC CÁCH XỬ LÝ LỖI.
	cards.saveToFile("my_cards")
}
