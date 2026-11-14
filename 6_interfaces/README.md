# 6. Interfaces — Giải thích dễ hiểu

Chạy một bài: `cd <folder> && go run main.go` (bài 2–5 cần mạng để gọi `http://google.com`)

| # | Folder | Học gì |
|---|---|---|
| 1 | [1_interfaces](1_interfaces/) | Interface là gì |
| 2 | [2_http_package](2_http_package/) | Gửi request lên web bằng `http.Get` |
| 3 | [3_read_function](3_read_function/) | Đọc nội dung trả về bằng `Read` |
| 4 | [4_writer_interface](4_writer_interface/) | Chuyển dữ liệu từ nguồn sang đích bằng `io.Copy` |
| 5 | [5_custom_writer](5_custom_writer/) | Tự làm một "đích" của riêng mình |

---

## Ý tưởng chính trong một câu

> **Interface là một bản mô tả công việc**: "Ai làm được việc X thì được nhận."
> Go không hỏi bạn *là ai*, chỉ hỏi bạn *làm được gì*.

### Ví dụ đời thường: ổ cắm điện

Ổ cắm trên tường không quan tâm bạn cắm **quạt, sạc điện thoại hay nồi cơm**.
Chỉ cần phích cắm **đúng hình dạng** (2 chân, đúng khoảng cách) là cắm được.

| Đời thường | Trong Go |
|---|---|
| Ổ cắm quy định hình dạng phích | `interface` quy định danh sách method |
| Quạt, sạc, nồi cơm | Các kiểu `struct` khác nhau |
| Phích cắm đúng hình dạng | Kiểu có đủ method, đúng tên, đúng tham số và kiểu trả về |
| Không cần đăng ký với nhà điện | Không cần viết `implements` — có method đúng là tự động hợp lệ |

---

## 1. Interface — "bản mô tả công việc"

### Hình dung

Công ty đăng tin tuyển: **"Cần người biết chào khách."**
- Anh người Anh nói *"Hi there!"* → đạt.
- Chị người Tây Ban Nha nói *"Hola!"* → đạt.

Quản lý chỉ cần gọi *"Chào khách đi!"*, không cần biết người đứng đó là ai. Mỗi người tự chào theo cách của mình.

### Code

```go
// Bản mô tả công việc: ai có method getGreeting() trả về string thì là "bot"
type bot interface {
    getGreeting() string
}

type englishBot struct{}   // anh người Anh
type spanishBot struct{}   // chị người Tây Ban Nha

func (englishBot) getGreeting() string { return "Hi there!" }
func (spanishBot) getGreeting() string { return "Hola!" }

// Quản lý: nhận BẤT KỲ ai làm được việc "bot"
func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

printGreeting(englishBot{})   // Hi there!
printGreeting(spanishBot{})   // Hola!
```

### Chuyện gì xảy ra

1. `englishBot` có method `getGreeting() string` → Go tự công nhận nó là `bot`.
2. `spanishBot` cũng có → cũng là `bot`.
3. `printGreeting` chỉ viết **một lần**, dùng được cho cả hai.
4. Khi gọi `b.getGreeting()`, Go tự chọn đúng phiên bản của kiểu thật đang được truyền vào.

### Nếu không có interface?

Phải viết hai hàm gần giống hệt nhau:

```go
func printGreetingEnglish(eb englishBot) { fmt.Println(eb.getGreeting()) }
func printGreetingSpanish(sb spanishBot) { fmt.Println(sb.getGreeting()) }
```

Thêm `frenchBot` lại phải viết thêm hàm thứ ba → interface giúp tránh việc lặp lại này.

### Nhớ

- `struct{}` là struct rỗng, chỉ dùng để gắn method lên.
- `func (englishBot) ...` không đặt tên receiver vì bên trong không dùng tới nó.

---

## 2. Gửi request lên web — `http.Get`

### Hình dung

Bạn **gửi thư** tới google.com và nhận về **một phong bì**.
- Bên ngoài phong bì ghi thông tin: trạng thái (200 OK), loại nội dung, ngày gửi…
- Bên trong là **lá thư thật** (nội dung trang web), nhưng **phong bì còn dán kín**.

### Code

```go
resp, err := http.Get("http://google.com")
if err != nil {                 // gửi thư thất bại (mất mạng, sai địa chỉ…)
    fmt.Println("Error:", err)
    os.Exit(1)
}
fmt.Println(resp)               // chỉ thấy thông tin ngoài phong bì
```

### Chuyện gì xảy ra

- In `resp` ra chỉ thấy status, header… **không thấy nội dung HTML**.
- Nội dung nằm trong `resp.Body` — "lá thư còn dán kín". Muốn xem phải **mở và đọc** (bài 3).

---

## 3. Đọc nội dung — `Read` và `io.Reader`

### Hình dung: rót nước từ bình ra cốc

- `resp.Body` là **bình nước** (nguồn dữ liệu).
- Bạn phải **chuẩn bị sẵn một cái cốc** đủ to.
- Gọi `Read` = **rót nước từ bình vào cốc**.

```
  resp.Body (bình)  ──Read──▶  bs (cốc 99999 byte)  ──string()──▶  in ra màn hình
```

### Code

```go
bs := make([]byte, 99999)   // chuẩn bị cái cốc chứa được 99999 byte
resp.Body.Read(bs)          // rót dữ liệu vào cốc
fmt.Println(string(bs))     // đổi byte thành chữ rồi in
```

### `io.Reader` là gì?

Là một interface (bản mô tả công việc) với đúng một yêu cầu:

```go
type Reader interface {
    Read(p []byte) (n int, err error)   // "rót dữ liệu vào cốc p, báo lại rót được bao nhiêu"
}
```

Ai có method `Read` như trên đều là **một cái bình** — dù là trang web, file trên ổ cứng, hay một chuỗi trong bộ nhớ.

### Lưu ý

- Cốc phải chuẩn bị **trước**; `Read` không tự làm cốc to ra.
- Một lần rót **có thể chưa hết bình** → bài 4 có cách tiện hơn.

---

## 4. Chuyển dữ liệu từ nguồn sang đích — `io.Copy` và `io.Writer`

### Hình dung: nối ống nước

Thay vì tự rót từng cốc, ta **nối một đường ống** từ bình (nguồn) thẳng tới bồn rửa (đích).
Nước tự chảy cho đến khi **hết bình**.

```
  resp.Body (nguồn, io.Reader)  ═══ io.Copy ═══▶  os.Stdout (đích, io.Writer = màn hình)
```

### Code

```go
io.Copy(os.Stdout, resp.Body)   // io.Copy(ĐÍCH, NGUỒN)
```

### Hai "bản mô tả công việc" ngược chiều nhau

| | Vai trò | Ví dụ đời thường | Ví dụ trong Go |
|---|---|---|---|
| `io.Reader` | **Cho ra** dữ liệu | Vòi nước, bình nước | `resp.Body`, file đang đọc |
| `io.Writer` | **Nhận vào** dữ liệu | Bồn rửa, xô chứa | `os.Stdout` (màn hình), file đang ghi |

```go
type Writer interface {
    Write(p []byte) (n int, err error)  // "nhận dữ liệu p, báo lại đã nhận bao nhiêu"
}
```

### Vì sao hay?

`io.Copy` không cần biết nguồn là gì, đích là gì — chỉ cần **nguồn biết `Read`, đích biết `Write`**.
Giống ống nước tiêu chuẩn: vòi nào cũng nối được, bồn nào cũng nối được.

- web → màn hình: `io.Copy(os.Stdout, resp.Body)`
- web → file: `io.Copy(file, resp.Body)`
- file → file: `io.Copy(fileMoi, fileCu)`

---

## 5. Tự làm một "đích" — custom Writer

### Hình dung

Bạn tự làm **một cái xô có gắn đồng hồ đo**: nước đổ vào thì vừa hứng, vừa ghi lại đã đổ bao nhiêu lít.
Chỉ cần xô có **miệng đúng kích cỡ ống tiêu chuẩn** (method `Write` đúng chữ ký) là nối vào `io.Copy` được.

### Code

```go
type logWriter struct{}          // cái xô tự chế

// Miệng xô đúng chuẩn io.Writer
func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))                              // hứng nước: in nội dung
    fmt.Println("Just wrote this many bytes:", len(bs))  // đồng hồ đo: in số byte
    return len(bs), nil                                  // báo: đã nhận đủ, không lỗi
}

io.Copy(logWriter{}, resp.Body)   // nối ống từ web vào xô tự chế
```

### Chuyện gì xảy ra

1. `io.Copy` đọc một đoạn dữ liệu từ `resp.Body`.
2. Gọi `logWriter.Write(đoạn đó)` → in nội dung và số byte.
3. Lặp lại đến khi hết dữ liệu → có thể thấy dòng "Just wrote…" in ra nhiều lần.

### Nhớ

- Chữ ký phải khớp **chính xác**: tên `Write` (viết hoa), nhận `[]byte`, trả về `(int, error)`.
- Phải trả về **đúng số byte đã nhận** (`len(bs)`). Báo thiếu mà không kèm lỗi → `io.Copy` báo lỗi `short write`.

---

## Liên tưởng với Python

| Go | Python tương tự | Ghi chú |
|---|---|---|
| `interface` + implement ngầm | **Duck typing**: "kêu như vịt, đi như vịt thì là vịt" | Go kiểm tra **lúc biên dịch**; Python chỉ phát hiện lỗi **lúc chạy** |
| `type bot interface {...}` | `class Bot(typing.Protocol): ...` | Cùng ý tưởng "xét theo hình dạng" |
| `type englishBot struct{}` | `class EnglishBot: pass` | — |
| `http.Get(url)` | `requests.get(url)` | Go trả về `(resp, err)` thay vì ném exception |
| `resp.Body` (`io.Reader`) | Object giống file, có `.read()` | — |
| `io.Copy(os.Stdout, r)` | `shutil.copyfileobj(r, sys.stdout.buffer)` | — |
| Custom `Write` | Class có `write()` rồi `print(..., file=obj)` | Go phải khớp chữ ký chính xác |
| `defer resp.Body.Close()` | `with urlopen(url) as resp:` | Đóng kết nối khi dùng xong |

---

## Tóm tắt bằng hình ảnh

| Khái niệm | Hình ảnh đời thường |
|---|---|
| `interface` | Ổ cắm điện / bản mô tả công việc |
| Kiểu thoả interface | Phích cắm đúng hình dạng / ứng viên làm được việc |
| `http.Get` | Gửi thư, nhận về phong bì dán kín |
| `io.Reader` | Bình nước, vòi nước — **cho ra** dữ liệu |
| `io.Writer` | Bồn rửa, xô — **nhận vào** dữ liệu |
| `Read(bs)` | Rót nước vào cốc chuẩn bị sẵn |
| `io.Copy(dst, src)` | Nối ống nước, nước tự chảy đến hết |
| Custom Writer | Xô tự chế có gắn đồng hồ đo |
