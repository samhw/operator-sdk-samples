package memcached

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main2() {
	MyFunc()
}

func MyFunc() string {

	 // ignore expired SSL certificates
     transCfg := &http.Transport{
             TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
     }

	url := "https://172.25.20.27:8444/v8/storage/logical-units/532"
	client := &http.Client{Transport: transCfg}
	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Set("X-Subsystem-User", "ms_vmware")
	req.Header.Set("X-Subsystem-Password", "Hitachi1")
	req.Header.Set("X-Subsystem-Type", "BLOCK")
	req.Header.Set("X-Management-IPs", "172.25.43.250")
	req.Header.Set("X-Storage-Id", "415056")
	
	response, err := client.Do(req)
    if err != nil {
        //fmt.Printf("The HTTP request failed with error %s\n", err)
        return fmt.Sprintf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        //fmt.Println(string(data))
        return fmt.Sprintf(string(data))
    }
}

