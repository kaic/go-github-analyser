package main

import "fmt"

type Command struct {
    Operation string
}

func HandleCommand(cmd Command) {
    switch cmd.Operation {
    case "analyse":
        fmt.Println("Analyzing...")
        // Add the analysis logic here
    default:
        fmt.Println("Unknown operation:", cmd.Operation)
    }
}     


func main() {
    var operation string
    fmt.Print("\n --- Supported Operations: --- \n\n ** Analyse ** \n\n - Fetch data from GitHub API for popular repositories \n - Analyze repositories across different programming languages '(Go, Python, TypeScript, etc.)\n - Extract key metrics (stars, forks, issues, etc.)\n - Generate insights on language trends and repository activity\n - Visualize data with charts and graphs\n - Export results in various formats (JSON, CSV, etc.)\n - Compare repositories by language and activity level\n - Generate reports with detailed analysis\n - Visualize trends over time with interactive charts\n - Compare languages by popularity and activity\n\nEnter the desired operation (eg. analyse): ")
    fmt.Scanln(&operation)
    cmd := Command{Operation: operation}
    HandleCommand(cmd)
}