package raw

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/subutai-io/gorjun/db"
	"github.com/subutai-io/gorjun/download"
	"github.com/subutai-io/gorjun/upload"
)

type RawItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Md5Sum      string `json:"md5Sum"`
	Version     string `json:"version"`
	Fingerprint string `json:"fingerprint"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		hash, owner := upload.Handler(w, r)
		_, header, _ := r.FormFile("file")
		db.Write(owner, hash, header.Filename, map[string]string{"type": "raw"})
		w.Write([]byte(hash))
		w.WriteHeader(http.StatusOK)
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	//raw-files download handler will be here
	download.Handler(w, r)
}

func Show(w http.ResponseWriter, r *http.Request) {
	//raw-files list handler will be here
	download.List("raw", w, r)
}

func List(w http.ResponseWriter, r *http.Request) {
	list := []RawItem{}
	for hash, _ := range db.List() {
		if info := db.Info(hash); info["type"] == "raw" {
			item := RawItem{
				ID:          "raw." + hash,
				Name:        info["name"],
				Fingerprint: info["owner"],
				Md5Sum:      hash,
			}
			item.Size, _ = strconv.ParseInt(info["size"], 10, 64)
			list = append(list, item)
		}
	}
	js, _ := json.Marshal(list)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(js)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" && len(upload.Delete(w, r)) != 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Removed"))
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
	return
}

func Info(w http.ResponseWriter, r *http.Request) {
	info := download.Info("raw", r)
	if len(info) == 0 {
		w.Write([]byte("Not found"))
	}
	w.Write(info)
}
