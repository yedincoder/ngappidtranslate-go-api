package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TranslateRequest struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type MyMemoryResponse struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req TranslateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Format request salah"}`, http.StatusBadRequest)
		return
	}

	safeText := url.QueryEscape(req.Text)
	apiURL := fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=%s|%s", safeText, req.Source, req.Target)

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, `{"error": "Gagal konek ke API"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var myMemResp MyMemoryResponse
	if err := json.Unmarshal(body, &myMemResp); err != nil {
		http.Error(w, `{"error": "Gagal parsing data"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"original":   req.Text,
		"translated": myMemResp.ResponseData.TranslatedText,
	})
}