# Modul OTA

## Struktur Direktori
- `ota.go`: Implementasi default layanan OTA
- `interfaces.go`: Definisi antarmuka layanan OTA
- `server.go`: Implementasi layanan HTTP OTA
- `README.md`: Dokumen penjelasan modul

## Cara Penggunaan
1. Import modul layanan OTA di program utama.
2. Inisialisasi layanan OTA.
3. Konfigurasikan parameter `UpdateURL`, tentukan alamat WebSocket.

## Penjelasan Antarmuka OTA
- `GET /api/ota/`: Mengembalikan status operasi antarmuka OTA dan alamat WebSocket.
- `POST /api/ota/`: Menerima permintaan perangkat, mengembalikan waktu server, informasi firmware, dan alamat WebSocket.

## Pengujian Antarmuka OTA (Apifox)

Anda dapat menggunakan [Apifox](https://apifox.com/) untuk menguji antarmuka OTA.

### 1. Pengujian Antarmuka GET
- Metode: GET
- URL: `http://localhost:8080/api/ota/`
- Respons yang diharapkan:
```json
{
  "status": "ok",
  "ws": "ws://localhost:8080/ws"
}
```

### 2. Pengujian Antarmuka POST
- Metode: POST
- URL: `http://localhost:8080/api/ota/`
- Tipe Body: JSON
- Contoh request body:
```json
{
  "device_id": "your_device_id"
}
```
- Respons yang diharapkan:
```json
{
  "server_time": "2024-01-01T12:00:00Z",
  "firmware": "v1.0.0",
  "ws": "ws://localhost:8080/ws"
}
```

### 3. Unduhan
- URL: `http://localhost:8080/ota_bin/{*.bin}`

### 4. Keterangan CORS
Layanan OTA sudah mendukung CORS, sehingga memudahkan frontend atau alat pihak ketiga untuk langsung memanggil.

Jika perlu menyesuaikan antarmuka atau konten respons lebih lanjut, silakan modifikasi `server.go` sesuai kebutuhan aktual.
