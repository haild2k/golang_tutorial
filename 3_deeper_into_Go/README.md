# 3. Deeper into Go — Tóm tắt cú pháp

Cả chương xây dựng dần một chương trình **bộ bài (deck)**: tạo bài, chia bài, lưu/đọc file, xáo bài và viết test.

Chạy một bài: `cd <folder> && go run .` · Chạy test (folder 8, 9): `go test` hoặc `go test -v`

## Liên tưởng với Python

| Go | Python tương tự | Khác biệt cần nhớ |
|---|---|---|
| slice `[]string{"a", "b"}` | list `["a", "b"]` | Slice chỉ chứa một kiểu; list chứa lẫn kiểu được |
| `s = append(s, x)` | `lst.append(x)` | Go trả về slice mới → **phải gán lại**; Python sửa tại chỗ |
| `for i, v := range s` | `for i, v in enumerate(lst)` | — |
| `for _, v := range s` | `for v in lst` | Go bắt buộc dùng `_` cho biến không dùng |
| `type deck []string` | `class Deck(list): ...` | Go không có kế thừa; chỉ tạo kiểu mới dựa trên kiểu cũ |
| `func (d deck) print()` | `def print(self):` | Receiver `d` ≈ `self`, nhưng khai báo **ngoài** kiểu |
| `func f() (deck, deck)` + `a, b := f()` | `return a, b` + `a, b = f()` | Go khai báo rõ kiểu từng giá trị trả về |
| `d[:n]`, `d[n:]` | `d[:n]`, `d[n:]` | Go **không** có chỉ số âm (`d[-1]`) và bước nhảy (`d[::2]`) |
| `strings.Join(s, ",")` | `",".join(s)` | — |
| `strings.Split(s, ",")` | `s.split(",")` | — |
| `[]byte(s)` / `string(bs)` | `s.encode()` / `bs.decode()` | — |
| `os.WriteFile(f, data, 0666)` | `open(f, "wb").write(data)` | Go truyền quyền file trực tiếp |
| `os.ReadFile(f)` | `open(f, "rb").read()` | — |
| `v, err := f()` + `if err != nil` | `try: ... except ...:` | Go không có exception; lỗi là **giá trị trả về** |
| `os.Exit(1)` | `sys.exit(1)` | — |
| `d[i], d[j] = d[j], d[i]` | `d[i], d[j] = d[j], d[i]` | Giống hệt |
| `rand.Intn(n)` | `random.randrange(n)` | Python có sẵn `random.shuffle(lst)` |
| file `deck_test.go`, hàm `TestXxx(t *testing.T)` | file `test_deck.py`, hàm `test_xxx()` (pytest) | — |
| `if cond { t.Errorf(...) }` | `assert cond, "..."` | `t.Errorf` báo lỗi nhưng **chạy tiếp**; `assert` dừng ngay |
| `go test` | `pytest` | — |

| # | Folder | Chủ đề chính |
|---|---|---|
| 1 | [1_slices_and_for_loops](1_slices_and_for_loops/) | Slice, `append`, `for ... range` |
| 2 | [2_custom_type_declarations](2_custom_type_declarations/) | Kiểu tự định nghĩa, receiver (method) |
| 3 | [3_create_new_deck](3_create_new_deck/) | Hàm tạo bộ bài, vòng lặp lồng nhau, `_` |
| 4 | [4_multiple_return_values](4_multiple_return_values/) | Hàm trả về nhiều giá trị, cắt slice |
| 5 | [5_join_slices](5_join_slices/) | Ép kiểu, `strings.Join` |
| 6 | [6_save_data_to_hard_drive](6_save_data_to_hard_drive/) | Ghi file, `[]byte`, kiểu `error` |
| 7 | [7_error_handle](7_error_handle/) | Đọc file, `if err != nil`, `os.Exit` |
| 8 | [8_assert_element_in_slices](8_assert_element_in_slices/) | Xáo bài (random, hoán đổi), viết test |
| 9 | [9_file_io](9_file_io/) | Test cho đọc/ghi file, dọn file tạm |

---

## 1. Slices và vòng lặp for

```go
cards := []string{"Ace of Diamonds", newCard()} // slice literal
cards = append(cards, "Six of Spades")          // append trả về slice MỚI → phải gán lại

for i, card := range cards {                    // i = chỉ số, card = giá trị
    fmt.Println(i, card)
}
```

- Slice giống mảng nhưng co giãn được; mọi phần tử cùng kiểu.
- `:=` vừa khai báo vừa gán, Go tự suy ra kiểu.

## 2. Khai báo kiểu tự định nghĩa (custom type)

```go
type deck []string            // kiểu mới, bên dưới là []string

func (d deck) print() {       // (d deck) là receiver → print là method của deck
    for i, card := range d {
        fmt.Println(i, card)
    }
}

cards := deck{"Ace of Diamonds"} // literal của kiểu deck
cards.print()                    // gọi method
```

- `deck` dùng được mọi thứ của slice (`append`, `range`, chỉ số) và gắn thêm method được.
- Các file cùng thư mục, cùng `package main` được gộp chung → `main.go` dùng được `deck` trong `deck.go`.

## 3. Tạo bộ bài mới

```go
func newDeck() deck {                 // hàm thường, trả về deck
    cards := deck{}                   // bộ bài rỗng
    for _, suit := range cardSuits {  // _ bỏ qua chỉ số không dùng
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit) // nối chuỗi bằng +
        }
    }
    return cards
}
```

- Go báo lỗi nếu khai báo biến mà không dùng → dùng `_` thay cho biến không cần.

## 4. Hàm trả về nhiều giá trị

```go
func deal(d deck, handSize int) (deck, deck) { // (deck, deck): trả về 2 giá trị
    return d[:handSize], d[handSize:]
}

hand, remainingCards := deal(cards, 5)          // nhận 2 giá trị
```

| Cú pháp | Ý nghĩa |
|---|---|
| `d[a:b]` | phần tử từ chỉ số `a` đến `b-1` |
| `d[:n]` | từ đầu đến `n-1` (= `d[0:n]`) |
| `d[n:]` | từ `n` đến hết |

- Số biến bên trái phải bằng số giá trị trả về; không cần giá trị nào thì dùng `_`.

## 5. Nối slice thành chuỗi

```go
import "strings"

func (d deck) toString() string {
    return strings.Join([]string(d), ",") // "Ace of Spades,Two of Spades,..."
}
```

- `KiểuMới(giá_trị)` là cú pháp ép kiểu: `[]string(d)` đổi `deck` về `[]string`.
- `strings.Join(slice, sep)` nối các phần tử, xen giữa là `sep`.

## 6. Lưu dữ liệu xuống ổ cứng

```go
import "io/ioutil"

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

- `[]byte("...")`: đổi chuỗi thành slice byte (file lưu dạng byte).
- `0666`: quyền file dạng bát phân — mọi người đọc/ghi được.
- Kiểu trả về `error`: `nil` = thành công.
- `ioutil` đã deprecated từ Go 1.16 → dùng `os.WriteFile` / `os.ReadFile` (cách dùng giống hệt).

## 7. Xử lý lỗi

```go
func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename) // trả về (nội dung []byte, error)
    if err != nil {                      // mẫu kiểm tra lỗi chuẩn của Go
        fmt.Println("Error:", err)
        os.Exit(1)                       // thoát, mã khác 0 = có lỗi
    }
    s := strings.Split(string(bs), ",")  // ngược lại với Join
    return deck(s)
}
```

- Go không có `try/catch`: lỗi là một giá trị trả về, kiểm tra ngay sau khi gọi hàm.
- `main.go` cố ý đọc file `"my"` không tồn tại để chạy vào nhánh lỗi.

## 8. Xáo bài và viết test (assert)

**Xáo bài — hoán đổi phần tử trong slice:**

```go
func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano()) // seed theo thời gian
    r := rand.New(source)
    for i := range d {                              // chỉ lấy chỉ số
        newPosition := r.Intn(len(d) - 1)           // số ngẫu nhiên 0..n-1
        d[i], d[newPosition] = d[newPosition], d[i] // hoán đổi trong 1 dòng
    }
}
```

- Không cần trả về: slice tham chiếu tới mảng bên dưới, sửa `d[i]` là sửa bộ bài gốc.
- Lưu ý: `r.Intn(len(d) - 1)` không bao giờ chọn lá cuối; `r.Intn(len(d))` xáo đều hơn.

**Viết test — file `deck_test.go`:**

```go
import "testing"

func TestNewDeck(t *testing.T) {   // tên bắt đầu bằng Test, nhận t *testing.T
    d := newDeck()
    if len(d) != 16 {               // Go không có assert → tự viết if
        t.Errorf("Expected deck length of 16, but got %v", len(d))
    }
}
```

- File test phải có đuôi `_test.go`; cần có `go.mod` (tạo bằng `go mod init <tên>`).
- `t.Errorf` đánh dấu test thất bại nhưng vẫn chạy tiếp; `%v` in giá trị dạng mặc định.

## 9. File I/O và test cho file

```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")                    // dọn file cũ còn sót
    deck := newDeck()
    deck.saveToFile("_decktesting")              // ghi
    loadedDeck := newDeckFromFile("_decktesting") // đọc lại
    if len(loadedDeck) != 16 {
        t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
    }
    os.Remove("_decktesting")                    // dọn file tạm sau test
}
```

- Luồng dữ liệu: `deck` → `toString` → `[]byte` → **file** → `[]byte` → `string` → `Split` → `deck`.
- `d[len(d)-1]` lấy phần tử cuối (Go không hỗ trợ chỉ số âm).
