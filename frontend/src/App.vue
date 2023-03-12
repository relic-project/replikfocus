

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
  <header>
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>

  </header>

  <div class="wrapper">
    <RouterView v-if="clicked" />
    <div v-if="!clicked">
      <h1>Click to connect</h1>
      <button @click="clicked = true; ReplikTime.connect()">Connect</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ReplikTime from './repliktime';
import { ref } from 'vue'


const clicked = ref(false);

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
  max-height: calc(100% - 70px);
  height: calc(100% - 70px);

  >* {
    max-height: calc(100%);
    height: calc(100% );
  }
}
</style>
