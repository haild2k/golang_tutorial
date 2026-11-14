// CHỦ ĐỀ: KHAI BÁO MAP
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

func main() {
	// >>> map[string]string{...}: MAP LITERAL — TẬP HỢP CÁC CẶP KEY: VALUE.
	// >>> map[KIỂU_KEY]KIỂU_VALUE: Ở ĐÂY KEY LÀ string, VALUE CŨNG LÀ string.
	// >>> MỌI KEY PHẢI CÙNG KIỂU, MỌI VALUE PHẢI CÙNG KIỂU; KEY KHÔNG ĐƯỢC TRÙNG NHAU.
	colors := map[string]string{
		// >>> MỖI DÒNG MỘT CẶP "key": "value", NGĂN CÁCH BẰNG DẤU PHẨY.
		"red": "#ff0000",
		// >>> VIẾT NHIỀU DÒNG THÌ CẶP CUỐI CŨNG PHẢI CÓ DẤU PHẨY.
		"green": "#4bf745",
	}

	// In cả map: map[green:#4bf745 red:#ff0000]
	// (fmt in map theo thứ tự key đã sắp xếp; còn khi duyệt bằng range thì thứ tự là ngẫu nhiên.)
	fmt.Println(colors)
}
