package fileServer

import (
    "net/http"
    "fmt"
    "io"
    "os"
)


const (
    UPLOAD_DIR = "/filePATH"
)

func POSTuploadHandler(w http.ResponseWriter, r *http.Request) {
    f, h, err := r.FormFile("image")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    filename := h.Filename
    fmt.Println("filename:", filename)

    defer f.Close() //文件流句柄

    t, err := os.Create(UPLOAD_DIR + "/" + filename)
    fmt.Println("filepath:", UPLOAD_DIR+"/"+filename)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer t.Close() //关闭临时文件句柄

    if _, err := io.Copy(t, f); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
}


func GETuploadHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<html><form method=\"POST\" action=\"/uploadp\" "+
        " enctype=\"multipart/form-data\">"+
        "Choose an image to upload: <input name=\"image\" type=\"file\" />"+
        "<input type=\"submit\" value=\"Upload\" />"+
        "</form></html>")
    return
}