// CHỦ ĐỀ: HÀM Read VÀ INTERFACE io.Reader
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> resp.Body CÓ KIỂU io.ReadCloser — INTERFACE GỘP TỪ 2 INTERFACE:
// >>>   type Reader interface { Read(p []byte) (n int, err error) }
// >>>   type Closer interface { Close() error }
// >>> NÊN Body CÓ METHOD Read: ĐỔ DỮ LIỆU VÀO SLICE BYTE MÀ TA TRUYỀN VÀO.

package main

import (
	"fmt"
	"net/http"
	"os"
)

// logWriter chưa dùng ở bài này (sẽ dùng ở bài 5_custom_writer).
type logWriter struct{}

func main() {
	// Gửi request GET, nhận response và lỗi.
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// >>> make([]byte, 99999): TẠO SLICE BYTE CÓ SẴN 99999 PHẦN TỬ (TOÀN SỐ 0) LÀM "THÙNG CHỨA".
	// >>> Read KHÔNG TỰ NỚI RỘNG SLICE, NÊN PHẢI TẠO SẴN ĐỦ CHỖ.
	bs := make([]byte, 99999)

	// >>> resp.Body.Read(bs): GHI DỮ LIỆU ĐỌC ĐƯỢC VÀO bs (SỬA TRỰC TIẾP SLICE).
	// >>> TRẢ VỀ (SỐ BYTE ĐÃ ĐỌC, error) — Ở ĐÂY ĐANG BỎ QUA CẢ HAI.
	// >>> MỘT LẦN Read CÓ THỂ CHƯA ĐỌC HẾT BODY; ĐỌC ĐỦ CẦN LẶP HOẶC DÙNG io.Copy / io.ReadAll (BÀI 4).
	resp.Body.Read(bs)

	// >>> string(bs): ĐỔI SLICE BYTE THÀNH CHUỖI ĐỂ IN NỘI DUNG HTML.
	fmt.Println(string(bs))

}
