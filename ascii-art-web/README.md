# ASCII Art Web 🎨

A web-based ASCII art generator written in Go.

This application allows users to:
- Enter text through a web form
- Choose a banner style
- Generate ASCII art output directly in the browser

---

## Features

- HTTP server using Go's `net/http`
- HTML templates
- Multiple ASCII banner styles
- Error handling
- Multi-line text support

---

## Project Structure

```bash
ascii-art-web/
│
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── templates/
│   └── index.html
│
├── main.go
├── banner.go
├── render.go
├── main_test.go
└── README.md
```

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd ascii-art-web
```

Run the server:

```bash
go run .
```

Server starts on:

```bash
http://localhost:7072/home
```

---

## Routes

### GET `/home`

Displays the home page.

Example:

```bash
http://localhost:7072/home
```

---

### POST `/ascii-art`

Processes the form input and generates ASCII art.

Form fields:
- `text`
- `banner`

Example:

```txt
Text: Hello
Banner: standard
```

---

## Supported Banners

- standard
- shadow
- thinkertoy

Banner files are stored inside:

```bash
banners/
```

---

## Error Handling

The application handles:

| Error | Status Code |
|------|------|
| Invalid route | 404 |
| Invalid request method | 400 |
| Missing form fields | 400 |
| Missing banner file | 404 |
| Template execution failure | 500 |

---

## Technologies Used

- Go
- HTML Templates
- net/http

---

## Learning Concepts

This project demonstrates:

- HTTP routing
- Request handling
- Form parsing
- Template rendering
- String manipulation
- File handling
- ASCII rendering logic
- Basic web server architecture

---

## Future Improvements 

- Add CSS themes
- Download ASCII art as text file
- Add color support
- Add live preview
- Add Docker support

---

## Author

Built with Golang by 
- Ayodeji Ajisegiri(team lead)
- Babatunde Abidemi