package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
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

var (
	dbgMu    sync.Mutex
	dbgLines []string
)

// dbg records a short diagnostic line (no secrets: paths and status only).
func dbg(format string, a ...any) {
	dbgMu.Lock()
	defer dbgMu.Unlock()
	dbgLines = append(dbgLines, time.Now().Format("15:04:05 ")+fmt.Sprintf(format, a...))
	if len(dbgLines) > 300 {
		dbgLines = dbgLines[len(dbgLines)-300:]
	}
}

func dbgDump() string {
	dbgMu.Lock()
	defer dbgMu.Unlock()
	return strings.Join(dbgLines, "\n") + "\n"
}
