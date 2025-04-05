# 🧪 TodoPad Performance Benchmarks

## 📝 Context

This repo contains simple Todo applications implemented in **Rust**, **Go**, **Python**, and **JavaScript** (excluded from benchmarking). Each app is a **Multi-Page Application (MPA)** using **HTMX** and **Tailwind CSS** for the frontend UI. Todo items are persisted in a **SQLite database**, and each app loads the list and renders HTML server-side.

The goal of this project is to **measure the performance of HTML rendering across different programming languages**, specifically how efficiently each can load and render the todo list from SQLite.

---

## ⚙️ Environment

All applications are containerized using **Docker**, and run with:

- **1 CPU**
- **250 MB memory**

Micro-benchmarking is done using the [`wrk`](https://github.com/wg/wrk) tool:

- Duration: **30 seconds**
- Threads: **12**
- Concurrent Connections: **400**

---

## 💻 How to Run the Containers (on macOS)

Each language implementation has a `scripts` folder to help you build and run the container.

For example, to run the **Rust** version:

```bash
cd ru-todopad
sh scripts/build.sh
sh scripts/start.sh
```

Replace `ru-todopad` with the desired folder (`go-todopad`, `py-todopad`, etc.).

---

## 📈 How to Run the Benchmark

Once the container is running (typically on `http://localhost:8080`), use the following command:

```bash
wrk -t12 -c400 -d30s --latency http://localhost:8080
```

---

## 📊 Results

### 🚀 Performance & Resource Usage Summary

| Language | Requests/sec | Total Requests | Avg Latency | P95 Latency | P99 Latency | Memory Usage |
|----------|---------------|----------------|--------------|---------------|---------------|----------------|
| **Rust**   | 10,578.37     | 317,644        | 37.69 ms     | 76.14 ms      | 80.72 ms      | 20.75 MiB      |
| **Go**     | 8,415.88      | 252,926        | 50.77 ms     | 98.54 ms      | 195.30 ms     | 102.00 MiB     |
| **Python** | 4,709.16      | 141,794        | 87.92 ms     | 101.11 ms     | 215.08 ms     | 199.10 MiB     |

> Note: JavaScript version is excluded from benchmarking in this round.

---

## 📬 Contributions

Feel free to submit PRs to add more language implementations or benchmark tooling improvements.

---

## 📄 License

MIT
