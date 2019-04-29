package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../config"
	"../model"
	"github.com/tidwall/gjson"
)

//======== function to send success http status 200
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

//======== function to send error
func ErrorWithString(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	var resps model.ResultData
	resps.Status = true
	resps.Message = "success"
	resps.Data = "Assalamu'alaikum"
	tmpBody, errs := json.MarshalIndent(resps, "", "  ")
	if errs != nil {
		log.Fatal(errs)
	}
	ResponseWithJSON(w, tmpBody, http.StatusOK)
}

// ===== function to run jar using  apache livy using http post
func RunLivy(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	//==== read data from param
	bodyBytes, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		ErrorWithString(w, "parameter not found", http.StatusServiceUnavailable)
		log.Println("Failed get parametes: ", err2)
		return
	}
	jsonparam := string(bodyBytes)
	table := gjson.Get(jsonparam, "table")
	zooKeeper := gjson.Get(jsonparam, "zooKeeper")
	hbaseMaster := gjson.Get(jsonparam, "hbaseMaster")
	pathCSV := gjson.Get(jsonparam, "pathCSV")

	//==== store argument for apache livy
	listarg := []string{table.String(), zooKeeper.String(), hbaseMaster.String(), pathCSV.String()}

	message = map[string]interface{}{
		"file":         config.GetConfig("pathjar"),   //path of jar file
		"className":    config.GetConfig("classname"), // name of main class from our spark apps
		"args":         listarg,                       // list argumen to our spark apps
		"driverMemory": "5G",                          // total memory for our spark apps
		"driverCores":  4,                             // num of cores for our spark apps
		"name":         "massiveprofiling",            //name for our saprk apps
	}

	//=== marshal json param to byte
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	//=== post data json param  spark to apache livy
	resp, err := http.Post(config.GetConfig("url"), "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytesRepresentation))

	//=== convert response body to string
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	tempsid := result["id"].(float64)
	var id int = int(tempsid)
	tempData := strconv.Itoa(id)

	//====== return id result data
	var resps model.ResultData
	resps.Status = true
	resps.Message = "success"
	resps.Data = "Alhamdulillah success running insert data to hbase with id " + tempData
	tmpBody, errs := json.MarshalIndent(resps, "", "  ")
	if errs != nil {
		log.Fatal(errs)
	}
	ResponseWithJSON(w, tmpBody, http.StatusOK)
}
