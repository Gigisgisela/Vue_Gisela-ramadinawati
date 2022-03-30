# vue helper

## Resume
navigasi adalah kegiatan berpindah dari satu halaman ke halaman lain.
navigasi beranak adalah sub-path dari path yang ada.
navigasi dinamis adalah dimana path tersebut memiliki nilai sembarang dan kita perlu mendefinisikan nilainya satu persatu.

layout adalah susunan tata letak,sedangkan layout templet pada vue dalah komponen yang dapat di pakai sebagai sususnan tataletak dasar yang membungkus masing masing halaman.
layout juga dapat diterapkan untuk mengisolasi logika tampilan viewport supaya kode di setiap halaman tetap bersih dan efisien.

penyimpanan global adalah sebuah metode untuk menyimpan variable yang dapat di akses dengan mudah di seluruh bagian aplikasi.
- state komponen 
- props
- store

vuex adalah pola manajemen penyimpanan atau store yang berbentuk pustaka untuk aplikasi vue.js.
vuex mutations berfungsi untuk memodfikasi nilai dari state yang tersimpan di dalam store mnggunakan sintaks mutation.

mutations hanya bertujuan untuk mengubah nilai variable yang ada di dalam store tanpa ada logika pengolahan muatan di dalam fungsi mutation tersebut.

vuex actions bertugas seagai pintu masuk perintah yang menghubungkan komponen dengan store,actions perlu memanggil fungsi di mutations untuk memodifikasi nilai state yang ada di store. selain itu, action juga tempat untuk melakukan komunikasi dgn API.

## task
membuat sebuah todo list yang di lengkapi dengan tombol hapus dan edit
![Screenshot (19)](https://user-images.githubusercontent.com/98401396/160293070-ef661634-0c4c-462f-a7cc-c7ad364afeb2.png)

setiap list yang ada jika kita klik list itu maka harus muncul sebuah deskripsi 
![Screenshot (20)](https://user-images.githubusercontent.com/98401396/160293158-5472d93b-29aa-4118-8a15-2d085c77ae8f.png)

setiap kita menekan tombol edit mak list itu akan berubah sesuai yang kita edit
![Screenshot (21)](https://user-images.githubusercontent.com/98401396/160293211-ba3bb9bc-f7a7-4e66-a74d-0ac90c7fd0c4.png)
![Screenshot (22)](https://user-images.githubusercontent.com/98401396/160293233-96dc2d98-2fce-4c51-bf74-dabcb58245a1.png)
