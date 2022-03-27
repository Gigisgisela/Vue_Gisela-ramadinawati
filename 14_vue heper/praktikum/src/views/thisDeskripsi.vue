<template>
  <div id="thisDeskripsi">
  <button class="btn m-5"><router-link to="/">kembali</router-link></button>
      <div class="card">
          <div class="card-header">
              <h4>
                  {{ getDataTodos[idTodos].todo }}
              </h4>
          </div>
          <div class="card-body">
                  {{ getDataTodos[idTodos].deskripsi }}
          </div>
          <div class="card-footer">
              <input type="text" v-model="newDeskripsi"> <button class="btn" @click="editDeskripsi(idTodos, getDataTodos[idTodos].todo)">Edit</button>
          </div>
      </div>
  </div>
</template>

<script>
export default {
    name: 'thisDeskripsi',
    data() {
        return {
            newDeskripsi: '',
            sendTodo: [],
           
        }
    },
    computed:{
        getDataTodos(){
            return this.$store.state.todos
        },
        idTodos(){
            return this.$route.params.id
        }
    },
    methods: {
        editDeskripsi(id,todo){

            if(this.newDeskripsi != ''){
                this.deskripsi = this.newDeskripsi;
                this.sendTodo = {
                   id: id,
                    deskripsi:this.deskripsi,
                    todo: todo,
                }
                this.$store.dispatch("editDeskripsi",this.sendTodo);
                this.newDeskripsi = ''
            }
        }
    }
}
</script>

<style>

</style>