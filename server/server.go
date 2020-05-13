package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Message struct {
	FileName     string
	FileData     string
	FileMaterial string
}

func expireFiles(fname []string) {
	time.Sleep(1 * time.Hour)
	for _, f := range fname {
		os.Remove("./models/" + f)
		fmt.Println("deleted" + f)
	}
}

func copyContentsTofile(content string, fname string) (string, error) {

	dec, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "failed to decode to file", err
	}

	f, err := os.Create(fname)
	if err != nil {
		return "error creating file", err
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "error writing file", err
	}

	return fname, nil
}

func (m Message) writeToFile() (string, error) {

	msg, err := copyContentsTofile(m.FileData, m.FileName+".obj")
	if err != nil {
		return "failed to create obj: " + msg, err
	}
	if m.FileMaterial != "" {

		msg, err = copyContentsTofile(m.FileMaterial, m.FileName+".mtl")
		if err != nil {
			return "failed to create mtl: " + msg, err
		}

	}
	return "success", nil
}

func usdz(w http.ResponseWriter, req *http.Request) {
	// read json
	decoder := json.NewDecoder(req.Body)
	var t Message
	var err error
	err = decoder.Decode(&t)
	if err != nil {
		log.Println("error parsing json")
		fmt.Fprint(w, "error parsing json")
		return
	}

	_, err = t.writeToFile()
	if err != nil {
		fmt.Fprint(w, "failed to create obj")
		return
	}

	// convert to obj
	commandArgs := []string{"-i", t.FileName + ".obj", "-o", "./models/" + t.FileName + ".gltf"}
	_, err = exec.Command("obj2gltf", commandArgs...).Output()

	defer os.Remove(t.FileName + ".obj")
	if t.FileMaterial != "" {

		defer os.Remove(t.FileName + ".mtl")
	}

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Fprint(w, "failed to convert to gltf")
		return
	}
	fmt.Fprint(w, "convert to gltf successful \n")

	// convert to usdz
	commandArgs = []string{"./models/" + t.FileName + ".gltf", "./models/" + t.FileName + ".usdz"}
	_, err = exec.Command("usd_from_gltf", commandArgs...).Output()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Fprint(w, "failed to convert to usdz")
		return
	}
	fmt.Fprint(w, "convert to usdz successful \n")

	fmt.Fprintf(w, t.FileName+".usdz")
	go expireFiles([]string{t.FileName + ".gltf", t.FileName + ".usdz"})
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", usdz)
	http.HandleFunc("/headers", headers)
	http.Handle("/models/", http.StripPrefix(strings.TrimRight("/models/", "/"), http.FileServer(http.Dir("models"))))
	http.ListenAndServe(":8090", nil)
}
