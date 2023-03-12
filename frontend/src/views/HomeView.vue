<template>
  <div class="container">
    <div v-if="!isConnecting && !isConnected" class="pulse-red">
      <h1>Ooops connection lost.</h1>
    </div>
    <div v-if="isConnecting && !isConnected" class="pulse-orange">
      <h1>Connecting {{ isConnecting }}</h1>
    </div>
    <div v-if="isConnected && !isConnecting"
      :class="{ 'bg-blue': task == 'work', 'pulse-green': task == 'break' || task == 'longbreak' }">
      <h1> {{ task }}</h1>
      <h1> {{ msg }}</h1>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  align-content: center;
  justify-content: center;
  align-items: center;
}

.bg {
  width: 100%;
  height: 100%;
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
  align-items: center;
  align-content: center;
  flex-direction: column;
}

.pulser {
  @extend .bg;
  animation: pulse-red 4s infinite;
}

.pulse-red {
  @extend .pulser;
  animation-name: pulse-red;
}

.pulse-orange {
  @extend .pulser;
  animation-name: pulse-orange;
}

.pulse-green {
  @extend .pulser;
  animation-name: pulse-green;
}

@keyframes pulse-red {
  0% {
    background-color: #F23F42;
  }

  50% {
    background-color: transparent;
  }

  100% {
    background-color: #F23F42;
  }
}

@keyframes pulse-orange {
  0% {
    background-color: #BF861C;
  }

  50% {
    background-color: transparent;
  }

  100% {
    background-color: #BF861C;
  }
}

@keyframes pulse-green {
  0% {
    background-color: #23A559;
  }

  50% {
    background-color: transparent;
  }

  100% {
    background-color: #23A559;
  }
}

// pulse blue
@keyframes pulse-blue {
  0% {
    background-color: #1E90FF;
  }

  50% {
    background-color: transparent;
  }

  100% {
    background-color: #1E90FF;
  }
}

.bg-blue {
  @extend .bg;
  background-color: #1E90FF;
  color: #fff;
}

.bg-green {
  @extend .bg;
  background-color: #23A559;
  color: #fff;
}
</style>

<script setup lang="ts">
import ReplikTime from "@/repliktime";
import { ref, watch } from "vue";

const isConnecting = ReplikTime.connecting;
const isConnected = ReplikTime.connected;
const task = ReplikTime.task;
const expire = ReplikTime.time;
const msg = ref("Hello World");

watch(isConnecting, (val) => {
  console.log("isConnecting", val);
  // isConnecting.value = val;
});

watch(isConnected, (val) => {
  console.log("isConnected", val);
  // isConnected.value = val;
});

watch(task, (val) => {
  console.log("task", val);
  // task.value = val;
});

watch(expire, (val) => {
  console.log("expire", val);
  // expire.value = val;
});

setInterval(() => {
  msg.value = beutifyTime(getExpire());
}, 1000);


function getExpire() {
  const x = Math.round(((expire.value - Date.now()) / 1000));
  return x > 0 ? x : 0;
}

function beutifyTime(time: number) {
  const minutes = Math.floor(time / 60);
  const seconds = time - minutes * 60;
  return `${minutes}:${seconds}`;
}

</script>