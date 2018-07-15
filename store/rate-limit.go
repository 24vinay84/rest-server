//Rate Limiting HTTP Requests (via http.HandlerFunc middleware)

package store

import (
    "log"
    "net/http"
    "strings"
    "time"
)

func limitRequest() {
   
    log.Println("********** Listening for Rate Limiting HTTP GET and POST Requests **********")
    go clearLastRequestsIPs()
    go clearBlockedIPs()
}

// Method allowed for Limit Rate access
var methods = [2]string{"GET", "POST"}
var limit = 5

// Stores last requests IPs
var lastRequestsIPs []string

// Block IP for 6 hours
var blockedIPs []string

func middleware(next http.HandlerFunc) http.HandlerFunc {
    
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		log.Println(" Rate limit method : " + r.Method)
		
		if allowedMethod(r.Method) == false {
			log.Println(" Rate limit ignore method : " + r.Method)
			return;
		}
		ipAddr := strings.Split(r.RemoteAddr, ":")[0]
		//log.Println(" Rate limit method : IP Address " + r.RemoteAddr)
        if existsBlockedIP(ipAddr) {
            http.Error(w, "", http.StatusTooManyRequests)
            return
        }
        // how many requests the current IP made in last 5 mins
        requestCounter := 0
        for _, ip := range lastRequestsIPs {
            if ip == ipAddr {
                requestCounter++
            }
        }
        if requestCounter >= limit {
			log.Println(" Rate limit method : Added to Block IP Address : " + ipAddr)
            blockedIPs = append(blockedIPs, ipAddr)
            http.Error(w, "", http.StatusTooManyRequests)
            return
        }
        lastRequestsIPs = append(lastRequestsIPs, ipAddr)

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

func existsBlockedIP(ipAddr string) bool {
    for _, ip := range blockedIPs {
        if ip == ipAddr {
            return true
        }
    }
    return false
}

func existsLastRequest(ipAddr string) bool {
    for _, ip := range lastRequestsIPs {
        if ip == ipAddr {
            return true
        }
    }
    return false
}

// Clears lastRequestsIPs array every 5 mins
func clearLastRequestsIPs() {
    for {
        lastRequestsIPs = []string{}
        time.Sleep(time.Minute * 5)
    }
}

// Clears blockedIPs array every 6 hours
func clearBlockedIPs() {
    for {
        blockedIPs = []string{}
        time.Sleep(time.Hour * 6)
    }
}