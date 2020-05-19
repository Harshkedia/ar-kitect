<template>
  <div class="home">
    <div class="input">
      <input
        v-if="!textFilled"
        v-model="modelName"
        class="input-text"
        type="text"
        placeholder="
      Input Model Name"
      >
      <button v-if="!textFilled" class="input-button" @click="load">Load Model</button>
    </div>
    <model-viewer
      v-if="textFilled"
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
      quick-look-browsers="safari chrome"
    >
      <button slot="ar-button" class="activate-ar">
        Activate AR
      </button>
    </model-viewer>
    <button v-if="textFilled" class="back-button" @click="load">Back</button>
  </div>
</template>

<script>
// @ is an alias to /src
import "@google/model-viewer";

export default {
  name: "Home",
  components: {},
  data() {
    return {
      textFilled: false,
      modelName: ""
    };
  },
  computed: {
    url() {
      return `https://ar.portfo.io/models/${this.modelName}.glb`;
    },
    urlIos() {
      return `https://ar.portfo.io/models/${this.modelName}.usdz`;
    }
  },
  methods: {
    load() {
      this.textFilled = !this.textFilled;
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
  height: 75%;
  /* border-style: solid;
  border-width: 5px;
  border-radius: 30px; */
  border-color: rgb(51, 51, 138);
  background-color: rgb(163, 163, 228);
}

.activate-ar {
  background-color: rgb(51, 51, 138);
  height: 50px;
  width: 100px;
  border-radius: 10px;
  color: white;
  font-size: 15px;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
}

.input {
  width: 200px;
  height: 40px;
  position: absolute;
  top: -5%;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
}

.input-text {
  width: 200px;
  height: 30px;
  font-size: 15px;
  background-color: rgb(51, 51, 138);
  color: white;
  padding: 5px;
  border-style: hidden;
  font-weight: bold;
  text-align: center;
}

::placeholder {
  color: rgb(104, 104, 175);
  opacity: 1; /* Firefox */
}
.input-button {
  width: 100px;
  height: 20px;
  margin: auto;
  margin-top: 10px;
  border-radius: 0px;
  background-color: rgb(51, 51, 138);
  border: none;
  color: white;
  text-align: center;
  font-size: 12px;
  padding-top: 5px;
  padding-bottom: 20px;
}

.back-button {
  position: absolute;
  top: 90%;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  width: 100px;
  height: 20px;
  margin: auto;
  margin-top: 10px;
  border-radius: 0px;
  background-color: rgb(51, 51, 138);
  border: none;
  color: white;
  text-align: center;
  font-size: 12px;
  padding-top: 5px;
  padding-bottom: 20px;
}
</style>
