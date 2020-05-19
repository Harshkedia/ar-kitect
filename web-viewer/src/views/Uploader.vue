/* eslint-disable no-plusplus */ /* eslint-disable func-names */
<template>
  <div id="file-drag-drop">
    <form ref="fileform">
      <span ref="formtext" class="drop-files">Drop your model here</span>
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
              this.$refs.fileform.style.backgroundColor = "grey";
            },
            false
          );
        }
      );

      this.$refs.fileform.addEventListener("drop", e => {
        e.dataTransfer.files.forEach(file => this.files.push(file));
        this.$refs.fileform.style.backgroundColor = "rgb(51, 51, 138)";
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

<style>
form {
  display: block;
  height: 400px;
  width: 400px;
  background: rgb(51, 51, 138);
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
  background-color: rgb(51, 51, 138);
  border: none;
  color: white;
  text-align: center;
  font-size: 12px;
  padding-top: 5px;
  padding-bottom: 5px;
}
</style>
