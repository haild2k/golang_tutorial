// CHỦ ĐỀ: GOROUTINE — CHẠY HÀM SONG SONG VỚI TỪ KHOÁ go (VÀ VÌ SAO CẦN CHANNEL)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> GOROUTINE LÀ MỘT "LUỒNG" NHẸ DO GO RUNTIME QUẢN LÝ. HÀM main CŨNG CHẠY TRONG MỘT GOROUTINE (MAIN ROUTINE).
// >>> KHI MAIN ROUTINE KẾT THÚC, CHƯƠNG TRÌNH THOÁT NGAY — MỌI GOROUTINE CON BỊ DỪNG DÙ CHƯA CHẠY XONG.

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

	for _, link := range links {
		// >>> go checkLink(link): TẠO GOROUTINE MỚI ĐỂ CHẠY checkLink, KHÔNG ĐỢI NÓ XONG MÀ CHẠY TIẾP VÒNG LẶP.
		// >>> TỪ KHOÁ go CHỈ ĐẶT TRƯỚC LỜI GỌI HÀM.
		go checkLink(link)
	}

	// >>> VẤN ĐỀ: VÒNG LẶP XONG GẦN NHƯ NGAY LẬP TỨC → main KẾT THÚC → CHƯƠNG TRÌNH THOÁT
	// >>> TRƯỚC KHI CÁC REQUEST KỊP PHẢN HỒI → THƯỜNG KHÔNG IN RA GÌ.
	// >>> GIẢI PHÁP: DÙNG CHANNEL ĐỂ main ĐỢI CÁC GOROUTINE BÁO KẾT QUẢ (BÀI 3).
}

// checkLink gửi request GET tới link và in ra link còn hoạt động hay không.
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
