# 7. Channels và Go Routines — Giải thích dễ hiểu

Chạy một bài: `cd <folder> && go run main.go` (cần mạng; bài 5 chạy mãi, dừng bằng `Ctrl+C`)

| # | Folder | Học gì |
|---|---|---|
| 1 | [1_site_status](1_site_status/) | Kiểm tra từng website một (chậm) |
| 2 | [2_channels](2_channels/) | Goroutine: làm nhiều việc cùng lúc |
| 3 | [3_channels_implementation](3_channels_implementation/) | Channel: cách các goroutine báo kết quả |
| 4 | [4_receiving_mesages](4_receiving_mesages/) | Đợi nhận đủ tất cả kết quả |
| 5 | [5_channels_gotcha](5_channels_gotcha/) | Kiểm tra liên tục và các lỗi hay gặp |

---

## Câu chuyện xuyên suốt: quán ăn

Cả chương có thể hình dung như **một quán ăn**:

| Trong quán ăn | Trong Go |
|---|---|
| **Chủ quán** | Hàm `main` (main goroutine) |
| **Đầu bếp được thuê thêm** | Goroutine (tạo bằng từ khoá `go`) |
| **Ô cửa chuyền món** giữa bếp và quầy | Channel |
| Đầu bếp đặt món lên ô cửa | Gửi: `c <- giá_trị` |
| Chủ quán lấy món từ ô cửa | Nhận: `<-c` |
| **Chủ quán về nhà → đóng cửa quán** | `main` kết thúc → chương trình thoát, mọi goroutine bị dừng |

Việc cần làm: kiểm tra 5 website (google, facebook, …) còn sống hay không. Mỗi website giống **một món ăn** cần nấu.

---

## 1. Làm từng việc một — chậm

### Hình dung

Quán chỉ có **một mình chủ quán**. Nấu xong món 1 mới nấu món 2, xong món 2 mới tới món 3…
Món nào lâu (website phản hồi chậm) thì cả hàng phải đợi.

```
Chủ quán: [google ████] [facebook ██████] [stackoverflow ███] [golang ██] [amazon █████]
          ─────────────────────────── thời gian ───────────────────────────────▶
```

### Code

```go
for _, link := range links {
    checkLink(link)          // đợi link này xong mới sang link sau
}
```

### Nhớ

- `http.Get` phải **đợi** website trả lời mới chạy tiếp (gọi là *blocking*).
- Tổng thời gian = cộng dồn thời gian từng website → chậm.

---

## 2. Goroutine — thuê thêm đầu bếp

### Hình dung

Chủ quán **thuê 5 đầu bếp**, mỗi người nấu một món, **tất cả nấu cùng lúc**.

```
Đầu bếp 1: [google ████]
Đầu bếp 2: [facebook ██████]
Đầu bếp 3: [stackoverflow ███]
Đầu bếp 4: [golang ██]
Đầu bếp 5: [amazon █████]
           ──── thời gian ────▶   (nhanh hơn nhiều)
```

### Code

```go
for _, link := range links {
    go checkLink(link)       // "go" = thuê một đầu bếp làm việc này, KHÔNG đứng đợi
}
// hết hàm main
```

### Nhưng chạy thử thì… không in ra gì!

Vì sao? Chủ quán thuê xong 5 đầu bếp thì **về nhà luôn và khoá cửa quán**.
Đầu bếp chưa kịp nấu xong món nào đã bị đuổi ra ngoài.

```
Chủ quán:  thuê, thuê, thuê, thuê, thuê → VỀ NHÀ (khoá cửa) ✖
Đầu bếp 1: [goo… ✖ bị dừng
Đầu bếp 2: [fac… ✖ bị dừng
```

### Nhớ

- `go f(x)` chạy `f` ở một goroutine mới và **không đợi** nó xong.
- `main` kết thúc là chương trình thoát, **không chờ** goroutine nào cả.
- → Cần một cách để chủ quán **đứng đợi** món ăn: đó là **channel**.

---

## 3. Channel — ô cửa chuyền món

### Hình dung

Giữa bếp và quầy có **một ô cửa nhỏ**. Đầu bếp nấu xong thì **đặt món lên ô cửa**. Chủ quán đứng ở quầy **chờ lấy món**.

Quy tắc của ô cửa này (*channel không bộ đệm*):
- Ô cửa **chỉ để vừa tay** — đầu bếp phải **đưa tận tay** chủ quán.
- Chủ quán ra lấy mà chưa có món → **đứng chờ**.
- Đầu bếp mang món ra mà chủ quán chưa tới → **đứng chờ**.

### Code

```go
c := make(chan string)            // làm một ô cửa, chỉ chuyền được "string"

for _, link := range links {
    go checkLink(link, c)          // thuê đầu bếp, chỉ cho họ ô cửa c
}

fmt.Println(<-c)                   // chủ quán đứng ở ô cửa, chờ lấy MỘT món

func checkLink(link string, c chan string) {
    // ... nấu xong ...
    c <- "Yep its up"              // đầu bếp đặt món lên ô cửa
}
```

### Ba ký hiệu cần thuộc

| Ký hiệu | Đọc là | Quán ăn |
|---|---|---|
| `make(chan string)` | tạo channel chở `string` | Làm một ô cửa chuyền món |
| `c <- "abc"` | gửi `"abc"` vào `c` | Đầu bếp đặt món lên ô cửa |
| `x := <-c` | nhận từ `c`, gán cho `x` | Chủ quán lấy món từ ô cửa |

> Mẹo nhớ: **mũi tên chỉ hướng dữ liệu chạy**. `c <- v`: dữ liệu chạy **vào** `c`. `<-c`: dữ liệu chạy **ra khỏi** `c`.

### Chạy thử thì chỉ thấy 1 kết quả!

Chủ quán chỉ **lấy đúng 1 món** (`<-c` một lần) rồi **về nhà**. 4 đầu bếp còn lại bị đuổi về.

```
Đầu bếp 4: [golang ██] ──▶ ô cửa ──▶ Chủ quán lấy 1 món → VỀ NHÀ ✖
Đầu bếp 1,2,3,5: … ✖ bị dừng
```

---

## 4. Đợi nhận đủ tất cả kết quả

### Hình dung

Chủ quán biết đã thuê **5 đầu bếp**, mỗi người làm **đúng 1 món** → đứng ở ô cửa **lấy đủ 5 món** rồi mới về.

### Code

```go
for i := 0; i < len(links); i++ {   // lặp 5 lần
    fmt.Println(<-c)                // mỗi lần lấy một món
}
```

### Chuyện gì xảy ra

```
Đầu bếp 4: golang xong trước   ──▶ chủ quán lấy món 1
Đầu bếp 3: stackoverflow xong  ──▶ chủ quán lấy món 2
Đầu bếp 1: google xong         ──▶ chủ quán lấy món 3
Đầu bếp 5: amazon xong         ──▶ chủ quán lấy món 4
Đầu bếp 2: facebook xong       ──▶ chủ quán lấy món 5 → về nhà ✔
```

### Nhớ

- **Gửi bao nhiêu thì nhận bấy nhiêu.**
- Món nào nấu xong trước thì tới trước — **không theo thứ tự** trong danh sách.
- Nếu chủ quán chờ món thứ 6 mà không còn đầu bếp nào → đứng chờ mãi. Go phát hiện và báo lỗi:
  `fatal error: all goroutines are asleep - deadlock!`

---

## 5. Kiểm tra liên tục và các lỗi hay gặp

### Hình dung

Quán chuyển sang **phục vụ liên tục**: mỗi khi nhận một món từ ô cửa, chủ quán **giao lại việc cho một đầu bếp**: "Nghỉ 5 giây rồi nấu lại món này nhé."
Như vậy mỗi website được kiểm tra lại cứ mỗi 5 giây, mãi mãi.

### Code

```go
for l := range c {                 // cứ có món ở ô cửa là lấy, lặp mãi
    go func(link string) {         // thuê một đầu bếp tạm cho việc này
        time.Sleep(5 * time.Second) // đầu bếp nghỉ 5 giây
        checkLink(link, c)          // rồi kiểm tra lại link
    }(l)                           // đưa tờ giấy ghi tên link cho đầu bếp
}
```

Đầu bếp gửi lại **chính tên link** (`c <- link`) để chủ quán biết món nào vừa xong.

### Cú pháp mới

| Cú pháp | Ý nghĩa | Quán ăn |
|---|---|---|
| `for l := range c` | Lặp mãi: mỗi lần channel có giá trị thì gán vào `l` | Chủ quán đứng ở ô cửa cả ngày |
| `func(link string) {...}(l)` | Hàm không tên, định nghĩa và gọi ngay | Viết nhanh một tờ giao việc rồi đưa luôn |
| `5 * time.Second` | Khoảng thời gian 5 giây | — |

### Lỗi 1: chủ quán đi ngủ

❌ **Sai**: đặt `time.Sleep` ngay trong vòng lặp của `main`:

```go
for l := range c {
    time.Sleep(5 * time.Second)   // CHỦ QUÁN ngủ 5 giây
    go checkLink(l, c)
}
```

Chủ quán ngủ thì **không ai đứng ở ô cửa**. Các đầu bếp khác nấu xong cũng phải **xếp hàng chờ** → cả quán chậm lại.

✅ **Đúng**: để **đầu bếp** nghỉ 5 giây (Sleep nằm trong goroutine con). Chủ quán luôn sẵn sàng ở ô cửa.

### Lỗi 2: cùng nhìn một tấm bảng

❌ **Sai**: đầu bếp đọc trực tiếp biến `l`:

```go
for l := range c {
    go func() {
        time.Sleep(5 * time.Second)
        checkLink(l, c)           // đọc l SAU 5 giây
    }()
}
```

Hình dung `l` là **tấm bảng trắng treo ở quầy**, chủ quán liên tục xoá đi viết tên món mới.
Đầu bếp ngủ dậy sau 5 giây mới **ngẩng lên nhìn bảng** → lúc đó bảng đã ghi món khác → nấu nhầm món.

✅ **Đúng**: chủ quán **đưa cho mỗi đầu bếp một tờ giấy riêng** ghi tên món — chính là truyền `l` vào tham số `link`:

```go
go func(link string) { ... }(l)   // mỗi đầu bếp giữ bản sao riêng
```

> Từ Go 1.22, mỗi vòng lặp tự có biến riêng nên ít bị lỗi này hơn. Nhưng truyền qua tham số vẫn là cách rõ ràng nhất.

> **Nguyên tắc vàng:** đừng để nhiều goroutine cùng nhìn một biến. Hãy **đưa bản sao** (tham số) hoặc **chuyền qua ô cửa** (channel).

---

## Liên tưởng với Python

| Go | Python tương tự | Ghi chú |
|---|---|---|
| Goroutine | `threading.Thread` / `asyncio` task | Goroutine rất nhẹ (tạo hàng nghìn cái được) và chạy song song thật trên nhiều CPU |
| `go f(x)` | `threading.Thread(target=f, args=(x,)).start()` | Go chỉ cần một từ khoá |
| `main` thoát → goroutine dừng | Giống **daemon thread** | Thread thường trong Python thì chương trình **đợi** nó xong |
| Channel | `queue.Queue()` | Channel không bộ đệm bắt "đưa tận tay"; `Queue()` thì đặt vào là đi luôn |
| `c <- v` / `v := <-c` | `q.put(v)` / `v = q.get()` | — |
| `for v := range c` | `while True: v = q.get()` | — |
| `go func(link string){...}(l)` | Hàm lồng `def` / `lambda` | — |
| Lỗi 2 (cùng nhìn một biến) | Lỗi *late binding* của `lambda` trong vòng lặp | Cách sửa giống nhau: truyền qua tham số (`lambda l=l: ...`) |
| Báo `deadlock!` | `q.get()` treo mãi | Go tự phát hiện, Python thì không |

---

## Tóm tắt bằng hình ảnh

| Khái niệm | Hình ảnh |
|---|---|
| `main` | Chủ quán |
| `go f()` | Thuê đầu bếp làm việc `f`, không đứng đợi |
| `main` kết thúc | Chủ quán về nhà, khoá cửa — mọi đầu bếp phải nghỉ |
| `make(chan T)` | Làm ô cửa chuyền món loại `T` |
| `c <- v` | Đầu bếp đặt món lên ô cửa |
| `<-c` | Chủ quán lấy món (chưa có thì đứng chờ) |
| Nhận đủ `len(links)` lần | Thuê bao nhiêu đầu bếp thì lấy bấy nhiêu món |
| `for l := range c` | Chủ quán đứng ở ô cửa cả ngày |
| Sleep trong goroutine con | Đầu bếp nghỉ, chủ quán vẫn trực |
| Truyền tham số cho func literal | Đưa mỗi đầu bếp một tờ giấy ghi món riêng |
