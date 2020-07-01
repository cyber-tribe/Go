package main

import (
	"io/ioutil"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// 画面を描画する関数
	title := r.URL.Path[6:]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} // 空のPage構造体を作成
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

type Page struct {
	// wikiのページを保存する構造体
	Title string
	Body  []byte
}

func (p *Page) save() error {
	// ページの内容をテキストファイルに保存するPage型のメソッド
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	// ページをテキストファイルから読み込む関数
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) // []byteあるいはerrorを返す
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
