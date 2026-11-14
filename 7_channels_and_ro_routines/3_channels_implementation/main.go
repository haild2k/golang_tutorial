// CHỦ ĐỀ: TẠO VÀ DÙNG CHANNEL ĐỂ GOROUTINE GIAO TIẾP
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> CHANNEL LÀ "ỐNG DẪN" ĐỂ CÁC GOROUTINE GỬI/NHẬN DỮ LIỆU VỚI NHAU. MỖI CHANNEL CHỈ CHỞ MỘT KIỂU DỮ LIỆU.
// >>>   c <- giá_trị   : GỬI giá_trị VÀO CHANNEL.
// >>>   x := <-c       : NHẬN MỘT GIÁ TRỊ TỪ CHANNEL, GÁN CHO x.
// >>>   fmt.Println(<-c): NHẬN VÀ DÙNG LUÔN.

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

	// >>> make(chan string): TẠO CHANNEL CHỞ GIÁ TRỊ KIỂU string.
	// >>> ĐÂY LÀ CHANNEL KHÔNG CÓ BỘ ĐỆM (UNBUFFERED): GỬI VÀ NHẬN PHẢI "GẶP NHAU" MỚI HOÀN TẤT.
	c := make(chan string)

	for _, link := range links {
		// >>> TRUYỀN CHANNEL c VÀO GOROUTINE ĐỂ NÓ GỬI KẾT QUẢ VỀ main.
		go checkLink(link, c)
	}

	// >>> <-c LÀ LỆNH BLOCKING: main DỪNG Ở ĐÂY ĐỢI ĐẾN KHI CÓ GOROUTINE GỬI VÀO c.
	// >>> NHƯNG CHỈ NHẬN 1 LẦN → NHẬN XONG TIN ĐẦU TIÊN LÀ main KẾT THÚC,
	// >>> 4 GOROUTINE CÒN LẠI BỊ DỪNG → CHỈ THẤY KẾT QUẢ CỦA LINK PHẢN HỒI NHANH NHẤT. (SỬA Ở BÀI 4.)
	fmt.Println(<-c)
}

// >>> c chan string: THAM SỐ KIỂU CHANNEL CHỞ string.
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// >>> c <- "...": GỬI CHUỖI VÀO CHANNEL ĐỂ BÁO CHO main.
		// >>> VỚI CHANNEL KHÔNG BỘ ĐỆM, GOROUTINE BỊ CHẶN Ở ĐÂY CHO ĐẾN KHI CÓ NGƯỜI NHẬN.
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	// >>> GỬI THÔNG ĐIỆP "LINK CÒN SỐNG" VÀO CHANNEL.
	c <- "Yep its up"
}
