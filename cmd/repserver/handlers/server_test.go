package handlers

import (
	"encoding/hex"
	"net/http"
	log "github.com/repbin/repbin/deferconsole"
	"testing"
	"time"
)

func TestGenPostHandler(t *testing.T) {
	pubKey, _ := hex.DecodeString("39d8913ab046428e409cf1fa7cee6f63c1f6bf701356a44a8c8c2559bdb2526f")
	privKey, _ := hex.DecodeString("20a2633e422090a4f4a102f8e3d112f2b4378dbd9957e8c892067fc09239d36c39d8913ab046428e409cf1fa7cee6f63c1f6bf701356a44a8c8c2559bdb2526f")

	log.SetMinLevel(log.LevelDebug)
	ms, err := New("/tmp/repbin/", pubKey, privKey)
	if err != nil {
		t.Fatalf("New: %s", err)
	}
	enforceTimeOuts = false
	debug = true
	ms.NotifyDuration = 0
	ms.FetchDuration = 0
	ms.LoadPeers()
	ms.NotifyPeers()
	ms.FetchPeers()
	http.HandleFunc("/id", ms.ServeID)
	http.HandleFunc("/keyindex", ms.GetKeyIndex)
	http.HandleFunc("/globalindex", ms.GetGlobalIndex)
	http.HandleFunc("/post", ms.GenPostHandler(false))
	http.HandleFunc("/local/post", ms.GenPostHandler(true))
	http.HandleFunc("/fetch", ms.Fetch)
	http.HandleFunc("/notify", ms.GetNotify)
	http.HandleFunc("/delete", ms.Delete)
	go http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second / 100)
	time.Sleep(time.Second * 10)
}
