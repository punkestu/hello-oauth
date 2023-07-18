# OAuth

## Review
Pada dasarnya oauth adalah sebuah standar yang dapat digunakan untuk autentikasi user menggunakan akun dari platform maupun sosial media yang menerapkan standar ini.
## Platform
Ada banyak sekali platform yang menyediakan oauth sebagai metode autentikasi antara lain:
- Google
- Github
- Facebook
- dll

## Alur Kerja
Seperti disebutkan diatas bahwa oauth merupakan sebuah standar sehingga walaupun kita menggunakan berbagai platform sebagai oauth, alur kerja yang digunakan tetap serupa yaitu:
1. Login dengan platform yang kita inginkan
2. Platform memberikan ACCESS TOKEN
3. Kita menggunakan ACCESS TOKEN untuk mengakses data pada platform

## Alur Pembuatan
1. Registrasi aplikasi pada platform yang ingin kita gunakan
2. Dapatkan ID dan SECRET KEY
3. Tambahkan REDIRECT URI ketika autentikasi berhasil dilakukan (untuk Facebook pada mode development kita tidak perlu menambahkan REDIRECT URI ini)
4. Pastikan SCOPES sudah ditambahkan dan sesuai dengan kebutuhan kita
5. --- Mulai Code ---
6. Siapkan dan simpan ID dan SECRET KEY yang sudah didapatkan sebelumnya
7. Buatlah sebuah redirect link ke arah ENDPOINT yang sudah disediakan oleh platform yang ingin kita gunakan (baca dokumentasinya)
8. Tambahkan parameter client_id, scopes, redirect_uri serta response_type (ini kita isi "code" saja)
9. Untuk keamanan kita juga bisa menambahkan parameter state pada URL (ini opsional)
10. Siapkan handler dengan ENDPOINT sesuai dengan REDIRECT URI yang kita kirimkan
11. Pada handler ini kita akan mendapatkan data berupa CODE dan STATE (jika kita menambahkan parameter state pada url autentikasi) yang dapat kita cek terlebih dahulu untuk keamanan
12. Selanjutnya kita bisa mengambil token dari CODE yang kita dapatkan
13. Untuk mengambil token kita dapat mengirimkan request ke platform yang kita gunakan untuk autentikasi sebelumnya dengan data sesuai yang diminta oleh platform tersebut (biasanya CODE, CLIENT_ID, dan CLIENT_SECRET)
14. Dari request ini kita akan mendapatkan data berupa ACCESS TOKEN dan EXPIRED TIME
15. Untuk mengambil data user (atau data lainnya yg kita butuhkan) kita hanya perlu mengirimkan request ke API yang disediakan oleh platform yang kita gunakan untuk autentikasi sebelumnya.

## Implementasi
[x] Google
[x] Github
[x] Facebook
