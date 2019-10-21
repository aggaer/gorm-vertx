package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm/src/model"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/poster", getPoster)
	_ = http.ListenAndServe(":10002", nil)
}

//noinspection SpellCheckingInspection
func getPoster(w http.ResponseWriter, req *http.Request) {
	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", "root:aggaer520@tcp(127.0.0.1:3306)/pintuan?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()
	key, _ := strconv.ParseInt(values["sharingKey"][0], 10, 64)
	resp := &model.Record{}
	db.Where(&model.Record{SharingKey: key}).First(resp)
	_ = json.NewEncoder(w).Encode(resp)
}
