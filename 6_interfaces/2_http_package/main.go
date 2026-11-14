// CHỦ ĐỀ: PACKAGE net/http — GỬI HTTP REQUEST
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import (
	"fmt"
	// >>> net/http: THƯ VIỆN CHUẨN ĐỂ GỬI REQUEST (CLIENT) VÀ TẠO WEB SERVER.
	"net/http"
	"os"
)

// logWriter chưa dùng ở bài này (sẽ dùng ở bài 5_custom_writer).
// Go cho phép khai báo kiểu không dùng; chỉ biến/import không dùng mới bị báo lỗi.
type logWriter struct{}

func main() {
	// >>> http.Get(url): GỬI REQUEST GET, TRẢ VỀ 2 GIÁ TRỊ:
	// >>>   - resp (*http.Response): CON TRỎ TỚI STRUCT CHỨA STATUS, HEADER, BODY...
	// >>>   - err (error): KHÁC nil NẾU LỖI MẠNG, SAI URL...
	resp, err := http.Get("http://google.com")

	// Kiểm tra lỗi theo mẫu chuẩn: có lỗi thì in ra và thoát.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// >>> IN CẢ STRUCT Response: CHỈ THẤY STATUS, HEADER... CÒN Body CHỈ HIỆN ĐỊA CHỈ, KHÔNG THẤY NỘI DUNG HTML.
	// >>> LÝ DO: Body CÓ KIỂU io.ReadCloser (MỘT INTERFACE) → PHẢI "ĐỌC" NÓ (XEM BÀI 3, 4).
	fmt.Println(resp)

}
