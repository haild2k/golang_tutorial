# 5. Maps — Tóm tắt cú pháp

Chương này dùng **map** để lưu các cặp `key → value` (tên màu → mã hex): tạo map, thêm/xoá phần tử, duyệt map và so sánh với struct.

Chạy một bài: `cd <folder> && go run main.go`

## Liên tưởng với Python

| Go | Python tương tự | Khác biệt cần nhớ |
|---|---|---|
| `map[string]string{"red": "#ff0000"}` | dict `{"red": "#ff0000"}` | Go: mọi key cùng kiểu, mọi value cùng kiểu; Python trộn được |
| `make(map[int]string)` | `{}` hoặc `dict()` | Map `nil` (`var m map[...]`) ghi vào sẽ panic — Python không có khái niệm này |
| `m[k] = v` | `d[k] = v` | Giống nhau |
| `delete(m, k)` | `del d[k]` | `del` báo `KeyError` nếu key không có; `delete` thì không |
| `m["không_có"]` → zero value (`""`) | `d["không_có"]` → `KeyError` | Go giống `defaultdict` hoặc `d.get(k, "")` |
| `v, ok := m[k]` | `k in d` / `d.get(k)` | — |
| `for k, v := range m` | `for k, v in d.items()` | — |
| `for k := range m` | `for k in d` | — |
| Thứ tự duyệt **ngẫu nhiên** | Giữ **thứ tự thêm vào** (Python 3.7+) | Đừng mang thói quen Python sang Go |
| `len(m)` | `len(d)` | Giống nhau |
| map vs struct | dict vs dataclass | Cùng ý tưởng: dict/map cho dữ liệu động, dataclass/struct cho cấu trúc cố định |

| # | Folder | Chủ đề chính |
|---|---|---|
| 1 | [1_map](1_map/) | Khai báo map, map literal |
| 2 | [2_manipulating_map](2_manipulating_map/) | `make`, thêm phần tử, `delete` |
| 3 | [3_iterating_over_maps](3_iterating_over_maps/) | Duyệt map bằng `range`, map làm tham số |
| 4 | [4_differences_map_and_structs](4_differences_map_and_structs/) | Khác nhau giữa map và struct |

---

## 1. Khai báo map

```go
colors := map[string]string{   // map[KIỂU_KEY]KIỂU_VALUE
    "red":   "#ff0000",
    "green": "#4bf745",         // viết nhiều dòng → cặp cuối CŨNG phải có dấu phẩy
}
fmt.Println(colors)             // map[green:#4bf745 red:#ff0000]
```

- Mọi key cùng kiểu, mọi value cùng kiểu; key không được trùng.

## 2. Thao tác với map

```go
colors := make(map[int]string) // tạo map rỗng
colors[10] = "#ffffff"         // thêm (hoặc ghi đè nếu key đã có)
delete(colors, 10)             // xoá theo key
fmt.Println(colors)            // map[]
```

| Cách tạo | Kết quả |
|---|---|
| `make(map[int]string)` | map rỗng, ghi được |
| `map[int]string{}` | map rỗng, ghi được (tương đương `make`) |
| `var m map[int]string` | map `nil`: đọc được, **ghi vào sẽ panic** |

- Map chỉ truy cập bằng `m[key]`, không dùng dấu chấm như struct.
- `delete` với key không tồn tại thì không làm gì, không báo lỗi.
- Đọc key không có → trả về zero value của kiểu value (vd `""`). Muốn biết key có tồn tại không: `hex, ok := colors["red"]` (`ok` là `bool`).

## 3. Duyệt map

```go
func printMap(c map[string]string) {   // tham số kiểu map
    for color, hex := range c {         // color = key, hex = value
        fmt.Println("Hex code for", color, "is", hex)
    }
}
```

| Cú pháp | Lấy gì |
|---|---|
| `for k, v := range m` | cả key và value |
| `for k := range m` | chỉ key |
| `for _, v := range m` | chỉ value |

- **Thứ tự duyệt map là ngẫu nhiên**, mỗi lần chạy có thể khác → không được dựa vào thứ tự.

## 4. Map khác struct thế nào

```go
colors["yellow"] = "laksjdf"   // thêm key lúc chạy — struct không làm được
delete(colors, "yellow")       // xoá key lúc chạy — struct không làm được
printMap(colors)               // map là kiểu tham chiếu → hàm nhận CÙNG map
```

| | Map | Struct |
|---|---|---|
| Kiểu dữ liệu | Mọi key cùng kiểu, mọi value cùng kiểu | Mỗi trường có thể khác kiểu |
| Key / trường | Thêm/xoá được lúc chạy | Cố định khi viết code |
| Duyệt bằng `range` | Được | Không |
| Truy cập | `m["key"]` | `s.field` |
| Truyền vào hàm | Kiểu tham chiếu → sửa được map gốc | Kiểu giá trị → bản sao, muốn sửa gốc phải dùng con trỏ |
| Dùng khi | Tập dữ liệu cùng loại, chưa biết trước số lượng | Mô tả một "vật" có nhiều thuộc tính khác nhau (vd `person`) |
