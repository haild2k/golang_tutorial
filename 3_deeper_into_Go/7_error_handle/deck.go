// CHỦ ĐỀ: XỬ LÝ LỖI (ERROR HANDLING)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	"io/ioutil" // đọc/ghi file (từ Go 1.16 nên dùng os.ReadFile / os.WriteFile)
	// >>> PACKAGE os: TƯƠNG TÁC VỚI HỆ ĐIỀU HÀNH, Ở ĐÂY DÙNG os.Exit ĐỂ THOÁT CHƯƠNG TRÌNH KHI LỖI.
	"os"
	"strings"
)

// deck là kiểu tự định nghĩa, bên dưới là slice chuỗi.
type deck []string

// newDeck tạo bộ bài 16 lá (4 chất x 4 giá trị).
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// print in từng lá bài kèm chỉ số.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal chia bộ bài thành 2 phần: handSize lá đầu và phần còn lại.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString nối các lá bài thành một chuỗi, cách nhau bằng dấu phẩy.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile ghi bộ bài ra file, trả về error (nil nếu thành công).
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// newDeckFromFile đọc file và dựng lại bộ bài từ nội dung file.
func newDeckFromFile(filename string) deck {
	// >>> bs, err := ioutil.ReadFile(filename): HÀM TRẢ VỀ 2 GIÁ TRỊ.
	// >>>   - bs ([]byte): NỘI DUNG FILE.
	// >>>   - err (error): nil NẾU ĐỌC THÀNH CÔNG, KHÁC nil NẾU LỖI (VD: FILE KHÔNG TỒN TẠI).
	// >>> ĐÂY LÀ MẪU CHUẨN CỦA GO: LỖI ĐƯỢC TRẢ VỀ NHƯ MỘT GIÁ TRỊ BÌNH THƯỜNG, KHÔNG DÙNG try/catch.
	bs, err := ioutil.ReadFile(filename)

	// >>> if err != nil { ... }: KIỂM TRA LỖI NGAY SAU KHI GỌI HÀM. CÓ LỖI THÌ XỬ LÝ TRONG KHỐI NÀY.
	if err != nil {
		// >>> CÁCH XỬ LÝ 1: IN THÔNG BÁO LỖI RA MÀN HÌNH.
		fmt.Println("Error:", err)
		// >>> CÁCH XỬ LÝ 2: os.Exit(1) DỪNG CHƯƠNG TRÌNH NGAY. MÃ KHÁC 0 BÁO CHO HỆ ĐIỀU HÀNH LÀ CÓ LỖI.
		// >>> (CÁCH KHÁC: TRẢ VỀ newDeck() MẶC ĐỊNH, HOẶC TRẢ error LÊN CHO HÀM GỌI XỬ LÝ.)
		os.Exit(1)
	}

	// string(bs): ép slice byte về chuỗi.
	// strings.Split(chuỗi, ","): tách chuỗi theo dấu phẩy → []string (ngược lại với strings.Join).
	s := strings.Split(string(bs), ",")
	// deck(s): ép []string về kiểu deck rồi trả về.
	return deck(s)
}
