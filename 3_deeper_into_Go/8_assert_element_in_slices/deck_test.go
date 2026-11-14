// CHỦ ĐỀ: XÁO TRỘN / HOÁN ĐỔI PHẦN TỬ TRONG SLICE VÀ KIỂM TRA (ASSERT) BẰNG TEST
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// >>> FILE TEST PHẢI CÓ ĐUÔI _test.go. CHẠY BẰNG LỆNH: go test (CẦN CÓ go.mod; TẠO BẰNG go mod init <tên>).

// File test cùng package main nên gọi được trực tiếp newDeck().
package main

import (
	// >>> PACKAGE testing: THƯ VIỆN TEST CHUẨN CỦA GO.
	"testing"
)

// >>> func TestNewDeck(t *testing.T): HÀM TEST PHẢI BẮT ĐẦU BẰNG "Test" + CHỮ HOA,
// >>> VÀ NHẬN DUY NHẤT THAM SỐ t KIỂU *testing.T (CON TRỎ) ĐỂ BÁO KẾT QUẢ TEST.
func TestNewDeck(t *testing.T) {
	// Tạo bộ bài cần kiểm tra.
	d := newDeck()

	// >>> GO KHÔNG CÓ HÀM assert: TA TỰ VIẾT if <ĐIỀU KIỆN SAI> { t.Errorf(...) }.
	// >>> len(d): ĐỘ DÀI SLICE. BỘ BÀI ĐÚNG PHẢI CÓ 16 LÁ.
	if len(d) != 16 {
		// >>> t.Errorf: ĐÁNH DẤU TEST THẤT BẠI VÀ IN THÔNG BÁO, NHƯNG VẪN CHẠY TIẾP CÁC DÒNG SAU.
		// >>> %v: CHỖ GIỮ CHỖ, ĐƯỢC THAY BẰNG GIÁ TRỊ len(d) Ở DẠNG MẶC ĐỊNH.
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
