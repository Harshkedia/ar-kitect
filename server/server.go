package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// const (
// 	fbx = ".fbx"
// 	obj = ".obj"
// )

type message struct {
	FileFormat  string
	FileContent http.Request
	FileNames   []string
}

func expireFiles(fnames []string) {
	time.Sleep(1 * time.Hour)
	var fname string
	for _, f := range fnames {
		fname = "./models/" + f
		if _, err := os.Stat(fname); err != nil {
			os.Remove(fname)
			fmt.Println(fname)
		}
	}
}

func (m *message) receiveFiles() (string, error) {
	m.FileNames = []string{}
	reader, err := m.FileContent.MultipartReader()
	if err != nil {
		log.Println(err)
		return "something wrong with multipart", err
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		defer part.Close()
		if part.FileName() == "" {
			continue
		}
		m.FileNames = append(m.FileNames, part.FileName())
		log.Println("filename: " + part.FileName())
		d, err := os.Create(part.FileName())
		if err != nil {
			// log.Fatal(err)
			return "failed to write file", err
		}
		defer d.Close()
		io.Copy(d, part)
	}
	return "success", nil
}

func (m *message) writeToFile() (string, error) {

	msg, err := m.receiveFiles()
	if err != nil {
		return msg, err
	}
	return "success", nil
}

func usdz(w http.ResponseWriter, req *http.Request) {
	// read json
	// decoder := json.NewDecoder(req.Body)
	var t message
	var err error
	t.FileContent = *req
	t.FileFormat = req.URL.Query().Get("mode")
	if t.FileFormat != "obj" {
		if t.FileFormat != "fbx" {
			fmt.Fprintf(w, "mode parameter invalid")
			return
		}
	}

	msg, err := t.writeToFile()
	if err != nil {
		fmt.Fprint(w, "failed to create obj"+msg+"\n")
		return
	} else {
		if len(t.FileNames) == 0 {
			fmt.Fprint(w, "missing attachments \n")
			return
		}
	}

	var commandArgs []string
	var fname string
	fname = t.FileNames[0]
	// fmt.Printf("fname: %s, FileNames %v, length: %d", fname, t.FileNames, len(t.FileNames))
	for _, fname := range t.FileNames {
		defer os.Remove(fname)
	}

	if t.FileFormat == "obj" {
		if !strings.HasSuffix(fname, ".obj") {
			fname = t.FileNames[1]
		}
		log.Println("converting fileformat obj")
		commandArgs = []string{"-i", fname, "-o", "./models/" + fname + ".glb"}
		_, err = exec.Command("obj2gltf", commandArgs...).Output()
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			fmt.Fprint(w, "failed to convert to gltf\n")
			return
		}
	} else if t.FileFormat == "fbx" {

		log.Println("converting fileformat fbx")
		commandArgs = []string{"--binary", "-i", fname, "-o", "./models/" + fname + ".glb"}
		_, err = exec.Command("./FBX2glTF", commandArgs...).Output()
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			fmt.Fprint(w, "failed to convert to gltf\n")
			return
		}
	}

	log.Println("convert to glb successful")

	// convert to usdz
	commandArgs = []string{"./models/" + fname + ".glb", "./models/" + fname + ".usdz"}
	_, err = exec.Command("usd_from_gltf", commandArgs...).Output()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Fprint(w, "failed to convert to usdz")
		return
	}
	log.Println("convert to usdz successful")

	fmt.Fprintf(w, fname+".usdz")
	go expireFiles([]string{fname + ".glb", fname + ".usdz"})
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
