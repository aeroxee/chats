package main

import (
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func main() {
	h := newHub()
	go h.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		roomid := r.URL.Query().Get("roomid")
		serveWS(h, roomid, w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.html")

		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		roomid := r.URL.Query().Get("roomid")

		filepath := path.Join("views", "room.html")
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userRandom := rand.Int31()
		user := strconv.FormatInt(int64(userRandom), 10)

		data := map[string]string{
			"roomid": roomid,
			"user":   user,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":8000", nil)
}
