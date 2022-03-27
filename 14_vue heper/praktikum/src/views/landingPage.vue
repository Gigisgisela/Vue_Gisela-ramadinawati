<template>
  <div id="todoList">
        <div class="card">
            <div class="card-header">
                <h4>Todo List</h4>            
            </div>
            <div class="card-body">
                <div class="list" v-for="(todo, index) in getTodos" :key="index">
                    <div class="link" v-if="todo.modeEdit === false">
                        <div @click="toDeskripsiPage(`${index}`)">
                             {{ index+1 }} {{ todo.todo }}
                        </div>
                       
                    </div>
                    <div class="link" v-else>
                        <input type="text" v-model="valueEditTodo">
                    </div>
                    <div id="btn" v-if="todo.modeEdit === false">
                        <button class="btn bg-danger" @click="deleteTodo(index)">Hapus</button>
                        <button class="btn" @click="onModeEdit(index, todo.todo, todo.deskripsi)">Edit</button>
                    </div>
                    <div id="btn" v-else>
                        <button class="btn bg-danger" @click="deleteTodo(index)">hapus</button>
                    </div>
                </div>

            </div>
            <div class="card-footer" v-if="modeEdit === false">
                <input type="text" v-model="newValueTodo"> <button class="btn" @click="addTodos">Tambahkan</button>
            </div>
            <div class="card-footer" v-else>
                <button class="btn" @click="editTodo(idTodo,todo,deskripsi)">Edit</button>
            </div>
        </div>
  </div>
</template>

<script>
export default {
    name: 'landingPage',
    computed: {
         getTodos() {
             return this.$store.state.todos
         } 
    },
    data() {
        return {
            newValueTodo: '',
            idTodo: 0,
            valueEditTodo: '',

            newDeskripsi: '',
            modeEdit: false,
        }
    },
    methods: {
        addTodos(){
            if(this.newValueTodo != ''){
                this.$store.dispatch("getDataTodos",this.newValueTodo)
                this.newValueTodo = ''
            }
        },
        deleteTodo(id){
            this.$store.dispatch("deleteTodos",id);
        },
        onModeEdit(id,todo,deskripsi){

            this.getTodos.splice(id,1,{todo:todo,deskripsi:deskripsi,modeEdit:true})
            this.idTodo = id
            this.todo = todo
            this.deskripsi = deskripsi
            this.modeEdit = true
        },
        editTodo(id,todo,deskripsi){
            todo = this.valueEditTodo
           const param ={
               id,todo,deskripsi
           }
           this.$store.dispatch("editTodo",param)
           this.modeEdit = false
           this.valueEditTodo = ''
        },
        toDeskripsiPage(id){
            this.$router.push(id);
        }
    }
}
</script>

<style>
body{
        padding:0;
        margin: 0;font-family: Arial, Helvetica, sans-serif;
        
    }
    #todoList{
        box-sizing: border-box;
        width:100%;
        padding-left: 5%;
        padding-right: 5%;
        padding-top: 20px;
    }
    .card{
        box-sizing: border-box;
        width:60%;
        box-shadow: 10px 10px 30px rgba(0,0,0,0.1);
        background-color: pink;
        border-radius: 10px;

    }
    .card-header{
        box-sizing: border-box;
        padding:5px;
        padding-left: 5%;
        padding-right: 5%;
        width: 100%;
        background-color: pink;
    }
    .card-body{
        box-sizing: border-box;
        width:100%;
        padding:10px;
        padding-left: 5%;
        padding-right: 5%;
        border-bottom: 1px solid rgb(206, 167, 167) ;
    }
    .list{
        box-sizing: border-box;
        margin-top:10px;
        margin-bottom: 10px;
        width:100%;
        display:flex;
        flex-flow: row wrap;
        justify-content:space-between; 
    }
    .card-footer{
        box-sizing: border-box;
        padding:5px;
        margin-top:5px;
        padding-bottom:10px;
        padding-left: 3%;
        padding-right: 3%;
        display:flex;
        flex-flow: row wrap;
        justify-content:space-between;
       
    }
    input{
        width:80%;
        padding:5px !important;
        border-radius: 1px;
        border: 1px solid  rgb(104, 92, 240)k;
    }
    h4{
        color:#fff !important;
        font-size: 20px;;
    }
    a{
        color: rgb(105, 0, 0) ;
        text-decoration: none !important;
    }
    #btn {
        width:20%;
        display: flex;
        flex-flow: row wrap;
        justify-content:space-between;
    }
    .text {
        width:70%;
    }
    .btn {
        padding:10px;
        border:1px solid #fff;
        border-radius: 3px;
        background-color: palevioletred;
        color:#ffff !important;
        cursor: pointer;


    }
    .bg-danger{
        background-color:rgb(241, 41, 41) !important;
    }
    label{
        color:blue;
        cursor: pointer;
    }
    label:hover{
        color:blueviolet;
    }
</style>