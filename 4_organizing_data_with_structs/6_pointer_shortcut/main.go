// CHỦ ĐỀ: CÚ PHÁP RÚT GỌN KHI DÙNG CON TRỎ (POINTER SHORTCUT)
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

// person nhúng contactInfo (embedding).
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Khởi tạo person kèm struct con contactInfo.
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// >>> SHORTCUT: KHÔNG CẦN VIẾT jimPointer := &jim RỒI jimPointer.updateName(...).
	// >>> jim LÀ GIÁ TRỊ person, NHƯNG updateName CẦN *person → GO TỰ ĐỘNG CHUYỂN THÀNH (&jim).updateName("jimmy").
	// >>> (CHỈ ÁP DỤNG KHI jim LÀ BIẾN CÓ ĐỊA CHỈ, KHÔNG ÁP DỤNG CHO GIÁ TRỊ TẠM NHƯ person{}.updateName(...).)
	jim.updateName("jimmy")

	// In ra: firstName đã đổi thành "jimmy".
	jim.print()
}

// >>> RECEIVER *person: GỌI ĐƯỢC BẰNG CẢ CON TRỎ (*person) LẪN BIẾN GIÁ TRỊ (person) NHỜ SHORTCUT Ở TRÊN.
func (pointerToPerson *person) updateName(newFirstName string) {
	// >>> CŨNG CÓ SHORTCUT: CÓ THỂ VIẾT GỌN pointerToPerson.firstName = newFirstName,
	// >>> GO TỰ HIỂU LÀ (*pointerToPerson).firstName.
	(*pointerToPerson).firstName = newFirstName
}

// print dùng value receiver vì chỉ đọc dữ liệu.
func (p person) print() {
	fmt.Printf("%+v", p)
}
