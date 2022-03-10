//clean Code
//  no 1
// kode berikut ini dituliskan tanpa mengikuti kebiasaan-kebiasaan penulis yang disarankan 

// berapa banyak kekurangan dalam penilisan kode tersebut?
// ada 2 kekurangan yang ada

// bagian mna saja terjadi kekurangan tersebut?
// 1. kekurangan ada padabagian pendeklarasian class 
// 2. pendeklarasian isi dari class userservice

// tuliskan alsan dari tiap kekurangan tersebut
// 1. code susah untuk dieja 
// 2. susah memahami kode

// no2

// JS

class kendaraan {
    var totalRoda = 0;
    var kecepatanPerJam = 0;
}

class mobil extends kendaraan {
    void berjalan () {
        tambahKecepatan(10);
    }

    tambahKecepatan(var kecepatanBaru){
        kecepatanPerJam +=  kecepatanBaru;
    }
}

void main() {
    mobilCepat = new mobil();
    mobilCepat.berjalan();
    

    mobilLamban = new mobil();
    mobilLamban.berjalan();
}