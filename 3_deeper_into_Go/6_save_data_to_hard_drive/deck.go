// CHỦ ĐỀ: LƯU DỮ LIỆU XUỐNG Ổ CỨNG (GHI FILE)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	// >>> PACKAGE io/ioutil CUNG CẤP HÀM ĐỌC/GHI FILE NHANH: WriteFile, ReadFile.
	// >>> TỪ GO 1.16, ioutil ĐÃ DEPRECATED → NÊN DÙNG os.WriteFile / os.ReadFile (CÚ PHÁP Y HỆT).
	"io/ioutil"
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

// >>> func (d deck) saveToFile(filename string) error: METHOD GHI BỘ BÀI RA FILE.
// >>> KIỂU TRẢ VỀ error: nil NẾU GHI THÀNH CÔNG, KHÁC nil NẾU CÓ LỖI (VD: KHÔNG CÓ QUYỀN GHI).
func (d deck) saveToFile(filename string) error {
	// >>> ioutil.WriteFile(tên_file, dữ_liệu []byte, quyền):
	// >>>   - []byte(d.toString()): ÉP CHUỖI THÀNH SLICE BYTE, VÌ FILE LƯU DỮ LIỆU DẠNG BYTE.
	// >>>   - 0666: QUYỀN FILE DẠNG BÁT PHÂN (OCTAL) → MỌI NGƯỜI ĐỀU ĐỌC/GHI ĐƯỢC.
	// >>>   - FILE CHƯA CÓ THÌ TẠO MỚI, ĐÃ CÓ THÌ GHI ĐÈ.
	// >>> return TRẢ THẲNG error CỦA WriteFile VỀ CHO NƠI GỌI XỬ LÝ.
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
