package main
import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	tingyun "github.com/TingYunAPM/go"
)

const (
	fcLogTailEndPrefix = "FC Invoke End RequestId: %s" // End of log tail mark
)

func invoke(w http.ResponseWriter, req *http.Request) {
	action, _ := tingyun.CreateAction("URI", req.URL.Path)
	defer action.Finish()
	requestID := req.Header.Get("x-fc-request-id")
	fmt.Println(fmt.Sprintf("FC Invoke Start RequestId: %s", requestID))
	defer func() {
		fmt.Println(fmt.Sprintf(fcLogTailEndPrefix, requestID))

	}()

	action, err := tingyun.CreateAction("URI", req.URL.Path)
	if err != nil {
		panic(err)
	}

	headerComponent := action.CreateComponent("header")
	n := rand.Intn(100) // n will be between 0 and 10
	fmt.Printf("Sleeping %d ms...\n", 400+n)
	time.Sleep(time.Duration(200+n) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	action.SetStatusCode(uint16(http.StatusOK))
	headerComponent.Finish()
	bodyComponent := action.CreateComponent("body")

	// your logic
	n = rand.Intn(100) // n will be between 0 and 10
	fmt.Printf("Sleeping %d ms...\n", 900+n)
	time.Sleep(time.Duration(900+n) * time.Millisecond)
	bodyComponent.Finish()
	action.Finish()
	time.Sleep(62 * time.Second)
	w.Write([]byte(fmt.Sprintf("Hello, golang  http invoke!")))
}

func main() {
	fmt.Println("FunctionCompute go runtime inited.")
	tingyun.AppInit("tingyun.json")
	defer tingyun.AppStop()

	http.HandleFunc("/2016-08-15/proxy/test-apm/tingyun-golang/invoke", invoke) // 如果不使用自定义域名，则 path 为 /2016-08-15/proxy/$serviceName/$functionName/a
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println(port)
	http.ListenAndServe(":"+port, nil)
}
