package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/stamm/dep_radar/src"
	"github.com/stamm/dep_radar/src/helpers"
)

func main() {
	recomendedFile := flag.String("recomended_file", "", "path to file with recomended versions of libraries")
	githubToken := flag.String("github_token", "", "token for github: maybe empty, but limit for requested applies")
	githubOrg := flag.String("github_org", "", "name for github organization")
	listen := flag.String("listen", ":8081", "on which ip and port http server will be listen to")
	flag.Parse()
	if *listen == "" {
		fmt.Println("parameter listen is empty")
		os.Exit(1)
	}
	if *githubOrg == "" {
		fmt.Println("parameter github_org is empty")
		os.Exit(1)
	}

	var recom src.MapRecomended
	if *recomendedFile != "" {
		raw, err := ioutil.ReadFile(*recomendedFile)
		if err != nil {
			fmt.Printf("error on read file %s: %s\n", *recomendedFile, err.Error())
			os.Exit(1)
		}
		err = json.Unmarshal(raw, &recom)
		if err != nil {
			fmt.Printf("error on unmarshal json: %s\n", err.Error())
			os.Exit(1)
		}
	}
	helpers.GithubOrg(*githubToken, *githubOrg, *listen, recom)
}
