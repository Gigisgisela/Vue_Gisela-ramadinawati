<template>

  <div id="app">
    <div class="isiContent">
      <h1>Todo List</h1>
      <table border="1" width="100%">
        <tr v-for="(message, index) in messages " :key="message.id">
          <td width="80%">
          {{ index + 1 }} {{ message }}
          </td>
          <td>
            <button @click="deletMessage(index)">Hapus</button>
            <button @click="editMessage(index, message)">Edit</button> 
          </td>
        
        </tr>
      </table>
      <div class="inputmessage" v-if="editing == true">
        <input type="text" v-model="message" @keyup.enter="update">
        <button @click="update">Update </button>
      </div>
      <div class="inputmessage" v-else>
        <input type="text" v-model="message" @keyup.enter="addList(message)">
        <button @click="addList(message)">Tambahkan</button>
      </div>
      

      
    </div>

    
  </div>
</template>

<script>
export default {
  name: 'App',
  data(){
    return {
      message: "",
      messages: [],
      editing : false,
      selectIndex : null

    }
  },
  methods:{
    addList(message){
      
      if (message==""){
        alert("data yang di masukan tidak boleh kosong")
      } else{
        this.messages.push(this.message)
        this.message=""
      }
    },
    deletMessage(index){
      this.messages.splice(index, 1)
    },
    editMessage(index, message){
      this.editing = true
      this.message = message
      this.selectIndex = index
    },
    update() {
      this.messages.splice(this.selectIndex, 1, this.message)
      this.message= ""
      this.editing= false
    }
  }
  

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;

}
.isiContent {
  width: 80%;
  margin: 0 auto;
}

.inputmessage input {
  width: 80%;
}

.inputmessage button {
  width: 19%;
}
</style>
