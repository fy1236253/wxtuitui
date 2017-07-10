package http

import (
	"encoding/xml"
	"g"
	"log"
	"net/http"
	"path/filepath"
)

// Start 路由相关的启动
func Start() {
	// 静态资源请求
	ConfigWebHTTP()
	ConfigAPIRoutes()
	ConfigWechatRoutes()
	Config3rdWechatRoutes()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(filepath.Join(g.Root, "/public"))).ServeHTTP(w, r)
	})
	// start http server
	addr := g.Config().HTTP.Listen

	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("http.Start ok, listening on", addr)
	log.Fatalln(s.ListenAndServe())
}

//RenderText200 只返回200和描述
func RenderText200(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(200)
	w.Write([]byte(s))
}

//RenderXML 只返回200和描述
func RenderXML(w http.ResponseWriter, v interface{}) {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
	w.Write(bs)
}

//RenderText 只返回描述
func RenderText(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.Write([]byte(s))
}
