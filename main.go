package main

import (
	"net/http"
	"os"

	"github.com/oodles-noodles/dr-demo-telemetry-collector/internal/ops"
	"github.com/oodles-noodles/dr-demo-telemetry-collector/internal/store"
)

func main() {
	st, err := store.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		rows, err := st.TenantJobs(r.URL.Query().Get("tenant"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/artifacts", func(w http.ResponseWriter, r *http.Request) {
		data, err := ops.ReadArtifact(r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		w.Write(data)
	})

	http.HandleFunc("/maintenance", func(w http.ResponseWriter, r *http.Request) {
		out, err := ops.RunMaintenance(r.URL.Query().Get("task"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(out)
	})

	http.ListenAndServe(":8080", nil)
}
