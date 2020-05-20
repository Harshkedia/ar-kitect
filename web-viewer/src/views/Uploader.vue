/* eslint-disable no-plusplus */ /* eslint-disable func-names */
<template>
  <div id="file-drag-drop">
    <form ref="fileform">
      <span ref="formtext" class="drop-files">Drop Model Here</span>
    </form>
    <button class="upload-button" @click="upload">Upload File</button>
  </div>
</template>

<script>
import uploadFile from "../lib/BackendHandler";

export default {
  data() {
    return {
      dragAndDropCapable: false,
      files: []
    };
  },
  mounted() {
    this.dragAndDropCapable = this.determineDragAndDropCapable();

    if (this.dragAndDropCapable) {
      ["drag", "dragstart", "dragend", "dragover", "dragenter", "dragleave", "drop"].forEach(
        evt => {
          this.$refs.fileform.addEventListener(
            evt,
            e => {
              e.preventDefault();
              e.stopPropagation();
              this.$refs.fileform.style.backgroundColor = "#a3a9ac";
            },
            false
          );
        }
      );

      this.$refs.fileform.addEventListener("drop", e => {
        e.dataTransfer.files.forEach(file => this.files.push(file));
        this.$refs.fileform.style.backgroundColor = "#545759";
        this.$refs.formtext.innerHTML = `${this.files.length} File(s) Added`;
      });
    }
  },
  methods: {
    determineDragAndDropCapable() {
      const div = document.createElement("div");
      return (
        ("draggable" in div || ("ondragstart" in div && "ondrop" in div)) &&
        "FormData" in window &&
        "FileReader" in window
      );
    },
    upload() {
      uploadFile(this.files);
    }
  }
};
</script>

<style lang="scss">
form {
  display: block;
  height: 400px;
  width: 400px;
  background: $dark-gray;
  margin: auto;
  margin-top: 40px;
  text-align: center;
  line-height: 400px;
  color: white;
}

div.file-listing {
  width: 400px;
  margin: auto;
  padding: 10px;
  border-bottom: 1px solid #ddd;
}

.upload-button {
  margin-top: 10px;
  border-radius: 0px;
  background-color: $light-gray;
  border: none;
  color: black;
  text-align: center;
  font-size: 12px;
  padding-top: 10px;
  padding-bottom: 10px;
  padding: 10px;
}
</style>
