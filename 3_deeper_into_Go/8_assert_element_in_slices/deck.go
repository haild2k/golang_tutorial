// CHỦ ĐỀ: XÁO TRỘN / HOÁN ĐỔI PHẦN TỬ TRONG SLICE VÀ KIỂM TRA (ASSERT) BẰNG TEST
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	"io/ioutil" // đọc/ghi file (từ Go 1.16 nên dùng os.ReadFile / os.WriteFile)
	// >>> math/rand: SINH SỐ NGẪU NHIÊN, DÙNG ĐỂ CHỌN VỊ TRÍ HOÁN ĐỔI KHI XÁO BÀI.
	"math/rand"
	"os"
	"strings"
	// >>> time: LẤY THỜI GIAN HIỆN TẠI LÀM "HẠT GIỐNG" (SEED) CHO BỘ SINH SỐ NGẪU NHIÊN.
	"time"
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

// newDeckFromFile đọc file, có lỗi thì in ra và thoát; không lỗi thì tách chuỗi thành bộ bài.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// >>> func (d deck) shuffle(): METHOD XÁO TRỘN BỘ BÀI NGAY TRÊN CHÍNH SLICE d.
// >>> KHÔNG CẦN TRẢ VỀ GÌ: SLICE CHỨA CON TRỎ TỚI MẢNG BÊN DƯỚI, NÊN SỬA d[i] LÀ SỬA LUÔN BỘ BÀI GỐC.
func (d deck) shuffle() {
	// >>> time.Now().UnixNano(): THỜI GIAN HIỆN TẠI TÍNH BẰNG NANO GIÂY (int64), MỖI LẦN CHẠY MỘT KHÁC.
	// >>> rand.NewSource(seed): TẠO NGUỒN SỐ NGẪU NHIÊN TỪ SEED ĐÓ.
	source := rand.NewSource(time.Now().UnixNano())
	// >>> rand.New(source): TẠO BỘ SINH SỐ NGẪU NHIÊN r DÙNG NGUỒN TRÊN.
	// >>> NẾU KHÔNG ĐỔI SEED, LẦN CHẠY NÀO CŨNG RA CÙNG MỘT THỨ TỰ (TRƯỚC GO 1.20).
	r := rand.New(source)

	// >>> for i := range d: CHỈ LẤY CHỈ SỐ i (BỎ GIÁ TRỊ), DUYỆT TỪ 0 ĐẾN len(d)-1.
	for i := range d {
		// >>> r.Intn(n): SỐ NGẪU NHIÊN TỪ 0 ĐẾN n-1.
		// Lưu ý: Intn(len(d)-1) không bao giờ chọn chỉ số cuối; dùng r.Intn(len(d)) sẽ xáo đều hơn.
		newPosition := r.Intn(len(d) - 1)

		// >>> HOÁN ĐỔI 2 PHẦN TỬ TRONG 1 DÒNG (GÁN SONG SONG):
		// >>> VẾ PHẢI ĐƯỢC TÍNH XONG HẾT RỒI MỚI GÁN SANG VẾ TRÁI, NÊN KHÔNG CẦN BIẾN TẠM.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
