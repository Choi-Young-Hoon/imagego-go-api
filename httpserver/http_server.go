package httpserver

import (
	"context"
	"imagego-go-api/httpserver/handler/echo"
	"imagego-go-api/httpserver/handler/image"
	"imagego-go-api/httpserver/handler/login"
	"imagego-go-api/httpserver/handler/register"
	"imagego-go-api/httpserver/handler/upload"
	"imagego-go-api/httpserver/handler/upscale"
	"imagego-go-api/httpserver/jwt"
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

	httpMux.HandleFunc("/login", login.LoginHandler)
	httpMux.HandleFunc("/register", register.RegisterHandler)
	httpMux.HandleFunc("/echo", jwt.JwtVerifyMiddleware(echo.EchoHandler))
	httpMux.HandleFunc("/upload", jwt.JwtVerifyMiddleware(upload.UploadHandler))
	httpMux.HandleFunc("/image/all", jwt.JwtVerifyMiddleware(image.ImageAllHandler))
	httpMux.HandleFunc("/image/{number}", jwt.JwtVerifyMiddleware(image.ImageHandler))
	httpMux.HandleFunc("/upscale/{number}", jwt.JwtVerifyMiddleware(upscale.UpscaleHandler))

	hs.httpServer.Handler = httpMux
}
