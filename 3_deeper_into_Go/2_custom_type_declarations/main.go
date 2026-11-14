// CHỦ ĐỀ: KHAI BÁO KIỂU TỰ ĐỊNH NGHĨA (CUSTOM TYPE) VÀ RECEIVER
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// Trong Go, package là đơn vị phạm vi (scope), không phải file:
//   - Mọi file .go trong cùng một thư mục và cùng tên package được compiler gộp lại thành một package duy nhất.
//   - Những gì khai báo ở top-level (type, func, var, const) trong một file sẽ tự động dùng được
//     ở mọi file khác cùng package, không cần import.
//     Vì vậy main.go dùng được kiểu deck và method print() khai báo trong deck.go.
//   - Chạy bằng: go run main.go deck.go  (hoặc go run .)

package main

func main() {
	// >>> deck{...} LÀ LITERAL CỦA KIỂU TỰ ĐỊNH NGHĨA deck, CÚ PHÁP GIỐNG HỆT []string{...}.
	// >>> BIẾN cards CÓ KIỂU deck (KHÔNG PHẢI []string), NÊN DÙNG ĐƯỢC CÁC METHOD CỦA deck.
	cards := deck{"Ace of Diamonds", newCard()}

	// >>> append VẪN DÙNG ĐƯỢC VỚI deck VÌ KIỂU GỐC CỦA NÓ LÀ SLICE. KẾT QUẢ VẪN LÀ deck.
	cards = append(cards, "Six of Spades")

	// >>> GỌI METHOD QUA RECEIVER: cards.print() — cards ĐƯỢC TRUYỀN VÀO VỊ TRÍ d CỦA (d deck).
	cards.print()
}

// newCard trả về một lá bài dạng string.
func newCard() string {
	return "Five of Diamonds"
}
