# ✨ Layanan Backend Chatbot AI Xiaozhi (Versi Komersial)

Xiaozhi AI adalah robot interaksi suara yang menggabungkan model bahasa besar seperti Qwen dan DeepSeek, terhubung ke berbagai perangkat (ESP32, Android, Python, dll.) melalui protokol MCP untuk menghadirkan percakapan manusia-mesin yang efisien dan alami.

Proyek ini adalah layanan backend yang bertujuan menyediakan **solusi deployment tingkat komersial** — konkurensi tinggi, biaya rendah, fitur lengkap, dan siap pakai.

<p align="center">
  <img src="https://github.com/user-attachments/assets/aa1e2f26-92d3-4d16-a74a-68232f34cca3" alt="Xiaozhi Architecture" width="600">
</p>

Proyek ini awalnya berbasis [proyek open-source ESP32 dari Xia Ge](https://github.com/78/xiaozhi-esp32?tab=readme-ov-file) dan kini telah membentuk ekosistem lengkap dengan dukungan berbagai protokol klien.

---

## ✨ Keunggulan Utama

| Keunggulan              | Keterangan                                                                         |
| ----------------------- | ---------------------------------------------------------------------------------- |
| �� Konkurensi Tinggi    | Mendukung 3000+ pengguna online di satu mesin, dapat diskalakan hingga jutaan pengguna secara terdistribusi |
| 👥 Sistem Pengguna      | Kemampuan registrasi, login, dan manajemen izin pengguna yang lengkap              |
| 💰 Integrasi Pembayaran | Terhubung dengan sistem pembayaran untuk mendukung bisnis end-to-end               |
| 🛠️ Integrasi Model Fleksibel | Mendukung pemanggilan berbagai model besar melalui API, penyederhanaan deployment, dan dukungan deployment lokal kustom |
| 📈 Dukungan Komersial   | Menyediakan dukungan teknis 7×24 dan jaminan operasional                           |
| 🧠 Kompatibilitas Model | Mendukung ASR (Doubao), TTS (EdgeTTS), LLM (OpenAI, Ollama), analisis gambar (Zhipu), dll. |

---

## ✅ Daftar Fitur Versi Komunitas

* [x] Mendukung koneksi websocket
* [x] Mendukung percakapan suara format PCM / Opus
* [x] Mendukung model besar: ASR (Doubao streaming), TTS (EdgeTTS/Doubao), LLM (OpenAI API, Ollama)
* [x] Mendukung kontrol suara untuk mengaktifkan kamera dan mengenali gambar (Zhipu API)
* [x] Mendukung tiga mode percakapan: auto/manual/realtime, termasuk interupsi percakapan real-time
* [x] Mendukung koneksi klien ESP32 Xiaozhi, klien Python, dan klien Android tanpa verifikasi
* [x] Distribusi firmware OTA
* [x] Mendukung protokol MCP (klien / lokal / server), dapat terhubung ke peta Gaode, query cuaca, dll.
* [x] Mendukung kontrol suara untuk mengganti suara karakter
* [x] Mendukung kontrol suara untuk mengganti peran preset
* [x] Mendukung kontrol suara untuk memutar musik
* [x] Mendukung deployment layanan di satu mesin
* [x] Mendukung database lokal SQLite
* [x] Mendukung alur kerja Coze
* [x] Mendukung deployment Docker

## ✅ Daftar Fitur Versi Komersial
* [x] Semua fitur versi komunitas
* [x] Dukungan teknis dari tim pengembang
* [x] Pembaruan fitur inti gratis di masa mendatang
* [x] Panel admin versi komersial dengan lebih banyak opsi fitur
* [x] Mendukung manajemen multi-pengguna
* [x] Kustomisasi tampilan halaman selamat datang
* [x] Kustomisasi logo hak cipta dengan merek bisnis perusahaan Anda
* [x] Kustomisasi template peran Agent
* [x] Mendukung lebih banyak model
* [x] Mendukung dua protokol komunikasi: websocket dan MQTT+UDP
* [x] Mendukung pembuatan dan pengiriman TTS streaming
* [x] Mendukung kloning suara
* [x] Mendukung basis pengetahuan
* [x] Mendukung timbre kustom (cosyvoice2, indextts)
* [x] Mendukung pembaruan firmware melalui OTA
* [x] Mendukung alur kerja Coze
* [x] Mendukung alur kerja Dify
* [x] Optimasi mendalam kecepatan respons
* [x] Mendukung autentikasi pengguna, aktivasi dan pengikatan perangkat
* [x] Mendukung manajemen perangkat: unbind/disable
* [x] Mendukung unbind perangkat dari panel admin
* [x] Mendukung Agent kustom pengguna
* [x] Dukungan multi-bahasa internasional: Mandarin, Inggris, Jepang, Spanyol, Indonesia, dll.
* [x] Mendukung titik akses MCP
* [x] Mendukung database jaringan
* [x] Mendukung deployment terdistribusi
* [x] Mendukung deployment model besar lokal

Alamat uji coba / demo versi komersial:

https://xiaozhi.xf.bj.cn/login

---

## 🚀 Memulai Cepat

### 1. Unduh Versi Release

> Disarankan langsung mengunduh versi Release, tanpa perlu mengkonfigurasi lingkungan pengembangan:

👉 [Klik untuk ke halaman Releases](https://github.com/AnimeAIChat/xiaozhi-server-go/releases)

* Pilih versi yang sesuai dengan platform Anda (mis. Windows: `windows-amd64-server.exe`)
* `.upx.exe` adalah versi terkompresi, fungsionalitas sama, ukuran lebih kecil, cocok untuk deployment jarak jauh

---


### 2. Konfigurasi `.config.yaml`

* Disarankan menyalin `config.yaml` dan mengganti namanya menjadi `.config.yaml`
* Konfigurasikan field model, WebSocket, alamat OTA, dll. sesuai kebutuhan
* Tidak disarankan menghapus atau menambah struktur field secara manual

#### Konfigurasi Alamat WebSocket (Wajib)

```yaml
web:
  websocket: ws://your-server-ip:8000
```

Digunakan sebagai alamat koneksi yang dikirimkan layanan OTA ke klien. Klien ESP32 akan otomatis terhubung ke WS dari alamat ini, tanpa perlu konfigurasi manual.

Catatan: Jika melakukan debug di jaringan lokal, `your-server-ip` harus dikonfigurasi sebagai **IP komputer di jaringan lokal**, dan perangkat terminal serta komputer harus berada di segmen jaringan yang sama agar perangkat dapat terhubung ke layanan di komputer melalui alamat IP ini.

#### Konfigurasi Alamat OTA (Wajib)

```text
http://your-server-ip:8080/api/ota/
```

> Firmware ESP32 memiliki alamat OTA bawaan, pastikan alamat layanan ini tersedia. **Setelah layanan berjalan, masukkan alamat ini di browser untuk memverifikasi bahwa layanan dapat diakses**.

Perangkat ESP32 dapat mengubah alamat OTA di antarmuka jaringan, sehingga dapat berpindah layanan backend tanpa perlu mem-flash firmware ulang.

#### Konfigurasi ASR, LLM, TTS

Konfigurasikan layanan model terkait sesuai format file konfigurasi, usahakan tidak menambah atau mengurangi field.

---

## 💬 Konfigurasi Protokol MCP

Referensi: `src/core/mcp/README.md`

---

## 🧪 Instalasi dan Menjalankan dari Kode Sumber

### Prasyarat

* Go 1.24.2+
* Pengguna Windows perlu menginstal CGO dan library Opus (lihat di bawah)

```bash
git clone https://github.com/AnimeAIChat/xiaozhi-server-go.git
cd xiaozhi-server-go
cp config.yaml .config.yaml
```

---

### Instalasi Lingkungan Kompilasi Opus di Windows

Instal [MSYS2](https://www.msys2.org/), buka konsol MSYS2 MINGW64, kemudian masukkan perintah berikut:

```bash
pacman -Syu
pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-go mingw-w64-x86_64-opus
pacman -S mingw-w64-x86_64-pkg-config
```

Atur variabel lingkungan (untuk PowerShell atau variabel sistem):

```bash
set PKG_CONFIG_PATH=C:\msys64\mingw64\lib\pkgconfig
set CGO_ENABLED=1
```

Sebaiknya jalankan perintah "go run ./src/main.go" sekali di lingkungan MINGW64 untuk memastikan layanan berjalan dengan normal.

Jika pembaruan Go mod lambat, pertimbangkan untuk mengatur proxy Go dan beralih ke sumber mirror lokal.

---

### Menjalankan Proyek

```bash
go mod tidy
go run ./src/main.go
```

### Kompilasi Versi Rilis

```bash
go build -o xiaozhi-server.exe src/main.go
```

### Pengujian
* Disarankan menggunakan perangkat hardware ESP32 untuk pengujian, sehingga masalah kompatibilitas dapat diminimalkan
* Disarankan menggunakan klien Android Xuanfeng Xiaozhi, cukup tambahkan alamat OTA layanan lokal di antarmuka pengaturan. Versi Android diterbitkan di halaman Release, dapat memilih versi terbaru
  <img width="221" height="470" alt="image" src="https://github.com/user-attachments/assets/145a6612-8397-439b-9429-325855a99101" />

  [xiaozhi-0.0.6.apk](https://github.com/AnimeAIChat/xiaozhi-server-go/releases/download/v0.1.0/xiaozhi-0.0.6.apk)
* Dapat menggunakan klien lain yang kompatibel dengan protokol Xiaozhi untuk pengujian
---

## 📚 Dokumentasi Swagger

* Buka browser dan akses: `http://localhost:8080/swagger/index.html`

### Memperbarui Dokumentasi Swagger (Jalankan setiap kali memodifikasi API)

```bash
cd src
swag init -g main.go
```

---

## ☁️ Panduan Deployment Kode Sumber di CentOS

> Lihat dokumentasi: [Panduan Instalasi CentOS 8](docs/Centos_Guide.md)

---

## Deployment di Lingkungan Docker

1. Siapkan file `docker-compose.yml`, `.config.yaml`, dan file program biner

👉 [Klik untuk ke halaman Releases](https://github.com/AnimeAIChat/xiaozhi-server-go/releases) untuk mengunduh file program biner

* Pilih versi yang sesuai dengan platform Anda (secara default menggunakan Linux: `linux-amd64-server-upx`; jika menggunakan versi lain, perlu memodifikasi docker-compose.yml)

2. Tempatkan ketiga file dalam satu direktori, konfigurasikan `docker-compose.yml` dan `.config.yaml`

3. Jalankan `docker compose up -d`

---

## 💬 Dukungan Komunitas


Selamat datang untuk mengirimkan Issue, PR, atau saran fitur baru!

<img src="https://github.com/Eric0308/assert/blob/main/xiaozhi/qr.jpg" width="450" alt="QR Code Grup WeChat"> 
<img src="https://github.com/user-attachments/assets/074c6aec-cfb5-4a68-8fc2-2d08679e366b" width="450" alt="QR Code Grup QQ">
---

## 🛠️ Pengembangan Kustom

Kami menerima berbagai proyek pengembangan kustom. Jika Anda memiliki kebutuhan khusus, silakan hubungi kami melalui WeChat untuk berdiskusi.

<img src="https://github.com/user-attachments/assets/e2639bc3-a58a-472f-9e72-b9363f9e79a3" width="450" alt="QR Code Pemilik Grup">

## 📄 Lisensi

Repositori ini mengikuti `Xiaozhi-server-go Open Source License` (versi yang ditingkatkan berbasis Apache 2.0)
