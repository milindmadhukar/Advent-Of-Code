package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func generateDay(day int) error {
	dayDir := fmt.Sprintf("day%02d", day)
	dayFile := fmt.Sprintf("%s.go", dayDir)
	dayPath := filepath.Join(dayDir, dayFile)

	if _, err := os.Stat(dayDir); err == nil {
		return fmt.Errorf("directory %s already exists", dayDir)
	}

	if err := os.MkdirAll(dayDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	templateContent, err := os.ReadFile("TEMPLATE")
	if err != nil {
		return fmt.Errorf("failed to read TEMPLATE: %w", err)
	}

	content := string(templateContent)
	content = strings.ReplaceAll(content, "day01", dayDir)
	content = strings.ReplaceAll(content, "2025, 1", fmt.Sprintf("2025, %d", day))

	if err := os.WriteFile(dayPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write day file: %w", err)
	}

	if err := updateSolutionsFile(day, dayDir); err != nil {
		return fmt.Errorf("failed to update solutions.go: %w", err)
	}

	fmt.Printf("Successfully generated %s\n", dayDir)
	fmt.Printf("Created: %s\n", dayPath)
	fmt.Printf("Updated: solutions/solutions.go\n")

	return nil
}

func updateSolutionsFile(day int, dayDir string) error {
	solutionsPath := "solutions/solutions.go"

	content, err := os.ReadFile(solutionsPath)
	if err != nil {
		return fmt.Errorf("failed to read solutions.go: %w", err)
	}

	solutionsContent := string(content)

	importLine := fmt.Sprintf("\t\"github.com/milindmadhukar/Advent-Of-Code/2025/golang/%s\"", dayDir)
	importRegex := regexp.MustCompile(`(import \(\n(?:.*\n)*?)(\t"github\.com/milindmadhukar/Advent-Of-Code/2025/golang/models"\n\))`)

	if importRegex.MatchString(solutionsContent) {
		solutionsContent = importRegex.ReplaceAllString(solutionsContent, fmt.Sprintf("$1%s\n$2", importLine))
	} else {
		return fmt.Errorf("could not find import section in solutions.go")
	}

	caseLine := fmt.Sprintf("\tcase %d:\n\t\tsolution = %s.Solve()", day, dayDir)

	caseRegex := regexp.MustCompile(`(\tcase \d+:\n\t\tsolution = day\d+\.Solve\(\)\n)(\t\})`)

	if caseRegex.MatchString(solutionsContent) {
		solutionsContent = caseRegex.ReplaceAllString(solutionsContent, fmt.Sprintf("$1%s\n$2", caseLine))
	} else {
		return fmt.Errorf("could not find switch case section in solutions.go")
	}

	if err := os.WriteFile(solutionsPath, []byte(solutionsContent), 0644); err != nil {
		return fmt.Errorf("failed to write solutions.go: %w", err)
	}

	return nil
}
