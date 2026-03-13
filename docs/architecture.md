# Overview

**SecureNote** is an end-to-end encrypted pastebin. Content is encrypted client-side using _AES-256-GCM_ via the _Web Crypto API_. Each paste has its own key, which can be accessed either through the URL fragment (for sharing) or via a master key derived from the user's password (for account access). The server only stores the ciphertext and never sees plaintext, paste keys, or master keys. URL fragments are never sent to the server per [RFC 3986](https://www.rfc-editor.org/rfc/rfc3986#section-3.5).

# Technology Stack

The server is written in [Go](https://go.dev/) using the standard library's [`net/http`](https://pkg.go.dev/net/http) package with no framework. Templates are rendered server-side using the [`html/template`](https://pkg.go.dev/html/template) package, with [HTMX](https://htmx.org/) handling dynamic interactions. Client-side encryption uses the [Web Crypto API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Crypto_API) for _AES-256-GCM_ encryption and PBKDF2 key derivation. [SQLite](https://sqlite.org/) for the database, accessed via [SQLC](https://sqlc.dev/) for type-safe query generation and [golang-migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4) for schema migrations. Authentication uses [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) password hashing and cookie-based sessions. The application is containerised with a multi-stage [Docker](https://www.docker.com/) build and uses [GitHub Actions](https://github.com/features/actions) for CI.

| Layer                  | Technology                                                                                      |
| ---------------------- | ----------------------------------------------------------------------------------------------- |
| Language               | [Go](https://go.dev/)                                                                           |
| Templating             | Go [`html/template`](https://pkg.go.dev/html/template)                                          |
| Interactivity          | [HTMX](https://htmx.org/)                                                                       |
| Client-side encryption | WebCrypto API (AES-256-GCM, PBKDF2)                                                             |
| Database               | [SQLite](https://sqlite.org/) via [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) |
| Query Generation       | [SQLC](https://sqlc.dev/)                                                                       |
| Migrations             | [`golang-migrate`](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)                     |
| Authentication         | [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt), cookie-based sessions                  |
| Containerisation       | [Docker](https://www.docker.com/)                                                               |
| CI                     | [GitHub Actions](https://github.com/features/actions)                                           |
| Dev tooling            | [Air](https://github.com/air-verse/air)                                                         |

# Project Management

Development is tracked using GitHub Issues and follows a docs-first TDD workflow. Each feature starts as an issue, isdeveloped on a branch following the [conventional branch naming scheme](https://conventional-branch.github.io/), and is merged via squash merge PR that references and closes the issue.

# Project Structure

```
.
├── cmd/server/main.go          # Application entry point
├── Dockerfile                  #
├── docs
│   ├── api.md
│   └── architecture.md
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   ├── db.go
│   │   ├── models.go
│   │   ├── pastes.sql.go
│   │   └── users.sql.go
│   ├── handlers
│   │   ├── handlers.go
│   │   ├── login.go
│   │   ├── login_test.go
│   │   ├── logout.go
│   │   ├── logout_test.go
│   │   ├── paste.go
│   │   ├── paste_test.go
│   │   ├── register.go
│   │   └── register_test.go
│   └── middleware
│       ├── auth.go
│       ├── csrf.go
│       ├── log.go
│       └── middleware.go
├── LICENSE
├── migrations
│   ├── 000001_init.down.sql
│   └── 000001_init.up.sql
├── README.md
├── sql
│   ├── pastes.sql
│   └── users.sql
├── sqlc.yaml
├── static
│   ├── encrypt_paste.js
│   ├── index.js
│   └── styles.css
└── templates
    ├── fragments
    │   ├── login_res.html
    │   ├── paste_card.html
    │   └── register_res.html
    └── pages
        ├── account.html
        ├── index.html
        ├── login.html
        ├── paste.html
        └── register.html
```

# Request Flow

```mermaid
sequenceDiagram
```

# Data Model

```mermaid
erDiagram
```

# Handler Convention

# Template Organisation

# Error Handling
