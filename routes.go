package main

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
        "IndexHandler",
        "GET",
        "/",
        IndexHandler,
    },
    Route{
        "FileHandler",
        "GET",
        "/files/{fileName}",
        FileHandler,
    },
    Route{
        "SaveHandler",
        "POST",
        "/save",
		SaveHandler,
    },
}