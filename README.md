# Go-reloaded
## Project Overview
This project is a Go-based text completion/editing/auto-correction tool that processes an input text file according to specified modification rules and generates an output file with the modified text. The tool performs various text manipulations such as converting hexadecimal and binary numbers to decimal, changing word cases, and formatting punctuations.

## Objectives
1. Use existing functions from previous repositories to create a text editing tool
2. Implement various text modifications, including:
  - Converting hexadecimal and binary numbers to decimal
- Changing case of words (uppercase, lowercase, capitalized)
- Correcting punctuation and spacing
- Replacing "a" with "an" when necessary
3. Write unit tests for the tool
4. Correct and audit other students' projects

## Allowed Packages
Standard Go packages are allowed

## How to Run
1. Clone the repository:
```bash
 git clone https://learn.zone01kisumu.ke/git/allkamau/go-reloaded.git
```
2. Navigate to the project directory:
```bash
 cd go-reloaded
 ```
 3. Install any necessary dependencies (if applicable):
 ```bash
go mod tidy
```
4. Run the tool using go run . <input_file> <output_file>
5. Example: 
```bash
go run . sample.txt result.txt
```

## Example

Given the following sample.txt:
```bash
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```
Running the command:
```bash
go run main.go sample.txt result.txt
```
will process sample.txt and produce result.txt with the modifications applied.
- cat result.txt
```bash
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```
## Modification Rules

The tool applies the following modifications:

   1. Hexadecimal Conversion: Every instance of (hex) replaces the preceding word with its decimal equivalent.
   2. Binary Conversion: Every instance of (bin) replaces the preceding word with its decimal equivalent.
   3.  Uppercase Conversion: Every instance of (up) converts the preceding word to uppercase.
   4.  Lowercase Conversion: Every instance of (low) converts the preceding word to lowercase.
   5.  Capitalization: Every instance of (cap) capitalizes the preceding word.
   6.  Punctuation Formatting:
-  Move punctuation marks to the end of the preceding word if they are not already there.
 -   Format punctuation marks such as ,, ., !, ?, :, and ; with proper spacing.
 -   Handle special cases like ellipses (...) and exclamation marks followed by a question mark (!?).

## Development

If you wish to contribute or modify the code:
1. Create a feature branch:
```bash
git checkout -b <feature-branch>
```
2. Make your changes and commit them:
```bash
git add .
git commit -m "Describe your changes"
```
3. Push the branch and create a pull request:
```bash
git push origin <feature-branch>
```
## Testing

The project currently does not have automated tests. You can manually test the functionality by creating input files with various cases and checking the output files.
##  Learning Outcomes
This project will help you learn about:
- The Go file system (fs) API
- String and numbers manipulation
