package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
    // Definisikan URL server
    url := "http://localhost:8080/banjir" // Ganti dengan URL yang sesuai dengan server Anda

    // Lakukan permintaan ke server setiap 15 detik
    for {
        // Lakukan permintaan GET ke server
        response, err := http.Get(url)
        if err != nil {
            fmt.Println("Error saat melakukan permintaan ke server:", err)
            continue
        }
        defer response.Body.Close()

        // Baca responsenya
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Println("Error saat membaca responsenya:", err)
            continue
        }

        // Tampilkan responsenya di terminal
        fmt.Println("Respons dari server:", string(body))

        // Tunggu 15 detik sebelum melakukan permintaan kembali
        time.Sleep(15 * time.Second)
    }
}
