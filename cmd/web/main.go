package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.abdou-salama-001.net/internal/models"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	addr := flag.String("addr", ":4000", "tcp port number")
	dsn := flag.String("dsn", "web:0000@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// start - session manager
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true
	// session manager

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		sessionManager: sessionManager,
	}
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		MinVersion:       tls.VersionTLS12, // Only accept TLS 1.2 or newer
	}
	srv := &http.Server{
		Addr:           *addr,
		MaxHeaderBytes: 524288, // add max header size , 1 mb by def.
		ErrorLog:       errorLog,
		Handler:        app.routes(),
		TLSConfig:      tlsConfig,
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	infoLog.Printf("Starting server on %s ", *addr)

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

/*

session steps
1 - setup third party to handle session management
2 - make table sessions + index for it
3 - add sessionManager property to app struct
4 - define session manager + inject it into the app definition
5 - modifiy router to dynamic handle .
6 - modify creatPost [post] handler to put message in context of response in specific handler
	-Stores flash message in session in db data blob
7 - send data to front end (modify template)
8 - path the data flash inside newTemplateData helper
		Retrieves AND removes flash message
*/

/* self generated tls

use generate.cert.go which is go utility used to generate public and private key (tls)
main > use  srv.ListenAndServeTLS() instead of ListenAndServe and path keys in it
main > set session.cookie.security to true
now run on https , A big plus of using HTTPS is that — if a client supports HTTP/2 connections — Go’s HTTPS
server will automatically upgrade the connection to use HTTP/2.

$ echo 'tls/' >> .gitignore
  print tls  to.  git ignore


*/
