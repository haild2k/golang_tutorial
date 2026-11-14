// CHỦ ĐỀ: NHẬN NHIỀU THÔNG ĐIỆP TỪ CHANNEL
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	"net/http"
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

	// Tạo channel chở string (không bộ đệm).
	c := make(chan string)

	// Mỗi link chạy trong một goroutine riêng, cùng gửi kết quả vào c.
	for _, link := range links {
		go checkLink(link, c)
	}

	// >>> GỬI BAO NHIÊU THÌ NHẬN BẤY NHIÊU: CÓ len(links) GOROUTINE, MỖI CÁI GỬI ĐÚNG 1 LẦN
	// >>> → NHẬN len(links) LẦN THÌ main CHỜ ĐƯỢC TẤT CẢ GOROUTINE CHẠY XONG.
	// >>> for i := 0; i < n; i++: VÒNG LẶP FOR KIỂU C (KHỞI TẠO; ĐIỀU KIỆN; BƯỚC TĂNG).
	for i := 0; i < len(links); i++ {
		// >>> MỖI <-c CHẶN main CHO ĐẾN KHI CÓ 1 THÔNG ĐIỆP MỚI.
		// >>> THỨ TỰ NHẬN = THỨ TỰ GOROUTINE GỬI XONG (LINK NÀO PHẢN HỒI NHANH THÌ ĐẾN TRƯỚC), KHÔNG THEO THỨ TỰ links.
		// >>> NẾU NHẬN NHIỀU HƠN SỐ LẦN GỬI, main SẼ ĐỢI MÃI → LỖI "all goroutines are asleep - deadlock!".
		fmt.Println(<-c)
	}
}

// checkLink kiểm tra link rồi gửi một thông điệp vào channel c.
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think" // gửi kết quả vào channel
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up" // gửi kết quả vào channel
}
