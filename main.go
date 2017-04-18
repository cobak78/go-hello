package main
 
import(
	"fmt"
	"net/http"
	"runtime"
)
 
func indexHandler( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS,runtime.GOARCH)
}

func pingHandler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
 
func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8083",nil)
}