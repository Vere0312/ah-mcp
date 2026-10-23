package main

import (
	"net/http"
	"sync"
)

// sharedOAuth holds the temporary OAuth proxy handler while a login flow is
// active and no separate callback port is used (AH_CALLBACK_PORT=0).
var (
	sharedMu    sync.RWMutex
	sharedOAuth http.Handler
)

func setSharedOAuth(h http.Handler) {
	sharedMu.Lock()
	sharedOAuth = h
	sharedMu.Unlock()
}

func getSharedOAuth() http.Handler {
	sharedMu.RLock()
	defer sharedMu.RUnlock()
	return sharedOAuth
}
