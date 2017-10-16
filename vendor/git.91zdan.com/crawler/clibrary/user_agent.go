package clibrary

import "net/http"

func userAgentNexus6p(r *http.Request) {
	if r.Header.Get("User-Agent") == "" {
		r.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Mobile Safari/537.36")
	}
}
