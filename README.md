# Currency Converter

A currency conversion web app with an **HTML frontend** and **Go backend**. The backend proxies live exchange rates from the [Frankfurter API](https://www.frankfurter.dev/) (free, no API key).

## Features

- Convert between 30+ currencies
- Live exchange rates (backend fetches from Frankfurter)
- Swap “from” and “to” currencies with one click
- Responsive layout

## Project layout

- `index.html` — Frontend (HTML, CSS, JavaScript)
- `main.go` — Go HTTP server: serves the app and `/api/currencies`, `/api/latest`
- `go.mod` — Go module

## How to run

**Using the Go backend (recommended):**

From the project folder:

```bash
go run .
```

Then open **http://localhost:8080**.

To use another port:

```bash
PORT=3000 go run .
```

**Without the backend (frontend only):**  
Opening `index.html` directly will not work for conversion, because the app expects the Go server’s `/api` endpoints. Use `go run .` and visit http://localhost:8080.

**Using Docker:**

```bash
# Build
docker build -t currency-converter .

# Run
docker run -p 8080:8080 currency-converter
```

Then open **http://localhost:8080**.

**Push to Docker Hub (or another registry):**

```bash
# Log in (use your Docker Hub username)
docker login

# Tag with your Docker Hub username
docker tag currency-converter YOUR_USERNAME/currency-converter:latest

# Push
docker push YOUR_USERNAME/currency-converter:latest
```

Replace `YOUR_USERNAME` with your Docker Hub username. Others can then run:

```bash
docker pull YOUR_USERNAME/currency-converter:latest
docker run -p 8080:8080 YOUR_USERNAME/currency-converter:latest
```

## Usage

1. Enter an amount.
2. Choose “From” and “To” currencies.
3. Click **Convert** or use the swap button to flip currencies.

Rates are updated daily by the Frankfurter API; the Go server proxies requests to it.
