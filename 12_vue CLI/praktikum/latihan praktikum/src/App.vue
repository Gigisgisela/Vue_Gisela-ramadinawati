<template>
  <div id="app">
    <ol>
      <li v-for="(messege, index) in messeges" :key="messege.id">
        {{ messege }}
        <button v-on:click="deleteList(index)">Hapus</button>
      </li>
      
    </ol>
    
    <p v-if="messeges.length == 0">List kegiatan anda masih kosong</p>

    <input type="text" v-on:keyup.enter="addList(messege)" v-model="messege">
    <button v-on:click="addList(messege)">Tambahkan</button>

    <p v-if="messeges.length >= 4">Hebat!</p>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      messege : "",
      messeges : []
    }
  },
  methods : {
    addList(messege) {
      if (messege == "") {
        alert("Data yang dimasukan tidak boleh kosong")
      } else {
        this.messeges.push(this.messege)
        this.messege = ""
      }
      
    },
    deleteList(index) {
      this.messeges.splice(index, 1)
    }
  }
}
</script>

<style>

</style>