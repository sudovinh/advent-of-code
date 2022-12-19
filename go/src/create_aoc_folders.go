package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
)

func syncProblem(year, day int) error {
    log.Println(fmt.Sprintf("Syncing %d Advent of Code for day: %d", year, day))
    // Download the problem page for the current day
    resp, err := http.Get(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Read the problem page
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    // Check if the problem page has a "404" status code
    if resp.StatusCode == 404 {
        // No problem is available for this day
        return nil
    }

    // Extract the problem descriptions from the page
    parts := strings.Split(string(body), "---")
    if len(parts) < 3 {
        return fmt.Errorf("Invalid problem page")
    }
    problem := strings.TrimSpace(parts[1]) + "\n\n" + strings.TrimSpace(parts[2])


    // Convert the problem descriptions to Markdown format
    markdown := fmt.Sprintf("# Day %d\n\n%s\n", day, problem)

    // Write the problem descriptions to a Markdown file
    err = ioutil.WriteFile(fmt.Sprintf("%d/day%d/README.md", year, day), []byte(markdown), 0777)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Check if a year was provided as a command-line argument
    if len(os.Args) < 2 {
        log.Fatal("Please provide a year (YYYY) as a command-line argument")
    }

    // Parse the year from the command-line argument
    year, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal("Invalid year: ", err)
    }

    // Check if the sync-only flag is set
    syncOnly := false
    if len(os.Args) > 2 && os.Args[2] == "--sync-only" {
        syncOnly = true
    }

    if !syncOnly {
        // Create the base directory for the Advent of Code event
        log.Println(fmt.Sprintf("Creating base directory for year: %d", year))
        err = os.Mkdir(fmt.Sprintf("%d", year), 0777)
        if err != nil {
            log.Fatal("Failed to create base directory: ", err)
        }
    }

    // Sync the problems for each day of the event
    for day := 1; day <= 25; day++ {
        if !syncOnly {
            // Create the directory for the current day
            err = os.Mkdir(fmt.Sprintf("%d/day%d", year, day), 0777)
            if err != nil {
                log.Fatal("Failed to create directory for day ", day, ": ", err)
            }
        }

        // Sync the problem for the current day
        err = syncProblem(year, day)
        if err != nil {
            log.Fatal("Failed to sync problem for day ", day, ": ", err)
        }
    }
}
