package main

import (
	"fmt"     // In dữ liệu ra màn hình
	"strings" // Các hàm xử lý chuỗi (Join, Split, ...)
)

// deck là kiểu tự định nghĩa, bên dưới là một slice chuỗi ([]string).
// Mỗi phần tử là một lá bài, ví dụ "Ace of Spades".
// Khai báo kiểu riêng giúp ta gắn thêm method (receiver function) cho nó.
type deck []string

// newDeck tạo và trả về một bộ bài mới gồm mọi tổ hợp chất x giá trị.
// Với 4 chất và 4 giá trị, bộ bài có 16 lá.
func newDeck() deck {
	cards := deck{} // Khởi tạo bộ bài rỗng

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // Các chất
	cardValues := []string{"Ace", "Two", "Three", "Four"}          // Các giá trị

	// Hai vòng lặp lồng nhau: với mỗi chất, ghép lần lượt từng giá trị.
	// Dấu _ bỏ qua chỉ số (index) vì ta không dùng đến.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append trả về slice mới, nên phải gán lại cho cards
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// print là method của kiểu deck (receiver d).
// Mọi biến kiểu deck đều gọi được: cards.print().
// Hàm in từng lá bài kèm chỉ số của nó.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal chia bộ bài d thành 2 phần và trả về cả hai (multiple return values):
//   - phần 1: handSize lá đầu tiên (bài trên tay)
//   - phần 2: các lá còn lại
//
// Cú pháp slice d[a:b] lấy phần tử từ chỉ số a đến b-1.
// d[:handSize] tương đương d[0:handSize]; d[handSize:] lấy từ handSize đến hết.
// Lưu ý: nếu handSize > len(d) thì chương trình sẽ panic.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString chuyển bộ bài thành một chuỗi duy nhất, các lá cách nhau bởi dấu phẩy.
// Ví dụ: "Ace of Spades,Two of Spades,..."
//
// []string(d) chuyển kiểu deck về []string, đúng kiểu tham số mà strings.Join cần.
// Chuỗi này hữu ích khi muốn lưu bộ bài ra file (chuyển tiếp sang []byte).
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
