<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">AR Viewer</router-link>
      <router-link v-if="isMobile" to="/upload">AR Uploader</router-link>
    </div>
    <transition name="fade">
      <router-view />
    </transition>
  </div>
</template>

<script>
export default {
  computed: {
    isMobile() {
      if (
        /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
      ) {
        return false;
      }
      return true;
    }
  }
};
</script>
<style>
body,
html {
  padding: 0;
  margin: 0;
  width: 100%;
  min-height: 100vh;
}
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  box-sizing: border-box;
  margin: 0 auto;
  font-family: Arial, Helvetica, sans-serif;
}

#nav {
  padding: 30px;
}
#nav a {
  font-weight: bold;
  color: black;
  text-decoration: none;
  margin: 20px;
  padding: 10px;
  color: black;
}

#nav a.router-link-exact-active {
  color: rgb(51, 51, 138);
}

.fade-enter-active,
.fade-leave-active {
  transition-property: opacity;
  transition-duration: 0.25s;
}
.fade-enter-active {
  transition-delay: 0.25s;
}
.fade-enter,
.fade-leave-active {
  opacity: 0;
}
</style>
