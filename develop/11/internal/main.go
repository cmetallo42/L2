package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	idatabase "calendar/internal/database"
	iserver "calendar/internal/server"
)

type (
	Configuration struct {
		Host  string

		Database idatabase.Configuration
	}
)

func Main(configuration *Configuration) (err error) {
	d, err := idatabase.NewDatabase(&configuration.Database)
	if err != nil {
		return
	}
	defer d.Close()

	http, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	s := iserver.NewServer(logger(handlerInit(d)))
	defer s.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	errs := make(chan error)

	go func() {
		errs <- s.Serve(http)
	}()

	select {
	case err = <-errs:
	case <-signals:
	}

	return
}

func logger(handler http.HandlerFunc) (middlewared http.HandlerFunc) {
	middlewared = func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.String(), r.Form.Encode())

		handler(w, r);
	}

	return
}

func handlerInit(d *idatabase.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		re := RespErr{}
		if r.Method == "POST" {
			user_id := r.Form.Get("user_id")
			date := r.Form.Get("date")
			name := r.Form.Get("name")
			if r.URL.Path == "/create_event" {
				d.Pool.QueryRow(d.DCTX, "INSERT INTO event (user_id, name, event_date) VALUES ($1, $2, $3)",
					user_id, name, date,
				)
			} else if r.URL.Path == "/update_event" {
				new_user := r.Form.Get("new_user")
				new_date := r.Form.Get("new_date")
				new_name := r.Form.Get("new_name")
				d.Pool.QueryRow(d.DCTX, "UPDATE event SET user_id = $1, name = $2, event_date = $3 WHERE (user_id = $4 AND name = $5 AND event_date = $6)",
					new_user, new_name, new_date, user_id, name, date)
			} else if r.URL.Path == "/delete_event " {
				d.Pool.QueryRow(d.DCTX, "DELETE FROM event  WHERE (user_id = $1 AND name = $2 AND event_date = $3)",
					user_id, date, name,
				)
			} else {
				w.WriteHeader(503)
				re.error = fmt.Errorf("path %s not found", r.URL.Path)
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
		} else if r.Method == "GET" {
			values := r.URL.Query()
			user_id := values.Get("user_id")
			date := values.Get("date")
			if user_id == "" || date == "" {
				w.WriteHeader(400)
				re.error = fmt.Errorf("wrong values")
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
			period := 0
			if r.URL.Path == "/events_for_day" {
				period = 1
			} else if r.URL.Path == "/events_for_week" {
				period = 7
			} else if r.URL.Path == "/events_for_month" {
				period = 30
			} else {
				w.WriteHeader(503)
				re.error = fmt.Errorf("path %s not found", r.URL.Path)
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
			newDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				w.WriteHeader(500)
				re.error = err
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
			newDate = newDate.Add(time.Hour * 24 * time.Duration(period))
			rows, err := d.Pool.Query(d.DCTX, "SELECT user_id, name, event_date FROM event WHERE user_id = $1 AND event_date BETWEEN $2 AND $3",
				user_id, date, newDate,
			)
			if err != nil {
				w.WriteHeader(500)
				re.error = err
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
			events := []Event{}
			event := Event{}
			for rows.Next() {
				err = rows.Scan(&event)
				if err != nil {
					w.WriteHeader(500)
					re.error = err
					log.Println(re.error)
					fmt.Fprintf(w, "%s", re.error.Error())
					return
				}
				events = append(events, event)
			}
			data, err := json.Marshal(events)
			if err != nil {
				w.WriteHeader(500)
				re.error = err
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
			_, err = w.Write(data)
			if err != nil {
				w.WriteHeader(500)
				re.error = err
				log.Println(re.error)
				fmt.Fprintf(w, "%s", re.error.Error())
				return
			}
		} else {
			w.WriteHeader(503)
			re.error = fmt.Errorf("method not found")
			log.Println(re.error)
			fmt.Fprintf(w, "%s", re.error.Error())
			return
		}
	}
}

type (
	RespErr struct {
		error error
	}

	Event struct {
		User_id    int64     `json:"user_id" db:"user_id"`
		Name       string    `json:"name" db:"name"`
		Event_date time.Time `json:"date" db:"event_date"`
	}
)

func WriteError(err error) {
}
