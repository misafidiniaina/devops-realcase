package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type input struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type API struct{ db *pgxpool.Pool }

func New(db *pgxpool.Pool) *API { return &API{db: db} }
func (a *API) Register(m *http.ServeMux) {
	m.HandleFunc("/health", a.health)
	m.HandleFunc("/metrics", a.metrics)
	m.HandleFunc("/api/tasks", a.tasks)
	m.HandleFunc("/api/tasks/", a.task)
}
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func (a *API) health(w http.ResponseWriter, _ *http.Request) {
	status := "ok"
	if a.db != nil && a.db.Ping(w.Context()) != nil {
		status = "degraded"
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": status})
}
func (a *API) metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	_, _ = w.Write([]byte("# HELP app_up Application availability\n# TYPE app_up gauge\napp_up 1\n"))
}
func (a *API) tasks(w http.ResponseWriter, r *http.Request) {
	if a.db == nil {
		writeJSON(w, 503, map[string]string{"error": "database is not configured"})
		return
	}
	switch r.Method {
	case http.MethodGet:
		a.list(w, r)
	case http.MethodPost:
		a.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (a *API) task(w http.ResponseWriter, r *http.Request) {
	if a.db == nil {
		writeJSON(w, 503, map[string]string{"error": "database is not configured"})
		return
	}
	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/api/tasks/"), 10, 64)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid task id"})
		return
	}
	switch r.Method {
	case http.MethodGet:
		a.get(w, r, id)
	case http.MethodPut:
		a.update(w, r, id)
	case http.MethodDelete:
		a.delete(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func decode(r *http.Request) (input, error) {
	var in input
	err := json.NewDecoder(r.Body).Decode(&in)
	if err == nil && strings.TrimSpace(in.Title) == "" {
		err = fmt.Errorf("title is required")
	}
	if in.Status == "" {
		in.Status = "todo"
	}
	return in, err
}
func (a *API) list(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `SELECT id,title,description,status,created_at,updated_at FROM tasks ORDER BY id DESC`)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()
	result := []Task{}
	for rows.Next() {
		var t Task
		if err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		result = append(result, t)
	}
	writeJSON(w, 200, result)
}
func (a *API) get(w http.ResponseWriter, r *http.Request, id int64) {
	var t Task
	err := a.db.QueryRow(r.Context(), `SELECT id,title,description,status,created_at,updated_at FROM tasks WHERE id=$1`, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err == pgx.ErrNoRows {
		writeJSON(w, 404, map[string]string{"error": "task not found"})
		return
	}
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, t)
}
func (a *API) create(w http.ResponseWriter, r *http.Request) {
	in, err := decode(r)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	var t Task
	err = a.db.QueryRow(r.Context(), `INSERT INTO tasks(title,description,status) VALUES($1,$2,$3) RETURNING id,title,description,status,created_at,updated_at`, in.Title, in.Description, in.Status).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, t)
}
func (a *API) update(w http.ResponseWriter, r *http.Request, id int64) {
	in, err := decode(r)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	var t Task
	err = a.db.QueryRow(r.Context(), `UPDATE tasks SET title=$1,description=$2,status=$3,updated_at=now() WHERE id=$4 RETURNING id,title,description,status,created_at,updated_at`, in.Title, in.Description, in.Status, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err == pgx.ErrNoRows {
		writeJSON(w, 404, map[string]string{"error": "task not found"})
		return
	}
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, t)
}
func (a *API) delete(w http.ResponseWriter, r *http.Request, id int64) {
	tag, err := a.db.Exec(r.Context(), `DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if tag.RowsAffected() == 0 {
		writeJSON(w, 404, map[string]string{"error": "task not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
