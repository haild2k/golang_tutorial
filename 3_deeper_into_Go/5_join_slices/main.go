// CHỦ ĐỀ: NỐI CÁC PHẦN TỬ SLICE THÀNH MỘT CHUỖI (strings.Join)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// main là điểm bắt đầu của chương trình.
func main() {
	// Tạo một bộ bài mới (16 lá) bằng hàm newDeck trong deck.go
	cards := newDeck()

	// >>> cards.toString() GỘP 16 LÁ BÀI THÀNH MỘT CHUỖI, CÁC LÁ CÁCH NHAU BẰNG DẤU PHẨY, RỒI IN RA.
	fmt.Println(cards.toString())
}
