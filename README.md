# Golang tutorial
## My note
-Chương trình chính run tại package main 
-Tự khai báo type
-Tự khai báo function cho type, mọi biến cùng kiểu đều dùng được
-Reveiver function (method) là hàm gắn vào 1 kiểu dữ liệu cụ thể => biến hàm thành method của type (giống như Class)
## Hello World trong Go

Nội dung file [main.go](main.go):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```
go run main.go 

# 3_deeper_into_Go
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

## Variable Declarations

Mục tiêu: trong hàm `main`, tạo một biến chứa chuỗi đại diện cho một lá bài, rồi in nó ra bằng `fmt`.

```go
package main

import "fmt"

func main() {
    var card string = "Ace of Spades"
    fmt.Println(card)
}
```

> Mẹo: khi lưu file, VS Code (extension Go) tự thêm `import "fmt"` vì nhận ra ta đang dùng package `fmt`.

### Phân tích `var card string = "Ace of Spades"`

| Phần | Ý nghĩa |
|---|---|
| `var` | Viết tắt của *variable*, báo hiệu sắp tạo một biến mới |
| `card` | Tên biến |
| `string` | Báo cho trình biên dịch Go rằng **chỉ giá trị kiểu string** được gán cho biến này |
| `= "Ace of Spades"` | Giá trị gán cho biến |

### Go là ngôn ngữ kiểu tĩnh (statically typed)

- JavaScript, Ruby, Python là ngôn ngữ **kiểu động**: một biến có thể chứa số rồi sau đó lại chứa chuỗi.
  ```js
  var number = 123;
  number = "abcd"; // JavaScript chấp nhận
  ```
- Go (cũng như Java, C++) là ngôn ngữ **kiểu tĩnh**: mỗi biến có một kiểu cố định. Biến `card` kiểu `string` thì luôn chứa string, gán số vào sẽ bị lỗi biên dịch.

### Một số kiểu cơ bản

| Kiểu | Mô tả | Ví dụ |
|---|---|---|
| `bool` | Đúng/sai | `true`, `false` |
| `string` | Chuỗi ký tự | `"Hi!"`, `"How's it going?"` |
| `int` | Số nguyên | `0`, `-10000`, `99999` |
| `float64` | Số có phần thập phân | `10.000001`, `0.00009`, `-100.003` |

Đây chưa phải danh sách đầy đủ, nhưng là những kiểu dùng nhiều nhất trong khóa học. Các kiểu phức tạp hơn đều được xây dựng dựa trên những kiểu cơ bản này.

### Cách viết ngắn gọn `:=`

Với cách viết đầy đủ ở trên, editor sẽ cảnh báo kiểu *"should omit type string from declaration of var card; it will be inferred from the right-hand side"*. Nghĩa là Go có thể **tự suy ra kiểu** từ giá trị bên phải, nên có thể viết gọn:

```go
var card string = "Ace of Spades" // cách viết đầy đủ
card := "Ace of Spades"           // cách viết ngắn, tương đương 100%
```

Cả hai đều tạo biến `card` kiểu `string`. Với `:=`, trình biên dịch tự suy luận kiểu dựa trên giá trị được gán.

### Khai báo mới vs gán lại

`:=` **chỉ dùng khi khai báo biến mới**. Khi gán giá trị mới cho biến đã có, chỉ dùng `=`:

```go
card := "Ace of Spades"    // khai báo + khởi tạo
card = "Five of Diamonds"  // gán lại, KHÔNG dùng :=
```

Đây là lỗi rất hay gặp khi mới học. Nếu viết `card := "Five of Diamonds"` lần nữa, trình biên dịch báo lỗi:

```
no new variables on left side of :=
```

### Tóm tắt

| Cú pháp | Khi nào dùng |
|---|---|
| `var card string = "Ace of Spades"` | Khai báo đầy đủ, ghi rõ kiểu |
| `card := "Ace of Spades"` | Khai báo ngắn, Go tự suy ra kiểu (chỉ khi tạo biến mới) |
| `card = "Five of Diamonds"` | Gán giá trị mới cho biến đã tồn tại |

## Functions and Return Types

Thay vì gán trực tiếp `"Ace of Spades"` cho `card`, ta tách việc tạo lá bài ra một **hàm riêng** và gọi nó từ `main`.

```go
package main

import "fmt"

func main() {
    card := newCard()
    fmt.Println(card)
}

func newCard() string {
    return "Five of Diamonds"
}
```

### Khai báo kiểu trả về

Nếu viết `func newCard() { return "Five of Diamonds" }` mà không ghi kiểu trả về, trình biên dịch báo lỗi:

```
too many return values
    have (string)
    want ()
```

Nghĩa là hàm được khai báo **không trả về gì**, nhưng thân hàm lại trả về một string. Cách sửa: thêm kiểu trả về ngay sau dấu `)`.

### Phân tích `func newCard() string`

| Phần | Ý nghĩa |
|---|---|
| `func` | Từ khóa khai báo hàm mới |
| `newCard` | Tên hàm |
| `()` | Danh sách đối số (ở đây rỗng) |
| `string` | Kiểu trả về: mỗi lần gọi, hàm **luôn trả về một giá trị kiểu string** |

Go bắt buộc phải ghi rõ kiểu dữ liệu mà hàm trả về.

### Suy luận kiểu với `:=`

Vì Go biết `newCard()` luôn trả về `string`, câu lệnh `card := newCard()` vẫn tự suy ra `card` có kiểu `string`. Di chuột qua `newCard` hoặc `card` trong VS Code sẽ thấy kiểu tương ứng.

### Đổi sang kiểu trả về khác

Kiểu trả về có thể là bất kỳ kiểu nào: `int`, `float64`, `bool`, ... Ví dụ đổi sang `int`:

```go
func newCard() int {
    return "Five of Diamonds" // lỗi: cannot use "Five of Diamonds" (untyped string constant) as int value in return statement
}
```

Sửa bằng cách trả về một số nguyên:

```go
func newCard() int {
    return 12
}
```

Lúc này `card := newCard()` sẽ được suy ra là kiểu `int`. Trong dự án Cards ta vẫn dùng `string`, vì mỗi lá bài là một chuỗi.

> Mẹo: thông báo lỗi của Go thường mô tả khá rõ vấn đề. Hãy đọc kỹ trước khi đi tra Google.

### Tóm tắt

- Khai báo hàm riêng bằng `func tenHam() kieuTraVe { ... }`.
- Kiểu trả về ghi ngay sau `()` và phải khớp với giá trị trong `return`.
- `:=` suy ra kiểu biến từ kiểu trả về của hàm được gọi.

## Slices and For Loops

Một bộ bài có tới 52 lá, nên ta cần một cấu trúc dữ liệu chứa **danh sách nhiều bản ghi**.

### Array và Slice

Go có hai cấu trúc dữ liệu cơ bản để chứa danh sách:

| | Array (mảng) | Slice (lát cắt) |
|---|---|---|
| Độ dài | **Cố định**, khai báo lúc tạo (ví dụ đúng 5 phần tử) | **Thay đổi được**, có thể thêm/bớt phần tử |
| Bản chất | Cấu trúc nguyên thủy, đơn giản | Cũng là mảng nhưng có thêm nhiều tính năng |

Điểm chung: mọi phần tử trong array hoặc slice **phải cùng một kiểu**.

```go
[]string{"Five of Diamonds", "Three of Diamonds", "Five of Diamonds"} // hợp lệ: toàn string
[]string{"Five of Spades", 10}                                       // lỗi: lẫn string và int
```

### Khai báo slice

```go
cards := []string{"Ace of Diamonds", newCard()}
```

| Phần | Ý nghĩa |
|---|---|
| `[]` | Báo hiệu đây là một slice |
| `string` | Kiểu của mọi phần tử trong slice. Đọc là "slice of type string" |
| `{ ... }` | Danh sách phần tử ban đầu, cách nhau bởi dấu phẩy |

Phần tử có thể là giá trị viết tay (`"Ace of Diamonds"`) hoặc kết quả trả về của hàm (`newCard()`), miễn là đúng kiểu `string`.

> Lưu ý: khai báo biến mà không dùng sẽ bị lỗi `declared and not used: cards`. Go không cho phép biến thừa, nên cần dùng nó, ví dụ `fmt.Println(cards)`.

### Thêm phần tử: `append`

```go
cards = append(cards, "Six of Spades")
```

- `append` nhận slice hiện có và phần tử mới, rồi thêm phần tử vào **cuối** slice.
- **Quan trọng:** `append` **không sửa slice cũ** mà trả về **một slice mới**, nên phải gán lại cho `cards`. Chủ đề này sẽ được nói kỹ hơn ở phần sau của khóa học.

### Duyệt slice bằng `for ... range`

```go
for i, card := range cards {
    fmt.Println(i, card)
}
```

| Phần | Ý nghĩa |
|---|---|
| `i` | Chỉ số (index) của phần tử hiện tại, bắt đầu từ 0 |
| `card` | Giá trị của phần tử hiện tại |
| `range cards` | Lấy slice `cards` và lặp qua từng phần tử |
| `{ ... }` | Thân vòng lặp, chạy một lần cho mỗi phần tử |

**Vì sao dùng `:=` trong vòng lặp?** Mỗi lần lặp, `i` và `card` của lần trước bị bỏ đi và được khai báo lại với giá trị mới, nên dùng cú pháp khai báo `:=`.

### Ví dụ đầy đủ

```go
package main

import "fmt"

func main() {
    cards := []string{"Ace of Diamonds", newCard()}
    cards = append(cards, "Six of Spades")

    for i, card := range cards {
        fmt.Println(i, card)
    }
}

func newCard() string {
    return "Five of Diamonds"
}
```

Chạy `go run main.go` trong thư mục `cards`:

```
0 Ace of Diamonds
1 Five of Diamonds
2 Six of Spades
```

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `[]string{"a", "b"}` | Tạo slice kiểu string với các phần tử ban đầu |
| `cards = append(cards, "c")` | Thêm phần tử vào cuối, trả về slice mới |
| `for i, card := range cards { }` | Lặp qua từng phần tử, lấy index và giá trị |

## OO Approach vs Go Approach

Trước khi bắt tay vào dự án Cards, hãy so sánh cách tiếp cận của một ngôn ngữ hướng đối tượng (Python, Ruby, Java) với cách làm trong Go.

> Go **không phải** ngôn ngữ hướng đối tượng: Go **không có class**.

### Cách tiếp cận hướng đối tượng (OOP)

- Tạo một **class `Deck`**: bản thiết kế mô tả một bộ bài, gồm thuộc tính và phương thức.
- Từ class đó tạo ra một **instance** (thể hiện) của bộ bài.
- Instance có thuộc tính `cards` là một mảng các string.
- Instance có các phương thức như `print`, `shuffle`, `saveToFile` để thao tác trên danh sách lá bài.

```python
class Deck:
    def __init__(self):
        self.cards = []

    def print(self): ...
    def shuffle(self): ...
    def save_to_file(self): ...
```

### Cách tiếp cận của Go

- Go có sẵn các kiểu cơ bản (`string`, `int`, `float64`, ...) và cả array, slice.
- Ta **định nghĩa một kiểu mới** tên `deck`, về bản chất là **một slice of string** (`[]string`). Mỗi phần tử là một lá bài.
- Để thêm chức năng cho kiểu `deck`, ta viết các **hàm có receiver** (*functions with a receiver*). Khái niệm receiver sẽ được giải thích ở phần sau.

```go
type deck []string // kiểu mới dựa trên slice of string

func (d deck) print() { ... } // hàm có receiver, gắn với kiểu deck
```

Ý tưởng chính: **lấy một kiểu có sẵn (slice) rồi mở rộng chức năng của nó** bằng kiểu mới và các hàm đi kèm. Đây là mẫu rất phổ biến, sẽ gặp lại nhiều lần trong các chương trình Go.

### So sánh

| | OOP (Python, Ruby, Java) | Go |
|---|---|---|
| Mô tả bộ bài | Class `Deck` | Kiểu mới `type deck []string` |
| Dữ liệu | Thuộc tính `cards` trong instance | Chính slice `deck` |
| Chức năng | Phương thức của class | Hàm có receiver gắn với kiểu `deck` |

### Cấu trúc dự án Cards

| File | Vai trò |
|---|---|
| `main.go` | Chứa hàm `main`, tạo và thao tác với bộ bài |
| `deck.go` | Mô tả bộ bài là gì và hoạt động thế nào (kiểu `deck` và các hàm của nó) |
| `deck_test.go` | Chứa các bài test cho bộ bài |

`deck.go` vẫn thuộc package `main`, nên dòng đầu tiên của file phải là:

```go
package main
```

## Custom Type Declarations

Ta định nghĩa kiểu mới `deck`. Về bản chất nó là một slice of string, nhưng vì là kiểu riêng nên ta có thể viết các hàm **chỉ dành cho `deck`**.

### Khai báo kiểu `deck`

Trong file `deck.go`:

```go
package main

// Create a new type of 'deck'
// which is a slice of strings
type deck []string
```

- `type deck []string`: tạo kiểu mới tên `deck`, có mọi hành vi của `[]string`.
- Nếu quen OOP, có thể hình dung `deck` như "mở rộng" hoặc "kế thừa" `[]string`. Go không dùng các thuật ngữ đó, đây chỉ là cách liên tưởng cho dễ hiểu.

### Thay `[]string` bằng `deck`

Vì `deck` và `[]string` tương đương, trong `main.go` ta thay:

```go
cards := []string{"Ace of Diamonds", newCard()} // trước
cards := deck{"Ace of Diamonds", newCard()}     // sau
```

### Chạy chương trình nhiều file

Giờ chương trình gồm hai file, nên phải truyền cả hai cho `go run`:

```bash
go run main.go deck.go
```

Kết quả in ra vẫn như trước, chứng tỏ thay `[]string` bằng `deck` không làm thay đổi hành vi.

### Vì sao cần kiểu riêng? Hàm `print` gắn với `deck`

Mục tiêu: viết một hàm thuộc kiểu `deck`, khi gọi sẽ in ra từng lá bài. Thêm vào `deck.go`:

```go
package main

import "fmt"

type deck []string

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

- Điểm lạ: giữa `func` và tên hàm có `(d deck)`. Đây gọi là **receiver**.
- Thân hàm vẫn là vòng lặp `for ... range` như trước, nhưng lặp trên `d` thay vì `cards`.

Trong `main.go`, bỏ vòng lặp `for` và gọi hàm theo cú pháp `bien.tenHam()`:

```go
func main() {
    cards := deck{"Ace of Diamonds", newCard()}
    cards = append(cards, "Six of Spades")

    cards.print()
}
```

Chạy `go run main.go deck.go`:

```
0 Ace of Diamonds
1 Five of Diamonds
2 Six of Spades
```

Kết quả giống hệt vòng lặp cũ. Receiver hoạt động thế nào sẽ được giải thích kỹ ở phần tiếp theo.

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `type deck []string` | Tạo kiểu mới `deck` dựa trên slice of string |
| `deck{"a", "b"}` | Tạo giá trị kiểu `deck` |
| `func (d deck) print() { }` | Hàm có receiver, gắn với kiểu `deck` |
| `cards.print()` | Gọi hàm `print` trên biến `cards` kiểu `deck` |
| `go run main.go deck.go` | Biên dịch và chạy chương trình gồm nhiều file |

## Receiver Functions

Ở phần trước, hàm `print` có thêm một cú pháp mới đặt trước tên hàm: **receiver**.

```go
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

### Receiver dùng để làm gì?

Receiver **gắn một phương thức vào một kiểu**. Nhờ `(d deck)`, **mọi biến kiểu `deck`** trong chương trình đều có phương thức `print`.

```go
cards := deck{"Ace of Diamonds", newCard()} // cards có kiểu deck
cards.print()                               // nên gọi được cards.print()
```

### Phân tích `func (d deck) print()`

| Phần | Ý nghĩa |
|---|---|
| `func` | Từ khóa khai báo hàm |
| `(d deck)` | Receiver |
| `deck` | Kiểu được gắn phương thức: mọi biến kiểu `deck` đều gọi được `print` |
| `d` | Tham chiếu tới **chính biến `deck` đang gọi hàm** |
| `print` | Tên phương thức |

Khi gọi `cards.print()`, có thể hình dung biến `cards` được truyền vào hàm `print` dưới tên `d`. Nói cách khác, bên trong hàm, `d` **chính là** `cards`. Nếu đổi tên `d` thành `cards` thì thân hàm trông y hệt vòng lặp `for` cũ trong `main`.

### So với `this` / `self`

- Nếu quen Python, Ruby, JavaScript, có thể hình dung `d` giống `self` hoặc `this`.
- Nhưng theo quy ước, Go **không bao giờ** đặt tên receiver là `this` hay `self`.

### Quy ước đặt tên receiver

- Go thường đặt tên receiver bằng **1–2 chữ cái viết tắt** của kiểu. Kiểu `deck` thì dùng `d`.
- Đây chỉ là quy ước, không bắt buộc: đặt `myDeck`, `thisDeck` vẫn chạy được.
- Ở ngôn ngữ khác thường khuyên tránh tên biến một chữ cái, nhưng với receiver thì đây là quy ước của Go và khóa học sẽ theo quy ước này.

### Liên hệ với OOP

Có thể **tạm hình dung** quan hệ giữa kiểu `deck` và phương thức `print` giống class `Deck` và phương thức của nó: mọi "instance" của `deck` đều có `print`. Tuy vậy, Go không phải ngôn ngữ hướng đối tượng và không dùng các thuật ngữ class hay instance.

Đây là mẫu rất phổ biến trong Go: **định nghĩa kiểu mới dựa trên kiểu có sẵn, rồi thêm phương thức cho nó bằng receiver**.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Receiver `(d deck)` | Gắn phương thức vào kiểu `deck` |
| `d` | Biến `deck` đang gọi phương thức, tương tự `this`/`self` |
| Tên receiver | Theo quy ước là 1–2 chữ cái viết tắt của kiểu, không dùng `this`/`self` |
| `cards.print()` | Gọi phương thức trên biến có kiểu tương ứng |

## Creating a New Deck

Mục tiêu: viết hàm `newDeck` để **tạo và trả về một bộ bài mới** (kiểu `deck`).

### Khai báo hàm

```go
func newDeck() deck {
    // ...
}
```

- `deck` sau `()` là **kiểu trả về**: gọi `newDeck()` sẽ nhận về một giá trị kiểu `deck`.
- **Không cần receiver**: hàm này dùng để *tạo* bộ bài, nên lúc gọi ta chưa có bộ bài nào. Receiver chỉ dùng khi muốn gọi theo kiểu `bien.tenHam()` trên một giá trị đã có.

### Ý tưởng: tạo bằng vòng lặp thay vì gõ tay 52 lá

Gõ tay đủ 52 lá (`"Ace of Spades"`, `"Two of Spades"`, ...) rất tẻ nhạt. Cách thông minh hơn:

1. Tạo một `deck` rỗng.
2. Tạo hai slice: một chứa các **chất** (suit), một chứa các **giá trị** (value).
3. Dùng hai vòng lặp `for` lồng nhau. Với mỗi cặp chất và giá trị, ghép thành một lá bài rồi `append` vào bộ bài.
4. Trả về bộ bài.

```go
func newDeck() deck {
    cards := deck{} // deck rỗng

    cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }

    return cards
}
```

- `deck{}`: tạo giá trị kiểu `deck` không có phần tử nào.
- `cardSuits` và `cardValues` là `[]string` chứ **không phải** `deck`. Chúng chỉ là "nửa lá bài", không đại diện cho lá bài thật.
- `cardValues` tạm thời chỉ có Ace đến Four cho gọn khi thử nghiệm (đủ bộ thì phải tới King).
- `value + " of " + suit`: dùng `+` để nối chuỗi, ví dụ `"Ace of Spades"`.

Phần này chủ yếu là kiến thức lập trình cơ bản (khai báo danh sách, lặp, nối chuỗi). Go khá "vanilla", nên kinh nghiệm từ ngôn ngữ khác áp dụng được ngay.

### Dấu gạch dưới `_` cho biến không dùng

Nếu viết `for i, suit := range cardSuits` và `for j, value := range cardValues`, Go báo lỗi vì `i` và `j` được khai báo nhưng không dùng:

```
declared and not used: i
declared and not used: j
```

Cách sửa: thay biến không dùng bằng **`_`**. Dấu `_` nghĩa là "biết ở đây có một giá trị, nhưng ta không quan tâm và không dùng nó".

### Dùng `newDeck` trong `main`

Xóa hàm `newCard` và đoạn `append` thủ công trong `main.go`, thay bằng:

```go
func main() {
    cards := newDeck()
    cards.print()
}
```

Chạy `go run main.go deck.go`:

```
0 Ace of Spades
1 Two of Spades
2 Three of Spades
3 Four of Spades
4 Ace of Diamonds
...
15 Four of Clubs
```

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `func newDeck() deck` | Hàm không có receiver, trả về kiểu `deck` |
| `deck{}` | Tạo một `deck` rỗng |
| `for _, suit := range cardSuits` | Lặp và bỏ qua index bằng `_` |
| `value + " of " + suit` | Nối chuỗi bằng `+` |
| `cards := newDeck()` | Tạo bộ bài mới bằng cách gọi hàm |

## Slice Range Syntax

Tiếp theo là hàm `deal`: lấy một bộ bài có sẵn và **chia ra một "tay bài"** với số lá cho trước, ví dụ chia 5, 6 hoặc 7 lá từ bộ 52 lá.

```
Bộ bài ban đầu (4 lá)          Chia tay 3 lá
┌─────────────────────┐        ┌───────────────┐   ┌──────────┐
│ A♠ │ 2♠ │ 3♠ │ 4♠   │  ───►  │ A♠ │ 2♠ │ 3♠  │   │ 4♠       │
└─────────────────────┘        └───────────────┘   └──────────┘
                                   hand (deck)      remaining (deck)
```

Lưu ý: dù gọi là "tay bài", **kiểu dữ liệu vẫn là `deck`**. Ta không đổi kiểu, chỉ **tách một slice thành hai slice nhỏ hơn**. Để làm vậy, cần hiểu thêm cách slice hoạt động.

### Truy cập phần tử theo chỉ số

Slice **đánh chỉ số từ 0**, giống nhiều ngôn ngữ khác. Truy cập một phần tử bằng `tenSlice[chiSo]`:

```go
fruits := []string{"apple", "banana", "grape", "orange"}

fruits[0] // "apple"
fruits[3] // "orange"
```

| Chỉ số | 0 | 1 | 2 | 3 |
|---|---|---|---|---|
| Giá trị | apple | banana | grape | orange |

### Lấy một khoảng: `slice[start:end]`

Go có sẵn cú pháp để lấy một **khoảng con** của slice:

```go
fruits[startIndexIncluding : upToNotIncluding]
```

- Số **bên trái** `:` là chỉ số bắt đầu, **có bao gồm**.
- Số **bên phải** `:` là chỉ số kết thúc, **không bao gồm**.

```go
fruits[0:2] // ["apple", "banana"]: lấy 0 và 1, KHÔNG lấy 2
```

### Bỏ trống một bên

Có thể bỏ số ở một bên dấu `:`:

| Cú pháp | Ý nghĩa | Kết quả |
|---|---|---|
| `fruits[0:2]` | Từ 0 đến trước 2 | `["apple", "banana"]` |
| `fruits[:2]` | Từ **đầu slice** đến trước 2 | `["apple", "banana"]` |
| `fruits[2:]` | Từ 2 đến **cuối slice** | `["grape", "orange"]` |

### Áp dụng cho hàm `deal`

Hàm `deal` sẽ nhận một số nguyên `handSize` là số lá muốn chia. Khi đó:

```go
cards[:handSize] // tay bài: từ đầu đến trước handSize
cards[handSize:] // phần còn lại: từ handSize đến cuối
```

Ví dụ với `cards = ["Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades"]` và `handSize = 3`:

| Biểu thức | Tương đương | Kết quả |
|---|---|---|
| `cards[:handSize]` | `cards[:3]` → chỉ số 0, 1, 2 | `["Ace of Spades", "Two of Spades", "Three of Spades"]` |
| `cards[handSize:]` | `cards[3:]` → chỉ số 3 đến hết | `["Four of Spades"]` |

Hai biểu thức này chia bộ bài thành đúng hai phần, không trùng và không sót lá nào. Phần tiếp theo sẽ viết code cho hàm `deal`.

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `s[i]` | Phần tử ở chỉ số `i` (bắt đầu từ 0) |
| `s[a:b]` | Từ `a` (bao gồm) đến `b` (không bao gồm) |
| `s[:b]` | Từ đầu đến trước `b` |
| `s[a:]` | Từ `a` đến cuối |

## Multiple Return Values

Giờ viết hàm `deal` dựa trên cú pháp lấy khoảng của slice ở phần trước.

### Chữ ký hàm `deal`

Trong `deck.go`:

```go
func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}
```

| Phần | Ý nghĩa |
|---|---|
| `d deck` | Đối số thứ nhất: bộ bài cần chia, kiểu `deck` |
| `handSize int` | Đối số thứ hai: số lá muốn chia, kiểu `int` |
| `(deck, deck)` | Hàm **trả về hai giá trị**, cả hai đều kiểu `deck` |
| `d[:handSize]` | Giá trị trả về thứ nhất: tay bài |
| `d[handSize:]` | Giá trị trả về thứ hai: phần còn lại của bộ bài |

### Khai báo đối số

- Mỗi đối số viết theo dạng **`tenBien kieu`**, ví dụ `d deck` hay `handSize int`.
- Go kiểm tra kiểu khi gọi hàm: `deal("blah", 1)` sẽ **lỗi** vì đối số đầu là `string` chứ không phải `deck`.
- Tên `d` vẫn theo quy ước viết tắt. Nếu đặt là `deck` thì dễ lẫn với tên kiểu `deck`. Có thể đặt `cards` nếu muốn rõ hơn.

### Trả về nhiều giá trị

Go cho phép một hàm trả về **nhiều giá trị**:

- Khai báo kiểu trả về trong **ngoặc đơn**, cách nhau bởi dấu phẩy: `(deck, deck)`.
- Câu lệnh `return` cũng liệt kê các giá trị cách nhau bởi dấu phẩy: `return d[:handSize], d[handSize:]`.

Hàm `deal` không có receiver mà nhận bộ bài qua đối số. Ta gọi `deal(cards, 5)` chứ không gọi `cards.deal(5)`.

### Nhận nhiều giá trị trả về

Trong `main.go`:

```go
func main() {
    cards := newDeck()

    hand, remainingCards := deal(cards, 5)

    hand.print()
    remainingCards.print()
}
```

- `hand, remainingCards := deal(cards, 5)`: giá trị trả về **thứ nhất** gán cho `hand`, giá trị **thứ hai** gán cho `remainingCards`.
- Cả hai đều là biến mới nên dùng `:=`, và đều có kiểu `deck`, nên gọi được `.print()`.
- `deal` khai báo trong `deck.go` nhưng `main.go` gọi được **mà không cần import**, vì hai file **cùng package `main`**.

### Kết quả

Chạy `go run main.go deck.go`:

```
0 Ace of Spades
1 Two of Spades
2 Three of Spades
3 Four of Spades
4 Ace of Diamonds
0 Two of Diamonds
1 Three of Diamonds
...
10 Four of Clubs
```

Chỉ số chạy từ 0 đến 4 (tay bài 5 lá), rồi **bắt đầu lại từ 0** cho 11 lá còn lại. Điều này cho thấy bộ bài đã được tách thành hai `deck` riêng biệt.

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `func f(a T1, b T2)` | Khai báo đối số dạng `tên kiểu` |
| `func f() (T1, T2)` | Hàm trả về hai giá trị |
| `return x, y` | Trả về nhiều giá trị, cách nhau bởi dấu phẩy |
| `a, b := f()` | Nhận nhiều giá trị trả về vào nhiều biến |

## Byte Slices

Tạm bỏ qua `shuffle` (sẽ quay lại sau) để làm `saveToFile`: **lưu bộ bài xuống ổ cứng**, để sau này `newDeckFromFile` đọc lại và tạo bộ bài từ file đó.

### Tìm package trong thư viện chuẩn

Khi cần làm việc với phần cứng hoặc hệ điều hành (đọc/ghi file, ...), hãy nghĩ ngay tới **thư viện chuẩn** của Go (https://pkg.go.dev/std).

Package dùng trong khóa học là **`io/ioutil`**: một package nhỏ chứa các thao tác phổ biến với file. Trong đó có hàm **`WriteFile`** để tạo mới hoặc ghi đè file trên ổ cứng:

```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

| Đối số | Kiểu | Ý nghĩa |
|---|---|---|
| `filename` | `string` | Tên file cần ghi |
| `data` | `[]byte` | Dữ liệu thô cần ghi vào file |
| `perm` | `os.FileMode` | Quyền (permission) dùng khi phải tạo file mới |

> Ghi chú: từ Go 1.16, `io/ioutil` đã **deprecated**. Hàm tương đương hiện nay là `os.WriteFile`, có cùng các đối số.

### `[]byte` là gì?

- Thấy `[]` thì nghĩ ngay tới **slice**. `[]byte` đọc là "**byte slice**" (slice of byte).
- Có thể hiểu đơn giản: **byte slice = một chuỗi ký tự (string)** được biểu diễn theo cách **thân thiện với máy tính**.
- Mỗi phần tử trong byte slice là **mã ASCII** của một ký tự.

Ví dụ chuỗi `"Hi there!"`:

| Ký tự | `H` | `i` | ` ` | `t` | `h` | `e` | `r` | `e` | `!` |
|---|---|---|---|---|---|---|---|---|---|
| Byte (thập phân) | 72 | 105 | 32 | 116 | 104 | 101 | 114 | 101 | 33 |

```go
[]byte("Hi there!") // [72 105 32 116 104 101 114 101 33]
```

Tra bảng ASCII (cột Decimal) sẽ thấy 72 là `H`, 105 là `i`, 32 là dấu cách, ...

### Vì sao cần quan tâm?

- Lập trình viên thường làm việc với **string** vì con người đọc được. Nhìn một dãy số thì khó biết ngay nó là chuỗi gì.
- Nhưng nhiều hàm trong **thư viện chuẩn** lại yêu cầu đầu vào là `[]byte`, như `WriteFile` ở trên.
- **Không có `[]byte` thì không gọi được `WriteFile`.**

Vậy bài toán đặt ra: **chuyển một `deck` (slice of string) thành `[]byte` như thế nào?** Phần tiếp theo sẽ giải quyết việc này.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `io/ioutil` | Package chuẩn chứa các thao tác file phổ biến (nay dùng `os`) |
| `WriteFile(filename, data, perm)` | Ghi dữ liệu `[]byte` vào file |
| `[]byte` | Byte slice: chuỗi biểu diễn bằng các mã ASCII |
| Bài toán tiếp theo | Đổi `deck` thành `[]byte` để ghi ra file |

## Deck to String

Mục tiêu: biến `deck` thành `[]byte` để truyền vào `WriteFile`.

### Chuyển đổi kiểu (Type Conversion)

Type conversion dùng để **biến một giá trị thuộc kiểu này thành giá trị thuộc kiểu khác**. Cú pháp: viết **kiểu muốn có**, rồi đặt **giá trị đang có** trong cặp ngoặc đơn.

```go
greeting := "Hi there!"
fmt.Println([]byte(greeting)) // [72 105 32 116 104 101 114 101 33]
```

Ta sẽ gặp type conversion rất nhiều, vì thường xuyên có dữ liệu thuộc một kiểu nhưng hàm có sẵn lại đòi kiểu khác.

### Từ `deck` đến `[]byte`

`deck` thực chất là `[]string` (đã khai báo `type deck []string` trong `deck.go`), nên quy trình là:

```
deck  ->  []string  ->  string (nối tất cả lại)  ->  []byte
```

Mẹo khi làm việc với ngôn ngữ kiểu tĩnh: xác định **kiểu đang có** (điểm A) và **kiểu muốn có** (điểm B), rồi tìm đường đi từ A đến B.

### Hàm helper `toString`

Phần khó nhất là biến `deck` thành `string`, nên tách ra một hàm helper riêng thay vì nhồi hết vào `saveToFile`. Lý do: sau này có thể còn cần chuyển `deck` thành `string` ở chỗ khác; còn từ `string` sang `[]byte` thì đã quá dễ.

- Kiểu trả về: `string`.
- Dùng **receiver** (thay vì truyền `deck` làm đối số), vì `toString` gắn liền tự nhiên với kiểu `deck`, giống `print`. (Chuyện khi nào dùng receiver, khi nào không sẽ bàn sau.)

```go
func (d deck) toString() string {
	// biến d thành string — hoàn thiện ở phần sau
}
```

Khi đó có thể gọi: `cards.toString()`.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Type conversion | `KiểuMuốn(giáTrị)`, ví dụ `[]byte("Hi there!")` |
| Quy trình | `deck` → `[]string` → `string` → `[]byte` |
| `func (d deck) toString() string` | Helper có receiver, trả về chuỗi biểu diễn bộ bài |
| Phần tiếp theo | Cài đặt cách nối `[]string` thành một `string` |

## Joining a Slice of Strings

Hoàn thiện hàm `toString` đã khai báo ở phần trước.

### Bước 1: `deck` → `[]string`

Dùng type conversion để chuyển receiver `d` về `[]string`:

```go
[]string(d)
```

Chuyển được vì `deck` vốn được định nghĩa từ `[]string` (`type deck []string`), nên có thể tự do chuyển ngược về kiểu gốc.

### Bước 2: `[]string` → `string` với `strings.Join`

Ý tưởng: nối mọi phần tử, đặt dấu phẩy ở giữa. Ví dụ `["red", "yellow", "blue"]` → `"red,yellow,blue"`.

Không cần tự viết vòng lặp: package **`strings`** trong thư viện chuẩn (https://pkg.go.dev/strings) có sẵn hàm **`Join`**:

```go
func Join(elems []string, sep string) string
```

| Đối số | Kiểu | Ý nghĩa |
|---|---|---|
| `elems` | `[]string` | Slice các chuỗi cần nối |
| `sep` | `string` | Dấu phân tách (separator) đặt giữa các phần tử |

### Import nhiều package

Khi cần nhiều package, bọc câu lệnh `import` bằng cặp ngoặc đơn, mỗi package một dòng, **không có dấu phẩy** ngăn cách:

```go
import (
	"fmt"
	"strings"
)
```

### Code hoàn chỉnh

`deck.go`:

```go
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

`main.go`:

```go
func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
```

Chạy `go run main.go deck.go`, kết quả là **một chuỗi duy nhất** chứa các lá bài cách nhau bằng dấu phẩy:

```
Ace of Spades,Two of Spades,Three of Spades,...
```

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `[]string(d)` | Chuyển `deck` về kiểu gốc `[]string` |
| `strings.Join(elems, sep)` | Nối slice chuỗi thành một chuỗi, ngăn cách bởi `sep` |
| `import ( ... )` | Import nhiều package, mỗi package một dòng, không dấu phẩy |
| Phần tiếp theo | Chuyển chuỗi thành `[]byte` và gọi `WriteFile` trong `saveToFile` |

## Saving Data to the Hard Drive

Đã có `toString`, giờ dùng `ioutil.WriteFile` để viết hàm `saveToFile`.

### Thiết kế chữ ký hàm

| Thành phần | Lựa chọn | Lý do |
|---|---|---|
| Receiver | `d deck` | Để gọi được `cards.saveToFile(...)` |
| Đối số | `filename string` | Cho người dùng tự chọn tên file, rồi truyền tiếp vào `WriteFile` |
| Kiểu trả về | `error` | `WriteFile` trả về `error` nếu ghi file thất bại; ta trả nguyên lỗi đó ra ngoài (xử lý lỗi sẽ học kỹ sau) |

`error` là một **kiểu có sẵn** trong Go, đại diện cho một lỗi xảy ra trong chương trình.

### Code hoàn chỉnh

`deck.go`:

```go
import (
	"fmt"
	"io/ioutil"
	"strings"
)

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

Giải thích các đối số truyền vào `WriteFile`:

| Đối số | Giá trị | Ý nghĩa |
|---|---|---|
| `filename` | `filename` | Tên file nhận từ người gọi |
| `data` | `[]byte(d.toString())` | `deck` → `string` → `[]byte` |
| `perm` | `0666` | Quyền mặc định khi file chưa tồn tại: **ai cũng đọc và ghi được** |

> Lưu ý: dòng này có 3 cặp ngoặc lồng nhau (`toString()`, type conversion `[]byte(...)`, lời gọi `WriteFile(...)`), kiểm tra kỹ để đóng mở cho đủ.

### Package con (sub-package)

- `ioutil` là **package con** nằm trong `io`, nên đường dẫn import là `"io/ioutil"`. Tài liệu của `ioutil` cũng ghi rõ cách import này.
- VS Code (với extension Go) **tự thêm câu lệnh import** khi lưu file nếu phát hiện đang dùng package chưa import.

### Chạy thử

`main.go`:

```go
func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
```

Chạy `go run main.go deck.go`. Sau khi chạy, trong thư mục hiện tại xuất hiện file mới `my_cards`. Mở ra là một **file văn bản thuần** chứa danh sách lá bài, cách nhau bởi dấu phẩy.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `func (d deck) saveToFile(filename string) error` | Lưu bộ bài ra file, trả về lỗi nếu có |
| `[]byte(d.toString())` | Chuỗi hoá bộ bài rồi chuyển thành byte slice |
| `0666` | Quyền file: mọi người đều đọc/ghi được |
| `"io/ioutil"` | Import package con `ioutil` nằm trong `io` |
| Phần tiếp theo | Viết hàm đọc bộ bài từ file (`newDeckFromFile`) |

## Reading From the Hard Drive

Viết hàm ngược lại với `saveToFile`: đọc file từ ổ cứng và biến nó trở lại thành `deck`.

### `ioutil.ReadFile`

```go
func ReadFile(filename string) ([]byte, error)
```

- Nhận vào **tên file** (`string`).
- Trả về **2 giá trị**: `[]byte` (nội dung file) và `error`.
- Thấy `[]byte` thì nghĩ ngay đến **string**: ở đây là chuỗi các lá bài cách nhau bởi dấu phẩy. Việc tiếp theo là tách chuỗi theo dấu phẩy để dựng lại `deck`.

### Thiết kế chữ ký hàm `newDeckFromFile`

| Thành phần | Lựa chọn | Lý do |
|---|---|---|
| Receiver | Không có | Lúc gọi hàm này ta **chưa có** bộ bài nào; gọi nó chính là để lấy một bộ bài |
| Đối số | `filename string` | Tên file cần mở |
| Kiểu trả về | `deck` | Kết quả mong đợi là một bộ bài |

### Giá trị `error` và `nil`

- Nếu `ReadFile` gặp sự cố, `err` sẽ **chứa giá trị lỗi**.
- Nếu mọi thứ thành công, `err` sẽ là **`nil`**: nghĩa là "không có giá trị".

Mẫu rất phổ biến trong Go: **ngay sau khi nhận về một `error`, kiểm tra nó bằng `if`**:

```go
if err != nil {
	// xử lý lỗi
}
```

Câu lệnh `if` trong Go giống các ngôn ngữ khác: từ khoá `if`, điều kiện, rồi khối `{ }` chỉ chạy khi điều kiện đúng.

### Xử lý lỗi thế nào?

Cách xử lý tuỳ vào từng lỗi cụ thể. Hãy tự hỏi: **"Nếu có gì sai ở đây, mình thực sự muốn chuyện gì xảy ra?"**

Ví dụ, file không tồn tại thì có 2 lựa chọn:

| Lựa chọn | Hành động |
|---|---|
| 1 | In lỗi ra, rồi trả về `newDeck()` để người gọi vẫn có bộ bài dùng được |
| 2 | In lỗi ra, rồi **thoát hẳn chương trình** |

Khoá học chọn **lựa chọn 2**: không đọc được file nghĩa là có gì đó sai nghiêm trọng (tưởng có bộ bài ở đó nhưng không có), nên dừng chương trình.

### Package `os` và `os.Exit`

- Package **`os`** cung cấp giao diện làm việc với hệ điều hành, **chạy như nhau trên Windows, Linux, macOS**.
- **`os.Exit(code)`** thoát chương trình ngay lập tức:
  - `0`: chương trình chạy thành công.
  - Khác `0`: báo có lỗi xảy ra.

### Code đến thời điểm này

`deck.go`:

```go
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Lựa chọn 2: in lỗi và thoát chương trình
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs chứa dữ liệu đọc được -> biến thành deck (phần sau)
}
```

- `bs` (byte slice) và `err` là biến mới, nên dùng `:=`.
- In `"Error:"` trước giá trị lỗi để nhìn là biết ngay đó là thông báo lỗi.
- Nhớ import `os` (VS Code tự thêm khi lưu file).

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `ioutil.ReadFile(filename)` | Đọc file, trả về `([]byte, error)` |
| `nil` | "Không có giá trị"; `err == nil` nghĩa là không có lỗi |
| `if err != nil { ... }` | Mẫu kiểm tra lỗi phổ biến nhất trong Go |
| `os.Exit(1)` | Thoát chương trình với mã lỗi khác 0 |
| Phần tiếp theo | Biến `bs` (`[]byte`) thành `deck` |

## Error Handling

Hoàn thiện `newDeckFromFile`: sau khi qua bước kiểm tra lỗi, `bs` chắc chắn có dữ liệu, giờ biến nó thành `deck`.

### Quy trình ngược với `saveToFile`

```
saveToFile:       deck  ->  []string  ->  string  ->  []byte
newDeckFromFile:  []byte  ->  string  ->  []string  ->  deck
```

`bs` là byte slice biểu diễn mọi ký tự của chuỗi lá bài, mỗi lá cách nhau bởi dấu phẩy.

### Bước 1: `[]byte` → `string`

```go
string(bs) // "Ace of Spades,Two of Spades,Three of Spades,..."
```

### Bước 2: `string` → `[]string` với `strings.Split`

Ngược với `strings.Join`, package `strings` có hàm **`Split`**:

```go
func Split(s, sep string) []string
```

| Đối số | Kiểu | Ý nghĩa |
|---|---|---|
| `s` | `string` | Chuỗi cần tách |
| `sep` | `string` | Dấu phân tách, ở đây là `","` |

```go
s := strings.Split(string(bs), ",")
```

> Tên biến 1 chữ cái (`s`, `bs`, `d`, ...) rất phổ biến trong Go, đây là quy ước của ngôn ngữ.

### Bước 3: `[]string` → `deck`

```go
deck(s)
```

Chuyển được vì `deck` được định nghĩa từ `[]string`. Có thể hình dung: lấy một `[]string` và **gắn thêm cho nó các hàm** của `deck`.

### Code hoàn chỉnh

`deck.go`:

```go
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
```

`main.go`:

```go
func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}
```

### Chạy thử

- Với file `my_cards` đã lưu trước đó: chạy `go run main.go deck.go` sẽ in ra toàn bộ lá bài.
- Với file **không tồn tại** (ví dụ `"my"`), lỗi được xử lý đúng như mong muốn:

```
Error: open my: no such file or directory
exit status 1
```

`exit status 1` đến từ `os.Exit(1)`, báo chương trình kết thúc do có lỗi.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `string(bs)` | Chuyển byte slice thành string |
| `strings.Split(s, sep)` | Tách chuỗi theo `sep`, trả về `[]string` (ngược với `Join`) |
| `deck(s)` | Chuyển `[]string` thành `deck` |
| `if err != nil` | Mẫu xử lý lỗi phổ biến trong Go |
| Phần tiếp theo | Viết hàm `shuffle` để xáo trộn bộ bài |

## Shuffling a Deck

Hàm lớn cuối cùng: `shuffle`, **sắp xếp ngẫu nhiên thứ tự** các lá bài trong bộ bài.

```
Trước:  [Ace, Two, Three, Four]
Sau:    [Three, Four, Ace, Two]   (ngẫu nhiên)
```

### Thuật toán

Khoá học tự viết logic xáo trộn như sau:

1. Lặp qua từng phần tử (theo **index**) của slice.
2. Với mỗi index, sinh một **số ngẫu nhiên** trong khoảng từ 0 đến độ dài slice.
3. **Hoán đổi** lá bài ở index hiện tại với lá bài ở vị trí ngẫu nhiên vừa sinh.

Ví dụ với `[Ace, Two, Three, Four]`:

- `i = 0` (Ace), sinh được `3` → đổi Ace với Four → `[Four, Two, Three, Ace]`
- `i = 1` (Two), sinh được `0` → đổi Two với Four → `[Two, Four, Three, Ace]`
- ... lặp tiếp đến hết slice.

Có thể không phải cách xáo trộn tốt nhất, nhưng chắc chắn làm thứ tự trở nên ngẫu nhiên.

> Ghi chú: từ Go 1.10 đã có sẵn `rand.Shuffle`, nhưng khoá học tự viết để học cách làm.

### Sinh số ngẫu nhiên: `math/rand`

- Package **`rand`** là package con của **`math`**, import bằng `"math/rand"`.
- Hàm **`rand.Intn(n)`** trả về một **số nguyên** ngẫu nhiên trong khoảng `[0, n)`.
- Vì kết quả là số nguyên nên dùng thẳng làm index của slice được.

### `len`: độ dài slice

`len(slice)` trả về một `int` là **số phần tử** của slice.

### Thiết kế chữ ký hàm

| Thành phần | Lựa chọn | Lý do |
|---|---|---|
| Receiver | `d deck` | Để gọi `cards.shuffle()` |
| Đối số | Không có | Chỉ làm việc trên chính bộ bài |
| Kiểu trả về | Không có | Chỉ đổi thứ tự bên trong bộ bài, không có gì để trả về |

### Code

`deck.go`:

```go
import "math/rand" // thêm vào khối import

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```

- **`for i := range d`**: chỉ lấy **index**, không lấy phần tử, vì ở đây chỉ cần index. `i` là tên viết tắt quen dùng cho index.
- **`newPosition`**: tên dài để dễ hiểu (theo quy ước Go thì thường đặt ngắn hơn).
- **Hoán đổi một dòng**: `d[i], d[newPosition] = d[newPosition], d[i]`. Gán giá trị ở `newPosition` cho `i`, đồng thời gán giá trị ở `i` cho `newPosition`.

> Ghi chú: `rand.Intn(len(d) - 1)` trả về số trong `[0, len(d)-1)`, nên **index cuối cùng không bao giờ được chọn** làm vị trí ngẫu nhiên. Dùng `rand.Intn(len(d))` mới đúng là chọn trong toàn bộ slice.

`main.go`:

```go
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
```

### Điều kỳ lạ

Chạy `go run main.go deck.go` thì thứ tự lá bài **có vẻ ngẫu nhiên**. Nhưng **chạy lại nhiều lần thì kết quả lần nào cũng giống hệt nhau** (ví dụ 4 lá cuối lần nào cũng là Two of Spades, Two of Clubs, Four of Diamonds, Three of Diamonds).

Ngẫu nhiên thật thì không thể lần nào cũng ra cùng một thứ tự, nên `shuffle` vẫn chưa đúng như mong đợi. Phần tiếp theo sẽ tìm hiểu nguyên nhân.

> Ghi chú: từ **Go 1.20**, `math/rand` tự khởi tạo seed ngẫu nhiên khi chương trình chạy, nên với Go mới sẽ **không còn thấy** hiện tượng này.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `"math/rand"` | Package con của `math` dùng để sinh số ngẫu nhiên |
| `rand.Intn(n)` | Số nguyên ngẫu nhiên trong `[0, n)` |
| `len(slice)` | Độ dài của slice |
| `for i := range d` | Lặp chỉ lấy index |
| `a[i], a[j] = a[j], a[i]` | Hoán đổi hai phần tử trong một dòng |
| Phần tiếp theo | Tìm hiểu vì sao lần chạy nào cũng ra cùng một thứ tự |

## Random Number Generation

Sửa lỗi của `shuffle`: lần chạy nào bộ bài cũng bị xáo trộn theo **cùng một thứ tự**.

### Nguyên nhân: bộ sinh số giả ngẫu nhiên và seed

- Logic vòng lặp và hoán đổi **không sai**. Vấn đề nằm ở bộ sinh số ngẫu nhiên.
- `math/rand` là bộ sinh số **giả ngẫu nhiên** (pseudo-random): dãy số nó tạo ra phụ thuộc vào một giá trị **seed** (hạt giống), có thể hình dung là "nguồn ngẫu nhiên".
- Cùng seed thì **luôn ra cùng một dãy số**.
- Theo mặc định, bộ sinh **luôn dùng cùng một seed**, nên mỗi lần chạy lại chương trình ta nhận đúng dãy số cũ, và bộ bài bị xáo theo đúng thứ tự cũ.

**Cách sửa:** mỗi lần chương trình chạy, tạo một **seed mới** rồi dùng nó để sinh số ngẫu nhiên.

### Lần theo tài liệu `math/rand`

Phần này khó tự tìm ra lần đầu. Phần lớn việc học Go là **học cách lần theo tài liệu** của thư viện chuẩn:

1. Kiểu **`rand.Rand`**: một nguồn sinh số ngẫu nhiên mà ta **tự chỉ định được seed**.
2. Tạo `Rand` bằng **`rand.New(src Source) *Rand`**, cần truyền vào một `Source`.
3. Tạo `Source` bằng **`rand.NewSource(seed int64) Source`**, cần một số `int64`.
4. Vậy chỉ cần một số `int64` **khác nhau ở mỗi lần chạy**.

### Lấy seed từ thời gian hiện tại: package `time`

- **`time.Now()`** trả về một giá trị kiểu `time.Time` là thời điểm hiện tại.
- Kiểu `Time` có hàm **`UnixNano()`** trả về `int64`: số nano giây tính từ mốc Unix (1/1/1970).
- Mỗi lần chạy chương trình thời điểm đều khác, nên seed khác và dãy số ngẫu nhiên cũng khác.

### Code

`deck.go`:

```go
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```

- `source`: nguồn ngẫu nhiên, seed là thời gian hiện tại tính bằng nano giây.
- `r`: bộ sinh số ngẫu nhiên kiểu `*rand.Rand` tạo từ `source`.
- Trong vòng lặp, gọi **`r.Intn(...)`** thay cho `rand.Intn(...)`. Kiểu `Rand` cũng có hàm `Intn`: nhận một `int`, trả về một `int`.

### Chạy thử

Chạy `go run main.go deck.go` nhiều lần: thứ tự các lá bài **khác nhau mỗi lần chạy**.

> Ghi chú với Go mới:
> - Từ **Go 1.20**, các hàm cấp package như `rand.Intn` đã **tự seed ngẫu nhiên**, nên không cần bước này nữa (và `rand.Seed` đã bị deprecated).
> - Cách `rand.New(rand.NewSource(...))` vẫn hữu ích khi cần **một bộ sinh riêng** hoặc cần seed cố định để tái tạo kết quả (ví dụ khi test).

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Pseudo-random | Dãy số "ngẫu nhiên" được quyết định hoàn toàn bởi seed |
| Seed | Giá trị khởi đầu; cùng seed thì cùng dãy số |
| `rand.NewSource(seed int64)` | Tạo `Source` từ một seed |
| `rand.New(source)` | Tạo bộ sinh số ngẫu nhiên `*rand.Rand` |
| `time.Now().UnixNano()` | Thời gian hiện tại dạng `int64` (nano giây), dùng làm seed |
| `r.Intn(n)` | Sinh số ngẫu nhiên trong `[0, n)` bằng bộ sinh `r` |

## Testing With Go

Các hàm chính của dự án Cards đã xong. Việc quan trọng cuối cùng là **viết test** để đảm bảo các hàm trong `deck.go` luôn chạy đúng.

### 3 điều cần biết về test trong Go

1. **Không giống các framework test khác** (RSpec, Mocha, Jasmine, Selenium, ...). Go chỉ cung cấp **một bộ hàm rất nhỏ** để hỗ trợ test, còn lại ta **viết code Go bình thường** để kiểm tra code có chạy đúng không.
2. **Tạo file test:** tạo file mới trong thư mục dự án, tên **kết thúc bằng `_test.go`**. Ví dụ, để test code trong `deck.go` thì tạo `deck_test.go`.
3. **Chạy test:** chạy lệnh **`go test`** để chạy toàn bộ test trong package.

### Chạy `go test` khi chưa có test

Có thể chạy ngay cả khi chưa có file test nào; Go sẽ báo là **không tìm thấy file test**:

```
$ go test
?   	cards	[no test files]
```

> Ghi chú: với Go mới (dùng module), thư mục dự án cần có `go.mod` (tạo bằng `go mod init cards`) thì `go test` mới chạy được.

### Tạo `deck_test.go`

Mọi file trong thư mục dự án, **kể cả file test**, đều phải khai báo package ở dòng đầu:

```go
package main
```

### Tích hợp với VS Code

Extension Go của VS Code nhận ra file test và hiển thị các **link bấm được** ngay trên code:

| Link | Ý nghĩa |
|---|---|
| `run package tests` | Chạy tất cả test của package |
| `run file tests` | Chạy tất cả test trong file hiện tại |
| `run test` (trên từng hàm) | Chỉ chạy một test |

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Test trong Go | Bộ hàm hỗ trợ rất nhỏ, chủ yếu viết code Go bình thường |
| `xxx_test.go` | Quy ước đặt tên file test |
| `go test` | Chạy tất cả test trong package |
| `package main` | File test cũng phải khai báo package |
| Phần tiếp theo | Viết test đầu tiên |

## Writing Useful Tests

### Test cái gì?

Câu hỏi muôn thuở với mọi framework test: **"Làm sao biết cần test cái gì?"** Với Go, câu trả lời khá đơn giản: hãy tự hỏi **"Mình thực sự quan tâm điều gì ở hàm này?"**, rồi viết vài **khẳng định (assertion) đơn giản** cho những điều đó.

Ví dụ `newDeck` có khá nhiều logic (tạo deck rỗng, danh sách suit/value, vòng `for` lồng nhau), nên dễ bị người khác vô tình làm hỏng. Ba điều dễ kiểm tra:

1. Bộ bài có **đúng số lượng lá** (4 suit × 4 value = **16** lá).
2. Lá **đầu tiên** là **Ace of Spades**.
3. Lá **cuối cùng** là **Four of Clubs**.

Nếu cả 3 đều đúng thì nhiều khả năng `newDeck` đang hoạt động như mong đợi.

### Cấu trúc file test

Mỗi hàm cần test có một hàm test tương ứng, tên là **`Test` + tên hàm**:

| Hàm trong `deck.go` | Hàm test trong `deck_test.go` | Kiểm tra |
|---|---|---|
| `newDeck` | `TestNewDeck` | Số lá, lá đầu, lá cuối |
| `saveToFile` + `newDeckFromFile` | `TestSaveToDeckAndNewDeckFromFile` | Hai hàm liên quan chặt chẽ nên test chung |

Quy trình bên trong một hàm test:

1. Tạo bộ bài mới.
2. Dùng `if` kiểm tra kết quả có đúng như mong đợi không.
3. Nếu sai, **báo cho test handler** biết test này thất bại.

### Viết test đầu tiên

`deck_test.go`:

```go
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
```

- **`Test` viết hoa chữ `T`**: khác các hàm trong `deck.go` (đều viết thường). Lý do sẽ được giải thích ở dự án sau.
- **`t *testing.T`**: hàm test được **test runner của Go tự động gọi** với đối số `t`. `t` chính là **test handler**: có gì sai thì báo qua `t`. (Ý nghĩa dấu `*` sẽ học sau.)
- **`t.Errorf(...)`**: báo test thất bại kèm thông báo lỗi.
- Nhớ `import "testing"` (VS Code tự thêm khi lưu file).

### Format string và `%v`

`Errorf` nhận một **format string** (chuỗi định dạng): các giá trị truyền thêm phía sau sẽ được **chèn vào chuỗi** tại vị trí các ký hiệu như **`%v`**.

```go
t.Errorf("Expected deck length of 16, but got %v", len(d))
// -> Expected deck length of 16, but got 16
```

Nếu truyền giá trị thêm mà không có `%v` trong chuỗi, VS Code sẽ hiện **cảnh báo gạch xanh**.

### Chạy test

- Chạy `go test` trong terminal, hoặc bấm link **`run test`** của VS Code (tiện khi dự án có nhiều test và chỉ muốn chạy một test).
- Test đúng thì in ra **`PASS`**.
- **Kiểm tra lại chính test**: thử đổi `16` thành số sai (ví dụ `2000`) để xem test có báo **`FAIL`** không, rồi sửa lại.

```
$ go test
--- FAIL: TestNewDeck (0.00s)
    deck_test.go:9: Expected deck length of 16, but got 16
FAIL
```

> Mẹo: thông báo lỗi cũng cần **mô tả đúng**. Khi đổi giá trị kỳ vọng, nhớ sửa luôn số trong thông báo.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Test cái gì | Những điều mình quan tâm, diễn đạt bằng các khẳng định đơn giản |
| `func TestXxx(t *testing.T)` | Hàm test, được test runner tự gọi |
| `t` (`*testing.T`) | Test handler, dùng để báo lỗi |
| `t.Errorf(format, ...)` | Báo test thất bại với thông báo định dạng |
| `%v` | Chỗ chèn giá trị vào format string |
| Phần tiếp theo | Viết thêm test kiểm tra lá đầu và lá cuối |

## Asserting Elements in a Slice

Viết tiếp `TestNewDeck` với 2 kiểm tra còn lại: **lá đầu tiên** và **lá cuối cùng**.

> Lưu ý: thứ tự suit/value của bạn có thể khác, nên lá đầu và lá cuối cũng có thể khác. Ở đây lá đầu là **Ace of Spades**, lá cuối là **Four of Clubs**.

### Truy cập phần tử đầu và cuối

| Phần tử | Cú pháp |
|---|---|
| Đầu tiên | `d[0]` |
| Cuối cùng | `d[len(d)-1]` |

Truy cập theo index giống ký hiệu mảng ở nhiều ngôn ngữ khác (ở đây là slice, nhưng cách dùng như nhau).

### Code hoàn chỉnh

`deck_test.go`:

```go
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
```

Chạy `go test` → `PASS`.

### Go không đếm số "test case"

- Khác các framework khác (thường báo kiểu "đã chạy 30 test, 2 test lỗi"), **Go không biết** bên trong `TestNewDeck` có 3 kiểm tra riêng.
- Go chỉ biết **có một hàm `TestNewDeck`**: chạy hàm đó, và nếu có `t.Errorf` nào được gọi thì hàm test đó **thất bại**.
- Vì vậy kết quả không hiện số lượng test đã chạy.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `d[0]` | Phần tử đầu tiên của slice |
| `d[len(d)-1]` | Phần tử cuối cùng của slice |
| Nhiều `if` trong một hàm test | Mỗi `if` là một khẳng định, nhưng Go chỉ tính theo **hàm** test |
| Phần tiếp theo | Test `saveToFile` và `newDeckFromFile` |

## Testing File IO

Test cuối cùng của dự án: kiểm tra `saveToFile` và `newDeckFromFile`, tức là **lưu bộ bài ra file rồi đọc lại** được.

### Vấn đề: dọn dẹp file tạm

Giả sử test tạo bộ bài và lưu ra file thành công, nhưng khi **đọc lại thì bị lỗi và test dừng giữa chừng**. File tạm vừa tạo **vẫn nằm lại trên ổ cứng**, vì test chưa kịp xoá nó.

- **Framework test của Go không tự dọn dẹp:** nó không biết ta đã ghi file nào và không tự xoá giúp.
- **Mọi việc dọn dẹp phải tự làm thủ công.**

### Cách giải quyết

Luôn dùng **cùng một tên file tạm** (`_decktesting`) và xoá nó **cả trước lẫn sau** khi test:

1. **Xoá** mọi file tên `_decktesting` trong thư mục hiện tại (phòng khi lần chạy trước bị lỗi và để sót lại), để test luôn bắt đầu từ **môi trường sạch**.
2. Tạo bộ bài mới.
3. Lưu vào file `_decktesting`.
4. Đọc lại từ file `_decktesting`.
5. Kiểm tra độ dài bộ bài.
6. **Xoá** file `_decktesting` một lần nữa.

> Dấu `_` ở đầu tên file không có ý nghĩa đặc biệt, chỉ để làm rõ đây là file tạm (và khó trùng với file quan trọng nào đó).

### Xoá file: `os.Remove`

```go
func Remove(name string) error
```

- Xoá file (hoặc thư mục rỗng) có tên `name`.
- Trả về `error`. Trong test này ta **bỏ qua lỗi**: nếu file chưa tồn tại thì cũng chẳng sao.

### Đặt tên hàm test

```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T)
```

- Luôn bắt đầu bằng **`Test`** (chữ `T` viết hoa).
- Tên dài nhưng **chứa đúng tên các hàm được test**, để sau này khi sửa `newDeckFromFile` chỉ cần copy tên hàm và tìm trong file test là ra ngay các test liên quan.

### Code

`deck_test.go`:

```go
import (
	"os"
	"testing"
)

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
```

- Biến `loadedDeck` đặt **tên đầy đủ** thay vì một chữ cái, vì khó có tên ngắn nào vừa rõ nghĩa (đặt `d2` thì khó hiểu).
- Chạy `go test` → `PASS`. Đổi `16` thành `160` để chắc test có **`FAIL`**, rồi sửa lại.

> Ghi chú với Go mới:
> - `t.TempDir()` (Go 1.15+) tạo thư mục tạm và **tự xoá** sau khi test xong; `t.Cleanup(func)` (Go 1.14+) đăng ký hàm dọn dẹp luôn được chạy khi test kết thúc. Hai cách này an toàn hơn xoá file thủ công.
> - Nếu đọc file lỗi, `newDeckFromFile` gọi `os.Exit(1)`, làm **dừng luôn cả tiến trình test**. Đây là lý do test cần chuẩn bị môi trường sạch ngay từ đầu.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Go test không tự dọn dẹp | Phải tự xoá file tạm trước và sau test |
| `os.Remove(name)` | Xoá file, trả về `error` |
| `_decktesting` | Tên file tạm cố định dùng trong test |
| `TestSaveToDeckAndNewDeckFromFile` | Tên test chứa tên các hàm được test để dễ tìm |
| Phần tiếp theo | Ôn lại dự án Cards rồi chuyển sang chủ đề mới |

## Project Review

Hoàn thành dự án lớn đầu tiên với Go. Phần này ôn lại các điểm chính và nêu ra vài điều **vẫn còn bí ẩn**, sẽ đào sâu ở các phần sau.

### Ôn lại `deck.go`

| Điểm chính | Ghi nhớ |
|---|---|
| `package main` | Dòng đầu tiên của **mọi file** trong thư mục dự án |
| `import ( ... )` | Import nhiều package: một lệnh `import`, bọc trong ngoặc đơn, mỗi package một dòng, **không dấu phẩy** |
| `type deck []string` | Kiểu tùy chỉnh: một lớp mỏng bọc `[]string`, vẫn hoạt động như slice nhưng **gắn thêm được hàm riêng** |
| `func newDeck() deck` | Tạo bộ bài mới từ đầu; kiểu trả về `deck` đặt **sau tên hàm** |
| `func (d deck) print()` | **Receiver**: mọi giá trị kiểu `deck` đều gọi được `.print()`; rất phổ biến trong code Go và tài liệu |
| Slice | Truy cập giống mảng ở ngôn ngữ khác, nhưng có thêm tính năng như **cắt khoảng** `[start:end]` để lấy một phần slice |

> Slice trong Go còn **một số khác biệt quan trọng** so với mảng truyền thống ở các ngôn ngữ khác, sẽ tìm hiểu sau.

### Điều kỳ lạ 1: vì sao `deal` không dùng receiver?

Hầu hết các hàm đều dùng receiver `deck`, riêng `deal` lại nhận `deck` làm **đối số**:

```go
func deal(d deck, handSize int) (deck, deck)

hand, remainingCards := deal(cards, 5)
```

Nếu viết bằng receiver:

```go
func (d deck) deal(handSize int) (deck, deck)

cards.deal(5)
```

Thì `cards.deal(5)` dễ gây hiểu lầm: đọc vào sẽ tưởng là **lấy 5 lá ra khỏi `cards`**, tức là **slice `cards` bị sửa đổi** và chỉ còn ít hơn 5 lá. Thực tế `deal` **không sửa** `cards`, mà trả về hai slice mới.

Không có quy tắc nào bắt buộc phải làm thế; đây chỉ là chọn cách viết **tránh gây hiểu lầm**. Câu hỏi "hàm có **sửa đổi** dữ liệu gốc hay không" chính là trọng tâm của chủ đề lớn tiếp theo.

### Điều kỳ lạ 2: dấu `*` trong `*testing.T`

```go
func TestNewDeck(t *testing.T)
```

Dấu `*` trước `testing.T` trông khá bí ẩn. Nó gắn liền với cùng chủ đề trên, sẽ được giải thích ở phần sau.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Kiểu tùy chỉnh `deck` | Bọc `[]string` để gắn thêm hàm riêng |
| Receiver | Cho phép gọi `value.method()` trên mọi giá trị của kiểu |
| `deal(d, n)` thay vì `d.deal(n)` | Tránh hiểu lầm là `deal` sửa đổi bộ bài gốc |
| `*testing.T` | Dấu `*` còn bí ẩn, sẽ học ở phần sau |
| Phần tiếp theo | Đi sâu vào chủ đề phức tạp hơn: có sửa đổi dữ liệu gốc hay không |

# 4_organizing_data_with_structs

# Struct:

## Structs in Go

### Vấn đề của dự án Cards

Trong dự án Cards, mỗi lá bài là một `string` (ví dụ `"Ace of Spades"`). Nếu mở rộng thành game poker, blackjack..., ta sẽ cần hỏi **lá này chất gì, giá trị bao nhiêu?** Với `string`, phải tách chuỗi theo `" of "` rồi lấy ra giá trị và chất: làm được nhưng **rất khó xử** và gắn chặt chương trình vào việc "lá bài là chuỗi".

### Struct là gì?

**Struct** (viết tắt của *structure*) là một cấu trúc dữ liệu gom **nhiều thuộc tính có liên quan** lại với nhau.

```go
type card struct {
    suit  string // "Spades"
    value string // "Ace"
}
```

Mỗi giá trị kiểu `card` có `suit` và `value` riêng, truy cập trực tiếp mà không cần thao tác chuỗi.

### So sánh với ngôn ngữ khác

Ở mức khái quát (không hoàn toàn chính xác), struct giống:

| Ngôn ngữ | Tương đương |
|---|---|
| JavaScript | Object đơn giản |
| Ruby | Hash |
| Python | Dictionary |

### Chuẩn bị dự án

Tạo thư mục mới `structs` và file `main.go` bên trong để thử nghiệm struct qua vài ví dụ nhỏ.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Lá bài là `string` | Muốn lấy chất/giá trị phải tách chuỗi, khó mở rộng |
| Struct | Gom các thuộc tính liên quan thành một kiểu dữ liệu |
| `card{suit, value}` | Cách biểu diễn lá bài tốt hơn |
| Phần tiếp theo | Viết code ví dụ với struct trong `structs/main.go` |

## Defining Structs

Tạo struct gồm 2 bước: **(1) định nghĩa** các trường và kiểu của chúng, **(2) tạo giá trị** theo định nghĩa đó (tạo bao nhiêu cũng được).

```go
package main

type person struct {
    firstName string
    lastName  string
}

func main() {}
```

- Cú pháp giống `type deck []string`: khai báo **kiểu tùy chỉnh mới** `person`.
- Các trường **không ngăn cách** bằng dấu phẩy/chấm phẩy: chỉ `tên kiểu` trên mỗi dòng.

## Declaring Structs

### Cách 1: theo thứ tự trường (không khuyến khích)

```go
alex := person{"Alex", "Anderson"}
```

Go gán theo **thứ tự định nghĩa** trường. Nếu sau này đổi thứ tự `firstName`/`lastName` thì dữ liệu bị gán sai mà không báo lỗi.

### Cách 2: chỉ rõ tên trường (khuyến khích)

```go
alex := person{firstName: "Alex", lastName: "Anderson"}
fmt.Println(alex) // {Alex Anderson}
```

Không phụ thuộc thứ tự trường, an toàn khi thêm/đổi trường.

## Updating Struct Values

### Cách 3: `var` và zero value

```go
var alex person
fmt.Println(alex)         // { }
fmt.Printf("%+v", alex)   // {firstName: lastName:}
```

Khi không gán giá trị, Go dùng **zero value** cho từng trường (không phải `nil`/`undefined` như JS, Ruby):

| Kiểu | Zero value |
|---|---|
| `string` | `""` |
| `int`, `float64` | `0` |
| `bool` | `false` |

`%+v` in ra **cả tên trường lẫn giá trị**.

### Cập nhật trường bằng dấu chấm

```go
alex.firstName = "Alex"
alex.lastName = "Anderson"
```

## Embedding Structs

Một struct có thể chứa **struct khác** làm trường (không chỉ kiểu cơ bản):

```go
type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contact   contactInfo
}

jim := person{
    firstName: "Jim",
    lastName:  "Party",
    contact: contactInfo{
        email:   "jim@gmail.com",
        zipCode: 94000,
    },
}
fmt.Printf("%+v", jim)
```

> Khai báo nhiều dòng: **mỗi dòng phải có dấu phẩy, kể cả dòng cuối** (dễ quên nếu quen JS/Ruby).

`contactInfo` vẫn dùng độc lập được, hoặc nhúng vào struct khác (ví dụ `business`).

## Structs with Receiver Functions

### Nhúng không cần tên trường

```go
type person struct {
    firstName string
    lastName  string
    contactInfo // tương đương: contactInfo contactInfo
}

jim := person{
    // ...
    contactInfo: contactInfo{ /* ... */ },
}
```

Chỉ ghi kiểu thì tên trường **tự lấy theo tên kiểu**. Điều này sẽ quan trọng khi nói về tái sử dụng code sau này.

### Receiver với struct

```go
func (p person) print() {
    fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
}

jim.updateName("Jimmy")
jim.print() // firstName vẫn là Jim!
```

Không lỗi, nhưng **tên không được cập nhật**. Lý do nằm ở chủ đề tiếp theo: con trỏ (pointer).

## Pass By Value

**RAM** có thể hình dung như dãy ô nhớ, mỗi ô có **địa chỉ** và chứa dữ liệu. `jim` trỏ đến ô chứa struct, ví dụ địa chỉ `0001`.

Go là ngôn ngữ **truyền theo giá trị (pass by value)**: khi truyền giá trị vào hàm (kể cả receiver), Go **sao chép** dữ liệu sang ô nhớ mới, hàm làm việc với **bản sao**.

| Địa chỉ | Giá trị |
|---|---|
| `0001` | `person{firstName: "Jim"...}` ← `jim` |
| `0002` | bản sao `person{...}` ← `p` trong `updateName` |

→ `p.firstName = ...` chỉ sửa bản sao, `jim` gốc không đổi.

## Structs with Pointers

Sửa code để cập nhật được bản gốc:

```go
jimPointer := &jim
jimPointer.updateName("Jimmy")
jim.print() // firstName: Jimmy

func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}
```

## Pointer Operations

| Toán tử | Ý nghĩa |
|---|---|
| `&value` | Lấy **địa chỉ bộ nhớ** của giá trị (tạo con trỏ) |
| `*pointer` | Lấy **giá trị** tại địa chỉ mà con trỏ trỏ tới |

- `jimPointer := &jim`: `jimPointer` chứa địa chỉ (ví dụ `0xc000010200`), không phải struct.
- `(*pointerToPerson)`: biến địa chỉ trở lại thành struct `jim` thật, nên sửa trường là sửa bản gốc.

### `*` trước kiểu vs `*` trước con trỏ (dễ nhầm nhất)

| Vị trí | Ví dụ | Ý nghĩa |
|---|---|---|
| Trước **kiểu** | `*person` | **Mô tả kiểu**: con trỏ tới `person` (không phải toán tử) |
| Trước **biến con trỏ** | `*pointerToPerson` | **Toán tử**: lấy giá trị tại địa chỉ |

### Quy tắc cần nhớ

- Có **địa chỉ** → thành giá trị bằng `*address`.
- Có **giá trị** → thành địa chỉ bằng `&value`.

## Pointer Shortcut

Nếu receiver là kiểu con trỏ (`*person`), Go cho phép gọi trực tiếp trên **giá trị** hoặc **con trỏ**, tự động lấy địa chỉ:

```go
jim.updateName("Jimmy") // Go tự hiểu là (&jim).updateName("Jimmy")
```

Không cần bước `jimPointer := &jim`. Cách viết này rất phổ biến trong code Go thực tế.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `type person struct {...}` | Định nghĩa struct, trường không có dấu phẩy |
| `person{firstName: "..."}` | Tạo struct, nên ghi rõ tên trường |
| `var p person` | Struct với zero value (`""`, `0`, `false`) |
| `%+v` | In struct kèm tên trường |
| Struct lồng nhau | Trường có kiểu là struct khác; dòng cuối vẫn cần dấu phẩy |
| Pass by value | Truyền vào hàm là truyền **bản sao** |
| `&` / `*` | Lấy địa chỉ / lấy giá trị tại địa chỉ |
| Receiver `*person` | Cho phép hàm sửa struct gốc |

# Map

## Maps

**Map** là tập hợp các cặp **key–value**, tương tự `hash` trong Ruby, object trong JavaScript hay `dict` trong Python.

- Mọi **key** phải **cùng một kiểu**, mọi **value** cũng phải **cùng một kiểu**.
- Kiểu của key và kiểu của value **không bắt buộc giống nhau** (ví dụ `map[string]int`).

### Khai báo bằng literal

```go
colors := map[string]string{
	"red":   "#ff0000",
	"green": "#4bf745",
}
fmt.Println(colors) // map[green:#4bf745 red:#ff0000]
```

| Phần | Ý nghĩa |
|---|---|
| `map[string]` | Kiểu của key là `string` |
| `string` (sau `]`) | Kiểu của value là `string` |
| `"red": "#ff0000",` | Một cặp key–value; **mỗi cặp đều phải có dấu phẩy**, kể cả cặp cuối |

## Manipulating Maps

### Ba cách tạo map

```go
colors := map[string]string{"red": "#ff0000"} // 1. literal, có sẵn giá trị
var colors map[string]string                  // 2. var → zero value
colors := make(map[string]string)             // 3. hàm có sẵn make → map rỗng
```

- Zero value của map in ra là `map[]` (rỗng).
- **Lưu ý:** map khai báo bằng `var` thực chất là **`nil` map**: đọc thì được, nhưng **gán key mới sẽ panic** (`assignment to entry in nil map`). Muốn thêm phần tử sau này thì dùng `make` hoặc literal.

### Thêm / sửa phần tử: cú pháp `[]`

```go
colors := make(map[string]string)
colors["white"] = "#ffffff"
```

- Map **không dùng dấu chấm** như struct (`colors.white` là sai), vì key có kiểu, không phải tên trường. Ví dụ với `map[int]string` thì viết `colors[10] = "#ffffff"`.

### Xóa phần tử: `delete`

```go
delete(colors, 10) // delete(tênMap, key)
```

## Iterating Over Maps

Duyệt map bằng `for ... range`, giống slice nhưng nhận về **key và value** thay vì index và value:

```go
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}
```

- Truyền map vào hàm: khai báo đối số dạng `c map[string]string`.
- `color, hex` là biến mới mỗi vòng lặp nên dùng `:=`.
- Thứ tự duyệt map **không cố định**, mỗi lần chạy có thể khác.

## Differences Between Maps and Structs

| | Map | Struct |
|---|---|---|
| Kiểu key | Mọi key cùng một kiểu | Không có "kiểu key", tên trường cố định |
| Kiểu value | Mọi value cùng một kiểu | Mỗi trường có thể khác kiểu |
| Truy cập | `m["key"]` | `s.field` |
| Duyệt qua (iterate) | Được, key được đánh chỉ mục | Không duyệt được các trường |
| Tập key | **Không cần biết** lúc biên dịch, thêm/xóa tùy ý | Phải khai báo **hết** lúc biên dịch |
| Bản chất | **Reference type**: truyền vào hàm, sửa map trong hàm thì map gốc cũng đổi | **Value type**: truyền vào hàm là bản sao |
| Dùng khi | Biểu diễn **tập các giá trị liên quan chặt chẽ** cùng loại (ví dụ tên màu → mã hex) | Biểu diễn **một "thứ"** có nhiều thuộc tính khác nhau (ví dụ `person`) |

- Nếu không biết trước tập key sẽ gồm những gì → nghĩ tới **map**.
- Nếu tập key cố định và biết trước (luôn chỉ có `red`, `green`, `white`) → cân nhắc **struct**.
- Trong code Go thực tế, **struct được dùng nhiều hơn map**, nhưng còn tùy loại ứng dụng.

### Tóm tắt

| Cú pháp | Ý nghĩa |
|---|---|
| `map[K]V{k: v,}` | Tạo map bằng literal |
| `make(map[K]V)` | Tạo map rỗng, dùng được ngay |
| `var m map[K]V` | `nil` map, không gán được phần tử |
| `m[k] = v` | Thêm hoặc sửa phần tử |
| `delete(m, k)` | Xóa phần tử theo key |
| `for k, v := range m { }` | Duyệt qua mọi cặp key–value |

# Interface

## Purpose of Interfaces

Hai điều đã biết:

- Mọi giá trị trong Go đều **có kiểu**, kể cả khi dùng `:=` để Go tự suy ra.
- Mọi hàm đều phải **chỉ rõ kiểu** của đối số hoặc receiver.

**Vấn đề:** hàm `shuffle` trong dự án Cards gắn với kiểu `deck`, nhưng logic bên trong không quan tâm đó là bộ bài, `[]string`, `[]int` hay `[]float64`. Nếu mỗi kiểu phải viết một `shuffle` riêng thì code sẽ bị **lặp lại** rất nhiều, dù logic giống hệt nhau.

→ **Interface** giúp **tái sử dụng code** cho nhiều kiểu khác nhau.

### Chương trình mẫu: hai chat bot

| | `englishBot` | `spanishBot` |
|---|---|---|
| `getGreeting() string` | Logic **riêng** cho tiếng Anh, trả về `"Hi There!"` | Logic **riêng** cho tiếng Tây Ban Nha, trả về `"Hola!"` |
| `printGreeting(...)` | `fmt.Println(eb.getGreeting())` | `fmt.Println(sb.getGreeting())` |

- `getGreeting`: cách cài đặt **khác nhau** cho từng bot.
- `printGreeting`: logic **giống hệt nhau**, chỉ khác kiểu đối số → nên dùng chung một hàm.

## Problems Without Interfaces

Cách viết "ngây thơ", chưa dùng interface:

```go
type englishBot struct{} // struct không có trường
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// logic rất riêng để tạo lời chào tiếng Anh
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) { fmt.Println(eb.getGreeting()) }
func printGreeting(sb spanishBot) { fmt.Println(sb.getGreeting()) } // lỗi: printGreeting redeclared
```

- `englishBot{}`: tạo giá trị của struct không có trường.
- Nếu receiver **không dùng** trong thân hàm, có thể **bỏ tên biến**, chỉ giữ kiểu: `func (englishBot) getGreeting()`.
- **Go không hỗ trợ overloading** (khác Java, C#): không thể khai báo hai hàm cùng tên dù kiểu đối số khác nhau.

## Interfaces in Practice

Refactor: xóa hai hàm `printGreeting`, thêm interface `bot` và một `printGreeting` duy nhất.

```go
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { return "Hi There!" }
func (spanishBot) getGreeting() string { return "Hola!" }
```

Chạy `go run main.go`:

```
Hi There!
Hola!
```

### Interface hoạt động thế nào?

`type bot interface { getGreeting() string }` có thể hiểu là:

> "Mọi kiểu trong chương trình có hàm `getGreeting()` trả về `string` thì **cũng được coi là kiểu `bot`**."

- `englishBot` và `spanishBot` đều có `getGreeting() string` → **tự động** trở thành `bot` (không cần khai báo `implements` như Java).
- `printGreeting(b bot)` nhận bất kỳ giá trị nào thỏa interface `bot` → truyền được cả `eb` lẫn `sb`.
- Thêm bot mới (Pháp, Ý, ...) chỉ cần cài `getGreeting() string`, không phải viết lại `printGreeting`.

## Rules of Interfaces

### Cú pháp đầy đủ

```go
type bot interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float64
	respondToUser(user) string
}
```

| Phần | Ý nghĩa |
|---|---|
| `type bot interface` | Khai báo kiểu interface mới tên `bot` |
| `getGreeting` | Tên hàm mà kiểu khác phải có |
| `(string, int)` | **Kiểu đối số** (chỉ ghi kiểu, không ghi tên biến) |
| `(string, error)` | **Kiểu trả về**, có thể nhiều giá trị |
| Nhiều dòng | Kiểu khác phải có **đủ tất cả** các hàm mới được coi là `bot` |

- Trong interface, mỗi dòng là `tênHàm(kiểuĐốiSố) kiểuTrảVề`: không có dấu phẩy cuối dòng, không có dấu `:`.

### Concrete type vs Interface type

| | Concrete type | Interface type |
|---|---|---|
| Ví dụ | `map`, `struct`, `int`, `string`, `englishBot` | `bot` |
| Tạo giá trị trực tiếp? | **Được** | **Không** |
| Vai trò | Chứa dữ liệu và cách cài đặt cụ thể | Mô tả tập hàm mà một kiểu phải có |

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Interface | Mô tả **tập hàm** (tên, kiểu đối số, kiểu trả về) |
| Thỏa interface | Kiểu nào có đủ các hàm đó thì **tự động** thuộc kiểu interface |
| Lợi ích | Viết một hàm nhận interface, dùng lại cho nhiều concrete type |
| Overloading | Go **không** hỗ trợ, dùng interface thay thế |
| Receiver không dùng | Có thể bỏ tên: `func (englishBot) getGreeting()` |

## Extra Interface Notes

| Lưu ý | Ý nghĩa |
|---|---|
| Interface **không phải generic type** | Không dùng interface để viết code "nhận mọi kiểu rồi tự xử lý" như generic trong Java/C#. Interface chỉ là công cụ giúp code gọn và ít lặp hơn. |
| Interface là **ngầm định (implicit)** | Không cần viết `englishBot implements bot`. Kiểu nào có đủ hàm thì Go tự coi là thỏa interface. |
| Interface là **hợp đồng (contract)** về kiểu | Chỉ kiểm tra **tên hàm, kiểu đối số, kiểu trả về**, không kiểm tra **logic** bên trong. |
| Interface **khó** | Bước 1 là **đọc hiểu** interface trong tài liệu thư viện chuẩn. Tự viết interface riêng để sau. |

> **Ghi chú:** khóa học được quay khi Go chưa có generics. Từ **Go 1.18** Go đã hỗ trợ generics (`func F[T any](x T)`), nhưng ý chính vẫn đúng: interface và generics là hai công cụ khác nhau.

### Implicit: vừa tiện vừa khó theo dõi

- **Tiện:** không cần viết boilerplate để liên kết `englishBot`, `spanishBot` với `bot`.
- **Khó:** nhìn vào một kiểu, khó biết ngay nó thỏa những interface nào. Viết sai kiểu trả về thì chỉ biết khi **compile báo lỗi**.

### Garbage in, garbage out

```go
func (englishBot) getGreeting() string {
	return "asdfghjkl" // vẫn là string → vẫn thỏa interface bot
}
```

- Code trên **vẫn compile được** vì vẫn trả về `string`, nhưng kết quả sai so với mong muốn (một lời chào).
- Interface giúp tái sử dụng code, **không** phải là bài test tự động, **không** ngăn được lỗi logic.

## The HTTP Package

Mục tiêu: viết chương trình gửi HTTP request tới `google.com` rồi in nội dung phản hồi (HTML) ra terminal.

Tạo thư mục mới `http`, file `main.go`. Tra tài liệu tại [pkg.go.dev/net/http](https://pkg.go.dev/net/http):

```go
func Get(url string) (resp *Response, err error)
```

- `http.Get` nhận URL, trả về **con trỏ** tới `Response` và một `error` (chỉ khác `nil` khi có lỗi).

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp) // cách "ngây thơ"
}
```

- URL phải ghi đầy đủ `http://`, không được chỉ ghi `google.com`.
- Kết quả `go run main.go`: in ra `200 OK`, `HTTP/1.1`, headers, ... nhưng **không thấy HTML** của trang. Khác với Ruby/JS/Python, in thẳng `resp` không cho ra nội dung body.

## Response Struct

`http.Response` là một **struct**:

```go
type Response struct {
	Status     string // "200 OK"
	StatusCode int    // 200
	Proto      string // "HTTP/1.0"
	// ...
	Body io.ReadCloser
	// ...
}
```

- 3 field đầu khớp với `200 OK 200 HTTP/1.1` thấy trên terminal.
- Nội dung HTML nằm trong field **`Body`**, nhưng kiểu của nó là `io.ReadCloser`, một **interface**, chứ không phải `string`.

Lần theo tài liệu:

```go
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

```
Response (struct)
 └── Body: io.ReadCloser (interface)
       ├── Reader (interface) → Read([]byte) (int, error)
       └── Closer (interface) → Close() error
```

## More Interface Syntax

### Interface làm kiểu của field trong struct

```go
type Response struct {
	Body io.ReadCloser
}
```

- Nghĩa là: field `Body` có thể gán **bất kỳ giá trị nào** miễn kiểu của nó thỏa interface `ReadCloser`.
- Ta có thể tự tạo một struct có hàm `Read` và `Close` đúng chữ ký, rồi gán nó vào `Body`. Interface ở đây mang lại sự **linh hoạt**.

### Nhúng (embed) interface vào interface

```go
type ReadCloser interface {
	Reader
	Closer
}
```

- Không liệt kê hàm, mà liệt kê **tên interface khác**: ghép nhiều interface thành một interface mới.
- Muốn thỏa `ReadCloser` thì phải thỏa **cả** `Reader` **và** `Closer`, tức là có đủ `Read(...)` và `Close()`.

## Interface Review

Nhắc lại ví dụ bot:

- `englishBot`, `spanishBot` đều có `getGreeting() string`.
- Interface `bot` nói: kiểu nào có `getGreeting() string` thì cũng là `bot`.
- Nhờ đó chỉ cần **một** hàm `printGreeting(b bot)` dùng chung cho cả hai.

> Lợi ích cốt lõi của interface: **tái sử dụng code**. Với `io.Reader`, lợi ích này lúc đầu chưa thấy rõ, phần dưới giải thích.

## The Reader Interface

### Vấn đề nếu không có interface (thế giới giả tưởng)

Chương trình có thể nhận dữ liệu từ nhiều nguồn, mỗi nguồn trả về một kiểu khác nhau:

| Nguồn input | Kiểu giả tưởng |
|---|---|
| HTTP request body | `[]flargen` |
| File text trên ổ cứng | `[]string` |
| Ảnh | `jpegne` |
| User gõ ở command line | `[]byte` |
| Cảm biến analog | `[]float` |

→ Mỗi kiểu lại phải viết một hàm xử lý riêng (`printHTTP`, `printFile`, `printImage`, ...), dù logic bên trong **giống nhau**. Đây chính là vấn đề đã gặp với `printGreeting` ở ví dụ bot.

### Giải pháp: `io.Reader`

```
HTTP body ──┐
Text file ──┤
Image ──────┼──► io.Reader ──► []byte ──► các hàm dùng chung
Terminal ───┤
Sensor ─────┘
```

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

- Mọi nguồn input chỉ cần cài đặt `Reader` là có một **điểm tiếp xúc chung**.
- Đầu ra chung là **`[]byte`** (có thể coi như một chuỗi).
- Viết hàm làm việc với `io.Reader` / `[]byte` **một lần**, dùng được cho **mọi** nguồn input.
- `io.Reader` xuất hiện ở khắp nơi trong thư viện chuẩn của Go.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `http.Get(url)` | Trả về `*http.Response` và `error` |
| `resp.Body` | Kiểu `io.ReadCloser` (interface), chứa nội dung phản hồi |
| Interface trong struct | Field nhận bất kỳ giá trị nào thỏa interface đó |
| Embed interface | `ReadCloser` = `Reader` + `Closer` |
| `io.Reader` | "Bộ chuyển đổi" đưa mọi nguồn input về `[]byte` chung |
| Interface là contract | Kiểm tra chữ ký hàm, không kiểm tra logic |

## More on the Reader Interface

Mọi nguồn input đưa vào chương trình đều cài đặt `Reader`, ví dụ `resp.Body`. Phần việc của ta chỉ là **dùng** `Reader` để lấy dữ liệu ra, không phải tự viết hàm `Read`.

### Hàm `Read` thực sự làm gì?

```go
Read(p []byte) (n int, err error)
```

```
Code của ta                              resp.Body (cài đặt Reader)
┌──────────────────┐   truyền []byte     ┌──────────────────────────┐
│ tạo []byte rỗng  │ ──────────────────► │ Read(p []byte)           │
│                  │                     │ lấy dữ liệu thô (HTML)   │
│ []byte giờ có    │ ◄────────────────── │ ghi vào p                │
│ dữ liệu          │   n, err            └──────────────────────────┘
└──────────────────┘
```

- **Ngược với trực giác:** ở nhiều ngôn ngữ, `read()` **trả về** dữ liệu. Ở Go, ta **truyền vào** một `[]byte`, còn `Read` **ghi** dữ liệu vào đó.
- Làm được vì slice là kiểu tham chiếu: hàm sửa phần tử trong slice thì slice bên ngoài cũng thấy thay đổi (xem phần Pass By Value / Pointers).
- Giá trị trả về **không phải dữ liệu**, chỉ là thông tin về lần đọc:

| Giá trị trả về | Ý nghĩa |
|---|---|
| `n int` | Số byte đã đọc vào `p` |
| `err error` | Lỗi nếu có (hết dữ liệu thì là `io.EOF`) |

## Working with the Read Function

```go
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
```

### `make([]byte, 99999)`

- `make` là hàm có sẵn của Go: tham số 1 là **kiểu slice**, tham số 2 là **số phần tử** khởi tạo sẵn (đều là giá trị 0).
- Khác cú pháp cũ `[]byte{...}` (liệt kê phần tử).
- **Tại sao cần 99999 phần tử?** `Read` **không tự mở rộng** slice. Nó chỉ đọc cho tới khi slice đầy. Slice rỗng (`len = 0`) thì `Read` coi như đã đầy và dừng ngay.
- 99999 là con số "đoán chừng đủ lớn" để chứa toàn bộ HTML.

### Kết quả

- `string(bs)` chuyển `[]byte` thành `string` để in.
- `go run main.go` in ra toàn bộ HTML của trang Google (`<!doctype html><html>...`).

> **Lưu ý thực tế:** một lần gọi `Read` có thể **không đọc hết** dữ liệu. Code thật thường dùng `io.ReadAll(resp.Body)` hoặc `io.Copy` (phần dưới), và nhớ `defer resp.Body.Close()`.

## The Writer Interface

Rút gọn đoạn code trên thành **một dòng**:

```go
io.Copy(os.Stdout, resp.Body)
```

Kết quả y hệt: in toàn bộ HTML ra terminal. Để hiểu dòng này, cần biết interface **`Writer`**, là chiều **ngược lại** của `Reader`:

```
                Reader                              Writer
Nguồn bên ngoài ──────► []byte (trong chương trình) ──────► Đích bên ngoài
(HTTP body, file,                                          (HTTP request đi, file,
 ảnh, terminal...)                                          ảnh, terminal...)
```

| Interface | Hướng dữ liệu |
|---|---|
| `Reader` | Từ **ngoài vào** chương trình, thành `[]byte` |
| `Writer` | Từ `[]byte` trong chương trình **ra ngoài** |

Muốn xuất dữ liệu ra đâu đó, ta tìm trong thư viện chuẩn một giá trị cài đặt `Writer` rồi dùng nó.

## The io.Copy Function

### `io.Writer`

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- Chữ ký giống `Read`, nhưng lần này `p` là **dữ liệu đầu vào**: `Write` lấy dữ liệu trong `p` và gửi ra một kênh output.

### `io.Copy`

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

| Tham số | Kiểu | Trong ví dụ |
|---|---|---|
| `dst` | thứ cài đặt `Writer` | `os.Stdout` |
| `src` | thứ cài đặt `Reader` | `resp.Body` |

- `os.Stdout` là một **biến** do package `os` export, có kiểu `*os.File`.
- `*os.File` có hàm `Write([]byte) (int, error)`, nên **thỏa** interface `Writer` và truyền được làm tham số đầu của `io.Copy`.
- `io.Copy` = lấy dữ liệu từ `Reader`, chuyển sang `Writer`.

## The Implementation of io.Copy

Trong VS Code, giữ **Cmd** (macOS) / **Ctrl** (Windows) rồi click vào `Copy` để xem mã nguồn:

```go
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}

func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// ...
	if buf == nil {
		size := 32 * 1024
		// ...
		buf = make([]byte, size) // giống bs := make([]byte, 99999)
	}
	for {
		nr, er := src.Read(buf) // Reader ghi dữ liệu vào buf
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr]) // đẩy dữ liệu trong buf sang Writer
			// ... cộng dồn written, kiểm tra lỗi
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
```

- Tạo buffer `[]byte` cỡ **32 KB**.
- **Vòng lặp:** `src.Read(buf)` đọc một phần dữ liệu vào buffer, rồi `dst.Write(buf[:nr])` ghi phần đó ra đích. Lặp đến khi gặp `EOF`.
- Nhờ vòng lặp, `io.Copy` đọc được **toàn bộ** dữ liệu mà không cần đoán kích thước như `make([]byte, 99999)`.

```
src (Reader) ──Read──► buf []byte (32KB) ──Write──► dst (Writer)
     ▲                                                  │
     └──────────────── lặp đến EOF ─────────────────────┘
```

## A Custom Writer

Tự tạo kiểu cài đặt `Writer` rồi truyền vào `io.Copy`:

```go
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
```

- `logWriter` có `Write([]byte) (int, error)` nên **tự động** thỏa `io.Writer` (implicit).
- Theo tài liệu, `n` phải là **số byte đã xử lý** trong `p`. Ở đây xử lý hết nên trả về `len(bs)`, lỗi là `nil`.
- Kết quả: in HTML, sau đó là dòng `Just wrote this many bytes: 10403` (con số tùy lần chạy).

### Garbage in, garbage out (lần nữa)

```go
func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil // đúng chữ ký, sai logic
}
```

- Vẫn **compile và chạy được**, nhưng **không in ra gì cả**.
- Interface chỉ đảm bảo **đúng kiểu**, không đảm bảo **đúng hành vi**. Phải đọc tài liệu để cài đặt đúng.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `Read(p []byte)` | Ghi dữ liệu **vào** `p`, trả về số byte đọc được và lỗi |
| `make([]byte, n)` | Tạo slice có sẵn `n` phần tử. `Read` không tự mở rộng slice |
| `io.Writer` | `Write(p []byte) (int, error)`: đưa dữ liệu trong `p` **ra ngoài** |
| `io.Copy(dst, src)` | Lặp `Read` từ `src` rồi `Write` sang `dst` qua buffer 32 KB |
| `os.Stdout` | Biến kiểu `*os.File`, thỏa `Writer`, ghi ra terminal |
| Custom Writer | Kiểu bất kỳ có `Write([]byte) (int, error)` |
| Interface là contract | Chỉ kiểm tra chữ ký, logic sai vẫn compile được |

# Channels and Go Routines

Channel và go routine là hai công cụ của Go để viết chương trình **đồng thời** (concurrent).

## Bài toán: Status Checker

Viết chương trình kiểm tra một danh sách website có đang hoạt động không bằng cách gửi HTTP request tới từng trang.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
```

- URL phải có cả giao thức (`http://`) vì `http.Get` cần URL đầy đủ.
- `_` bỏ qua index trong vòng lặp `range`, và bỏ qua `resp` vì chỉ cần biết có `err` hay không.
- `fmt.Println(link, "is up!")` nhận nhiều đối số nên không cần nối chuỗi.

## Vấn đề: chạy tuần tự (serial)

- Kết quả in ra **đúng thứ tự** trong slice, giữa mỗi dòng có một khoảng trễ.
- `http.Get` là **lệnh chặn (blocking call)**: chương trình đứng chờ response rồi mới sang link tiếp theo.
- Nếu có hàng nghìn link, mỗi link chỉ được kiểm tra lại sau khi chạy hết cả danh sách.
- Giải pháp: gửi **tất cả request cùng lúc**, response nào về trước thì in trước.

```
Tuần tự:  google ──chờ──► facebook ──chờ──► stackoverflow ──chờ──► ...
Đồng thời: google ─┐
           facebook ─┤ (gửi cùng lúc, in kết quả khi có response)
           stackoverflow ─┘
```

## Go Routine

- Khi chạy chương trình, Go tự tạo **một go routine chính (main routine)**, thực thi code từng dòng.
- Thêm từ khóa `go` trước lời gọi hàm để chạy hàm đó trong **một go routine mới**:

```go
for _, link := range links {
	go checkLink(link)
}
```

- Khi go routine con gặp lệnh chặn (`http.Get`), quyền điều khiển trả về main routine để nó tiếp tục vòng lặp và tạo go routine tiếp theo.
- `go` luôn đặt **trước một lời gọi hàm**.

## Go Scheduler

- **Go Scheduler** quyết định go routine nào được chạy.
- Với **một lõi CPU**: dù có nhiều go routine, tại một thời điểm chỉ **một** go routine chạy. Scheduler chuyển qua lại rất nhanh khi routine kết thúc hoặc gặp lệnh chặn.
- Với **nhiều lõi CPU**: mỗi lõi chạy một go routine, nên nhiều đoạn code chạy **thực sự cùng lúc**.

## Concurrency is not Parallelism

| | Concurrency (đồng thời) | Parallelism (song song) |
|---|---|---|
| Ý nghĩa | Có nhiều go routine, lên lịch xen kẽ nhau | Nhiều việc chạy **đúng cùng một thời điểm** |
| Cần gì | Chỉ cần 1 lõi CPU | Cần nhiều lõi CPU vật lý |
| Ví dụ | 1 lõi chuyển qua lại giữa các routine | Lõi 1 chạy routine A, lõi 2 chạy routine B cùng lúc |

## Main routine và Child routine

- **Main routine**: tạo tự động khi chạy chương trình, thực thi `main`.
- **Child routine**: các routine tạo bằng từ khóa `go`.
- Hai loại này **không ngang hàng nhau**. Điểm khác biệt này gây ra lỗi ngay khi bắt đầu dùng go routine (sẽ tìm hiểu ở phần tiếp theo).

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Blocking call | Lệnh mất thời gian, routine phải đứng chờ (vd `http.Get`) |
| Go routine | Một "động cơ" thực thi code của một hàm |
| `go f()` | Chạy `f` trong một go routine mới |
| Go Scheduler | Chọn routine nào chạy, chuyển đổi khi gặp lệnh chặn |
| Concurrency | Nhiều routine xen kẽ, có thể chỉ trên 1 lõi |
| Parallelism | Nhiều routine chạy cùng lúc trên nhiều lõi |
| Main vs child routine | Main tạo tự động, child tạo bằng `go`, không ngang hàng |

## Lỗi: main routine thoát sớm

- Chỉ thêm `go checkLink(link)` rồi chạy: **không có gì được in ra**.
- Lý do: **chỉ main routine quyết định khi nào chương trình kết thúc**.
- Main routine tạo xong các child routine, hết vòng `for`, không còn code để chạy, nên thoát luôn. Nó **không chờ** các child routine đang `http.Get`.
- Khi main routine thoát, toàn bộ chương trình dừng và các child routine bị hủy giữa chừng.

```
main:   tạo G1, G2, G3, G4, G5 ──► hết code ──► THOÁT
G1..G5: đang chờ http.Get ...          ✗ bị hủy
```

## Channel

- **Channel** là cách **duy nhất** để các go routine giao tiếp với nhau.
- Hình dung như tin nhắn: một routine gửi dữ liệu vào channel, routine khác có channel đó sẽ nhận được.
- Channel là một giá trị như mọi giá trị khác (struct, slice, int...), có thể truyền đi khắp chương trình.
- Channel **có kiểu**: `chan string` chỉ gửi/nhận được `string`. Gửi kiểu khác (vd `float64`) sẽ báo lỗi kiểu.

### Tạo và truyền channel

```go
c := make(chan string)

for _, link := range links {
	go checkLink(link, c)
}
```

- Biến `c` tuân theo quy tắc phạm vi bình thường: chỉ dùng được trong `main`, nên phải **truyền vào hàm** làm đối số.
- Tham số của hàm cũng phải khai báo kiểu dữ liệu của channel:

```go
func checkLink(link string, c chan string) { ... }
```

### Cú pháp gửi và nhận

| Cú pháp | Ý nghĩa |
|---|---|
| `channel <- 5` | Gửi giá trị `5` vào channel |
| `myNumber <- channel` | Chờ nhận một giá trị từ channel rồi gán vào `myNumber` |
| `fmt.Println(<-channel)` | Chờ nhận một giá trị rồi truyền thẳng vào hàm |

> Lưu ý: cú pháp gán đầy đủ trong Go là `myNumber := <-channel`.

### Áp dụng vào checkLink

```go
func main() {
	// ...
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
```

- Kết quả: chỉ in ra **một** link (link phản hồi nhanh nhất) cùng một tin nhắn từ channel, rồi chương trình thoát.

## Nhận từ channel là lệnh chặn

- `<-c` là một **lệnh chặn (blocking call)**, giống `http.Get`. Main routine bị tạm dừng cho đến khi có tin nhắn gửi vào channel.
- Diễn biến:
  1. Main routine tạo hết các child routine rồi dừng ở `<-c` để chờ.
  2. Routine nào có response trước (vd google) sẽ gửi tin nhắn vào `c`.
  3. Main routine thức dậy, nhận tin nhắn, in ra.
  4. Không còn code để chạy, nên main routine thoát. Các routine còn lại bị bỏ dở.

```
main:   tạo G1..G5 ──► chờ <-c ......... nhận ──► in ──► THOÁT
G1 (google):  http.Get ......... c <- "..." ┘
G2..G5:       http.Get ........................ ✗ bị hủy
```

### Thử nghiệm: nhiều lệnh nhận

- Viết `fmt.Println(<-c)` **2 lần**: nhận 2 tin nhắn rồi thoát.
- Viết **5 lần** (bằng số link): nhận đủ 5 tin nhắn, mọi routine đều xong.
- Viết **6 lần** (nhiều hơn số link): in 5 tin nhắn rồi chương trình **bị treo**, vì main routine chờ mãi tin nhắn thứ 6 không bao giờ tới.
- Kết luận: số lần nhận phải khớp với số tin nhắn được gửi. Viết lặp lại nhiều `Println` cũng không thực tế khi có nhiều link (cách giải quyết ở phần tiếp theo).

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| Main routine thoát sớm | Chương trình kết thúc khi main hết code, không chờ child routine |
| Channel | Kênh giao tiếp duy nhất giữa các go routine |
| `make(chan string)` | Tạo channel chỉ truyền giá trị kiểu `string` |
| `c chan string` | Khai báo tham số kiểu channel trong hàm |
| `c <- v` | Gửi `v` vào channel |
| `v := <-c` / `f(<-c)` | Chờ nhận giá trị từ channel |
| Nhận là lệnh chặn | Routine dừng ở `<-c` cho đến khi có tin nhắn |
| Nhận nhiều hơn số gửi | Chương trình bị treo vì chờ mãi |

## Nhận tin nhắn bằng vòng lặp for

- Thay vì viết lặp lại nhiều dòng `fmt.Println(<-c)`, dùng vòng lặp `for` kiểu C (không dùng `range`), lặp đúng bằng số link:

```go
for i := 0; i < len(links); i++ {
	fmt.Println(<-c)
}
```

- Vòng lặp **không chạy hết ngay lập tức**: mỗi vòng dừng ở `<-c` cho đến khi có tin nhắn rồi mới sang vòng tiếp theo.
- Nhận đủ `len(links)` tin nhắn thì main routine hết code và chương trình thoát.

## Kiểm tra link liên tục (lặp vô hạn)

- Yêu cầu mới: kiểm tra mỗi link **liên tục**, giống một trình theo dõi trạng thái. Routine của link nào xong thì tạo ngay routine mới để kiểm tra lại link đó.
- Cách làm: gửi **chính link** vào channel thay vì một chuỗi cố định, main routine nhận link đó và tạo routine mới.

```go
func main() {
	// ...
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
```

- `for { }` là **vòng lặp vô hạn**. Nó không quay hàng triệu lần mỗi giây vì `<-c` vẫn là lệnh chặn: chỉ lặp khi nhận được một link.
- `<-c` dùng được làm đối số `string`, vì Go biết channel `chan string` sẽ trả về một `string`.
- Tại mọi thời điểm chỉ có khoảng `len(links)` request đang chạy.
- Vấn đề: các trang bị gọi liên tục, quá nhanh. Nên có một khoảng dừng giữa các lần kiểm tra.

## Cú pháp range với channel

- Viết lại vòng lặp vô hạn cho dễ đọc hơn. Hai cách viết dưới đây **tương đương nhau**:

```go
for {
	go checkLink(<-c, c)
}

for l := range c {
	go checkLink(l, c)
}
```

- `for l := range c`: chờ channel trả về một giá trị, gán vào `l`, chạy thân vòng lặp, rồi lại chờ giá trị tiếp theo.
- Người đọc code thấy ngay mục đích: "chạy thân vòng lặp mỗi khi channel có giá trị mới".

## Tạm dừng với time.Sleep

- Gói `time` có hàm `time.Sleep(d Duration)`: tạm dừng **go routine hiện tại** ít nhất một khoảng `d`.
- `time.Second` là hằng số kiểu `Duration` (1 giây). `Duration` thực chất là `int64` nên có thể nhân lên:

```go
time.Sleep(5 * time.Second) // dừng 5 giây
```

### Đặt Sleep ở đâu?

| Vị trí | Vấn đề |
|---|---|
| Trong main routine (thân `for l := range c`) | Main bị chặn 5 giây mỗi vòng, không nhận được tin nhắn khác. Tin nhắn bị dồn lại trong channel, tối đa chỉ tạo **1 routine mỗi 5 giây** |
| Đầu hàm `checkLink` | Chạy được, nhưng lần kiểm tra đầu tiên cũng bị chờ 5 giây. `checkLink` nên kiểm tra **ngay**, không nên tự dừng |
| Trong một function literal chạy bằng `go` | Đúng: main luôn sẵn sàng nhận tin nhắn, `checkLink` vẫn làm đúng một việc |

## Function literal

- **Function literal** là hàm không có tên, giống *anonymous function* trong JavaScript hay *lambda* trong Python, C#, Ruby.
- Dùng để bọc một đoạn code nhỏ rồi chạy nó trong một go routine:

```go
for l := range c {
	go func() {
		time.Sleep(5 * time.Second)
		checkLink(l, c)
	}()
}
```

- Cặp `()` ở cuối dùng để **gọi** function literal. Từ khóa `go` luôn cần một **lời gọi hàm** phía sau.
- Editor hiện cảnh báo: *loop variable l captured by func literal*.

## Lỗi: biến được capture bởi function literal

- Khi chạy: 5 lần kiểm tra đầu đúng, nhưng sau 5 giây, các lần sau **đều là cùng một link** (vd toàn `facebook.com`).
- Nguyên nhân: function literal tham chiếu tới biến `l` ở **phạm vi bên ngoài**. Main routine và các child routine cùng trỏ tới **một địa chỉ bộ nhớ**.
- Main routine liên tục gán giá trị mới cho `l` mỗi khi nhận từ channel. Đến lúc child routine thức dậy sau 5 giây, `l` đã bị đổi sang link khác.

```
Sai (dùng chung biến):          Đúng (truyền đối số):
main  ──► l ──► [0xA0] "fb"      main  ──► l ──► [0xA0] "fb"
child ──► l ──┘                  child ──► link ──► [0xB0] "google" (bản sao)
```

### Cách sửa: truyền vào làm đối số

- Go là ngôn ngữ **pass by value**: truyền đối số vào hàm sẽ **tạo một bản sao** trong bộ nhớ, nên child routine có giá trị riêng, ổn định.

```go
for l := range c {
	go func(link string) {
		time.Sleep(5 * time.Second)
		checkLink(link, c)
	}(l)
}
```

- `(l)` ở cuối: gọi function literal và truyền `l` vào. Tham số đặt tên là `link` để khỏi nhầm với `l` bên ngoài.
- Sau khi sửa: mỗi lần kiểm tra là một link khác nhau, và theo thời gian các lần kiểm tra lệch nhau dần.

> Lưu ý: từ **Go 1.22**, mỗi vòng lặp `for` tạo một biến mới, nên lỗi này không còn xảy ra với biến vòng lặp nữa. Tuy vậy, truyền dữ liệu qua đối số vẫn là cách rõ ràng và an toàn hơn.

### Nguyên tắc

- **Không bao giờ** truy cập cùng một biến từ nhiều go routine khác nhau.
- Chỉ chia sẻ dữ liệu với child routine bằng cách **truyền đối số** hoặc **giao tiếp qua channel**.

### Tóm tắt

| Khái niệm | Ý nghĩa |
|---|---|
| `for i := 0; i < n; i++` | Vòng lặp kiểu C, dùng để nhận đúng `n` tin nhắn |
| `for { }` | Vòng lặp vô hạn |
| `go checkLink(<-c, c)` | Nhận link từ channel rồi kiểm tra lại link đó |
| `for l := range c` | Lặp mỗi khi channel có giá trị mới |
| `time.Sleep(5 * time.Second)` | Tạm dừng go routine hiện tại 5 giây |
| Sleep trong main routine | Sai: main bị chặn, không nhận được tin nhắn |
| Function literal | Hàm không tên: `func(...) { ... }(...)` |
| Capture biến bên ngoài | Nhiều routine dùng chung biến, giá trị bị thay đổi |
| Truyền đối số | Tạo bản sao, mỗi routine có giá trị riêng |
