
# Hello World trong Go

Nội dung file [main.go](main.go):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```

## Giải thích cú pháp

### 1. `package main`
- Mọi file Go đều thuộc về một **package**, khai báo ở dòng đầu tiên.
- `main` là package đặc biệt: báo cho Go biết đây là một **chương trình thực thi được** (executable), không phải thư viện.
- Nếu đặt tên package khác (ví dụ `package utils`), `go build` sẽ tạo ra thư viện dùng chung thay vì file chạy được.

### 2. `import "fmt"`
- Từ khóa `import` dùng để nạp package khác để sử dụng trong file này.
- `fmt` (format) là package có sẵn trong thư viện chuẩn, cung cấp các hàm vào/ra có định dạng (in ra màn hình, đọc input, ...).
- Go không cho phép import mà không dùng: nếu import `fmt` nhưng không gọi, trình biên dịch sẽ báo lỗi.
- Import nhiều package:
  ```go
  import (
      "fmt"
      "math"
  )
  ```

### 3. `func main() { ... }`
- `func` là từ khóa khai báo **hàm**.
- `main` là tên hàm; cùng với `package main`, đây là **điểm bắt đầu** của chương trình – Go tự động gọi hàm này khi chạy.
- `()` là danh sách tham số (ở đây rỗng); `main` không nhận tham số và không trả về giá trị.
- `{ }` bao thân hàm. Dấu `{` bắt buộc nằm cùng dòng với `func main()`.

### 4. `fmt.Println("Hi there!")`
- `fmt.Println` là gọi hàm `Println` của package `fmt`, dạng `tên_package.TênHàm`.
- Hàm `Println` in các đối số ra màn hình và tự thêm ký tự xuống dòng.
- `"Hi there!"` là chuỗi (string) đặt trong dấu nháy kép.
- Tên bắt đầu bằng **chữ hoa** (`Println`) nghĩa là *exported* – có thể dùng từ package khác. Tên chữ thường thì chỉ dùng nội bộ package.
- Go không cần dấu `;` cuối câu lệnh (trình biên dịch tự thêm).

## Chạy chương trình

```bash
cd 2_simple_start/helloworld
go run main.go     # biên dịch và chạy ngay
go build main.go   # chỉ biên dịch, tạo file thực thi
```

Kết quả:

```
Hi there!
```


## Vì sao có cả `go run` và `go build`?

Go là ngôn ngữ **biên dịch (compiled)**: mã nguồn `.go` phải được dịch thành **file nhị phân** (mã máy) thì máy mới chạy được. Hai lệnh phục vụ hai mục đích khác nhau:

| | `go run main.go` | `go build main.go` |
|---|---|---|
| Việc làm | Biên dịch, chạy ngay, rồi xóa file tạm | Chỉ biên dịch, giữ lại file thực thi (`main`) |
| Dùng khi | Phát triển, thử nhanh | Đóng gói, triển khai (deploy) |
| Kết quả | Chỉ thấy output | Có file chạy độc lập, ví dụ `./main` |

- `go run` thực chất là build vào thư mục tạm, chạy file đó, rồi xóa. Nó chỉ là lối tắt tiện cho lúc phát triển.
- File build ra là binary **độc lập**: máy chủ không cần cài Go, chỉ cần copy file sang là chạy.
- Có thể build chéo nền tảng, ví dụ `GOOS=windows go build`.

## So sánh với React

Có điểm giống về ý tưởng, nhưng bản chất khác nhau.

| | Go | React |
|---|---|---|
| Lệnh dev | `go run` | `npm start` / `npm run dev` |
| Lệnh đóng gói | `go build` | `npm run build` |
| `build` tạo ra | Binary mã máy | JS/CSS/HTML đã tối ưu (minify, bundle) |
| Chạy bởi | Hệ điều hành | Trình duyệt (hoặc Node) |

- **Giống:** đều có chế độ dev để thử nhanh và chế độ build để ra bản chạy thật.
- **Khác về bản chất:** JavaScript là ngôn ngữ thông dịch nên React vẫn chạy được mà không cần build; build chỉ để tối ưu và chuyển JSX/TypeScript sang JS. Với Go, biên dịch là bắt buộc, `go run` chỉ giấu bước đó đi.
- **Khác về dev server:** `npm start` có hot reload, còn `go run` chạy xong là thoát và không tự reload. Muốn tự reload thì dùng thêm công cụ như `air`.

## Vì sao có cả `go run` và `go build`?

Go là ngôn ngữ **biên dịch (compiled)**: mã nguồn `.go` phải được dịch thành **file nhị phân** (mã máy) thì máy mới chạy được. Hai lệnh phục vụ hai mục đích khác nhau:

| | `go run main.go` | `go build main.go` |
|---|---|---|
| Việc làm | Biên dịch, chạy ngay, rồi xóa file tạm | Chỉ biên dịch, giữ lại file thực thi (`main`) |
| Dùng khi | Phát triển, thử nhanh | Đóng gói, triển khai (deploy) |
| Kết quả | Chỉ thấy output | Có file chạy độc lập, ví dụ `./main` |

- `go run` thực chất là build vào thư mục tạm, chạy file đó, rồi xóa. Nó chỉ là lối tắt tiện cho lúc phát triển.
- File build ra là binary **độc lập**: máy chủ không cần cài Go, chỉ cần copy file sang là chạy.
- Có thể build chéo nền tảng, ví dụ `GOOS=windows go build`.

## So sánh với React

Có điểm giống về ý tưởng, nhưng bản chất khác nhau.

| | Go | React |
|---|---|---|
| Lệnh dev | `go run` | `npm start` / `npm run dev` |
| Lệnh đóng gói | `go build` | `npm run build` |
| `build` tạo ra | Binary mã máy | JS/CSS/HTML đã tối ưu (minify, bundle) |
| Chạy bởi | Hệ điều hành | Trình duyệt (hoặc Node) |

- **Giống:** đều có chế độ dev để thử nhanh và chế độ build để ra bản chạy thật.
- **Khác về bản chất:** JavaScript là ngôn ngữ thông dịch nên React vẫn chạy được mà không cần build; build chỉ để tối ưu và chuyển JSX/TypeScript sang JS. Với Go, biên dịch là bắt buộc, `go run` chỉ giấu bước đó đi.
- **Khác về dev server:** `npm start` có hot reload, còn `go run` chạy xong là thoát và không tự reload. Muốn tự reload thì dùng thêm công cụ như `air`.


## Go CLI

Go CLI là bộ lệnh `go` dùng để làm việc với mã nguồn Go. Các lệnh thường dùng:

| Lệnh | Chức năng |
|---|---|
| `go build` | Biên dịch một nhóm file mã nguồn Go (tạo file thực thi, không chạy) |
| `go run` | Biên dịch và chạy một hoặc hai file |
| `go fmt` | Định dạng lại code trong mọi file ở thư mục hiện tại |
| `go install` | Biên dịch và "cài đặt" một package |
| `go get` | Tải mã nguồn thô của package do người khác viết |
| `go test` | Chạy các bài test của project hiện tại |

Ví dụ:

```bash
go build main.go         # biên dịch, tạo file thực thi
go run main.go           # biên dịch và chạy
go fmt ./...             # format toàn bộ code
go install               # build và cài binary vào $GOPATH/bin
go get github.com/gin-gonic/gin   # tải package bên thứ ba
go test ./...            # chạy toàn bộ test
```

## Go CLI

Go CLI là bộ lệnh `go` dùng để làm việc với mã nguồn Go. Các lệnh thường dùng:

| Lệnh | Chức năng |
|---|---|
| `go build` | Biên dịch một nhóm file mã nguồn Go (tạo file thực thi, không chạy) |
| `go run` | Biên dịch và chạy một hoặc hai file |
| `go fmt` | Định dạng lại code trong mọi file ở thư mục hiện tại |
| `go install` | Biên dịch và "cài đặt" một package |
| `go get` | Tải mã nguồn thô của package do người khác viết |
| `go test` | Chạy các bài test của project hiện tại |

Ví dụ:

```bash
go build main.go         # biên dịch, tạo file thực thi
go run main.go           # biên dịch và chạy
go fmt ./...             # format toàn bộ code
go install               # build và cài binary vào $GOPATH/bin
go get github.com/gin-gonic/gin   # tải package bên thứ ba
go test ./...            # chạy toàn bộ test
```


## Go Packages

**Package** là một tập hợp các file mã nguồn `.go` dùng chung, có thể xem như một dự án hoặc không gian làm việc. Một package có thể gồm nhiều file, và **dòng đầu tiên của mỗi file** phải khai báo package mà nó thuộc về (ví dụ `package main`).

### Hai loại package

| Loại | Mục đích | Khi `go build` |
|---|---|---|
| **Executable** (thực thi) | Viết chương trình chạy được, làm việc thực tế | Tạo ra file thực thi (ví dụ `main`) |
| **Reusable** (tái sử dụng) | Thư viện, hàm trợ giúp, dependency dùng lại cho nhiều dự án | Không tạo file thực thi |

### Phân biệt bằng tên package

- `package main` → package **thực thi**. Đây là tên đặc biệt, chỉ dùng khi muốn tạo ra file chạy được.
- Bất kỳ tên nào khác (ví dụ `package apple`) → package **tái sử dụng**, `go build` sẽ không sinh ra file thực thi.
- Package thực thi bắt buộc phải có hàm `func main()`, là điểm bắt đầu của chương trình.

Trong khóa học này chủ yếu dùng `package main`, vì ta muốn tạo ra chương trình để chạy và thử nghiệm ngay. Khi cần viết thư viện để chia sẻ cho người khác dùng, ta đặt tên package chuyên biệt hơn.

## Go Packages

**Package** là một tập hợp các file mã nguồn `.go` dùng chung, có thể xem như một dự án hoặc không gian làm việc. Một package có thể gồm nhiều file, và **dòng đầu tiên của mỗi file** phải khai báo package mà nó thuộc về (ví dụ `package main`).

### Hai loại package

| Loại | Mục đích | Khi `go build` |
|---|---|---|
| **Executable** (thực thi) | Viết chương trình chạy được, làm việc thực tế | Tạo ra file thực thi (ví dụ `main`) |
| **Reusable** (tái sử dụng) | Thư viện, hàm trợ giúp, dependency dùng lại cho nhiều dự án | Không tạo file thực thi |

### Phân biệt bằng tên package

- `package main` → package **thực thi**. Đây là tên đặc biệt, chỉ dùng khi muốn tạo ra file chạy được.
- Bất kỳ tên nào khác (ví dụ `package apple`) → package **tái sử dụng**, `go build` sẽ không sinh ra file thực thi.
- Package thực thi bắt buộc phải có hàm `func main()`, là điểm bắt đầu của chương trình.

Trong khóa học này chủ yếu dùng `package main`, vì ta muốn tạo ra chương trình để chạy và thử nghiệm ngay. Khi cần viết thư viện để chia sẻ cho người khác dùng, ta đặt tên package chuyên biệt hơn.


## Import Statements

Câu lệnh `import` dùng để **truy cập mã nằm trong một package khác** từ package hiện tại. Ví dụ `import "fmt"` cấp cho package của ta quyền dùng toàn bộ hàm trong package `fmt`.

- `fmt` (viết tắt của *format*) là package thuộc **thư viện chuẩn (standard library)**, đi kèm sẵn với Go. Nó dùng để in thông tin ra terminal, hỗ trợ debug, ...
- Có thể import nhiều package chuẩn, ví dụ `"math"` để dùng package toán học.
- Không chỉ giới hạn ở thư viện chuẩn: ta có thể import package do kỹ sư khác viết và công bố (các package **reusable**).

```go
import (
    "fmt"
    "math"
)
```

Có thể hình dung `package main` nằm ở trung tâm, xung quanh là các package khác (`fmt`, `math`, ...). `import` chính là liên kết từ `main` tới từng package đó.

### Tài liệu thư viện chuẩn

Danh sách và tài liệu chính thức của tất cả package chuẩn: https://golang.org/pkg (nay là https://pkg.go.dev/std). Ta sẽ tra cứu tài liệu này rất thường xuyên trong quá trình học.

## Function và File Organization

### Khai báo hàm (`func`)

`func` là viết tắt của *function*, tương tự hàm trong Ruby, Python, JavaScript. Cú pháp:

```go
func tenHam(danhSachDoiSo) {
    // thân hàm
}
```

1. Từ khóa `func` để khai báo hàm mới.
2. Tên hàm (ví dụ `main`).
3. Cặp ngoặc đơn `()` chứa danh sách đối số truyền vào hàm.
4. Cặp ngoặc nhọn `{ }` chứa thân hàm.

### Cách tổ chức một file Go

Mọi file Go đều theo cùng một thứ tự:

```go
package main          // 1. Khai báo package của file

import "fmt"          // 2. Import các package cần dùng (chuẩn hoặc tự tạo)

func main() {         // 3. Phần thân: hàm, khai báo biến, logic xử lý
    fmt.Println("Hi there!")
}
```

| Thứ tự | Phần | Vai trò |
|---|---|---|
| 1 | `package ...` | File thuộc package nào |
| 2 | `import ...` | Liệt kê các package khác cần truy cập (ví dụ `fmt`, `io`, `os`, hoặc package tự viết) |
| 3 | Nội dung file | Các hàm, khai báo biến và logic thực sự làm việc |

## Function và File Organization

### Khai báo hàm (`func`)

`func` là viết tắt của *function*, tương tự hàm trong Ruby, Python, JavaScript. Cú pháp:

```go
func tenHam(danhSachDoiSo) {
    // thân hàm
}
```

1. Từ khóa `func` để khai báo hàm mới.
2. Tên hàm (ví dụ `main`).
3. Cặp ngoặc đơn `()` chứa danh sách đối số truyền vào hàm.
4. Cặp ngoặc nhọn `{ }` chứa thân hàm.

### Cách tổ chức một file Go

Mọi file Go đều theo cùng một thứ tự:

```go
package main          // 1. Khai báo package của file

import "fmt"          // 2. Import các package cần dùng (chuẩn hoặc tự tạo)

func main() {         // 3. Phần thân: hàm, khai báo biến, logic xử lý
    fmt.Println("Hi there!")
}
```

| Thứ tự | Phần | Vai trò |
|---|---|---|
| 1 | `package ...` | File thuộc package nào |
| 2 | `import ...` | Liệt kê các package khác cần truy cập (ví dụ `fmt`, `io`, `os`, hoặc package tự viết) |
| 3 | Nội dung file | Các hàm, khai báo biến và logic thực sự làm việc |


## Project Overview: Cards

Dự án thực tế đầu tiên là một package mô phỏng việc thao tác với **bộ bài tây** (loại bài dùng để chơi blackjack, gin rummy, poker, ...), ví dụ: át chủ bài, hai cơ, ba rô, ...

### Các chức năng cần xây dựng

| Hàm | Mục đích |
|---|---|
| `newDeck` | Tạo một bộ bài mới (52 lá). Hình dung như một mảng/danh sách chuỗi, mỗi phần tử là một lá bài |
| `print` | In ra danh sách tất cả các lá bài trong bộ bài |
| `shuffle` | Xáo trộn các lá bài |
| `deal` | Chia một "tay bài" từ bộ bài |
| `saveToFile` | Lưu bộ bài vào một file trên ổ cứng |
| `newDeckFromFile` | Đọc lại bộ bài đã lưu từ ổ cứng |

### Ví dụ về `deal`

Với bộ bài 52 lá, nếu chia một tay 5 lá thì kết quả gồm:
- Một danh sách **5 lá** (tay bài vừa chia).
- Một danh sách **47 lá** còn lại trong bộ bài.

### Ghi chú

- Ở thời điểm này chưa học về mảng (array), slice hay chuỗi trong Go, nên sẽ làm quen dần với các khái niệm đó ngay trong quá trình làm dự án.
- Bước đầu tiên: tạo một **thư mục mới** chứa package cho dự án cards.
- Qua dự án này sẽ học thêm về cách viết code và các mẫu thiết kế (design pattern) phổ biến trong Go.
