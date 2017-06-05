package goroute

import (
	"net/http"
	"net/url"
)

//Param gets the named parameter from the request
func Param(r *http.Request, key string) string {
	return r.URL.Query().Get(":" + key)
}

func addParam(r *http.Request, key, value string) {
	q := url.QueryEscape(":"+key) + "=" + url.QueryEscape(value)
	if r.URL.RawQuery != "" {
		r.URL.RawQuery += "&" + q
	} else {
		r.URL.RawQuery += q
	}
}
