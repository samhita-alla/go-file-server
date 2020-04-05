package main

import (
	"log"
	"net/http"
	"fmt"
    "os"
	"io/ioutil"
    "path/filepath"
)

// upload logic
func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<h3>Uploaded File: %+v\n</h3>", handler.Filename)
    fmt.Fprintf(w, "<h3>File Size: %+v\n</h3>", handler.Size)
    fmt.Fprintf(w, "<h3>MIME Header: %+v\n</h3>", handler.Header)   // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    _, err_dir := os.Stat("temp-folder")
 
    if os.IsNotExist(err_dir) {
        errDir := os.MkdirAll("temp-folder", 0755)
        if errDir != nil {
            log.Fatal(err)
        }
 
    }

    tempFile, err := os.Create(filepath.Join("temp-folder", filepath.Base(handler.Filename)))

    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "<h3>Successfully Uploaded File\n</h3>")
   }

func main() {
	http.HandleFunc("/upload", uploadFile)
	log.Printf("Serving . on HTTP port: 3000\n")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
