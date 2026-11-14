// CHỦ ĐỀ: CẬP NHẬT GIÁ TRỊ CÁC TRƯỜNG CỦA STRUCT
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// person là struct có 2 trường kiểu string.
type person struct {
	firstName string
	lastName  string
}

func main() {
	// >>> var alex person: KHAI BÁO BIẾN STRUCT MÀ KHÔNG GÁN GIÁ TRỊ.
	// >>> MỌI TRƯỜNG NHẬN "ZERO VALUE" THEO KIỂU: string → "", int → 0, bool → false.
	var alex person

	// >>> biến.trường = giá_trị: DÙNG DẤU CHẤM ĐỂ TRUY CẬP VÀ GÁN LẠI TỪNG TRƯỜNG.
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// >>> %+v TRONG Printf: IN STRUCT KÈM TÊN TRƯỜNG → {firstName:Alex lastName:Anderson}.
	// >>> (%v CHỈ IN GIÁ TRỊ → {Alex Anderson}.)
	// Printf không tự xuống dòng; thêm "\n" nếu cần.
	fmt.Printf("%+v", alex)
}
