# Go GitHub Repository Analyzer

A fast and concurrent Go application that analyzes popular GitHub repositories across multiple programming languages and generates insightful reports.

## 🎯 What This Project Does

This tool will:

- **Fetch** data from GitHub API for popular repositories
- **Analyze** repositories across different programming languages (Go, Python, TypeScript, etc.)
- **Generate** multiple report formats (JSON, HTML, TXT)
- **Visualize** trends with interactive charts
- **Compare** languages by popularity and activity

## ✨ Key Features

### 🚀 **Performance & Concurrency**

- Uses Go goroutines for parallel API requests
- Processes 500+ repositories in ~30 seconds
- Built-in rate limiting to respect GitHub API limits
- Zero external dependencies (uses only Go standard library)

### 📊 **Rich Data Analysis**

- **Language Statistics**: Count, total stars, average popularity per language
- **Top Repositories**: Ranked list of most popular projects
- **Trend Reports**: Interactive HTML dashboard with charts
- **Complete Dataset**: Full JSON export for custom analysis

### 📋 **Multiple Output Formats**

- `language_stats.json` - Aggregated statistics by programming language
- `top_repos.txt` - Human-readable ranking of top repositories
- `trends_report.html` - Interactive web dashboard with Chart.js visualizations
- `full_data.json` - Complete dataset for further analysis

## 🛠️ Getting Started

### Prerequisites

- Go 1.19+ installed
- Internet connection
- No GitHub token required (uses public API)

### Quick Start

```bash
# 1. Create project directory
mkdir github-analyzer && cd github-analyzer

# 2. Initialize Go module
go mod init github-analyzer

# 3. Create main.go file with the provided code

# 4. Run the analyzer
go run main.go
```

### Expected Output

```
✓ 100 repositories for go
✓ 100 repositories for python
✓ 100 repositories for typescript
✓ 100 repositories for rust
✓ 100 repositories for javascript

Total: 500 repositories analyzed
Analysis complete! Generated files:
- language_stats.json
- top_repos.txt  
- trends_report.html
- full_data.json
```

## 📈 Sample Results

### Language Statistics (JSON)

```json
[
  {
    "language": "TypeScript",
    "count": 95,
    "total_stars": 147820,
    "avg_stars": 1556.53
  },
  {
    "language": "Go", 
    "count": 88,
    "total_stars": 98430,
    "avg_stars": 1118.52
  }
]
```

### Top Repositories (TXT)

```
🏆 TOP REPOSITORIES BY STARS
================================
1. microsoft/vscode (TypeScript)
   ⭐ 45,623 stars | 🍴 8,934 forks
   📝 Visual Studio Code - Code editing redefined

2. golang/go (Go)
   ⭐ 42,341 stars | 🍴 6,789 forks
   📝 The Go programming language
```

### Interactive Dashboard (HTML)

- Bar charts showing repository distribution by language
- Responsive design that works on desktop and mobile
- Real-time data visualization with Chart.js
- Professional styling for presentations

## 🎮 Use Cases

### **For Developers**

- Discover trending libraries and frameworks
- Find popular projects to contribute to
- Stay updated with technology trends

### **For Tech Teams**

- Make informed technology stack decisions
- Analyze market adoption of programming languages
- Research competitor projects and tools

### **For Students & Researchers**

- Study open source ecosystem trends
- Analyze programming language popularity
- Create data-driven presentations about software development

## ⚙️ Customization

### Change Languages to Analyze

```go
languages := []string{"go", "rust", "python", "java", "kotlin", "swift"}
```

### Adjust Search Criteria

```go
// Search for repositories created in the last year with 100+ stars
url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s+created:>2023-01-01+stars:>100&sort=stars&order=desc&per_page=100", language)
```

### Modify Rate Limiting

```go
time.Sleep(time.Second * 2) // Wait 2 seconds between requests
```

## 🔧 Technical Highlights

**Why Go?**

- **Fast**: Concurrent processing with goroutines
- **Simple**: Single binary, no runtime dependencies
- **Reliable**: Strong typing and built-in error handling
- **Efficient**: Low memory usage and fast execution

**Architecture Pattern:**

- Worker pool for concurrent API requests
- Channel-based communication between goroutines
- Rate limiting to respect API quotas
- Structured data processing with custom types

## 📝 Learning Objectives

This project teaches:

- **Go Concurrency**: Goroutines, channels, and sync patterns
- **HTTP APIs**: REST client implementation and JSON parsing
- **File I/O**: Writing data in multiple formats
- **Data Processing**: Aggregation, sorting, and statistical analysis
- **Error Handling**: Robust error management in Go
- **Rate Limiting**: Respecting external API constraints

## 🤔 Why Build This?

Perfect first Go project because it covers:

- Real-world API integration
- Practical concurrency usage
- Multiple output formats
- Data analysis concepts
- Professional code structure

---

**Ready to analyze the GitHub ecosystem? Let's build it! 🚀**
