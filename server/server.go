package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type FileSystem struct {
	fs http.FileSystem
}

type Message struct {
	FileName string
	FileData string
}

func (m Message) writeToFile() (string, error) {

	var msg string
	dec, err := base64.StdEncoding.DecodeString(m.FileData)
	if err != nil {
		msg = "failed to decode to file"
		return msg, err
	}

	f, err := os.Create(m.FileName + ".obj")
	if err != nil {
		return "error creating file", err
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "error writing file", err
	}

	// go to begginng of file
	f.Seek(0, 0)
	// output file contents
	io.Copy(os.Stdout, f)

	return m.FileName, nil
}

func usdz(w http.ResponseWriter, req *http.Request) {
	// read json
	decoder := json.NewDecoder(req.Body)
	var t Message
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("error parsing json")
		fmt.Fprint(w, "error parsing json")
		return
	}

	msg, err := t.writeToFile()
	if err != nil {
		fmt.Fprint(w, msg)
		return
	}

	// convert to obj
	commandArgs := []string{"-i", t.FileName + ".obj", "-o", "./models/" + t.FileName + ".gltf"}
	_, err = exec.Command("obj2gltf", commandArgs...).Output()
	if err != nil {
		// log.Fatal(err)
		fmt.Fprint(w, "failed to convert to gltf")
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "convert to gltf successful \n")

	// convert to usdz
	commandArgs = []string{"./models/" + t.FileName + ".gltf", "./models/" + t.FileName + ".usdz"}
	_, err = exec.Command("usd_from_gltf", commandArgs...).Output()
	if err != nil {
		// log.Fatal(err)
		fmt.Fprint(w, "failed to convert to usdz")
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "convert to usdz successful \n")

	fmt.Fprintf(w, t.FileName+".usdz")
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
