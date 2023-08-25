package main

import (
    "database/sql"
    "fmt"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/go-chi/chi"
)

// Defina a estrutura da rota
type Route struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Source      Location `json:"source"`
    Destination Location `json:"destination"`
}

type Location struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

var db *sql.DB

func main() {
    // Inicialize a conexão com o banco de dados
    var err error
    db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/routes")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    r := chi.NewRouter()

    r.Post("/api/routes", createRoute)
    r.Get("/api/routes", listRoutes)

    // Inicie o servidor na porta 8080
    port := ":8080"
    fmt.Printf("Listening on port %s...\n", port)
    http.ListenAndServe(port, r)
}

func createRoute(w http.ResponseWriter, r *http.Request) {
    var route Route

    if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    insertQuery := "INSERT INTO routes (name, source_lat, source_lng, dest_lat, dest_lng) VALUES (?, ?, ?, ?, ?)"
    _, err := db.Exec(insertQuery, route.Name, route.Source.Lat, route.Source.Lng, route.Destination.Lat, route.Destination.Lng)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func listRoutes(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, source_lat, source_lng, dest_lat, dest_lng FROM routes")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var routes []Route
    for rows.Next() {
        var route Route
        if err := rows.Scan(&route.ID, &route.Name, &route.Source.Lat, &route.Source.Lng, &route.Destination.Lat, &route.Destination.Lng); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        routes = append(routes, route)
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(routes); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
