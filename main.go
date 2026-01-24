package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stock int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomi", Harga: 3500, Stock: 10},
	{ID: 2, Nama: "sarimi", Harga: 3500, Stock: 10},
}

type Category struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi`
}

var category = []Category{
	{ID: 1, Nama: "Sembako", Deskripsi: "Sembilan Bahan Pokok"},
	{ID: 2, Nama: "Semen", Deskripsi: "Alat tukang"},
}

func getProdukbyID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk Belum ada", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti jad int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id dan ganti dari request body
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)

}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk", http.StatusBadRequest)
		return
	}

	// loop produk cari id dan index untuk dihapus
	for i, p := range produk {
		if p.ID == id {
			// buat slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses Delete",
			})
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)

}

func showKategori(w http.ResponseWriter, r *http.Request) {
	// ambil id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti ID menjadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid kategori ID", http.StatusBadRequest)
		return
	}

	// get kategory by id request

	for _, p := range category {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}

	}

	http.Error(w, "Kategory tidak ada", http.StatusNotFound)

}

func updateKategori(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ubah id jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid kategory id", http.StatusBadRequest)
		return
	}

	// get data dari request body
	var updateKategory Category
	err = json.NewDecoder(r.Body).Decode(&updateKategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop untuk id dan update
	for i := range category {
		if category[i].ID == id {
			updateKategory.ID = id
			category[i] = updateKategory

			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategory)
			return
		}
	}

	http.Error(w, "Kategory tidak ada", http.StatusNotFound)

}

func hapusKategori(w http.ResponseWriter, r *http.Request) {
	// get id dari url path
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ubah id jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// loop id produk untuk dihapus
	for i, p := range category {
		if p.ID == id {
			// buat slice baru dengan data sebelum dan sesudah index
			category = append(category[:i], category[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses hapus kategori",
			})

			return
		}
	}
	http.Error(w, "Tidak ada kategori", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			showKategori(w, r)
		} else if r.Method == "PUT" {
			updateKategori(w, r)
		} else if r.Method == "DELETE" {
			hapusKategori(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)

		} else if r.Method == "POST" {
			// baca dari body post
			var kategoryBaru Category
			err := json.NewDecoder(r.Body).Decode(&kategoryBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			kategoryBaru.ID = len(category) + 1
			category = append(category, kategoryBaru)

			w.Header().Set("Content-Type", "application-json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(kategoryBaru)

		}
	})

	// Get api/produk/{id}
	// PUT api/produk/{id}
	// DELETE api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukbyID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}

	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == ("GET") {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			// baca dari request
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "invalid Request", http.StatusBadRequest)
			}
			// masukkan data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
