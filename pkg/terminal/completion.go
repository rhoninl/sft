package terminal

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var commonCommands = []string{
	"ls", "cd", "pwd", "mkdir", "rm", "cp", "mv", "cat", "grep",
	"echo", "touch", "chmod", "chown", "find", "ps", "kill",
	"top", "df", "du", "tar", "zip", "unzip", "ssh", "scp",
	"curl", "wget", "ping", "netstat", "ifconfig", "clear",
	// Add more common commands...
}

func GetCompletions(partial string, currentDir string) []string {
	if partial == "" {
		return commonCommands
	}

	// If input starts with ./ or /, or contains /, try path completion
	if strings.HasPrefix(partial, "./") || strings.HasPrefix(partial, "/") || strings.Contains(partial, "/") {
		return getPathCompletions(partial, currentDir)
	}

	// Otherwise, try command completion
	return getCommandCompletions(partial)
}

func getCommandCompletions(partial string) []string {
	var completions []string

	// Special handling for Windows environment
	pathSep := ":"
	if runtime.GOOS == "windows" {
		pathSep = ";"
	}

	// Check executable files in PATH
	paths := strings.Split(os.Getenv("PATH"), pathSep)
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range files {
			name := file.Name()
			// Windows requires checking more extensions
			if runtime.GOOS == "windows" {
				nameWithoutExt := strings.TrimSuffix(name, filepath.Ext(name))
				if strings.HasPrefix(strings.ToLower(nameWithoutExt), strings.ToLower(partial)) {
					completions = append(completions, name)
				}
			} else if strings.HasPrefix(name, partial) {
				if info, err := file.Info(); err == nil && info.Mode()&0111 != 0 {
					completions = append(completions, name)
				}
			}
		}
	}

	// Add completion for common commands
	for _, cmd := range commonCommands {
		if strings.HasPrefix(cmd, partial) {
			completions = append(completions, cmd)
		}
	}

	return uniqueStrings(completions)
}

func getPathCompletions(partial string, currentDir string) []string {
	var completions []string

	// Handle Windows paths
	if runtime.GOOS == "windows" {
		partial = strings.ReplaceAll(partial, "\\", "/")
		currentDir = strings.ReplaceAll(currentDir, "\\", "/")
	}

	// Determine the directory to search
	searchDir := currentDir
	searchPrefix := partial

	if strings.HasPrefix(partial, "/") || (runtime.GOOS == "windows" && strings.Contains(partial, ":")) {
		if runtime.GOOS == "windows" && strings.Contains(partial, ":") {
			// Windows absolute path handling
			searchDir = filepath.Dir(partial)
			searchPrefix = filepath.Base(partial)
		} else {
			searchDir = "/"
			searchPrefix = partial[1:]
		}
	} else if strings.HasPrefix(partial, "./") {
		searchPrefix = partial[2:]
	}

	baseDir := filepath.Dir(filepath.Join(searchDir, searchPrefix))
	prefix := filepath.Base(searchPrefix)

	files, err := os.ReadDir(baseDir)
	if err != nil {
		return completions
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			fullPath := filepath.Join(baseDir, name)
			if file.IsDir() {
				fullPath += "/"
			}
			// Convert to relative path if applicable
			if strings.HasPrefix(partial, "./") {
				if rel, err := filepath.Rel(currentDir, fullPath); err == nil {
					fullPath = "./" + rel
				}
			}
			completions = append(completions, fullPath)
		}
	}

	return completions
}

func uniqueStrings(strs []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, str := range strs {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}
