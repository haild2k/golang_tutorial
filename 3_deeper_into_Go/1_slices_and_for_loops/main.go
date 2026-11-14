// CHỦ ĐỀ: SLICES VÀ VÒNG LẶP FOR
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

// Khai báo file này thuộc package main (package tạo ra chương trình chạy được).
package main

// Import package fmt để in dữ liệu ra màn hình.
import "fmt"

// main là điểm bắt đầu của chương trình.
func main() {
	// >>> []string{...} LÀ SLICE LITERAL: TẠO MỘT SLICE CHUỖI VÀ GÁN SẴN CÁC PHẦN TỬ.
	// >>> SLICE GIỐNG MẢNG NHƯNG CÓ THỂ CO GIÃN (KHÔNG CỐ ĐỊNH ĐỘ DÀI). MỌI PHẦN TỬ PHẢI CÙNG KIỂU.
	// Dấu := vừa khai báo biến vừa gán giá trị, Go tự suy ra kiểu ([]string).
	// newCard() được gọi và giá trị trả về trở thành phần tử thứ 2.
	cards := []string{"Ace of Diamonds", newCard()}

	// >>> append(slice, phần_tử) TRẢ VỀ MỘT SLICE MỚI ĐÃ THÊM PHẦN TỬ VÀO CUỐI.
	// >>> NÓ KHÔNG SỬA SLICE CŨ, NÊN PHẢI GÁN LẠI KẾT QUẢ: cards = append(...).
	cards = append(cards, "Six of Spades")

	// >>> for i, card := range cards: DUYỆT QUA TỪNG PHẦN TỬ CỦA SLICE.
	// >>> i LÀ CHỈ SỐ (BẮT ĐẦU TỪ 0), card LÀ GIÁ TRỊ PHẦN TỬ TẠI CHỈ SỐ ĐÓ.
	// >>> MỖI VÒNG LẶP, i VÀ card ĐƯỢC KHAI BÁO LẠI NÊN DÙNG := .
	for i, card := range cards {
		// In chỉ số và lá bài, cách nhau bởi khoảng trắng.
		fmt.Println(i, card)
	}
}

// newCard là hàm không nhận tham số, trả về một giá trị kiểu string.
// Cú pháp: func tên() kiểu_trả_về { ... }
func newCard() string {
	return "Five of Diamonds"
}
