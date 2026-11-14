// CHỦ ĐỀ: DUYỆT QUA MAP (ITERATING OVER MAPS)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

func main() {
	// Map literal: tên màu → mã hex.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Truyền map vào hàm để in.
	printMap(colors)
}

// >>> c map[string]string: THAM SỐ KIỂU MAP — GHI ĐẦY ĐỦ map[KIỂU_KEY]KIỂU_VALUE.
func printMap(c map[string]string) {
	// >>> for key, value := range map: DUYỆT TỪNG CẶP. color NHẬN KEY, hex NHẬN VALUE.
	// >>> THỨ TỰ DUYỆT MAP LÀ NGẪU NHIÊN, MỖI LẦN CHẠY CÓ THỂ KHÁC → KHÔNG ĐƯỢC DỰA VÀO THỨ TỰ.
	// >>> CHỈ CẦN KEY: for color := range c ; CHỈ CẦN VALUE: for _, hex := range c.
	for color, hex := range c {
		// Println tự thêm khoảng trắng giữa các tham số.
		fmt.Println("Hex code for", color, "is", hex)
	}
}
