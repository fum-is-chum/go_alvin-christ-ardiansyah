## (22) Compute Services
1. Salah satu cloud compute service yang tersedia yaitu Amazon EC2
2. Langkah Deploy ke EC2:
   - Membuat instance ec2, pilih image amazon linux 2
   - Mengenerate RSA key untuk koneksi ssh ke instance nantinya
   - Mengatur security group, edit inbound rule sesuai kebutuhan (contoh: aplikasi go memerlukan port 8000, maka tambah rule custom TCP port 8000)
   - Menggunakan elastic IP supaya IP Address publik tidak berubah-ubah
   - Build dan jalankan aplikasi Go
3. Langkah Setup Database:
   - Pilih service RDS pada console aws
   - Setup database (initial db name, db alias, username, password)
   - Set up EC2 Connection ke instance ec2 yang dibuat di step sebelumnya