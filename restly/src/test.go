package main 

import (
	"fmt"
	"net/http"
	"html/template"
	"os/exec"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
//	fmt.Fprintf(w, "Hi Kranthi!!!")
}

func uuid() string{
	out, err := exec.Command("uuidgen").Output()
	if err != nil{
		panic(err)
	}
	return string(out)
}

func readConfigFile(){
	f,err := os.Open("config.json")
	if err != nil{
		panic(err)
	}
	
	fmt.Println(f)
}

func uuid_urandom() string{
	b := make([]byte, 16)
	f,err := os.Open("/dev/urandom")
	if err != nil{
		panic(err)
	}
	f.Read(b)
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
//	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("bootstrap"))))
//	http.HandleFunc("/", handler)
//    http.ListenAndServe(":8080", nil)
//	fmt.Println(uuid())
//	fmt.Println(uuid_urandom())
	readConfigFile()
}

