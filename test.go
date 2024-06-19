import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    http.HandleFunc("/user", getUser)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUser(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")

    // Vulnerable SQL query
    query := fmt.Sprintf("SELECT id, name FROM users WHERE username = '%s'", username)

    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    row := db.QueryRow(query)
    var id int
    var name string
    if err := row.Scan(&id, &name); err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "User ID: %d, Name: %s", id, name)
}
