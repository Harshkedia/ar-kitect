package main

import (
	"crypto/tls"
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"ar-kitect/server/haikunator"

	"golang.org/x/crypto/acme/autocert"
)

var (
	OBJ          = "obj"
	FBX          = "fbx"
	GLTF         = "gltf"
	USDZ         = "usdz"
	OBJ_TO_GLTF  = "obj2gltf"
	FBX_TO_GLTF  = "./FBX2glTF"
	GLTF_TO_USDZ = "usd_from_gltf"
	CERT_DOMAIN  = "CERT_DOMAIN"
)

type message struct {
	FileFormat  string
	FileContent http.Request
	FileNames   []string
}

type middleware struct {
	handler http.Handler
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	m.handler.ServeHTTP(w, r)
}

func newMiddleware(h http.Handler) *middleware {
	return &middleware{h}
}

func expireFiles(fnames []string) {
	time.Sleep(1 * time.Hour)
	var fname string
	for _, f := range fnames {
		fname = "./models/" + f
		if _, err := os.Stat(fname); err != nil {
			_ = os.Remove(fname)
			fmt.Println(fname)
		}
	}
}

func (m *message) receiveFiles() (string, error) {
	namegen := haikunator.New(time.Now().UTC().UnixNano())
	randname := namegen.Haikunate()
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
		}
		if err != nil {
			log.Println(err)
			break
		}
		defer part.Close()
		if part.FileName() == "" {
			continue
		}
		thisfname := randname + extractFileNameWithoutExtension(part.FileName())
		m.FileNames = append(m.FileNames, thisfname)

		log.Printf("filename: %s", thisfname)
		file, err := os.Create(thisfname)
		if err != nil {
			return "failed to write file", err
		}
		defer file.Close()
		_, _ = io.Copy(file, part)
	}
	return "success", nil
}

func extractFileNameWithoutExtension(fname string) string {
	split := strings.Split(fname, ".")
	return strings.Join(split[:len(split)-1], ".")
}
func changeFileNameExtension(fname string, extn string) string {
	split := strings.Split(fname, ".")
	joined := strings.Join(split[:len(split)-1], ".")
	return fmt.Sprintf("%s.%s", joined, extn)
}

func (m *message) writeToFile() (string, error) {

	msg, err := m.receiveFiles()
	if err != nil {
		return msg, err
	}
	return "success", nil
}

func usdzHandler(w http.ResponseWriter, req *http.Request) {
	var t message
	t.FileContent = *req
	t.FileFormat = req.URL.Query().Get("mode")
	if t.FileFormat != OBJ {
		if t.FileFormat != FBX {
			log.Println("mode parameter invalid")
			_, _ = fmt.Fprintf(w, "mode parameter invalid")
			return
		}
	}

	msg, err := t.writeToFile()
	if err != nil {
		log.Println("failed to create obj")
		_, _ = fmt.Fprintln(w, "failed to create obj :"+msg)
		return
	}

	if len(t.FileNames) == 0 {
		log.Println("missing attachments")
		_, _ = fmt.Fprintln(w, "missing attachments")
		return
	}

	fname := t.FileNames[0]

	log.Printf("fname: %s, FileNames : %v, length: %d", fname, t.FileNames, len(t.FileNames))

	// remove received files
	for _, fnm := range t.FileNames {
		defer os.Remove(fnm)
	}

	if t.FileFormat == "obj" {
		ok := convertOBJtoGLTF(w, fname, t)
		if !ok {
			return
		}
	} else if t.FileFormat == "fbx" {
		ok := convertFBXtoGLTF(w, fname)
		if !ok {
			return
		}
	}
	log.Println("convert to gltf successful")

	ok := convertToUSDZ(w, fname)
	if !ok {
		return
	}

	log.Println("convert to usdz successful")

	go expireFiles([]string{
		fmt.Sprintf("%s.%s", fname, GLTF),
		fmt.Sprintf("%s.%s", fname, USDZ),
	})
}

func convertOBJtoGLTF(w http.ResponseWriter, fname string, t message) bool {
	var commandArgs []string
	if !strings.HasSuffix(fname, ".obj") {
		fname = t.FileNames[1]
	}
	log.Println("converting obj file")
	commandArgs = []string{"-i", fname, "-o", "./models/" + strings.TrimSuffix(fname, ".obj") + ".gltf"}
	_, err := exec.Command(OBJ_TO_GLTF, commandArgs...).Output()
	if err != nil {
		log.Println(err)
		_, _ = fmt.Fprintln(w, "failed to convert to gltf")
		return false
	}
	fname = strings.TrimSuffix(fname, ".obj")
	return true
}

func convertFBXtoGLTF(w http.ResponseWriter, fname string) bool {
	var commandArgs []string
	var msg []byte
	log.Println("converting file format fbx")
	commandArgs = []string{
		"--embed",
		"-i",
		fname,
		"-o",
		fmt.Sprintf("./models/%s", changeFileNameExtension(fname, GLTF)),
	}
	msg, err := exec.Command(FBX_TO_GLTF, commandArgs...).Output()
	if err != nil {
		log.Println(string(msg))
		_, _ = fmt.Fprintln(w, "failed to convert to gltf")
		return false
	}
	return true
}

func convertToUSDZ(w http.ResponseWriter, fname string) bool {
	var commandArgs []string
	commandArgs = []string{
		fmt.Sprintf("./models/%s.%s", fname, GLTF),
		fmt.Sprintf("./models/%s.%s", fname, USDZ),
	}
	_, err := exec.Command(GLTF_TO_USDZ, commandArgs...).Output()
	if err != nil {
		log.Println(err)
		_, _ = fmt.Fprint(w, "failed to convert to usdz")
		_, _ = fmt.Fprint(w, fname)
		return false
	}

	_, _ = fmt.Fprintln(w, fname)
	return true
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	certDomain, ok := os.LookupEnv(CERT_DOMAIN)
	// TODO: handle http also
	if !ok && certDomain != "" {
		log.Panicf("env variable %s not set", CERT_DOMAIN)
	}

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(certDomain), //Your domain here
		Cache:      autocert.DirCache("certs"),         //Folder for storing certificates
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", usdzHandler)
	mux.HandleFunc("/headers", headers)
	mux.Handle(
		"/models/",
		http.StripPrefix(
			strings.TrimRight("/models/", "/"),
			http.FileServer(http.Dir("models")),
		),
	)
	mainMux := newMiddleware(mux)
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
		Handler: mainMux,
	}
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
