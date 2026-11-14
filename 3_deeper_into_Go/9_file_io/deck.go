// CHỦ ĐỀ: ĐỌC/GHI FILE (FILE I/O) VÀ TEST CHO FILE I/O
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	// >>> io/ioutil: ĐỌC/GHI CẢ FILE MỘT LẦN (ReadFile, WriteFile).
	// >>> TỪ GO 1.16 ĐÃ DEPRECATED → THAY BẰNG os.ReadFile / os.WriteFile, CÁCH DÙNG Y HỆT.
	"io/ioutil"
	"math/rand" // sinh số ngẫu nhiên để xáo bài
	// >>> os: LÀM VIỆC VỚI HỆ ĐIỀU HÀNH (THOÁT CHƯƠNG TRÌNH, XOÁ FILE, ...).
	"os"
	"strings"
	"time" // lấy thời gian hiện tại làm seed
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

// >>> GHI FILE: deck → string (toString) → []byte → ioutil.WriteFile.
// >>> 0666 LÀ QUYỀN FILE (BÁT PHÂN): ĐỌC + GHI CHO MỌI NGƯỜI. TRẢ VỀ error (nil = THÀNH CÔNG).
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// >>> ĐỌC FILE: ioutil.ReadFile → []byte → string → strings.Split → deck. NGƯỢC LẠI VỚI saveToFile.
func newDeckFromFile(filename string) deck {
	// >>> ReadFile TRẢ VỀ (NỘI DUNG []byte, error).
	bs, err := ioutil.ReadFile(filename)
	// >>> KIỂM TRA LỖI ĐỌC FILE (VD: FILE KHÔNG TỒN TẠI) → IN LỖI VÀ THOÁT VỚI MÃ 1.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// >>> string(bs): BYTE → CHUỖI; strings.Split(..., ","): TÁCH THEO DẤU PHẨY THÀNH []string.
	s := strings.Split(string(bs), ",")
	// >>> deck(s): ÉP []string VỀ KIỂU deck.
	return deck(s)
}

// shuffle xáo trộn bộ bài ngay trên slice d.
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // seed theo thời gian hiện tại
	r := rand.New(source)                           // bộ sinh số ngẫu nhiên

	for i := range d {
		// Vị trí ngẫu nhiên từ 0 đến len(d)-2 (dùng len(d) sẽ xáo đều hơn).
		newPosition := r.Intn(len(d) - 1)

		// Hoán đổi 2 phần tử trong một dòng.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
