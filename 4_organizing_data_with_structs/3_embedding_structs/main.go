// CHỦ ĐỀ: NHÚNG STRUCT VÀO STRUCT (EMBEDDING STRUCTS)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

// contactInfo là struct chứa thông tin liên lạc.
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// >>> CHỈ GHI TÊN KIỂU, KHÔNG GHI TÊN TRƯỜNG → contactInfo ĐƯỢC NHÚNG (EMBEDDED) VÀO person.
	// >>> TÊN TRƯỜNG MẶC ĐỊNH CHÍNH LÀ TÊN KIỂU: contactInfo.
	// >>> CÁC TRƯỜNG CON ĐƯỢC "ĐẨY LÊN": jim.email TƯƠNG ĐƯƠNG jim.contactInfo.email.
	// >>> (CÁCH KHÁC: contact contactInfo — TRƯỜNG THƯỜNG CÓ TÊN, LÚC ĐÓ PHẢI GỌI jim.contact.email.)
	contactInfo
}

func main() {
	// >>> KHỞI TẠO STRUCT LỒNG NHAU: VIẾT NHIỀU DÒNG, MỖI TRƯỜNG MỘT DÒNG.
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// >>> TRƯỜNG NHÚNG KHỞI TẠO BẰNG TÊN KIỂU: contactInfo: contactInfo{...}.
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
			// >>> KHI VIẾT NHIỀU DÒNG, DÒNG CUỐI CŨNG PHẢI CÓ DẤU PHẨY, NẾU THIẾU SẼ BÁO LỖI BIÊN DỊCH.
		},
	}

	// In struct kèm tên trường, gồm cả struct con:
	// {firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:94000}}
	fmt.Printf("%+v", jim)
}
