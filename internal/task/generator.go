package task

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Data struct {
	Raw   string
	URL   string
	Level string
}

func (d *Data) DirName() string {
	return strings.ReplaceAll(d.Raw, "-", "_")
}

func (d *Data) PackageName() string {
	res := strings.ReplaceAll(d.Raw, "-", "")
	res = strings.ReplaceAll(res, "_", "")
	return strings.ToLower(res)
}

func (d *Data) CamelCaseRaw() string {
	parts := strings.Split(d.Raw, "-")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func ValidateLevel(level string) error {
	switch level {
	case "easy", "medium", "hard":
	default:
		return fmt.Errorf("unknown level: %s", level)
	}
	return nil
}

func ParseURL(leetcodeURL string, data *Data) error {
	parsedURL, err := url.Parse(leetcodeURL)
	if err != nil {
		return err
	}
	if parsedURL.Host != "leetcode.com" {
		return fmt.Errorf("host %s is not leetcode.com", parsedURL.Host)
	}
	if parsedURL.Scheme != "https" {
		return fmt.Errorf("wrong scheme: %s", parsedURL.Scheme)
	}

	re := regexp.MustCompile(`^/problems/([a-z0-9-]+)/?`)
	matches := re.FindStringSubmatch(parsedURL.Path)
	if len(matches) < 2 {
		return fmt.Errorf("URL has wrong path: %s", parsedURL.Path)
	}

	data.URL = leetcodeURL
	data.Raw = matches[1]

	return nil
}

func CreateScaffold(data *Data) error {
	dirName := data.DirName()
	dirPath := filepath.Join(data.Level, dirName)
	pkgName := data.PackageName()

	if _, err := os.Stat(dirPath); !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("directory %s already exists", dirPath)
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	// 1. Генерируем solution.go
	solutionContent := fmt.Sprintf("// %s\npackage %s\n", data.URL, pkgName)
	solutionPath := filepath.Join(dirPath, "solution.go")
	if err := os.WriteFile(solutionPath, []byte(solutionContent), 0644); err != nil {
		return fmt.Errorf("failed to write solution.go: %w", err)
	}

	// Формируем полный путь до тестируемого пакета
	// Например: github.com/nepridumalnik/leetgo/easy/greatest_common_divisor_of_strings
	targetImport := fmt.Sprintf("github.com/nepridumalnik/leetgo/%s/%s", data.Level, dirName)

	// 2. Генерируем solution_test.go с вашим форматом импортов
	testContent := fmt.Sprintf(`package %s_test

import (
	"fmt"
	"testing"

	"%s"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip bool
}

func Test_%s(t *testing.T) {
	tests := []testCase{}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}
		})
	}
}
`, pkgName, targetImport, data.CamelCaseRaw())

	testPath := filepath.Join(dirPath, "solution_test.go")
	if err := os.WriteFile(testPath, []byte(testContent), 0644); err != nil {
		return fmt.Errorf("failed to write solution_test.go: %w", err)
	}

	fmt.Printf("✓ Успешно создана структура в директории: %s/\n", dirPath)
	return nil
}
