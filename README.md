# Snippetbox

Snippetbox is a simple web application for managing code snippets. It allows users to create, view, and manage snippets.

## Features

- Create new snippets
- View existing snippets
- List latest snippets
- Session management
- Form validation



## Dependencies

- [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) - A lightweight, high-performance HTTP request router (mux) for Go.
- [github.com/justinas/alice](https://github.com/justinas/alice) - A middleware chaining library for Go.
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - A MySQL driver for Go's `database/sql` package.
- [github.com/alexedwards/scs/v2](https://github.com/alexedwards/scs) - A session management package for Go.
- [github.com/alexedwards/scs/mysqlstore](https://github.com/alexedwards/scs) - A MySQL session store for the `scs` session management package.

## Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/abdelrahmanShawki/snippetbox.git
    cd snippetbox
    ```

2. **Install dependencies:**

    ```sh
    go get -u github.com/julienschmidt/httprouter
    go get -u github.com/justinas/alice
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/alexedwards/scs/v2
    go get -u github.com/alexedwards/scs/mysqlstore
    ```

3. **Set up the MySQL database:**

    ```sql
    CREATE DATABASE snippetbox;
    USE snippetbox;

    CREATE TABLE snippets (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL,
        expires DATETIME NOT NULL
    );
    ```

4. **Run the application:**

    ```sh
    go run cmd/web/main.go
    ```

5. **Access the application:**

    Open your web browser and navigate to [http://localhost:4000]

## Usage

- **Create a new snippet:** Navigate to `/snippet/create` and fill out the form.
- **View a snippet:** Navigate to `/snippet/view/{id}` where `{id}` is the ID of the snippet.
- **List latest snippets:** Navigate to `/`.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
