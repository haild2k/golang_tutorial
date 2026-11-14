// CHỦ ĐỀ: HÀM TRẢ VỀ NHIỀU GIÁ TRỊ (MULTIPLE RETURN VALUES)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// deck là kiểu tự định nghĩa, bên dưới là slice chuỗi.
type deck []string

// newDeck tạo bộ bài 16 lá (4 chất x 4 giá trị).
func newDeck() deck {
	cards := deck{} // bộ bài rỗng

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Hai vòng lặp lồng nhau ghép từng giá trị với từng chất.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// print là method của deck: in từng lá bài kèm chỉ số.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// >>> func deal(d deck, handSize int) (deck, deck):
// >>>   - NHẬN 2 THAM SỐ: d KIỂU deck VÀ handSize KIỂU int.
// >>>   - (deck, deck) TRONG NGOẶC ĐƠN = HÀM TRẢ VỀ 2 GIÁ TRỊ, CẢ HAI ĐỀU KIỂU deck.
func deal(d deck, handSize int) (deck, deck) {
	// >>> return a, b: TRẢ VỀ NHIỀU GIÁ TRỊ, NGĂN CÁCH BẰNG DẤU PHẨY, ĐÚNG THỨ TỰ ĐÃ KHAI BÁO.
	// >>> d[:handSize] = d[0:handSize]: LẤY CÁC PHẦN TỬ TỪ CHỈ SỐ 0 ĐẾN handSize-1 (BÀI TRÊN TAY).
	// >>> d[handSize:]: LẤY TỪ CHỈ SỐ handSize ĐẾN HẾT (BÀI CÒN LẠI).
	// Lưu ý: nếu handSize > len(d) chương trình sẽ panic.
	return d[:handSize], d[handSize:]
}
