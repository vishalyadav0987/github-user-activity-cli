# 🐙 GitHub Activity CLI (Golang)

🔗 **Project URL:**
https://roadmap.sh/projects/github-user-activity

---

A lightweight and efficient **Command Line Interface (CLI)** application built in **Go (Golang)** to fetch and display a GitHub user's recent activity directly in the terminal.

---

## 🚀 Features

* 👤 Fetch GitHub user activity using username
* 📡 Uses GitHub public API (no authentication required)
* 🧾 Clean and readable terminal output
* 🔄 Supports multiple event types (Push, Issues, Stars, etc.)
* ⚡ Fast and lightweight CLI tool
* ❌ Graceful error handling (invalid user, API failure, etc.)
* 🧠 Clean Architecture (Domain + Application + Infrastructure)

---

## 📁 Project Structure


```
github-user-activity
├─ README.md
├─ cmd
│  └─ github-user-activity
│     └─ main.go
├─ go.mod
├─ interfaces
│  └─ cli
│     ├─ handler.go
│     ├─ output.go
│     └─ parser.go
└─ internal
   ├─ application
   │  └─ github-activity
   │     └─ service.go
   ├─ config
   ├─ domain
   │  └─ github-activity
   │     ├─ entity.go
   │     └─ errors.go
   └─ infrastructure
      ├─ client
      │  └─ github_client.go
      └─ github-activity
         └─ service_implementation.go


```

---

## ⚙️ Installation

### 1. Clone repo

```bash
git clone https://github.com/vishalyadav0987/github-activity-cli.git
cd github-activity-cli
```

---

### 2. Build binary

```bash
go build -o github-activity cmd/github-activity/main.go
```

---

### 3. Make it global (Mac/Linux)

```bash
sudo mv github-activity /usr/local/bin/
```

---

## 🧪 Usage

### ▶️ Fetch GitHub Activity

```bash
github-activity <username>
```

---

### 📌 Example

```bash
github-activity vishalyadav0987
```

---

## 🖥️ Example Output

```
- Pushed commits to vishalyadav0987/github-user-activity-cli
- Did CreateEvent on vishalyadav0987/github-user-activity-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/task-tracker-cli
- Did CreateEvent on vishalyadav0987/task-tracker-cli
- Pushed commits to vishalyadav0987/vishalyadav0987
- Pushed commits to vishalyadav0987/vishalyadav0987
- Pushed commits to vishalyadav0987/go-auth-ddd
- Pushed commits to vishalyadav0987/go-auth-ddd
- Did CreateEvent on vishalyadav0987/go-auth-ddd
```

---

## ⚠️ Error Handling

* ❌ Invalid username → "user not found"
* ❌ API failure → "failed to fetch data"
* ❌ No activity → "no recent activity found"
* ❌ Rate limit exceeded → "github api rate limit exceeded"

---

## 🧠 Tech Stack

* Go (Golang)
* GitHub REST API
* Clean Architecture (DDD inspired)
* CLI Interface
* Net/http package (no external libraries)

---

## 🏗️ Architecture Highlights

* 🔹 Domain-driven structure
* 🔹 Dependency Injection (manual)
* 🔹 Interface-based design
* 🔹 Separation of concerns (CLI / Service / Infra)

---

## 🔥 Future Improvements

* 🎯 Add caching for repeated users
* ⏱️ Add rate limit handling & retry logic
* 🎨 Colored and formatted CLI output
* 📊 Group events by type (Push / Star / Issue)
* 🔐 GitHub token support for higher API limits
* 🧪 Unit tests with mocked GitHub client

---

## 👨‍💻 Author

**Vishal Yadav**

---

## ⭐ If you like this project

Give it a ⭐ on GitHub! 🚀

