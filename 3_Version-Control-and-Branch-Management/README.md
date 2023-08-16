## (3) Version Control and Branch Management
1. Tool populer yang digunakan untuk VCS adalah git. Github adalah layanan git secara online; kita bisa membackup git lokal ke dalam github dan juga mempermudah kolaborasi dengan
developer yang lain
2. Pull request adalah sebuah fitur untuk memberitahu developer lain bahwa kita sudah membuat perubahan pada branch tertentu dan ingin menggabungkan perubahan ke dalam branch lain.
Juga bisa dipakai untuk tempat diskusi dan cross-check
3. Workflow git yang baik adalah dengan memiliki branch master, develop, dan fitur. Setiap ada fitur baru, maka dibuat suatu branch fitur baru yang akan di merge ke dalam develop.
Kemudian branch develop akan di merge ke dalam branch master. Jika terdapat bugs pada versi master dan perlu perbaikan yang segera, dibuat suatu branch hotfix yang boleh di merge
lansung ke dalam master.