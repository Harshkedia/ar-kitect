<template>
  <div id="home">
    <div class="input">
      <input
        v-if="!textFilled"
        v-model="modelName"
        class="input-text"
        type="text"
        placeholder="Search For Your Model"
        onfocus="this.placeholder = ''"
        onblur="this.placeholder = 'Search For Your Model'"
      >
      <br >
      <button v-if="!textFilled" class="input-button" @click="load">Load</button>
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
      alt="A 3D Model"
      :ios-src="urlIos"
      loading="eager"
      :poster="require(`../assets/loading.gif`)"
      quick-look-browsers="safari chrome"
      exposure="0.1"
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
    url: {
      get() {
        return "https://srv-file6.gofile.io/download/ki8pTg/Intermediate.glb";
      },
      set() {}
    },
    urlIos: {
      get() {
        return "https://srv-file18.gofile.io/download/t5CzrI/Intermediate.usdz";
      },
      set() {}
    }
  },
  mounted() {
    if (this.$route.query.name !== undefined) {
      this.textFilled = true;
      this.modelName = this.$route.query.name;
    }
  },
  methods: {
    load() {
      this.textFilled = !this.textFilled;
    }
  }
};
</script>

<style lang="scss" scoped>
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
  background-color: white;
}

.activate-ar {
  background-color: $navy;
  height: 50px;
  width: 100px;
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
  width: 80%;
  height: 40px;
  position: absolute;
  top: -40%;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
}

.input-label {
  color: white;
  font-weight: bolder;
  font-size: 2.5vh;
}

.input-text {
  width: 80%;
  height: 5vh;
  font-size: 20px;
  background-color: $dark-gray;
  color: white;
  padding: 5px;
  border-style: hidden;
  font-weight: bold;
  text-align: center;
}

::placeholder {
  color: $light-gray;
}

.input-button {
  width: 100px;
  height: 20px;
  margin: auto;
  margin-top: 4vh;
  border-radius: 0px;
  background-color: $light-gray;
  border: none;
  color: black;
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
  background-color: $light-gray;
  border: none;
  color: black;
  text-align: center;
  font-size: 12px;
  padding-top: 5px;
  padding-bottom: 20px;
}
</style>
