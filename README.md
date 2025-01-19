# Prompt Filler CLI (Under Development)

**Note:** This project is still under development and is not the final version. We are actively working on improving and adding more features.

## Overview

Prompt Filler CLI is a command-line interface tool written in Go that helps users fill out templates with placeholders efficiently. The tool scans a given text template for placeholders, prompts the user for input, and substitutes the placeholders with the provided values.

## Features

- Identifies placeholders within a text template.
- Interactive command-line interface for user input.
- Validates user input to ensure fields are not left empty.
- Outputs the completed text with all placeholders filled.

## Installation

1. Make sure you have Go installed on your system. If not, you can download it from [golang.org](https://golang.org/doc/install).
2. Clone the repository:

   ```bash
   git clone https://github.com/m7real/prompt-filler-cli.git
   ```

3. Navigate to the project directory:

   ```bash
   cd prompt-filler-cli
   ```

4. Build the executable:

   ```bash
   go build
   ```

## Usage

1. Run the CLI tool using the executable

2. Follow the prompts to paste the template and fill in each placeholder.

3. The completed template will be printed to the console.

## Example

Given a template like this:

```
[Your Name]
[Your Email Address]

Dear [Recipient's Name],

I am interested in the [Position] position at [Company Name].
```

The CLI will prompt you for each placeholder, and upon completion, output the filled version.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any bug fixes or enhancements.
