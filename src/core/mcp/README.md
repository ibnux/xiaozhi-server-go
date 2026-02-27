# Cara Penggunaan MCP

## ESP32 MCP
Untuk MCP ESP32, penambahan alat MCP kustom perlu dilakukan di frontend. Secara default mendukung 4 alat (self.get_device_status, self.audio_speaker.set_volume, self.screen.set_brightness, self.screen.set_theme).

Layanan tidak perlu mengubah konfigurasi, langsung didukung.

Setelah klien ESP32 terhubung ke layanan, log layanan akan mencetak informasi registrasi 4 alat di atas. Langsung bercakap-cakap, minta Xiaozhi mengatur volume, untuk menguji hasilnya.

## MCP Eksternal Server
Server mendukung pemanggilan MCP eksternal dengan mengonfigurasi file .mcp_server_settings.json di direktori root kode sumber / direktori tempat program biner (windows-amd64-server.exe atau linux-amd64-server-upx) berada. Formatnya adalah:
```
{
  "mcpServers": {
    "amap-maps": {
      "command": "npx",
      "args": [
          "-y",
          "@amap/amap-maps-mcp-server"
      ],
      "env": {
          "AMAP_MAPS_API_KEY": "api key Gaode Anda"
      }
    },
    "filesystem": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-filesystem",
        "path izin yang dikonfigurasi"
      ]
    },
     "playwright": {
      "command": "npx",
      "args": ["-y", "@executeautomation/playwright-mcp-server"],
      "des" : "jalankan 'npx playwright install' terlebih dahulu"
    },
    "windows-cli": {
      "command": "npx",
      "args": ["-y", "@simonb97/server-win-cli"]
    }
  }
}
```
Server perlu menginstal node untuk mendukung format MCP npx. Untuk format MCP lainnya silakan coba sendiri.

Saat ini hanya mendukung format MCP Stdio. Jika perlu menggunakan mode SSE, dapat mempertimbangkan menggunakan cara mcp-proxy, konfigurasinya sebagai berikut:

```
{
  "mcpServers": {
    "zapier": {
      "command": "mcp-proxy",
      "args": [
        "https://actions.zapier.com/mcp/****/sse"
      ]
    }
  }
}
```

Saat layanan dimulai, konfigurasi MCP akan dimuat secara otomatis dan pool sumber daya MCP akan dibuat terlebih dahulu. Amati log untuk memastikan apakah MCP berhasil dimuat.
