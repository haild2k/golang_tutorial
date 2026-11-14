// CHỦ ĐỀ: STRUCT VỚI RECEIVER FUNCTION (METHOD CỦA STRUCT)
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

	// >>> GỌI METHOD TRÊN STRUCT: biến.method(tham_số).
	// >>> KẾT QUẢ: firstName VẪN LÀ "Jim", KHÔNG ĐỔI THÀNH "jimmy" (XEM GIẢI THÍCH Ở updateName).
	jim.updateName("jimmy")
	jim.print()
}

// >>> func (p person) updateName(...): METHOD VỚI RECEIVER KIỂU GIÁ TRỊ (VALUE RECEIVER).
// >>> GO LÀ NGÔN NGỮ "PASS BY VALUE": KHI GỌI jim.updateName, p LÀ MỘT BẢN SAO CỦA jim.
// >>> GÁN p.firstName CHỈ SỬA BẢN SAO, jim GỐC KHÔNG ĐỔI → BÀI SAU DÙNG CON TRỎ ĐỂ SỬA ĐƯỢC.
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// >>> func (p person) print(): METHOD CHỈ ĐỌC DỮ LIỆU NÊN DÙNG VALUE RECEIVER LÀ ĐỦ.
func (p person) print() {
	fmt.Printf("%+v", p)
}
