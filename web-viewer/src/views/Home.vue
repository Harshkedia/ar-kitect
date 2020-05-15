<template>
  <div class="home">
    <model-viewer
      class="model"
      src="https://srv-file7.gofile.io/download/CE9GDK/Intermediate.obj.glb"
      ar
      ar-modes="webxr scene-viewer quick-look fallback"
      ar-scale="auto"
      magic-leap
      camera-controls
      alt="A 3D model of an astronaut"
      ios-src="https://srv-file7.gofile.io/download/CE9GDK/Intermediate.obj.usdz"
    />
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
      return `http://ec2-13-233-130-134.ap-south-1.compute.amazonaws.com/models/${this.$route.query.name}.glb`;
    },
    urlIos() {
      return `http://ec2-13-233-130-134.ap-south-1.compute.amazonaws.com/models/${this.$route.query.name}.usdz`;
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
  height: 800px;
  width: 800px;
  margin: 0 auto;
}
</style>
