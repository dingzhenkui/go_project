package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const BaseUploadPath = "./uploadfiles"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload2", upload)
	http.HandleFunc("/download2", handleDownload)
	err := http.ListenAndServe(":7373", nil)
	if err != nil {
		fmt.Println("服务器启动失败", err.Error())
		return
	}

}
func upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(32 << 20)
	//接收客户端传来的文件 uploadfile 与客户端保持一致
	file, handler, err := request.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//上传的文件保存在ppp路径下
	//ext := path.Ext(handler.Filename) //获取文件后缀
	//fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext
	fileNewName := handler.Filename

	f, err := os.OpenFile(BaseUploadPath+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintln(writer, "upload ok!"+fileNewName)
}
func handleDownload(w http.ResponseWriter, request *http.Request) {
	//文件上传只允许GET方法
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}
	//文件名
	filename := request.FormValue("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	log.Println("filename: " + filename)
	//打开文件
	file, err := os.Open(BaseUploadPath + "/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}
func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload2" method="post">
<input type="file" name="uploadfile">
<input type="hidden" name="token" value="{...{.}...}">
<input type="submit" value="upload">
</form>
<form enctype="multipart/form-data" action="/download2" method="get">
<input type="file" name="downloadfile">
<input type="hidden" name="token" value="{...{.}...}">
<input type="submit" value="download">
</form>
</body>
</html>
`
