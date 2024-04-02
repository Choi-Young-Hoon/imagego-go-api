package httpserver

import (
	"context"
	"imagego-go-api/httpserver/handler"
	"imagego-go-api/util"
	"net/http"
)

type HttpServer struct {
	Port string

	httpServer http.Server
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (hs *HttpServer) HttpStart() {
	hs.setServerInfo()
	hs.setHandler()

	err := hs.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (hs *HttpServer) HttpsStart(certFile, privKeyFile string) {
	hs.setHandler()

	/*
		err := http.ListenAndServeTLS(certFile, privKeyFile)
		if err != nil {
			panic(err)
		}
	*/
}

func (hs *HttpServer) HttpStop() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	if err := hs.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func (hs *HttpServer) setServerInfo() {
	hs.httpServer = http.Server{
		Addr: ":" + hs.Port,
	}
}

func (hs *HttpServer) setHandler() {
	// hs의 httpServer에 ServerMux에 핸들러 지정해서 등록
	httpMux := http.NewServeMux()

	httpMux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(util.GetServerConfig().ImageDir))))

	httpMux.HandleFunc("/login", handler.LoginHandler)
	httpMux.HandleFunc("/register", handler.RegisterHandler)
	httpMux.HandleFunc("/echo", handler.EchoHandler)
	httpMux.HandleFunc("/upload", handler.UploadHandler)
	httpMux.HandleFunc("/image/all", handler.ImageAllHandler)
	httpMux.HandleFunc("/image/{number}", handler.ImageHandler)

	hs.httpServer.Handler = httpMux
}
