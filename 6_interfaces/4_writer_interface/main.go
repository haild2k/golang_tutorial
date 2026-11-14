// CHỦ ĐỀ: INTERFACE io.Writer VÀ HÀM io.Copy
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> io.Writer LÀ INTERFACE NGƯỢC VỚI io.Reader:
// >>>   type Writer interface { Write(p []byte) (n int, err error) }
// >>> Reader: LẤY DỮ LIỆU TỪ NGUỒN RA []byte. Writer: ĐƯA []byte TỚI ĐÍCH (MÀN HÌNH, FILE, MẠNG...).

package main

import (
	"fmt"
	// >>> io: CHỨA CÁC INTERFACE Reader, Writer VÀ HÀM TIỆN ÍCH NHƯ io.Copy.
	"io"
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

	// >>> io.Copy(dst io.Writer, src io.Reader): ĐỌC HẾT DỮ LIỆU TỪ src VÀ GHI SANG dst.
	// >>>   - os.Stdout: MÀN HÌNH TERMINAL, KIỂU *os.File, CÓ METHOD Write → THOẢ io.Writer.
	// >>>   - resp.Body: CÓ METHOD Read → THOẢ io.Reader.
	// >>> io.Copy TỰ LẶP ĐẾN KHI ĐỌC HẾT, KHÔNG CẦN TẠO SẴN SLICE BYTE NHƯ BÀI 3.
	// >>> NHỜ INTERFACE, io.Copy DÙNG ĐƯỢC VỚI MỌI NGUỒN/ĐÍCH (FILE, MẠNG, BỘ NHỚ...).
	io.Copy(os.Stdout, resp.Body)

}
