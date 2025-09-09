# go-helper-pkg

Kumpulan helper utilities untuk mempermudah pengembangan aplikasi dengan bahasa **Go**.  
Paket ini berisi fungsi-fungsi kecil yang reusable sehingga dapat mempercepat proses coding tanpa harus menulis ulang boilerplate code.

## ✨ Fitur

- Fungsi untuk manipulasi string, date, dan number.  
- Validasi data umum.  
- Konversi tipe data dengan mudah.  
- Utilitas lain yang sering dibutuhkan pada aplikasi Go.  

## 📦 Instalasi

Jalankan perintah berikut untuk menambahkan dependency:

```bash
go get github.com/Ekofebri9/go-helper-pkg
```

Lalu import ke dalam project Go Anda:

```go
import "github.com/Ekofebri9/go-helper-pkg"
```

## 🚀 Contoh Penggunaan

```go
package main

import (
	"fmt"
	"github.com/Ekofebri9/go-helper-pkg"
)

func main() {
	result := helperpkg.SomeHelperFunction("contoh input")
	fmt.Println(result)
}
```

> **Catatan:** Ganti `SomeHelperFunction` dengan fungsi yang sesuai dari package ini.

## 📂 Struktur Project

```
go-helper-pkg/
├── stringutil/   # Helper untuk string
├── datetime/     # Helper untuk tanggal & waktu
├── conv/         # Konversi tipe data
└── ...
```

## 🤝 Kontribusi

Kontribusi sangat terbuka! Silakan fork repository ini, buat branch baru, lalu ajukan pull request.  

## 📄 Lisensi

Proyek ini dirilis di bawah lisensi **MIT**.  
Silakan gunakan, modifikasi, dan distribusikan sesuai kebutuhan.  
