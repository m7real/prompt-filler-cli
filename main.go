package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	coverLetter := `
[Your Name]  
[Your Address]  
[City, State, ZIP]  
[Your Email Address]  
[Your Phone Number]  

Dear [Recipient's Name],  

I hope this message finds you well. I am reaching out regarding [specific purpose or topic]. With my background in [relevant skills or field], I believe I can [specific contribution or goal].  

Thank you for your time and consideration. I look forward to discussing [topic or opportunity] further.  

Best regards,  
[Your Name]`

	re := regexp.MustCompile(`\[(.*?)\]`)

	matches := re.FindAllStringSubmatch(coverLetter, -1)

	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)

	for _, match := range matches {
		fmt.Println(match[1], ":")
		scanner.Scan()
		text := scanner.Text()

		inputs = append(inputs, text)
	}
	for idx, val := range inputs {
		coverLetter = strings.Replace(coverLetter, matches[idx][0], val, 1)
	}
	fmt.Println(coverLetter)
}
