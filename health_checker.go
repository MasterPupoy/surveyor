package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

// func main() {
// 	out, err := exec.Command("curl", "--unix-socket", "/var/run/docker.sock", "http:/v1.47/containers/json").CombinedOutput()

// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	fmt.Printf("%s\n", out)
// }

func main() {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	res, err := client.Get("http://v1.47/containers/json")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	bodyString := string(body)
	fmt.Printf("%s\n", bodyString)

}
