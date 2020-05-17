package fileServer

import (
    "net/http"
    "fmt"
    "io"
    "os"
)


const (
    //UPLOAD_DIR = "/app/filePATH"
    UPLOAD_DIR = "uploads"
)

func POSTuploadHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling /uploadp")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    f, h, err := r.FormFile("file")
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
    //http.Redirect(w, r, "/accounts/123", http.StatusFound)
}

func OPTIONSPOSTuploadHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Methods", "GET;POST")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // 由于前端的POST请求为json，所以后端这里特别注明一下
    //w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    //w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}


func GETuploadHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<html><form method=\"POST\" action=\"/uploadp\" "+
        " enctype=\"multipart/form-data\">"+
        "Choose an image to upload: <input name=\"image\" type=\"file\" />"+
        "<input type=\"submit\" value=\"Upload\" />"+
        "</form></html>")
    return
}