// CHỦ ĐỀ: KHAI BÁO KIỂU TỰ ĐỊNH NGHĨA (CUSTOM TYPE) VÀ RECEIVER
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

// Import fmt để dùng fmt.Println.
import "fmt"

// >>> type deck []string: KHAI BÁO KIỂU MỚI TÊN deck, BÊN DƯỚI LÀ MỘT SLICE CHUỖI.
// >>> deck CÓ MỌI KHẢ NĂNG CỦA []string (append, range, CHỈ SỐ...) VÀ CÓ THỂ GẮN THÊM METHOD.
type deck []string

// >>> func (d deck) print(): KHAI BÁO METHOD print CHO KIỂU deck.
// >>> (d deck) LÀ RECEIVER: MỌI BIẾN KIỂU deck ĐỀU GỌI ĐƯỢC cards.print().
// >>> BÊN TRONG HÀM, d CHÍNH LÀ BIẾN ĐÃ GỌI METHOD (GIỐNG this/self Ở NGÔN NGỮ KHÁC).
// >>> THEO QUY ƯỚC GO, TÊN RECEIVER NGẮN, 1-2 KÝ TỰ, THƯỜNG LÀ CHỮ CÁI ĐẦU CỦA KIỂU.
func (d deck) print() {
	// Duyệt từng lá bài trong d, in chỉ số và giá trị.
	for i, card := range d {
		fmt.Println(i, card)
	}
}
