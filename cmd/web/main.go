package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"snippetbox.abdou-salama-001.net/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil { // what is ping so
		return nil, err
	}

	return db, nil
}

func main() {

	addr := flag.String("addr", ":4000", "tcp port number")
	dsn := flag.String("dsn", "web:0000@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// creat two loggers for info and errors

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn) // flag of dsn
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog, // errorLog on lift point to some error logger by difintion and it is errorlogger we made
		infoLog:  infoLog,  //	to lines above
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s ", *addr)

	err = srv.ListenAndServe() // will not creat instance of our server instead we creat it with out values
	// err := http.ListenAndServe(*addr, mux) // note this here * becuase flag string return the refrence not the value
	// This is roughly what http.ListenAndServe does internally
	/*
		  func ListenAndServe(addr string, handler Handler) error {
		    server := &Server{
		        Addr:    addr,
		        Handler: handler,
		    }
		    return server.ListenAndServe()
			}
	*/
	errorLog.Fatal(err)
}

/* fmux := http.NewServeMux()
// mux.Handle("/", http.HandlerFunc(home))  http. gandler Func wrap the home function with the http serve interface


type home struct {
func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("This is my home page"))
}

mux := http.NewServeMux()
mux.Handle("/", &home{})

this also equivlant as this more manual implmenting :

	type Handler interface {
 ServeHTTP(ResponseWriter, *Request)
}
*/
