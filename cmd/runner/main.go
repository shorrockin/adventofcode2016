package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("expected 1 parameter which should be the name of the day package you want to create, i.e. 'day07', arg: %+v\n", args)
	}
	if !strings.HasPrefix(args[1], "day") {
		log.Fatalf("expected the argument parameter to start with the string day\n")
	}

	day := args[1]
	directory := fmt.Sprintf("./pkg/%v", day)
	fmt.Printf("creating pkg: %v\n", directory)

	info, err := os.Stat(directory)
	if err == nil {
		log.Fatalf("expected directory to not exist, but had os.Stat value: %+v\n", info)
	}
	if !os.IsNotExist(err) {
		log.Fatalf("expected directory be not existing, but was: %+v\n", err)
	}

	fmt.Println("  - generating solve.go")
	if err = template("./cmd/runner/solve_template.go.txt", fmt.Sprintf("%v/solve.go", directory), day); err != nil {
		log.Fatalf("unable to create solve.go: %v", err)
	}

	fmt.Println("  - generating solve_test.go")
	if err = template("./cmd/runner/solve_template_test.go.txt", fmt.Sprintf("%v/solve_test.go", directory), day); err != nil {
		log.Fatalf("unable to create solve_test.go: %v", err)
	}

	fmt.Println("  - generating input.example.txt")
	if err = template("./cmd/runner/input.example.txt", fmt.Sprintf("%v/input.example.txt", directory), day); err != nil {
		log.Fatalf("unable to create input.example.txt: %v", err)
	}
}

func template(src, dst, day string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	modifiedContent := string(content)
	modifiedContent = strings.ReplaceAll(modifiedContent, "{{DAY}}", day)

	dstDir := dst[:strings.LastIndex(dst, "/")]
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	err = os.WriteFile(dst, []byte(modifiedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write destination file: %w", err)
	}

	return nil
}
