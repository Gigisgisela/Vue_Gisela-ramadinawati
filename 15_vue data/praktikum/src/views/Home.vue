<template>
  <div class="home">
    <h1>Berita Terkini</h1>
    <div class="news" v-for="(message, index) in newsList" :key="index">
      <router-link :to="'/deskripsi/id' + index">
        <img :src="message.urlToImage" alt="gambar">

        <div class="author">
          <p v-if="newsList[index].author === null">anonymous</p>
          <p v-else>{{ message.author }}</p>
        </div>
        <div class="title">
          <h3> {{ message.title }}</h3>
        </div>
        <div class="publishTime">
          <p> {{ message.publishedAt}} </p>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script>



export default {
  name: 'HomeView',
  mounted(){
    this.$store.dispatch("fetchNews")
  },
  computed: {
    newsList(){
      return this.$store.state.list
    }
  }
  
}
</script>
<style scoped>
* {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  color: rgb(28, 27, 32);
}
h1 {
  text-align: center;
  margin: 20px 0;
}
.home {
  width: 80%;
  margin: 0 auto;
  padding: 20px;
}
.news {
  background-color: #f2f2f2;
  margin: 20px;
  padding: 20px;
  text-align: center;
}

a {
  text-decoration: none;
  color: rgb(28, 27, 32);

}
img {
  width: 40%;
}
.author {
  font-size: 14px;
  margin: 12px 0;
}
.title {
  font-size: 20px;
  margin: 12px;
}
.publishTime {
  font-size: 12px;
  margin-top: 12px;
}

</style>