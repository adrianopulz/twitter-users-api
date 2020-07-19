package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/adrianopulz/twitter-users-api/users"
)

func mapRoutes() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Main route.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Users API!"))
	})

	// RESTy routes for "users" resource
	r.Route("/users", func(r chi.Router) {
		// r.With(paginate).Get("/", listUsers)

		// r.Post("/", createUser)
		r.Get("/search/{userName}", searchUsers)

		// Subrouters:
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Get("/", getUser)
			r.Put("/", updateUser)
			r.Delete("/", deleteUser)
		})
	})

	return r
}

// UserCtx provide basic userID context to Get, Put and Delete.
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, convErr := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if convErr == nil {
			fmt.Printf("%d of type %T", userID, userID)
		}

		user, err := users.UsersService.GetUser(userID)
		if err != nil {
			http.Error(w, err.Msg, err.Code)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Return User by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(*users.User)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(user.Marshall()))
}

// Upadte a User
func deleteUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This method is not available yet.", http.StatusNotImplemented)
}

// Delete a User
func updateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This method is not available yet.", http.StatusNotImplemented)
}

func searchUsers(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "userName")
	users, err := users.UsersService.SearchUsers(userName)
	if err != nil {
		http.Error(w, err.Msg, err.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(users.Marshall()))
}
