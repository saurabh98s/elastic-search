package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)


	type ElasticResponse struct {
		Name        string `json:"name"`
		ClusterName string `json:"cluster_name"`
		ClusterUUID string `json:"cluster_uuid"`
		Version     struct {
			Number                           string    `json:"number"`
			BuildFlavor                      string    `json:"build_flavor"`
			BuildType                        string    `json:"build_type"`
			BuildHash                        string    `json:"build_hash"`
			BuildDate                        time.Time `json:"build_date"`
			BuildSnapshot                    bool      `json:"build_snapshot"`
			LuceneVersion                    string    `json:"lucene_version"`
			MinimumWireCompatibilityVersion  string    `json:"minimum_wire_compatibility_version"`
			MinimumIndexCompatibilityVersion string    `json:"minimum_index_compatibility_version"`
		} `json:"version"`
		Tagline string `json:"tagline"`
	}


func main() {
	resp,err:=http.Get("http://localhost:9200/")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	
	
	var result ElasticResponse
	err=json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Println(result)

	fmt.Println(result)

	

}