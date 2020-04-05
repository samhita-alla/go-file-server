# Download and Upload files using Go

* For creating a **file server** using Go (fileserver.go):

Use the command,

```
go build fileserver.go
```
Then run the file as `./fileserver`

Hit `localhost:8000` in your browser.

There it is! All the files present in the directory from where you would trigger this script would be displayed.

* For **uploading files** using Go (upload.go):

Use the command,

```
go build upload.go
```
Then run the file as `./upload`

Open the upload.html file in your browser. Upload a file of your choice. The uploaded files are then stored in temp-folder/ directory (present in the repo from where you would trigger this script). 

Inspired by - [file-upload](https://tutorialedge.net/golang/go-file-upload-tutorial/)
