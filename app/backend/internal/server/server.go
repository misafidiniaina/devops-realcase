package server

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misafidiniaina/cloud-platform-lab/internal/config"
	"github.com/misafidiniaina/cloud-platform-lab/internal/health"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
)

type Server struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Hostname    string `json:"hostname"`
	IPAddress   string `json:"ip_address"`
	Environment string `json:"environment"`
	Provider    string `json:"provider"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type Router struct {
	db       *pgxpool.Pool
	cfg      config.Config
	requests atomic.Uint64
}

func NewRouter(db *pgxpool.Pool, cfg config.Config) http.Handler {
	r := &Router{db: db, cfg: cfg}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Handler)
	mux.HandleFunc("/metrics", r.metrics)
	mux.HandleFunc("/api/v1/servers", r.servers)
	mux.HandleFunc("/api/v1/servers/", r.serverByID)
	return r.middleware(mux)
}
func (r *Router) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.requests.Add(1)
		w.Header().Set("Access-Control-Allow-Origin", r.cfg.CORSOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if req.Method == http.MethodOptions {
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, req)
	})
}
func (r *Router) metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	_, _ = w.Write([]byte("# HELP cloud_platform_http_requests_total Total HTTP requests received.\n# TYPE cloud_platform_http_requests_total counter\ncloud_platform_http_requests_total " + strconv.FormatUint(r.requests.Load(), 10) + "\n"))
}
func (r *Router) servers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.listServers(w, req)
	case http.MethodPost:
		r.createServer(w, req)
	default:
		writeError(w, 405, "method not allowed")
	}
}
func (r *Router) serverByID(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(strings.TrimPrefix(req.URL.Path, "/api/v1/servers/"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, 400, "invalid server id")
		return
	}
	switch req.Method {
	case http.MethodGet:
		r.getServer(w, req, id)
	case http.MethodPut:
		r.updateServer(w, req, id)
	case http.MethodDelete:
		r.deleteServer(w, req, id)
	default:
		writeError(w, 405, "method not allowed")
	}
}
func (r *Router) listServers(w http.ResponseWriter, req *http.Request) {
	rows, err := r.db.Query(req.Context(), `SELECT id,name,hostname,ip_address::text,environment,provider,created_at::text,updated_at::text FROM servers ORDER BY id DESC`)
	if err != nil {
		writeError(w, 500, "failed to list servers")
		return
	}
	defer rows.Close()
	out := make([]Server, 0)
	for rows.Next() {
		var s Server
		if err := rows.Scan(&s.ID, &s.Name, &s.Hostname, &s.IPAddress, &s.Environment, &s.Provider, &s.CreatedAt, &s.UpdatedAt); err != nil {
			writeError(w, 500, "failed to read servers")
			return
		}
		out = append(out, s)
	}
	if err := rows.Err(); err != nil {
		writeError(w, 500, "failed to read servers")
		return
	}
	writeJSON(w, 200, out)
}
func (r *Router) getServer(w http.ResponseWriter, req *http.Request, id int64) {
	var s Server
	err := r.db.QueryRow(req.Context(), `SELECT id,name,hostname,ip_address::text,environment,provider,created_at::text,updated_at::text FROM servers WHERE id=$1`, id).Scan(&s.ID, &s.Name, &s.Hostname, &s.IPAddress, &s.Environment, &s.Provider, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, 404, "server not found")
		return
	}
	if err != nil {
		writeError(w, 500, "database error")
		return
	}
	writeJSON(w, 200, s)
}
func (r *Router) createServer(w http.ResponseWriter, req *http.Request) {
	var in Server
	if err := json.NewDecoder(req.Body).Decode(&in); err != nil {
		writeError(w, 400, "invalid JSON")
		return
	}
	if err := validate(in); err != nil {
		writeError(w, 400, err.Error())
		return
	}
	var s Server
	err := r.db.QueryRow(req.Context(), `INSERT INTO servers(name,hostname,ip_address,environment,provider) VALUES($1,$2,$3,$4,$5) RETURNING id,name,hostname,ip_address::text,environment,provider,created_at::text,updated_at::text`, in.Name, in.Hostname, in.IPAddress, in.Environment, in.Provider).Scan(&s.ID, &s.Name, &s.Hostname, &s.IPAddress, &s.Environment, &s.Provider, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		writeError(w, 500, "failed to create server")
		return
	}
	writeJSON(w, 201, s)
}
func (r *Router) updateServer(w http.ResponseWriter, req *http.Request, id int64) {
	var in Server
	if err := json.NewDecoder(req.Body).Decode(&in); err != nil {
		writeError(w, 400, "invalid JSON")
		return
	}
	if err := validate(in); err != nil {
		writeError(w, 400, err.Error())
		return
	}
	var s Server
	err := r.db.QueryRow(req.Context(), `UPDATE servers SET name=$1,hostname=$2,ip_address=$3,environment=$4,provider=$5,updated_at=NOW() WHERE id=$6 RETURNING id,name,hostname,ip_address::text,environment,provider,created_at::text,updated_at::text`, in.Name, in.Hostname, in.IPAddress, in.Environment, in.Provider, id).Scan(&s.ID, &s.Name, &s.Hostname, &s.IPAddress, &s.Environment, &s.Provider, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, 404, "server not found")
		return
	}
	if err != nil {
		writeError(w, 500, "failed to update server")
		return
	}
	writeJSON(w, 200, s)
}
func (r *Router) deleteServer(w http.ResponseWriter, req *http.Request, id int64) {
	res, err := r.db.Exec(req.Context(), `DELETE FROM servers WHERE id=$1`, id)
	if err != nil {
		writeError(w, 500, "failed to delete server")
		return
	}
	if res.RowsAffected() == 0 {
		writeError(w, 404, "server not found")
		return
	}
	w.WriteHeader(204)
}
func validate(s Server) error {
	for n, v := range map[string]string{"name": s.Name, "hostname": s.Hostname, "ip_address": s.IPAddress, "environment": s.Environment, "provider": s.Provider} {
		if strings.TrimSpace(v) == "" {
			return errors.New(n + " is required")
		}
	}
	return nil
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
