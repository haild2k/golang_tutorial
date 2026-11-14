// CHỦ ĐỀ: INTERFACE
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// >>> type bot interface { ... }: KHAI BÁO INTERFACE TÊN bot — MỘT "HỢP ĐỒNG" GỒM DANH SÁCH METHOD.
// >>> KIỂU NÀO CÓ ĐỦ CÁC METHOD NÀY (ĐÚNG TÊN, THAM SỐ, KIỂU TRẢ VỀ) THÌ TỰ ĐỘNG LÀ MỘT bot.
type bot interface {
	// >>> getGreeting() string: KIỂU MUỐN LÀ bot PHẢI CÓ METHOD getGreeting KHÔNG THAM SỐ, TRẢ VỀ string.
	getGreeting() string
}

// struct{}: struct rỗng, không có trường nào — chỉ dùng để gắn method.
type englishBot struct{}
type spanishBot struct{}

func main() {
	// Tạo giá trị của 2 kiểu struct rỗng.
	eb := englishBot{}
	sb := spanishBot{}

	// >>> eb VÀ sb KHÁC KIỂU NHƯNG ĐỀU TRUYỀN ĐƯỢC VÀO printGreeting(b bot)
	// >>> VÌ CẢ HAI ĐỀU CÓ METHOD getGreeting() string → ĐỀU THOẢ INTERFACE bot.
	printGreeting(eb)
	printGreeting(sb)
}

// >>> b bot: THAM SỐ KIỂU INTERFACE — NHẬN BẤT KỲ KIỂU NÀO THOẢ bot.
// >>> NHỜ VẬY CHỈ CẦN VIẾT 1 HÀM, KHÔNG PHẢI VIẾT RIÊNG printGreeting CHO TỪNG LOẠI BOT.
func printGreeting(b bot) {
	// >>> GỌI METHOD QUA INTERFACE: GO TỰ CHỌN getGreeting CỦA KIỂU THẬT (englishBot HAY spanishBot).
	fmt.Println(b.getGreeting())
} 

// >>> IMPLEMENT NGẦM (IMPLICIT): KHÔNG CẦN VIẾT "englishBot implements bot" — CHỈ CẦN CÓ METHOD ĐÚNG.
// >>> (englishBot) KHÔNG ĐẶT TÊN RECEIVER VÌ BÊN TRONG KHÔNG DÙNG ĐẾN NÓ.
func (englishBot) getGreeting() string {
	return "Hi there!"
}

// >>> spanishBot CŨNG CÓ getGreeting() string → CŨNG LÀ MỘT bot.
func (spanishBot) getGreeting() string {
	return "Hola!"
}
