package mosaicassembly

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "io"
    "log"
)

type RouteInfo struct {
    Name        string
    Method      string
    Pattern     string
}

type ImagePosition struct {
    Pos string `json:"pos"`
    Image string `json:"image"`
}

type RouteInfos []RouteInfo
type ImagePositions []ImagePosition

var routeInfos = RouteInfos{
    RouteInfo{
        "Index",
        "GET",
        "/",
    },
    RouteInfo{
        "ProcessMosaicImage",
        "POST",
        "/mosaic",        
    },
    RouteInfo{
        "DownloadMosaicImage",
        "GET",
        "/mosaic/{imageId}",
    },
}

func Index(w http.ResponseWriter, r *http.Request) {
       if err := json.NewEncoder(w).Encode(routeInfos); err != nil {
        panic(err)
    }
}

func ProcessMosaicImage(w http.ResponseWriter, r *http.Request) {   
    dec := json.NewDecoder(r.Body)
        var i ImagePositions
        err := dec.Decode(&i)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        // fmt.Printf("%s: %s\n", i.Pos, i.Image)
    fmt.Printf("Received request: %v", i)
    assembleMosaic(i)

    fmt.Fprintln(w, json.NewEncoder(w).Encode(i))
}

func DownloadMosaicImage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}