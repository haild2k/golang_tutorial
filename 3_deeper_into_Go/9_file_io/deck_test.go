// CHỦ ĐỀ: ĐỌC/GHI FILE (FILE I/O) VÀ TEST CHO FILE I/O
// Quy ước comment:
//   - Comment VIẾT HOA, bắt đầu bằng ">>>": cú pháp thuộc chủ đề của bài này.
//   - Comment chữ thường: giải thích chung.
//
// Chạy test bằng lệnh: go test  (hoặc go test -v để xem chi tiết từng test).

package main

import (
	// >>> os: DÙNG os.Remove ĐỂ XOÁ FILE TẠM TRƯỚC VÀ SAU KHI TEST.
	"os"
	"testing" // thư viện test chuẩn của Go
)

// TestNewDeck kiểm tra bộ bài mới: đủ 16 lá, lá đầu và lá cuối đúng thứ tự.
// Hàm test phải bắt đầu bằng "Test" và nhận tham số t *testing.T.
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Kiểm tra độ dài; sai thì t.Errorf đánh dấu test thất bại.
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Kiểm tra phần tử đầu tiên (chỉ số 0).
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// d[len(d)-1]: phần tử cuối cùng (Go không hỗ trợ chỉ số âm như d[-1]).
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// >>> TEST CHO FILE I/O: GHI BỘ BÀI RA FILE RỒI ĐỌC LẠI, KIỂM TRA DỮ LIỆU KHÔNG BỊ MẤT.
// Tên hàm dài nhưng mô tả rõ những gì được test.
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// >>> os.Remove(tên_file): XOÁ FILE. GỌI TRƯỚC KHI TEST ĐỂ DỌN FILE CŨ CÒN SÓT TỪ LẦN CHẠY LỖI TRƯỚC.
	// >>> FILE KHÔNG TỒN TẠI THÌ os.Remove TRẢ VỀ error, Ở ĐÂY BỎ QUA VÌ KHÔNG QUAN TRỌNG.
	os.Remove("_decktesting")

	// Biến tên deck trùng tên kiểu deck: hợp lệ nhưng che mất tên kiểu trong hàm này (nên tránh).
	deck := newDeck()
	// >>> GHI BỘ BÀI RA FILE TẠM "_decktesting".
	deck.saveToFile("_decktesting")

	// >>> ĐỌC LẠI TỪ FILE VỪA GHI ĐỂ DỰNG BỘ BÀI MỚI.
	loadedDeck := newDeckFromFile("_decktesting")

	// >>> BỘ BÀI ĐỌC TỪ FILE PHẢI ĐỦ 16 LÁ NHƯ LÚC GHI.
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// >>> DỌN DẸP: XOÁ FILE TẠM SAU KHI TEST XONG ĐỂ KHÔNG ĐỂ LẠI RÁC TRONG THƯ MỤC.
	os.Remove("_decktesting")
}
