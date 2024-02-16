package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
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

	http.HandleFunc("/xml", xmlGet)
	http.HandleFunc("/sampleXML", getDataXML)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

type Student struct {
	NoInduk string `json:"NoInduk"`
	Nama    string `json:"Nama"`
	Kelas   string `json:"Kelas,omitempty"`
}

type ResponseXML struct {
	Message string
	Data    []string `xml:"Names>Name"`
}

type Students []Student

func getData(w http.ResponseWriter, r *http.Request) {
	// data := strings.TrimPrefix(r.URL.RequestURI())
	// log.Println(r.URL.Query().Get("id") + "=" + r.URL.Query().Get("data"))
	log.Println(r.RequestURI)
}

func xmlGet(w http.ResponseWriter, r *http.Request) {
	response := ResponseXML{"Hello", []string{"World", "Sarkar"}}

	// Wraps the response to Response struct
	x, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	// Write
	w.Write(x)
}

func getDataXML(w http.ResponseWriter, r *http.Request) {
	httpposturl := "http://202.152.20.230:1010/simulator/bpjstk/inq_pu.php"
	postBody, _ := io.ReadAll(r.Body)

	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer([]byte(postBody)))
	request.Header.Add("Content-Type", "application/xml; charset=utf-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer response.Body.Close()
	// var hasil map[string]interface{}
	respDat, _ := io.ReadAll(response.Body)

	w.Write(respDat)

	// Mengurai XML ke dalam struktur data
	var envelope Envelope
	err := xml.Unmarshal([]byte(respDat), &envelope)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mengakses nilai masing-masing elemen
	addValues := envelope.Body.Response.Return.AddValues
	kode := envelope.Body.Response.Return.Kode
	msg := envelope.Body.Response.Return.Msg
	ret := envelope.Body.Response.Return.Ret
	blnTagihan := envelope.Body.Response.Return.BlnTagihan
	iuranJHT := envelope.Body.Response.Return.IuranJHT
	iuranJKK := envelope.Body.Response.Return.IuranJKK
	iuranJKM := envelope.Body.Response.Return.IuranJKM
	iuranJPK := envelope.Body.Response.Return.IuranJPK
	iuranJPN := envelope.Body.Response.Return.IuranJPN
	kodeDivisi := envelope.Body.Response.Return.KodeDivisi
	namaPerusahaan := envelope.Body.Response.Return.NamaPerusahaan
	noTagihan := envelope.Body.Response.Return.NoTagihan
	npp := envelope.Body.Response.Return.Npp
	totalBPJSK := envelope.Body.Response.Return.TotalBPJSK
	totalBPJSTK := envelope.Body.Response.Return.TotalBPJSTK
	totalIuran := envelope.Body.Response.Return.TotalIuran

	// Menampilkan nilai masing-masing elemen
	fmt.Println("addValues:", addValues)
	fmt.Println("kode:", kode)
	fmt.Println("msg:", msg)
	fmt.Println("ret:", ret)
	fmt.Println("blnTagihan:", blnTagihan)
	fmt.Println("iuranJHT:", iuranJHT)
	fmt.Println("iuranJKK:", iuranJKK)
	fmt.Println("iuranJKM:", iuranJKM)
	fmt.Println("iuranJPK:", iuranJPK)
	fmt.Println("iuranJPN:", iuranJPN)
	fmt.Println("kodeDivisi:", kodeDivisi)
	fmt.Println("namaPerusahaan:", namaPerusahaan)
	fmt.Println("noTagihan:", noTagihan)
	fmt.Println("npp:", npp)
	fmt.Println("totalBPJSK:", totalBPJSK)
	fmt.Println("totalBPJSTK:", totalBPJSTK)
	fmt.Println("totalIuran:", totalIuran)

}

func setStudent() (data []Student) {
	dataKelas := Students{
		Student{
			NoInduk: "1234",
			Nama:    "Adi",
			Kelas:   "",
		},
		Student{
			NoInduk: "3456",
			Nama:    "Budi",
			Kelas:   "",
		}}

	dataKelas = append(dataKelas, Student{
		NoInduk: "7890",
		Nama:    "Rudi",
		Kelas:   "",
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
	var data []Student
	for _, val := range setStudent() {
		studentData := Student{
			NoInduk: val.NoInduk,
			Nama:    val.Nama,
		}
		data = append(data, studentData)
	}
	json.NewEncoder(w).Encode(data)
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
			http.Error(w, "Bad Data", http.StatusOK)
			return
		} else {
			siswa := setStudent()
			siswa = append(siswa, student)
			log.Printf("Data : %v", siswa)
			stringJson := map[string]interface{}{
				"rcode": "1234",
				"pesan": "Sukses",
				"data":  setStudent(),
			}
			json.NewEncoder(w).Encode(stringJson)
		}

	} else {
		log.Printf("Invalid Method")
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
