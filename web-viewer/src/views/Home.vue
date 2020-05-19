<template>
  <div class="home">
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
      <button
        slot="ar-button"
        style="background-color: white; border-radius: 4px; border: none; position: absolute; top: 16px; right: 16px; "
      >
        👋 Activate AR
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
  width: 50%;
  height: 500px;
  margin: 0 auto;
}
</style>
