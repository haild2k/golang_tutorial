// CHỦ ĐỀ: KHAI BÁO STRUCT
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// >>> type person struct { ... }: KHAI BÁO KIỂU MỚI TÊN person LÀ MỘT STRUCT.
// >>> STRUCT LÀ TẬP HỢP CÁC TRƯỜNG (FIELD) CÓ TÊN, MỖI TRƯỜNG CÓ KIỂU RIÊNG (GIỐNG OBJECT/RECORD).
type person struct {
	// >>> MỖI DÒNG LÀ MỘT TRƯỜNG: tên_trường kiểu. KHÔNG CÓ DẤU PHẨY GIỮA CÁC TRƯỜNG.
	firstName string
	lastName  string
}

func main() {
	// >>> person{firstName: "Alex", lastName: "Anderson"}: STRUCT LITERAL, GÁN GIÁ TRỊ THEO TÊN TRƯỜNG.
	// >>> CÓ THỂ VIẾT THEO THỨ TỰ person{"Alex", "Anderson"} NHƯNG DỄ SAI KHI THÊM/ĐỔI TRƯỜNG → NÊN GHI TÊN.
	alex := person{firstName: "Alex", lastName: "Anderson"}

	// In struct ra màn hình; fmt.Println chỉ in giá trị: {Alex Anderson}
	fmt.Println(alex)
}
