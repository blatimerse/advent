package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Passport struct {
	BirthYear      string `validate:"required,gte=1920,lte=2002"`
	IssueYear      string `validate:"required,gte=2010,lte=2020""`
	ExpirationYear string `validate:"required,gte=2020,lte=2030""`
	Height         string `validate:"required,"`
	HairColor      string `validate:"required"`
	EyeColor       string `validate:"required"`
	PassportID     string `validate:"required"`
	CountryID      string
}

var validate *validator.Validate

func main() {
	var in io.ReadCloser = os.Stdin

	validate = validator.New()
	if len(os.Args) > 1 {
		var err error
		in, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal("opening " + os.Args[1])
		}
		defer in.Close()
	}

	fileScanner := bufio.NewScanner(in)

	// start with single empty line
	lines := []string{""}
	count := 0

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if len(line) == 0 {
			// start a new line
			lines = append(lines, "")
			continue
		}

		lines[len(lines)-1] += " "+line
	}

	for _, line := range lines {
		pass := Passport{}
		for _, p := range strings.Split(line, " ") {
			pp := strings.SplitN(p, ":", 2)
			if len(pp) != 2 {
				continue
			}
			key, val := pp[0], pp[1]

			switch key {
			case "byr": pass.BirthYear = val
			case "iyr": pass.IssueYear = val
			case "eyr": pass.ExpirationYear = val
			case "hgt": pass.Height = val
			case "hcl": pass.HairColor = val
			case "ecl": pass.EyeColor = val
			case "pid": pass.PassportID = val
			case "cid": pass.CountryID = val
			default:
				log.Fatal("invalid key " + key)
			}
		}

		if err := validate.Struct(pass); err != nil {
			continue
		}
		count++
	}

	fmt.Println("Valid passports: ", count)

}
