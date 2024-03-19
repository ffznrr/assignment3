package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

// Struct untuk data JSON
type WeatherData struct {
	Water       int
	WaterStatus string
	Wind        int
	WindStatus  string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Generate angka acak untuk water dan wind
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Menentukan status water
		waterStatus := "Aman"
		if water <= 5 {
			waterStatus = "Aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "Siaga"
		} else {
			waterStatus = "Bahaya"
		}

		// Menentukan status wind
		windStatus := "Aman"
		if wind <= 6 {
			windStatus = "Aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "Siaga"
		} else {
			windStatus = "Bahaya"
		}

		// Membuat data WeatherData
		data := WeatherData{
			Water:       water,
			WaterStatus: waterStatus,
			Wind:        wind,
			WindStatus:  windStatus,
		}

		// Render template HTML dengan data
		tmpl.Execute(w, data)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
