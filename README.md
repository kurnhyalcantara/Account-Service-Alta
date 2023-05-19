# Account-Service-Alta

Ini adalah aplikasi Account-Service yang dikembangkan menggunakan bahasa pemrograman Golang dan MySQL sebagai database.

## Struktur Folder

Aplikasi ini memiliki struktur folder sebagai berikut:

- `controllers`: Folder ini berisi file-file yang bertanggung jawab untuk mengontrol aliran data dan logika aplikasi.
- `entities`: Folder ini berisi definisi entitas atau model yang digunakan dalam aplikasi.
- `main.go`: File ini merupakan file utama yang menjalankan aplikasi.

## Fitur

Aplikasi Account-Service memiliki beberapa fitur utama berikut:

1. Manajemen Akun:
   - Register: Pengguna dapat mendaftarkan akun baru dengan memberikan informasi yang diperlukan.
   - Login: Pengguna dapat melakukan otentikasi dengan menggunakan nomor handphone dan password yang terdaftar.
   - Edit Akun: Pengguna dapat mengubah informasi profil mereka
   - Hapus Akun: Pengguna dapat menghapus akun mereka jika diperlukan.

2. Top Up Saldo dan Transfer Dana:
   - Top Up Saldo: Pengguna dapat menambahkan saldo ke akun mereka.
   - Transfer Dana: Pengguna dapat mentransfer dana dari akun mereka ke akun lain yang terdaftar dalam sistem.

3. Melihat Riwayat Top Up dan Transfer:
   - Riwayat Top Up: Pengguna dapat melihat riwayat top up saldo yang telah dilakukan sebelumnya.
   - Riwayat Transfer: Pengguna dapat melihat riwayat transfer dana yang telah dilakukan sebelumnya baik yang masuk dan keluar dari dananya.
