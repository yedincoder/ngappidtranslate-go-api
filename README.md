# NgAppID Translate API 🚀

Backend API super ringan pakai **Go murni (Golang)** untuk memproses *request* terjemahan bahasa. Dibangun tanpa framework tambahan agar dapat berjalan langsung sebagai **Serverless Function** di Vercel.

API ini menggunakan layanan gratis dari [MyMemory Translation API](https://mymemory.translated.net/) dan sudah dilengkapi dengan konfigurasi CORS agar aman dipanggil dari aplikasi frontend (React / Mobile).

## 🚀 Cara Clone Repository

Pastikan Git sudah terinstal, lalu jalankan perintah ini di terminal:

    git clone https://github.com/yedincoder/ngappidtranslate-go-api.git
    cd ngappidtranslate-go-api

## 📁 Struktur Folder (Format Vercel)

    ├── api/
    │   └── translate.go    # Handler utama API (Serverless)
    ├── go.mod              
    ├── go.sum
    └── README.md

## 📡 API Reference

**Endpoint:** `POST /api/translate`  
**Headers:** `Content-Type: application/json`

**Body Request:**

    {
      "text": "Selamat Malam",
      "source": "id",
      "target": "en"
    }

**Response (200 OK):**

    {
      "original": "Selamat Malam",
      "translated": "Good Night"
    }

## 🌐 Cara Deploy ke Vercel

Cukup push repository ini ke GitHub, lalu import ke dashboard Vercel. Vercel akan otomatis mendeteksi folder api/ dan meng-compile file Go ini menjadi endpoint Serverless secara instan. Tidak perlu setting vercel.json!

---
**👨‍💻 Pengembang:** Yedin (YedinCoder)  
*Dibuat khusus untuk ekosistem NgAppID*