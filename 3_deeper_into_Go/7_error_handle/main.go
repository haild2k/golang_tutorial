// CHỦ ĐỀ: XỬ LÝ LỖI (ERROR HANDLING)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

func main() {
	// >>> CỐ Ý ĐỌC FILE "my" KHÔNG TỒN TẠI ĐỂ THỬ NHÁNH LỖI:
	// >>> CHƯƠNG TRÌNH IN "Error: open my: no such file or directory" RỒI THOÁT VỚI MÃ 1.
	// >>> ĐỔI THÀNH "my_cards" (FILE CÓ SẴN TRONG THƯ MỤC) ĐỂ THẤY TRƯỜNG HỢP ĐỌC THÀNH CÔNG.
	cards := newDeckFromFile("my")

	// In bộ bài đọc được (chỉ chạy tới đây nếu đọc file thành công).
	cards.print()
}
