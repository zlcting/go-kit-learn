package main

import (
	"fmt"
	"net/http"
	. "zlcting/gokit/demo/Services"

	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
)

func main() {
	user := UserService{}
	endp := GetUserEndpoint(user)
	serverHanlder := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)
	fmt.Println(EncodeUserResponse)

	r := mymux.NewRouter()
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHanlder)
		r.Methods("GET").Path("/health").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type", "application/json")
			rw.Write([]byte(`{"status":"ok"}`))
		})
	}

	//r.Handle(`/user/{uid:\d+}`, serverHanlder)

	res := http.ListenAndServe(":8080", r)
	fmt.Println(res)

}
