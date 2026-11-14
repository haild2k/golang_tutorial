// CHỦ ĐỀ: TẠO BỘ BÀI MỚI (HÀM newDeck, VÒNG LẶP LỒNG NHAU)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// deck là kiểu tự định nghĩa, bên dưới là slice chuỗi.
type deck []string

// >>> func newDeck() deck: HÀM KHÔNG NHẬN THAM SỐ, TRẢ VỀ GIÁ TRỊ KIỂU deck.
// >>> ĐÂY LÀ HÀM THƯỜNG (KHÔNG CÓ RECEIVER) VÌ LÚC GỌI CHƯA CÓ BỘ BÀI NÀO ĐỂ GẮN VÀO.
func newDeck() deck {
	// >>> deck{} TẠO MỘT BỘ BÀI RỖNG (0 PHẦN TỬ) ĐỂ LÁT NỮA append VÀO.
	cards := deck{}

	// >>> HAI SLICE NGUỒN: DANH SÁCH CHẤT VÀ DANH SÁCH GIÁ TRỊ.
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// >>> VÒNG LẶP LỒNG NHAU: VỚI MỖI CHẤT, DUYỆT HẾT CÁC GIÁ TRỊ → 4 x 4 = 16 LÁ.
	// >>> DẤU _ (BLANK IDENTIFIER) BỎ QUA CHỈ SỐ. GO BÁO LỖI NẾU KHAI BÁO BIẾN MÀ KHÔNG DÙNG,
	// >>> NÊN KHI KHÔNG CẦN CHỈ SỐ PHẢI VIẾT _ THAY CHO i.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// >>> value+" of "+suit: NỐI CHUỖI BẰNG TOÁN TỬ +, VÍ DỤ "Ace of Spades".
			// >>> append TRẢ VỀ SLICE MỚI NÊN PHẢI GÁN LẠI CHO cards.
			cards = append(cards, value+" of "+suit)
		}
	}
	// >>> TRẢ BỘ BÀI ĐÃ TẠO VỀ CHO NƠI GỌI HÀM.
	return cards

}

// print là method của deck: in từng lá bài kèm chỉ số.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
