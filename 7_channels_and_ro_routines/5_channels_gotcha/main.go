// CHỦ ĐỀ: KIỂM TRA LIÊN TỤC VỚI CHANNEL — CÁC "BẪY" (GOTCHA) HAY GẶP
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// Chương trình chạy mãi (cứ 5 giây kiểm tra lại mỗi link). Dừng bằng Ctrl+C.

package main

import (
	"fmt"
	"net/http"
	"time" // dùng time.Sleep để tạm dừng
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

	// Lần kiểm tra đầu tiên: mỗi link một goroutine.
	for _, link := range links {
		go checkLink(link, c)
	}

	// >>> for l := range c: CỨ MỖI LẦN CHANNEL CÓ GIÁ TRỊ MỚI THÌ GÁN VÀO l VÀ CHẠY THÂN VÒNG LẶP.
	// >>> TƯƠNG ĐƯƠNG for { l := <-c; ... }. VÒNG LẶP CHỈ DỪNG KHI CHANNEL BỊ close(c) — Ở ĐÂY KHÔNG BAO GIỜ close → CHẠY MÃI.
	for l := range c {
		// >>> go func(link string) { ... }(l): FUNCTION LITERAL (HÀM VÔ DANH / LAMBDA) CHẠY TRONG GOROUTINE MỚI.
		// >>>   - func(link string) {...}: ĐỊNH NGHĨA HÀM KHÔNG TÊN NGAY TẠI CHỖ.
		// >>>   - (l) Ở CUỐI: GỌI HÀM ĐÓ NGAY, TRUYỀN l VÀO THAM SỐ link.
		go func(link string) {
			// >>> GOTCHA 1: KHÔNG GỌI time.Sleep TRONG main ROUTINE — SẼ CHẶN VIỆC NHẬN TỪ CHANNEL CỦA MỌI LINK KHÁC.
			// >>> ĐẶT Sleep TRONG GOROUTINE CON → CHỈ LINK NÀY ĐỢI 5 GIÂY, main VẪN NHẬN TIẾP ĐƯỢC.
			// >>> 5 * time.Second: KIỂU time.Duration = 5 GIÂY.
			time.Sleep(5 * time.Second)
			// >>> GOTCHA 2: DÙNG link (THAM SỐ, BẢN SAO RIÊNG CỦA GOROUTINE NÀY), KHÔNG DÙNG TRỰC TIẾP l.
			// >>> l BỊ VÒNG LẶP GÁN LẠI LIÊN TỤC; NẾU GOROUTINE ĐỌC l SAU 5 GIÂY CÓ THỂ NHẬN NHẦM LINK KHÁC.
			// >>> (TỪ GO 1.22 MỖI VÒNG LẶP CÓ BIẾN RIÊNG NÊN BỚT LỖI NÀY, NHƯNG TRUYỀN QUA THAM SỐ VẪN LÀ CÁCH RÕ RÀNG NHẤT.)
			// >>> NGUYÊN TẮC: KHÔNG CHIA SẺ BIẾN GIỮA CÁC GOROUTINE — HÃY TRUYỀN QUA THAM SỐ HOẶC CHANNEL.
			checkLink(link, c)
		}(l)
	}
}

// checkLink kiểm tra link rồi gửi CHÍNH link đó vào channel,
// để main biết link nào vừa kiểm tra xong và lên lịch kiểm tra lại.
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// >>> GỬI link (THAY VÌ CÂU THÔNG BÁO) → main NHẬN ĐƯỢC LINK ĐỂ KIỂM TRA LẠI.
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
