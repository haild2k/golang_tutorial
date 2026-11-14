// CHỦ ĐỀ: CÁC PHÉP TOÁN CON TRỎ (& VÀ *)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> TÓM TẮT:
// >>>   &biến      → LẤY ĐỊA CHỈ BỘ NHỚ CỦA biến (TẠO CON TRỎ).
// >>>   *con_trỏ   → LẤY GIÁ TRỊ TẠI ĐỊA CHỈ MÀ CON TRỎ TRỎ TỚI (DEREFERENCE).
// >>>   *kiểu      → (TRONG KHAI BÁO KIỂU) "CON TRỎ TỚI kiểu", VD: *person.

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
	// >>> &jim: LẤY ĐỊA CHỈ CỦA jim. jimPointer CÓ KIỂU *person (CON TRỎ TỚI person).
	jimPointer := &jim

	// >>> GỌI METHOD CÓ POINTER RECEIVER QUA CON TRỎ: updateName NHẬN ĐỊA CHỈ CỦA jim,
	// >>> NÊN SỬA ĐƯỢC jim GỐC.
	jimPointer.updateName("jimmy")

	// In ra: firstName lúc này đã là "jimmy".
	jim.print()
}

// >>> func (pointerToPerson *person) ...: POINTER RECEIVER — RECEIVER CÓ KIỂU *person.
// >>> METHOD NHẬN ĐỊA CHỈ CỦA STRUCT (KHÔNG PHẢI BẢN SAO) NÊN THAY ĐỔI ĐƯỢC STRUCT GỐC.
func (pointerToPerson *person) updateName(newFirstName string) {
	// >>> (*pointerToPerson): DEREFERENCE — ĐI TỚI GIÁ TRỊ person TẠI ĐỊA CHỈ ĐÓ, RỒI GÁN TRƯỜNG firstName.
	// >>> PHẢI CÓ NGOẶC ĐƠN VÌ DẤU . ƯU TIÊN CAO HƠN DẤU *.
	(*pointerToPerson).firstName = newFirstName
}

// print dùng value receiver vì chỉ đọc dữ liệu.
func (p person) print() {
	fmt.Printf("%+v", p)
}
