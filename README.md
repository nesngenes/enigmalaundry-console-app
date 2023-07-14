Dokumentasi Project Enigma Laundry
Agnes Maria Anggelina


About Project: 


Project ini dibuat dengan menggunakan bahasa pemrograman Go dan PostgreSQL sebagai databasenya. 

Cara menjalankan Program Enigma Laundry:

A) Buka teriminal kemudian jalankan perintah go run main.go, kemudian akan tampil
pilihan tabel seperti di bawah ini: 

![gambar1](images/gambar1.png)


User dapat menginput angka 1 untuk menuju ke tabel customer, 2 untuk ke tabel service dan 3 untuk menuju ke tabel transaksi. 


B) JIka user menuju ke tabel master yaitu tabel Customer dan Service, maka user akan diberikan pilihan untuk VIEW, INSERT, UPDATE, dan DELETE data

![gambar2](images/gambar2.png)


C) Jika user menuju ke tabel transaksi, user hanya akan diberikan pilihan untuk VIEW dan INSERT data
![gambar3](images/gambar3.png)


D) Jika user menginput angka 1 untuk ke menu VIEW make akan ditampilkan data dari database dengan format tabel yang diperoleh dari package github.com/olekukonko/tablewriter
![gambar4](images/gambar4.png)



E) Jika user memilih menu INSERT, maka user dapat memasukkan data melalu console 
![gambar5](images/gambar5.png)



F) Jika user memilih menu UPDATE, maka user harus memasukkan id data yang 
ingin diedit
![gambar6](images/gambar6.png)

G) Jika user memilih menu DELETE, maka user juga harus memasukkan id data
yang ingin didelete

