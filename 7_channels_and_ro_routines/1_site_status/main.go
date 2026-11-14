// CHỦ ĐỀ: KIỂM TRA TRẠNG THÁI WEBSITE — CHẠY TUẦN TỰ (CHƯA DÙNG GOROUTINE)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	"net/http" // gửi HTTP request
)

func main() {
	// Danh sách các website cần kiểm tra.
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// >>> GỌI checkLink TUẦN TỰ (BLOCKING): PHẢI ĐỢI LINK NÀY XONG MỚI KIỂM TRA LINK TIẾP THEO.
	// >>> TỔNG THỜI GIAN = TỔNG THỜI GIAN CỦA TỪNG REQUEST → CHẬM. BÀI SAU DÙNG GOROUTINE ĐỂ CHẠY SONG SONG.
	for _, link := range links {
		checkLink(link)
	}

}

// checkLink gửi request GET tới link và in ra link còn hoạt động hay không.
func checkLink(link string) {
	// >>> http.Get BLOCK (CHẶN) CHƯƠNG TRÌNH CHO ĐẾN KHI CÓ PHẢN HỒI HOẶC LỖI.
	// Dấu _ bỏ qua response vì chỉ cần biết có lỗi hay không.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// return sớm để không in dòng "is up!" bên dưới.
		return
	}

	fmt.Println(link, "is up!")
}
