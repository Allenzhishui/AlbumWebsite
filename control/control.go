package control

import (
	"AlbumWebsite+sql/model"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//首页
func IndexView(w http.ResponseWriter, r *http.Request) {
	html := LoadHtml("./views/index.html")
	w.Write(html)
}

//上传页面
func UploadView(w http.ResponseWriter, r *http.Request) {
	html := LoadHtml("./views/upload.html")
	w.Write(html)
}

//图片上传
func ApiUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, h, err := r.FormFile("file")
	defer f.Close()
	if err != nil {
		io.WriteString(w, "上传失败")
		return
	}
	t := h.Header.Get("Content-Type")
	if !strings.Contains(t, "image") {
		io.WriteString(w, "文件类型错误")
		return
	}
	os.Mkdir("./static", 0666)
	now := time.Now()
	name := now.Format("2006-01-02150405") + path.Ext(h.Filename)
	outputfile, err := os.Create("./static/" + name)
	defer outputfile.Close()
	if err != nil {
		io.WriteString(w, "文件创建错误")
		return
	}
	io.Copy(outputfile, f)
	mod := model.Info{
		Name: h.Filename,
		Path: "/static/" + name,
		Note: r.Form.Get("note"),
		Unix: now.Unix(),
	}
	model.InfoAdd(&mod)
	http.Redirect(w, r, "/list", 302)
}

func DetailView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.FormValue("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod, _ := model.InfoGet(id)
	html := LoadHtml("./views/detail.html")
	data := time.Unix(mod.Unix, 0).Format("2006-01-02 15:04:05")
	html = bytes.Replace(html, []byte("@src"), []byte(mod.Path), 1)
	html = bytes.Replace(html, []byte("@note"), []byte(mod.Note), 1)
	html = bytes.Replace(html, []byte("@unix"), []byte(data), 1)
	w.Write(html)
}

func ApiList(w http.ResponseWriter, r *http.Request) {
	mods, _ := model.InfoList()
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func ListView(w http.ResponseWriter, r *http.Request) {
	html := LoadHtml("./views/list.html")
	w.Write(html)
}

func ApiDel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := model.InfoDel(id)
	if err != nil {
		io.WriteString(w, "删除失败")
		return
	}
	io.WriteString(w, "删除成功")
	return

}
func LoadHtml(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte("")
	}
	return buf
}
