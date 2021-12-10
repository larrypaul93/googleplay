package main

import (
	"time"

	gp "github.com/89z/googleplay"
)

func delivery(app string, ver int) (*gp.Delivery, error) {
	auth, err := getAuth()
	if err != nil {
		return nil, err
	}
	dev := new(gp.Device)
	dev.Android_ID = AndroidId
	// src, err := os.Open(cache + "/googleplay/device.json")
	// if err != nil {
	// 	return nil, err
	// }
	// defer src.Close()
	// if err := dev.Decode(src); err != nil {
	// 	return nil, err
	// }
	del, err := auth.Delivery(dev, app, ver)
	if err != nil {
		return nil, err
	}
	return del, nil
	// if err := download(del.DownloadURL, "", app, ver); err != nil {
	// 	return nil,err
	// }
	// for _, split := range del.SplitDeliveryData {
	// 	err := download(split.DownloadURL, split.ID, app, ver)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return nil
}

func details(app string) (*gp.Details, error) {
	auth, err := getAuth()
	if err != nil {
		return nil, err
	}
	dev := new(gp.Device)
	dev.Android_ID = AndroidId
	// src, err := os.Open(cache + "/googleplay/device.json")
	// if err != nil {
	// 	return nil, err
	// }
	// defer src.Close()
	// if err := dev.Decode(src); err != nil {
	// 	return nil, err
	// }
	return auth.Details(dev, app)
}

func device() (*gp.Device, error) {
	auth, err := getAuth()
	if err != nil {
		return nil, err
	}
	dev, err := gp.NewDevice()
	if err != nil {
		return nil, err
	}
	if err := auth.Upload(dev, gp.DefaultConfig); err != nil {
		return nil, err
	}
	// fmt.Printf("Sleeping %v for server to process\n", gp.Sleep)
	time.Sleep(gp.Sleep)
	// cache = filepath.Join(cache, "/googleplay/device.json")
	// write, err := os.Create(cache)
	// if err != nil {
	// 	return "", err
	// }
	// defer write.Close()
	// if err := dev.Encode(write); err != nil {
	// 	return "", err
	// }
	return dev, nil
}

func getAuth() (*gp.Auth, error) {
	tok := new(gp.Token)
	tok.Token = Token
	// var cache , err := os.UserCacheDir()
	// cache := TOKEN_DIR
	// if err != nil {
	// 	return nil, "", err
	// }
	// src, err := os.Open(cache + "/googleplay/token.json")
	// if err != nil {
	// 	return nil, "", err
	// }
	// defer src.Close()
	// if err := tok.Decode(src); err != nil {
	// 	return nil, "", err
	// }
	auth, err := tok.Auth()
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func purchase(app string) error {
	auth, err := getAuth()
	if err != nil {
		return err
	}
	dev := new(gp.Device)
	dev.Android_ID = AndroidId
	// src, err := os.Open(cache + "/googleplay/device.json")
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()
	// if err := dev.Decode(src); err != nil {
	// 	return err
	// }
	return auth.Purchase(dev, app)
}

func token(email, password string) (*gp.Token, error) {
	tok, err := gp.NewToken(email, password)
	return tok, err
	// if err != nil {
	// 	return "", err
	// }
	// cache := TOKEN_DIR
	// // cache, err := os.UserCacheDir()
	// // cache = TOKEN_DIR
	// if err != nil {
	// 	return "", err
	// }
	// cache = filepath.Join(cache, "googleplay")
	// os.Mkdir(cache, os.ModePerm)
	// cache = filepath.Join(cache, "token.json")
	// file, err := os.Create(cache)
	// if err != nil {
	// 	return "", err
	// }
	// defer file.Close()
	// if err := tok.Encode(file); err != nil {
	// 	return "", err
	// }
	// return cache, nil
}
