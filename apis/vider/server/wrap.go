package server

import (
	"bytes"
	"github.com/asaskevich/govalidator"
	"io/ioutil"
	"log"
	"net/http"
)

func UnSupportedMethod(w http.ResponseWriter, r *http.Request) {
	log.Printf("UnSupportedMethod ip:%s method:%s", r.RemoteAddr, r.Method)
	w.WriteHeader(404)
	_, err := w.Write([]byte("UnSupportedMethod"))
	if err != nil {
		log.Printf("w.Write err:%v", err)
	}
}
func UnSupportedContentType(w http.ResponseWriter, r *http.Request) {
	log.Printf("UnSupportedContentType ip:%s method:%s", r.RemoteAddr, r.Method)
	w.WriteHeader(404)
	_, err := w.Write([]byte("UnSupportedContentType"))
	if err != nil {
		log.Printf("w.Write err:%v", err)
	}
}
func UnSupportedJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("UnSupportedJson ip:%s method:%s", r.RemoteAddr, r.Method)
	w.WriteHeader(404)
	_, err := w.Write([]byte("UnSupportedJson"))
	if err != nil {
		log.Printf("w.Write err:%v", err)
	}
}
func UnSupportedJsonBody(w http.ResponseWriter, r *http.Request) {
	log.Printf("UnSupportedJsonBody ip:%s method:%s", r.RemoteAddr, r.Method)
	w.WriteHeader(404)
	_, err := w.Write([]byte("UnSupportedJsonBody"))
	if err != nil {
		log.Printf("w.Write err:%v", err)
	}
}

func JSONParse(jh JsonHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			UnSupportedContentType(w, r)
			return
		}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("ioutil.ReadAll Err:%v", err)
			UnSupportedJsonBody(w, r)
			return
		}
		if !govalidator.IsJSON(string(data)) {
			UnSupportedJson(w, r)
			return
		}
		jh(w, bytes.NewReader(data))
	}
}

func WrapRestful(
	post JsonHandler,
	delete http.HandlerFunc,
	put http.HandlerFunc,
	get http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request ip:%s method:%s", r.RemoteAddr, r.Method)
		switch r.Method {
		case "POST":
			JSONParse(post)(w, r)
		case "DELETE":
			delete(w, r)
		case "PUT":
			put(w, r)
		case "GET":
			get(w, r)
		default:
			UnSupportedMethod(w, r)
		}
	}
}
