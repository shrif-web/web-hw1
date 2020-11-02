// package main

// import (
//     "fmt"
//     "log"
//     "os"
//     "bufio"
//     "net/http"
//     "strconv"
//     "crypto/sha256"
//     "encoding/hex"
//     "encoding/json"
// )

// type SHA_Response struct {
//     SHA string
// }

// type Write_Response struct {
//     Line string
// }

// type Err struct {
//     Error string
// }

// type ShaInput struct{
//     Num1 string
//     Num2 string
// }

// func shaHandler(w http.ResponseWriter, r *http.Request) {

//     w.Header().Set("Access-Control-Allow-Origin", "*")
//     w.Header().Set("Content-Type", "aplication/json; charset=utf-8")
//     w.Header().Set("Connection", "keep-alive")
//     // if err := r.ParseForm(); err != nil {
//     //     e := &Err{
//     //         Error : "Error parsing form!",
//     //     }

//     //     w.Header().Set("Content-Type", "application/xml")
//     //     json.NewEncoder(w).Encode(e)
//     //     return
//     // }
//     // r.ParseForm()

//     var p ShaInput

//     // Try to decode the request body into the struct. If there is an error,
//     // respond to the client with the error message and a 400 status code.

//     decoder := json.NewDecoder(r.Body)

//     err := decoder.Decode(&p)
//     if err != nil {
//         panic(err)
//     }

//     log.Println(p)

//     // Do something with the Person struct...
//     // fmt.Fprintf(w, "Person Salam: %+v", p)

//     // fmt.Printf("form parser %+v\n", r.Form)
//     number1,err1 := strconv.ParseUint(p.Num1, 10, 64)
//     number2,err2 := strconv.ParseUint(p.Num2, 10, 64)

//     if (err1 != nil) || (err2 != nil){
//         e := &Err{
//             Error : "Incorrect input type!",
//         }

//         json.NewEncoder(w).Encode(e)
//         return
//     }

//     var sum = number1 + number2
//     stringToHash := strconv.Itoa(int(sum))
//     sha256Bytes := sha256.Sum256([]byte(stringToHash))

//     //fmt.Printf()

//     data := &SHA_Response{
//         SHA: hex.EncodeToString(sha256Bytes[:]),
//       }

//     json.NewEncoder(w).Encode(data)
// }

// func writeHandler(w http.ResponseWriter, r *http.Request) {

//     // if r.Method != "GET" {
//     //     http.Error(w, "Method is not supported.", http.StatusNotFound)
//     //     return
//     // }

//     if err := r.ParseForm(); err != nil {
//         e := &Err{
//             Error : "Error parsing form!",
//         }

//         json.NewEncoder(w).Encode(e)
//         return
//     }
//     r.ParseForm()

//     line,err1 := strconv.ParseUint(r.Form["line"][0], 10, 64)

//     if (err1 != nil) || (line < 1) || (line > 100){
//         e := &Err{
//             Error : "Incorrect input type or line number out of range!",
//         }

//         json.NewEncoder(w).Encode(e)
//         return
//     }

//     file, err := os.Open("./sample.txt")

//     if err != nil {
//         log.Fatalf("failed to open")

//     }

//     scanner := bufio.NewScanner(file)

//     scanner.Split(bufio.ScanLines)
//     var text []string

//     for scanner.Scan() {
//         text = append(text, scanner.Text())
//     }

//     file.Close()

//     line_res := &Write_Response{
//         Line : text[line-1],
//     }

//     json.NewEncoder(w).Encode(line_res)
//     return

// }

// func main() {

//     fileServer := http.FileServer(http.Dir("./static"))
//     http.Handle("/", fileServer)
//     http.HandleFunc("/go/sha256", shaHandler)
//     http.HandleFunc("/go/write", writeHandler)

//     fmt.Printf("Starting server at port 80\n")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatal(err)
//     }
// }

package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Numbers struct {
	Num1 string
	Num2 string
}

type SHA_Response struct {
	SHA string
}

type Line_req struct {
	Line_str_number string
}

type Write_Response struct {
	Line string
}

type Err struct {
	Error string
}

func shaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "aplication/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	var nums Numbers

	//fmt.Println(r.Body);

	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//log.Fatal("Invalid request!");
		e := &Err{
			Error: "Error parsing form!",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
		return
	}

	//fmt.Fprintf(w, "Numbers: %+v", nums)

	number1, err1 := strconv.ParseUint(nums.Num1, 10, 64)
	number2, err2 := strconv.ParseUint(nums.Num2, 10, 64)

	if (err1 != nil) || (err2 != nil) {
		e := &Err{
			Error: "Incorrect input type!",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
		return
	}

	var sum = number1 + number2
	stringToHash := strconv.Itoa(int(sum))
	sha256Bytes := sha256.Sum256([]byte(stringToHash))

	//fmt.Printf()

	data := &SHA_Response{
		SHA: hex.EncodeToString(sha256Bytes[:]),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {

	var line_str Line_req

	err := json.NewDecoder(r.Body).Decode(&line_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//log.Fatal("Invalid request!");
		e := &Err{
			Error: "Error parsing form!",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
		return
	}

	line, err1 := strconv.ParseUint(line_str.Line_str_number, 10, 64)

	if (err1 != nil) || (line < 1) || (line > 100) {
		e := &Err{
			Error: "Incorrect input type or line number out of range!",
		}

		json.NewEncoder(w).Encode(e)
		return
	}

	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	line_res := &Write_Response{
		Line: text[line-1],
	}

	json.NewEncoder(w).Encode(line_res)
	return

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/go/sha256", shaHandler)
	http.HandleFunc("/go/write", writeHandler)

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
