package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Label string `json:"label"`
	Nilai int64  `json:"nilai"`
}

type Response struct {
	Label  string  `json:"label"`
	Hasil  float64 `json:"hasil"`
	Satuan string  `json:"satuan"`
}

func konversiMBG(nilai int64) (float64, string) {
	hMBG := float64(nilai) / 1200000000000.0
	if hMBG/7 >= 1 {
		return hMBG / 7, "minggu"
	} else if hMBG >= 1 {
		return hMBG, "hari"
	} else if hMBG*24 >= 1 {
		return hMBG * 24, "jam"
	} else if hMBG*1440 >= 1 {
		return hMBG * 1440, "menit"
	}
	return hMBG * 86400, "detik"
}

func konversiHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Gunakan metode POST", http.StatusMethodNotAllowed)
		return
	}
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	hasil, satuan := konversiMBG(req.Nilai)

	response := Response{
		Label:  req.Label,
		Hasil:  hasil,
		Satuan: satuan,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/konversi", konversiHandler)
	println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// fufufafa kesukaan prabo
