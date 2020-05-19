const axios = require("axios");

export default function uploadFile(files) {
  let objFile;
  let mtlFile;

  files.forEach(file => {
    if (file.name.includes(".obj")) {
      objFile = file;
    } else if (file.name.includes(".mtl")) {
      mtlFile = file;
    }
  });

  const bodyFormData = new FormData();
  bodyFormData.append(".obj", objFile);
  bodyFormData.append(".mtl", mtlFile);

  console.log(bodyFormData.getAll(".obj"));

  axios
    .post("https://ar.portfo.io/?mode=obj", bodyFormData, {
      headers: {
        "Content-Type": "multipart/form-data"
      }
    })
    .then(res => {
      // handle success
      console.log(res);
    })
    .catch(res => {
      // handle error
      console.log(res);
    });

  // axios({
  //   method: "post",
  //   url: "https://ar.portfo.io//?mode=obj",
  //   data: bodyFormData,
  //   headers: { "Content-Type": "multipart/form-data" }
  // })
  //   .then(res => {
  //     // handle success
  //     console.log(res);
  //   })
  //   .catch(res => {
  //     // handle error
  //     console.log(res);
  //   });
}
