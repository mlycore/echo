package main
import(
	"fmt"
	"encoding/json"
	"os"
	"net/http"
)

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	return hostname
}

func sourceIP(r *http.Request) string {
	return r.RemoteAddr 
}

type Response struct{
	Hostname string `json:"hostname"`
	SourceIP string `json:"sourceIP"`
} 

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "health\n")
	})
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request){
		h := hostname()
		srcip := sourceIP(r)
		resp := &Response{
			Hostname: h,
			SourceIP: srcip,
		}
		body, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error\n")
			panic(err)
		}else {
			fmt.Fprintf(w, "%v\n", string(body))
		}
	})
	http.ListenAndServe(":8080", nil)
}