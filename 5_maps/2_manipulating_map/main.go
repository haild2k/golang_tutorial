// CHỦ ĐỀ: THAO TÁC VỚI MAP (TẠO, THÊM, XOÁ)
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.

package main

import "fmt"

func main() {

	// >>> make(map[int]string): TẠO MAP RỖNG, SẴN SÀNG ĐỂ THÊM PHẦN TỬ. KEY KIỂU int, VALUE KIỂU string.
	// >>> CÁC CÁCH TẠO MAP RỖNG KHÁC:
	// >>>   colors := map[int]string{}  → TƯƠNG ĐƯƠNG make.
	// >>>   var colors map[int]string   → MAP nil: ĐỌC ĐƯỢC NHƯNG GHI VÀO SẼ PANIC.
	colors := make(map[int]string)

	// >>> map[key] = value: THÊM CẶP MỚI (HOẶC GHI ĐÈ NẾU KEY ĐÃ CÓ).
	// >>> MAP CHỈ DÙNG ĐƯỢC CÚ PHÁP [ ], KHÔNG DÙNG ĐƯỢC DẤU CHẤM NHƯ STRUCT (colors.10 LÀ SAI).
	colors[10] = "#ffffff"

	// >>> delete(map, key): XOÁ CẶP CÓ KEY TƯƠNG ỨNG. KEY KHÔNG TỒN TẠI THÌ KHÔNG LÀM GÌ, KHÔNG BÁO LỖI.
	delete(colors, 10)

	// In map sau khi xoá: map[] (rỗng).
	fmt.Println(colors)
}
