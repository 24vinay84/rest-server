//Rate Limiting HTTP Requests (via http.HandlerFunc middleware)

package store

import (
    "log"
    "net/http"
    "strings"
)

func limitRequest() {
   
    log.Println("********** Listening for Rate Limiting HTTP GET and POST Requests **********")
}

// Method allowed for Limit Rate access
var methods = [2]string{"GET", "POST"}

var limit = int64(100) // Limitation count for 5 minutes

func middleware(next http.HandlerFunc) http.HandlerFunc {
    
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		log.Println(" Rate limit method : " + r.Method)
		
		if allowedMethod(r.Method) == false {
			log.Println(" Rate limit ignore method : " + r.Method)
			return;
		}
		ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	    
		if isBlockIP(ipAddr) {
            http.Error(w, "", http.StatusTooManyRequests)
            return
        }
        // how many requests the current IP made in last 5 mins
        requestCounter := getCount(ipAddr)
        if requestCounter >= limit {
			log.Println(" Rate limit method : Added to Block IP Address : " + ipAddr)
			
			blockIP(ipAddr)  // Block this IP
            
			http.Error(w, "", http.StatusTooManyRequests)
            return
        }

		incementCount(ipAddr) // Increment number of access by this IP
		
        // Don't cut the chain of middlewares
        if next == nil {
            http.DefaultServeMux.ServeHTTP(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func allowedMethod(method string) bool {
    for _, meth := range methods {
        if meth == method {
            return true
        }
    }
    return false
}
