# vue fundamental 

## Resume 

vue adalah framework dari javascript untuk mengembangkan tamilan website yang lebih interaktif dan dinamis 

kenapa vue?
- mudah membuat aplikasi frontend dan website
- mudah dipelajari dan menyenangkan 
- dokumentasi yang sangat lengkap dan rapi
- ramping dan cepat
- populer

vue instance-> setiap aplikasi vue dimulai dengan membuat instance vue baru dengan fungsi vue

vue data binding->vue.js menggunakan sintaks template berbasis html yang memungkinkan kita untuk secara deklaratif mengikat DOM yang dirender ke data instance vue yang mendasarinya
ada tiga jenis data binding pada vue:
- di dalam konten
- di atribut elemen
- elemen html

vue reactivity-. vue instance memiliki property bernama data jika value dari data ada yang berubah maka value yang ditampilkanya pada interface akan berubah otomatis tanpa harus dimuat ulang

vue directive-> atribut khusus yang diawali dengan v-
directve berfungsi untuk menjalankan suatu perintah atau ekspresi java script di dalam atribut.
macam vue directive:
- v-bind
- v-model
- v-if,v-else,v-else-if
- v-on
- v-for

khusus untuk v-bind dan v-on dapat disingkat
v-bind menjadi :
v-on menjadi @

1. direktive data binding :v-bind, v-model,v-htnl,v-text
2. direktive conditional rendering:v-if,v-else,v-else-if
3. direktive pengulangan : v-for
4. direktife event :v-on

komponen adalah vue instance yang dapat digunakan kembali dgn nama yang kita definisikan 

##TASK

membuat to do list sederhana 
![Screenshot (13)](https://user-images.githubusercontent.com/98401396/158540088-9f7600ef-21f3-4152-81e2-c54f032cd104.png)

