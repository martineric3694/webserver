package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Student struct {
// 	NoInduk string
// 	Nama    string
// 	Kelas   string
// }

// func setStudent() (data []Student) {
// 	dataKelas := []Student{
// 		Student{
// 			NoInduk: "1234",
// 			Nama:    "Adi",
// 			Kelas:   "X",
// 		},
// 		Student{
// 			NoInduk: "3456",
// 			Nama:    "Budi",
// 			Kelas:   "XI",
// 		}}

// 	dataKelas = append(dataKelas, Student{
// 		NoInduk: "7890",
// 		Nama:    "Rudi",
// 		Kelas:   "X",
// 	})
// 	data = dataKelas
// 	return data
// }

// func student(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(setStudent())

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(result)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }
