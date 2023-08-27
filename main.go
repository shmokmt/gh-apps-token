package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bradleyfalzon/ghinstallation/v2"
)

func main() {
	var appID int64
	var installationID int64

	flag.Int64Var(&appID, "app-id", 0, "App ID")
	flag.Int64Var(&installationID, "installation-id", 0, "Installation ID")

	flag.Parse()

	if appID == 0 {
		fmt.Fprintln(os.Stderr, "Error: app-id is required")
		os.Exit(1)
	}

	if installationID == 0 {
		fmt.Fprintln(os.Stderr, "Error: installation-id is required")
		os.Exit(1)
	}

	privateKey := os.Getenv("GITHUB_APPS_PRIVATE_KEY")

	if privateKey == "" {
		fmt.Fprintln(os.Stderr, "Error: GITHUB_APPS_PRIVATE_KEY is required")
		os.Exit(1)
	}

	tr := http.DefaultTransport

	itr, err := ghinstallation.New(tr, appID, installationID, []byte(privateKey))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	token, err := itr.Token(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(token)
}
