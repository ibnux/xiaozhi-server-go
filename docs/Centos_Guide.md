# Proses Instalasi Xiaozhi Server Versi Go di CentOS 8


# 1. Unduh dan Instal Go Versi 1.24
```Bash
cd /tmp
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
```

Kemudian atur variabel lingkungan:
```Bash
echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
```
Pengaturan di atas hanya berlaku sementara. Untuk berlaku permanen, perlu mengubah file /etc/profile

```Shell
vim /etc/profile
```

Tambahkan di bagian akhir:

```Bash
export GO111MODULE=on
export GOROOT=/usr/local/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Aktifkan variabel lingkungan:

```Shell
source /etc/profile
```

Verifikasi versi:
```Shell
go version
```
Beralih sumber Go ke proxy lokal:
```Shell
go env -w GOPROXY=https://goproxy.cn,direct
```


# 2. Unduh dan Instal Opus Versi Terbaru

### ✅ 1. Instal paket opus dan opus-devel
```
sudo dnf install opus opus-devel  # CentOS 8+
```
Jika menggunakan CentOS 7, gunakan yum:
```
sudo yum install opus opus-devel
```
Paket opus-devel berisi file .pc (untuk digunakan pkg-config), header file, link library, dll.

### ✅ 2. Pastikan pkg-config dapat menemukan opus.pc

Biasanya setelah instalasi akan otomatis ditempatkan di /usr/lib64/pkgconfig/opus.pc.

Anda dapat memverifikasi:
```
pkg-config --cflags opus
```
Jika tidak ada error berarti sudah ditemukan. Jika masih tidak ditemukan, coba atur variabel lingkungan:
```
export PKG_CONFIG_PATH=/usr/lib64/pkgconfig
```
Namun opus yang diinstal dengan cara ini bukan versi terbaru dan tidak dapat digunakan.

## ✅ Analisis Penyebab

- OPUS_GET_IN_DTX_REQUEST adalah makro yang relatif baru dalam libopus, versi lama libopus tidak mengandungnya.

- Kode binding qrtc/opus-go merujuk pada makro ini, sehingga jika sistem menginstal versi libopus yang lebih lama, akan terjadi error.

---

## ✅ Solusi

### Metode 1: Upgrade libopus (Direkomendasikan)
```
pkg-config --modversion opus
```
Jika lebih rendah dari 1.3.1 (seperti 1.1 atau 1.2), berarti versinya terlalu lama.

# Instal dependensi
```
sudo dnf install gcc make autoconf automake libtool
```
# Unduh dan kompilasi
```
cd /tmp
git clone https://github.com/xiph/opus.git
cd opus
./autogen.sh
./configure
make -j$(nproc)
sudo make install
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
```
Dapat ditambahkan ke .bashrc:
```
echo 'export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig' >> ~/.bashrc
```
```
source ~/.bashrc
```
Kemudian coba lagi:
```
go clean -modcache
go run ./src/main.go
```

# 3. Unduh Kode Sumber
```
git clone https://github.com/AnimeAIChat/xiaozhi-server-go
```
Setelah selesai diunduh, salin config.yaml ke .config.yaml, dan isi key dan informasi yang diperlukan dalam konfigurasi, untuk menghindari kebocoran kunci.

# 4. Menjalankan

Jalankan:

```
 go mod tidy
 go run ./src/main.go
```


