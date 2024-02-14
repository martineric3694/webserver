package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Apa Kabar?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/students", studentsOther)

	http.HandleFunc("/student", student)

	http.HandleFunc("/student/post", postStudent)
	http.HandleFunc("/student/postJSON", postStudentJSON)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

type Student struct {
	NoInduk string `json:NoInduk`
	Nama    string `json:Nama`
	Kelas   string `json:Kelas`
}

type Students []Student

func setStudent() (data []Student) {
	dataKelas := Students{
		Student{
			NoInduk: "1234",
			Nama:    "Adi",
			Kelas:   "X",
		},
		Student{
			NoInduk: "3456",
			Nama:    "Budi",
			Kelas:   "XI",
		}}

	dataKelas = append(dataKelas, Student{
		NoInduk: "7890",
		Nama:    "Rudi",
		Kelas:   "X",
	})
	data = dataKelas
	return data
}

func students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(setStudent())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func studentsOther(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(setStudent())
}

func student(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range setStudent() {
			if each.NoInduk == id {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func postStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Print(r.FormValue("NoInduk") + "=" + r.FormValue("Nama") + "=" + r.FormValue("Kelas"))

		noInduk := r.FormValue("NoInduk")
		name := r.FormValue("Nama")
		class := r.FormValue("Kelas")
		siswa := setStudent()
		siswa = append(siswa, Student{
			NoInduk: noInduk,
			Nama:    name,
			Kelas:   class,
		})

		json.NewEncoder(w).Encode(siswa)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func postStudentJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			log.Printf("Bad Request : %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if student.NoInduk == "" {
			log.Printf("Bad Request : Nomor Induk tidak boleh kosong")
			http.Error(w, "Bad Request : Nomor Induk tidak boleh kosong", http.StatusBadRequest)
			return
		} else {
			siswa := setStudent()
			siswa = append(siswa, student)
			log.Printf("Data : %v", siswa)
			json.NewEncoder(w).Encode(siswa)
		}

	} else {
		log.Printf("Invalid Method")
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
