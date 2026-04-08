package mock

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// StartMockServer starts a dummy API server that has BOLA vulnerabilities
func StartMockServer(port string) {
	http.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		// Mock logic: 
		// If path is /api/v1/users/8888 (Victim) and token is pentester_token_xyz (Attacker)
		token := r.Header.Get("Authorization")
		
		if strings.HasSuffix(r.URL.Path, "8888/profile") {
			if token == "Bearer pentester_token_xyz" {
				// BOLA VULNERABILITY! Attacker accessing Victim's profile
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Victim's Secret Profile Data"))
				return
			}
		} else if strings.HasSuffix(r.URL.Path, "8888/email") {
			if token == "Bearer pentester_token_xyz" {
				// PROTECTED: Email cannot be changed by attacker (Secure)
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403 Forbidden"))
				return
			}
		}
		
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/api/v1/posts/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if strings.HasSuffix(r.URL.Path, "8888") {
			if token == "Bearer pentester_token_xyz" && r.Method == "DELETE" {
				// BOLA VULNERABILITY! Attacker can delete Victim's post
				w.WriteHeader(http.StatusOK) 
				w.Write([]byte("Post Deleted"))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/api/v1/update-profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		body, _ := io.ReadAll(r.Body)
		// Check if the body contains victim's ID (8888) 
		// but the token belongs to Attacker (pentester_token_xyz)
		token := r.Header.Get("Authorization")
		if strings.Contains(string(body), "8888") && token == "Bearer pentester_token_xyz" {
			// DEEP BOLA VULNERABILITY FOUND IN BODY!
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status": "success", "message": "Profile updated for user 8888"}`))
			return
		}
		
		w.WriteHeader(http.StatusUnauthorized)
	})

	fmt.Printf("[*] Internal Mock API Server running on port %s for testing...\n", port)
	go http.ListenAndServe(port, nil)
}
