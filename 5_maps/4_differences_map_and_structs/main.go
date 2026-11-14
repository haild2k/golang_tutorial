// CHỦ ĐỀ: KHÁC NHAU GIỮA MAP VÀ STRUCT
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> SO SÁNH MAP VÀ STRUCT:
// >>>   MAP                                    | STRUCT
// >>>   - MỌI KEY CÙNG KIỂU, MỌI VALUE CÙNG KIỂU | - MỖI TRƯỜNG CÓ THỂ KHÁC KIỂU
// >>>   - KEY CÓ THỂ THÊM/XOÁ LÚC CHẠY          | - CÁC TRƯỜNG CỐ ĐỊNH TỪ LÚC VIẾT CODE
// >>>   - DUYỆT ĐƯỢC BẰNG range                | - KHÔNG DUYỆT ĐƯỢC CÁC TRƯỜNG BẰNG range
// >>>   - TRUY CẬP BẰNG m["key"]               | - TRUY CẬP BẰNG s.field
// >>>   - KIỂU THAM CHIẾU: TRUYỀN VÀO HÀM VẪN  | - KIỂU GIÁ TRỊ: TRUYỀN VÀO HÀM LÀ BẢN SAO,
// >>>     SỬA ĐƯỢC MAP GỐC                      |   MUỐN SỬA GỐC PHẢI DÙNG CON TRỎ
// >>>   - DÙNG CHO TẬP DỮ LIỆU CÙNG LOẠI,       | - DÙNG ĐỂ MÔ TẢ MỘT "VẬT" CÓ NHIỀU THUỘC TÍNH
// >>>     CHƯA BIẾT TRƯỚC SỐ LƯỢNG              |   KHÁC NHAU (VD: person)

package main

import "fmt"

func main() {
	// Map literal: tên màu → mã hex.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// >>> THÊM KEY MỚI LÚC CHẠY — ĐIỀU STRUCT KHÔNG LÀM ĐƯỢC (STRUCT KHÔNG THỂ THÊM TRƯỜNG MỚI).
	colors["yellow"] = "laksjdf"

	// >>> XOÁ KEY LÚC CHẠY — STRUCT CŨNG KHÔNG XOÁ ĐƯỢC TRƯỜNG.
	delete(colors, "yellow")

	// >>> MAP LÀ KIỂU THAM CHIẾU: printMap NHẬN CÙNG MỘT MAP (KHÔNG PHẢI BẢN SAO).
	// >>> NẾU printMap SỬA c THÌ colors Ở ĐÂY CŨNG ĐỔI THEO — KHÁC VỚI STRUCT TRUYỀN THEO GIÁ TRỊ.
	printMap(colors)
}

// printMap in từng cặp màu → mã hex.
func printMap(c map[string]string) {
	// >>> DUYỆT ĐƯỢC TOÀN BỘ KEY/VALUE BẰNG range — STRUCT KHÔNG LÀM ĐƯỢC NHƯ VẬY.
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
