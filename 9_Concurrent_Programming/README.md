## (9) Concurrent Programming
1. Parralel tidak sama dengan Concurrent. Parallel mengeksekusi program secara bersamaan pada core yang berbeda, Concurrent mengeksekusi program secara independen sehingga seolah-olah terlihat dieksekusi secara bersamaan
2. Keuntungan Concurrency pada Golang adalah memanfaatkan multi processor, sehingga concurrency pada Golang juga termasuk Parralelism
3. Untuk mendapatkan Concurrency, pada Go tersedia Goroutines dan Channels. Untuk menghindari race condition, dapat digunakan WaitGroups, Channel Blocking dan Mutex