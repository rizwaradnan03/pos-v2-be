<!-- CARA PENGGUNAAN -->
1. instalasi go lang di website: https://go.dev/dl/
2. melakukan setup environment pada file: .env
3. membuat database di postgresql bernama nexuz_be
4. eksekusi kode query sql (untuk aktivasi uuid): CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
5. menjalankan: go build -o myapp.exe cmd/main.go -> bertujuan untuk membuild golang dan siap untuk dijalankan
6. menjalankan file hasil build bernama myapp.exe