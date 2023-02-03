package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	cp "github.com/otiai10/copy"
)

func main() {
	// Usage: go run main.go -l golang -y 2022 -d 1

	// Parse command line flags
	var lang string
	var year int
	var day int
	flag.StringVar(&lang, "l", "golang", "Language")
	flag.IntVar(&year, "y", 2022, "Year")
	flag.IntVar(&day, "d", 1, "Day")
	flag.Parse()

	// Check that all the flags are set
	if lang == "" || year == 0 || day == 0 {
		log.Fatal("Missing flags")
	}

	// Check if the aocsession environment variable is setn
	aoctoken := os.Getenv("aocsession")
	if aoctoken == "" {
		log.Fatal("aocsession environment variable not set")
	}

	rootDir := getRelativeRoot()
	scriptDir := rootDir + "scripts/createDays/"

	// First we should if we have a template for the language
	_, err := os.Stat(fmt.Sprintf(scriptDir+"templates/%s", lang))
	if os.IsNotExist(err) {
		log.Fatal("No template for language: " + lang)
	}

	dayDir := fmt.Sprintf(rootDir+"%d/%s/day%d", year, lang, day)
	if _, err := os.Stat(dayDir); !os.IsNotExist(err) {
		log.Fatal("Directory already exists")
	}

	log.Printf("Creating directory: %s", dayDir)
	createDir(dayDir)

	// Copy the template to the correct directory
	log.Printf("Copying template to: %s", dayDir)
	err = cp.Copy(fmt.Sprintf(scriptDir+"templates/%s", lang), dayDir)
	if err != nil {
		log.Fatal(err)
	}

	replaceModFile(dayDir, year, lang, day)

	getDayInput(dayDir, aoctoken, year, day)

	parsePrompt(dayDir, aoctoken, year, day)

	log.Printf("Successfully created day %d for year %d in %s (%s)", day, year, lang, dayDir)
}

func getRelativeRoot() string {
	maxDepth := 10
	relativePath := ""

	for {
		if _, err := os.Stat(relativePath + ".git"); os.IsNotExist(err) {
			relativePath += "../"
		} else {
			break
		}

		maxDepth--
		if maxDepth == 0 {
			log.Fatal("Could not find root directory")
		}
	}

	return relativePath
}

func createDir(dir string) {
	segments := strings.Split(dir, "/")
	for i := 0; i < len(segments); i++ {
		if _, err := os.Stat(strings.Join(segments[:i+1], "/")); os.IsNotExist(err) {
			os.Mkdir(strings.Join(segments[:i+1], "/"), 0755)
		}
	}
}

func replaceModFile(dayDir string, year int, lang string, day int) {
	log.Printf("Replacing module path in go.mod")
	modContent, err := os.Open(dayDir + "/go.mod")
	if err != nil {
		log.Fatal(err)
	}

	defer modContent.Close()
	sc := bufio.NewScanner(modContent)
	lines := []string{}

	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "module") {
			// Replace the PATH part with the correct year, language and day
			line = strings.Replace(line, "PATH", fmt.Sprintf("%d/%s/day%d", year, lang, day), 1)
		}

		lines = append(lines, line)
	}

	// Write the new content to the file
	err = ioutil.WriteFile(dayDir+"/go.mod", []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getDayInput(dayDir string, aoctoken string, year int, day int) {
	log.Printf("Getting input for year %d, day %d", year, day)
	// Get the input for the day
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", aoctoken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if len(body) == 0 {
		log.Fatal("No input found")
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/input.txt", dayDir), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func parsePrompt(dayDir string, aoctoken string, year int, day int) {
	log.Printf("Getting prompts for year %d, day %d", year, day)
	// Get the prompt for the day
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", aoctoken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if len(body) == 0 {
		log.Fatal("No prompt found")
	}

	// Parse the prompts
	prompts := getPrompts(string(body))

	// Convert the prompt to markdown
	convertedPrompts := convertPrompts(prompts)

	// Write the prompt to the fil
	writePrompts(dayDir, convertedPrompts)
}

func getPrompts(content string) []string {
	// We want to use regex to get the prompts
	// Each prompt is wrapped in a <article> tag with a class of day-desc

	re := regexp.MustCompile(`<article class="day-desc">([\s\S]*?)<\/article>`)
	matches := re.FindAllStringSubmatch(content, -1)

	prompts := []string{}
	for _, match := range matches {
		prompts = append(prompts, match[1])
	}

	return prompts
}

func convertPrompts(prompts []string) string {
	markdown := md.NewConverter("", true, nil)

	convertedPrompts := ""
	for i, prompt := range prompts {
		converted, err := markdown.ConvertString(prompt)
		if err != nil {
			log.Fatal(err)
		}

		convertedPrompts += converted
		if i != len(prompts)-1 {
			convertedPrompts += "\n\n"
		}
	}

	return convertedPrompts
}

func writePrompts(dayDir string, promptText string) {
	readmeFile, err := os.Create(fmt.Sprintf("%s/README.md", dayDir))
	if err != nil {
		log.Fatal(err)
	}

	defer readmeFile.Close()

	_, err = readmeFile.WriteString(promptText)
	if err != nil {
		log.Fatal(err)
	}
}
