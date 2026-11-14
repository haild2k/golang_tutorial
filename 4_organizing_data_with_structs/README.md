# 4. Organizing Data with Structs — Tóm tắt cú pháp

Chương này dùng **struct** để gom dữ liệu liên quan (thông tin một người), rồi học cách sửa struct qua method và **con trỏ**.

Chạy một bài: `cd <folder> && go run main.go`

## Liên tưởng với Python

| Go | Python tương tự | Khác biệt cần nhớ |
|---|---|---|
| `type person struct { firstName string }` | `@dataclass class Person: first_name: str` | Kiểu của trường trong Go được kiểm tra lúc biên dịch |
| `person{firstName: "Alex"}` | `Person(first_name="Alex")` | Go không cần viết hàm khởi tạo `__init__` |
| `var alex person` (zero value) | không có tương đương | Python bắt buộc truyền giá trị hoặc đặt mặc định (`None`, `""`) |
| `alex.firstName = "Alex"` | `alex.first_name = "Alex"` | Go không thêm được trường mới lúc chạy; Python thì được |
| `fmt.Printf("%+v", alex)` | `print(alex)` (dataclass tự có `__repr__`) | — |
| Nhúng struct (`contactInfo` trong `person`) | Composition: `self.contact = ContactInfo(...)` | Go "đẩy" trường lên: `jim.email` thay vì `jim.contact.email` — trông giống kế thừa nhưng **không phải kế thừa** |
| `func (p person) updateName()` (value receiver) | không có tương đương | `p` là **bản sao** → sửa không ảnh hưởng gốc. Trong Python `self` luôn trỏ tới object gốc |
| `func (p *person) updateName()` (pointer receiver) | `def update_name(self):` bình thường | Đây mới là cách hoạt động giống `self` trong Python |
| `&jim`, `*p` (con trỏ) | không có | Python không có con trỏ; mọi biến đều là tham chiếu tới object |
| Truyền theo giá trị (copy struct) | Truyền theo tham chiếu object | Muốn copy trong Python phải dùng `copy.copy()` |

| # | Folder | Chủ đề chính |
|---|---|---|
| 1 | [1_declaring_structs](1_declaring_structs/) | Khai báo struct, struct literal |
| 2 | [2_update_struct_values](2_update_struct_values/) | Zero value, gán trường, `%+v` |
| 3 | [3_embedding_structs](3_embedding_structs/) | Nhúng struct vào struct |
| 4 | [4_structs_with_receiver_functions](4_structs_with_receiver_functions/) | Method của struct, value receiver |
| 5 | [5_pointer_operations](5_pointer_operations/) | `&`, `*`, pointer receiver |
| 6 | [6_pointer_shortcut](6_pointer_shortcut/) | Go tự lấy địa chỉ / tự dereference |

---

## 1. Khai báo struct

```go
type person struct {
    firstName string   // tên_trường kiểu — không có dấu phẩy giữa các trường
    lastName  string
}

alex := person{firstName: "Alex", lastName: "Anderson"} // gán theo tên trường
fmt.Println(alex)                                       // {Alex Anderson}
```

- Có thể viết theo thứ tự `person{"Alex", "Anderson"}` nhưng dễ sai khi thêm/đổi trường → nên ghi tên.

## 2. Cập nhật giá trị struct

```go
var alex person            // chưa gán → mọi trường nhận zero value
alex.firstName = "Alex"    // truy cập / gán trường bằng dấu chấm
alex.lastName = "Anderson"
fmt.Printf("%+v", alex)    // {firstName:Alex lastName:Anderson}
```

| Kiểu | Zero value |
|---|---|
| `string` | `""` |
| `int`, `float64` | `0` |
| `bool` | `false` |
| con trỏ, slice, map | `nil` |

| Định dạng | Kết quả |
|---|---|
| `%v` | `{Alex Anderson}` |
| `%+v` | `{firstName:Alex lastName:Anderson}` |

## 3. Nhúng struct (embedding)

```go
type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contactInfo          // chỉ ghi tên kiểu → nhúng; tên trường = tên kiểu
}

jim := person{
    firstName: "Jim",
    lastName:  "Party",
    contactInfo: contactInfo{
        email:   "jim@gmail.com",
        zipCode: 94000,  // viết nhiều dòng → dòng cuối CŨNG phải có dấu phẩy
    },
}
```

- Trường của struct nhúng được "đẩy lên": `jim.email` = `jim.contactInfo.email`.
- Nếu viết `contact contactInfo` (có tên trường) thì là trường thường → phải gọi `jim.contact.email`.

## 4. Struct với receiver function

```go
func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName   // chỉ sửa BẢN SAO
}

func (p person) print() {
    fmt.Printf("%+v", p)
}

jim.updateName("jimmy")
jim.print()                      // firstName vẫn là "Jim"
```

- Go **truyền theo giá trị (pass by value)**: value receiver `p` là bản sao của `jim`.
- Value receiver phù hợp cho method chỉ đọc (như `print`); muốn sửa struct → dùng con trỏ (bài 5).

## 5. Các phép toán con trỏ

| Cú pháp | Ý nghĩa |
|---|---|
| `&biến` | lấy địa chỉ của biến → tạo con trỏ |
| `*con_trỏ` | lấy giá trị tại địa chỉ (dereference) |
| `*person` | (trong khai báo kiểu) "con trỏ tới person" |

```go
jimPointer := &jim                 // kiểu *person
jimPointer.updateName("jimmy")

func (pointerToPerson *person) updateName(newFirstName string) { // pointer receiver
    (*pointerToPerson).firstName = newFirstName // sửa struct GỐC
}
```

- Cần ngoặc `(*p).firstName` vì dấu `.` ưu tiên cao hơn dấu `*`.
- Kết quả: `firstName` đổi thành `"jimmy"`.

## 6. Cú pháp rút gọn khi dùng con trỏ

```go
jim.updateName("jimmy")            // Go tự đổi thành (&jim).updateName("jimmy")

func (p *person) updateName(newFirstName string) {
    p.firstName = newFirstName     // Go tự hiểu là (*p).firstName
}
```

- Method có receiver `*person` gọi được bằng cả biến `person` lẫn con trỏ `*person`.
- Không áp dụng cho giá trị tạm không có địa chỉ, ví dụ `person{}.updateName(...)` sẽ lỗi biên dịch.

---

## Khi nào dùng value receiver hay pointer receiver?

| Value receiver `(p person)` | Pointer receiver `(p *person)` |
|---|---|
| Chỉ đọc dữ liệu | Cần sửa struct gốc |
| Struct nhỏ, sao chép rẻ | Struct lớn, tránh sao chép |
| — | Một method đã dùng con trỏ thì các method khác nên dùng con trỏ cho thống nhất |
