// CHỦ ĐỀ: TỰ VIẾT MỘT KIỂU THOẢ INTERFACE io.Writer (CUSTOM WRITER)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// >>> logWriter: STRUCT RỖNG SẼ ĐƯỢC GẮN METHOD Write ĐỂ TRỞ THÀNH MỘT io.Writer.
type logWriter struct{}

func main() {
	// Gửi request GET, nhận response và lỗi.
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Tạo giá trị logWriter.
	lw := logWriter{}

	// >>> TRUYỀN lw VÀO VỊ TRÍ io.Writer CỦA io.Copy — HỢP LỆ VÌ logWriter CÓ METHOD Write ĐÚNG CHỮ KÝ.
	// >>> io.Copy SẼ GỌI lw.Write(...) MỘT HOẶC NHIỀU LẦN, MỖI LẦN VỚI MỘT ĐOẠN DỮ LIỆU CỦA BODY.
	io.Copy(lw, resp.Body)
}

// >>> func (logWriter) Write(bs []byte) (int, error): IMPLEMENT io.Writer.
// >>> CHỮ KÝ PHẢI KHỚP CHÍNH XÁC VỚI INTERFACE: TÊN Write (VIẾT HOA), NHẬN []byte, TRẢ VỀ (int, error).
func (logWriter) Write(bs []byte) (int, error) {
	// >>> TỰ QUYẾT ĐỊNH "GHI" NGHĨA LÀ GÌ: Ở ĐÂY IN NỘI DUNG VÀ IN THÊM SỐ BYTE.
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	// >>> TRẢ VỀ ĐÚNG HỢP ĐỒNG: SỐ BYTE ĐÃ XỬ LÝ VÀ nil (KHÔNG LỖI).
	// >>> NẾU TRẢ VỀ SỐ NHỎ HƠN len(bs) MÀ KHÔNG KÈM error, io.Copy SẼ BÁO LỖI "short write".
	return len(bs), nil
}
