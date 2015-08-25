package mosaicassembly

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "ProcessMosaicImage",
        "POST",
        "/mosaic",
        ProcessMosaicImage,
    },
      Route{
        "ProcessMosaicImage",
        "OPTIONS",
        "/mosaic",
        ProcessMosaicImage,
    },
    Route{
        "DownloadMosaicImage",
        "GET",
        "/mosaic/{imageId}",
        DownloadMosaicImage,
    },
}
