

<template>
  <!--<header>
    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/about">About</RouterLink>
        </nav>
      </div>
    </header>-->
  <!--<header>
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>

  </header>-->

  <div class="wrapper">
    <RouterView v-if="clicked" />
    <div v-if="!clicked">
      <h1>Click to connect</h1>
      <input class="input-field" type="text" v-model="username" placeholder="Username"  />
      <button class="btn" @click="btnClick()">Connect</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ReplikTime from './repliktime';
import { ref } from 'vue'


const clicked = ref(false);
const username = ReplikTime.username;

function btnClick() {
  if (username.value.length < 3) {
    alert("Please enter a username");
    return;
  }
  clicked.value = true; 
  ReplikTime.connect().catch((err) => {
    console.log(err);
    clicked.value = false;
  })
}

</script>


<style scoped lang="scss">
// import css vars on base
header {
  // punt on top
  width: 100%;
  background-color: lighten(#181818, 5%);
  color: #fff;
  width: 100%;
  background-color: lighten(#181818, 5%);
  color: #fff;
  height: 70px;
  display: flex;
  align-items: center;

}

nav {
  justify-content: flex-start;
  align-items: center;
  display: flex;
  margin-left: 5px;

  a {
    margin-left: 15px;

    &.router-link-active {
      text-decoration: underline;
    }
  }

  // if a is current route    
}

.wrapper {
  display: block;
  max-height: 100%;
  height: 100%;

  >* {
    max-height: 100%;
    height: 100%;
  }
}

.input-field {
  min-height: 35px;
  min-width: 80px;
  
  border: 1px solid lighten(#181818, 5%);
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  background-color: #181818;
  color: #fff;
}

.btn {
  min-height: 35px;
  min-width: 80px;
  border: 0px solid lighten(#181818, 5%);
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  background-color: hsla(160, 100%, 37%, 1);
  color: #fff;
  font-size: 1.1rem;
}
</style>
