package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/89z/googleplay"
)

var Token, AndroidId string

func download(addr, id, app string, ver int) error {
	fmt.Println("GET", addr)
	res, err := http.Get(addr)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var name string
	if id != "" {
		name = fmt.Sprintf("%v-%v-%v.apk", app, id, ver)
	} else {
		name = fmt.Sprintf("%v-%v.apk", app, ver)
	}
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.ReadFrom(res.Body); err != nil {
		return err
	}
	return nil
}

func printJson(i interface{}) {
	s, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func main() {
	var (
		app, email, pass, proxy string
		purch, verbose          bool
		version                 int
	)

	flag.StringVar(&app, "a", "", "app")
	flag.StringVar(&proxy, "proxy", "", "Proxy")
	// flag.BoolVar(&dev, "d", false, "create device")
	flag.StringVar(&email, "e", "", "email")
	flag.StringVar(&pass, "p", "", "password")
	flag.StringVar(&Token, "t", "", "Auth Token")
	flag.StringVar(&AndroidId, "d", "", "Device ID")
	flag.BoolVar(
		&purch, "purchase", false,
		"Purchase app. Only needs to be done once per Google account.",
	)
	flag.IntVar(&version, "v", 0, "version")
	flag.BoolVar(&verbose, "verbose", false, "dump requests")
	flag.Parse()
	googleplay.Verbose = verbose
	if proxy != "" {
		os.Setenv("HTTP_PROXY", proxy)
	}
	switch {
	case email != "":
		token, err := token(email, pass)
		if err != nil {
			panic(err)
		}
		printJson(token)
		// fmt.Println("Create", cache)
	// case dev:
	// 	cache, err := device()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Create", cache)
	case app != "" && !purch && version == 0 && Token != "" && AndroidId != "":
		res, err := details(app)
		if err != nil {
			panic(err)
		}
		printJson(res)
	case app != "" && purch && Token != "" && AndroidId != "":
		err := purchase(app)
		if err != nil {
			panic(err)
		}
		fmt.Println("{\"status\":\"OK\"}")
	case app != "" && version != 0 && Token != "" && AndroidId != "":
		del, err := delivery(app, version)
		if err != nil {
			panic(err)
		}
		printJson(del)

	default:
		fmt.Println("googleplay [flags]")
		flag.PrintDefaults()
	}
}
