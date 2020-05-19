<template>
  <div class="home">
    <h1>AR Viewer</h1>
    <model-viewer
      class="model"
      :src="url"
      ar
      ar-modes="webxr scene-viewer quick-look fallback"
      ar-scale="auto"
      magic-leap
      camera-controls
      alt="A 3D model of an astronaut"
      :ios-src="urlIos"
      loading="eager"
      shadow-intensity="10"
      shadow-softness="1"
      :poster="require(`../assets/loading.gif`)"
    >
      <button slot="ar-button" class="activate-ar">
        Activate AR
      </button>
    </model-viewer>
  </div>
</template>

<script>
// @ is an alias to /src
import "@google/model-viewer";

export default {
  name: "Home",
  components: {},
  computed: {
    url() {
      return `https://ar.portfo.io/models/${this.$route.query.name}.glb`;
    },
    urlIos() {
      return `https://ar.portfo.io/models/${this.$route.query.name}.usdz`;
    }
  }
};
</script>

<style scoped>
.model {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  box-sizing: border-box;
  vertical-align: middle;
  position: absolute;
  margin-left: 5%;
  width: 90%;
  height: 85%;
  border-style: solid;
  border-width: 5px;
  border-radius: 30px;
  border-color: rgb(51, 51, 138);
  background-color: rgb(163, 163, 228);
}

.activate-ar {
  background-color: rgb(51, 51, 138);
  position: relative;
  height: 50px;
  width: 100px;
  margin-top: 50%;
  border-radius: 10px;
  color: white;
  font-size: 15px;
}
</style>
