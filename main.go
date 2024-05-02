package main

import (
	"fmt"
	"os"
	"strings"

	urlverifier "github.com/davidmytton/url-verifier"
	"github.com/gosuri/uitable"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: urlVerify <url> [<url>...]")
		os.Exit(1)
	}

	fmt.Printf("verifying the url(s)..")
	outTable := uitable.New()
	outTable.AddRow("URL", "Status")

	for _, url := range os.Args[1:] {
		fmt.Printf(".")
		verifier := urlverifier.NewVerifier()
		status := "none"
		verifier.EnableHTTPCheck()
		verifier.AllowHTTPCheckInternal()

		ret, err := verifier.Verify(url)
		if err != nil && !strings.Contains(err.Error(), "failed to verify certificate") {
			status = "ERR: " + err.Error()
		} else if err != nil && strings.Contains(err.Error(), "failed to verify certificate") {
			status = "OK - failed to verify certificate"
		} else if !ret.IsURL {
			status = "ERR: not valid"
		} else if ret.HTTP == nil {
			status = "ERR: not reachable"
		} else if ret.HTTP.IsSuccess {
			status = "OK - status code " + fmt.Sprintf("%d", ret.HTTP.StatusCode)
		} else if ret.HTTP.Reachable {
			status = "OK - but status code " + fmt.Sprintf("%d", ret.HTTP.StatusCode)
		}

		outTable.AddRow(url, status)
	}
	fmt.Printf("\n")
	fmt.Println(outTable)
}
