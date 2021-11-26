package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Passport struct {
	BirthYear      int    `validate:"gte=1920,lte=2002"`
	IssueYear      int    `validate:"gte=2010,lte=2020"`
	ExpirationYear int    `validate:"gte=2020,lte=2030"`
	Height         string `validate:"height"`
	HairColor      string `validate:"haircolor"`
	EyeColor       string `validate:"eyecolor"`
	PassportID     string `validate:"number,len=9"`
	CountryID      string
}

var validate *validator.Validate

func validateHeight(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	switch {
	case strings.HasSuffix(s, "cm"):
		cm, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			log.Fatal(err)
		}
		return cm >= 150 && cm <= 193

	case strings.HasSuffix(s, "in"):
		in, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			log.Fatal(err)
		}
		return in >= 59 && in <= 76

	default:

	}
	return false
}

func validateHairColor(fl validator.FieldLevel) bool {
	clr := fl.Field().String()
	if len(clr) != 7 || clr[0] != '#' {
		return false
	}
	_, err := strconv.ParseUint(clr[1:], 16, 24)
	return err == nil
}

func validateEyeColor(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}
	return true
}
func main() {
	var in io.ReadCloser = os.Stdin

	validate = validator.New()
	if err := validate.RegisterValidation("height", validateHeight); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("eyecolor", validateEyeColor); err != nil {
		log.Fatal(err)
	}
	if err := validate.RegisterValidation("haircolor", validateHairColor); err != nil {
		log.Fatal(err)
	}

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

		lines[len(lines)-1] += " " + line
	}

	for _, line := range lines {
		pass := Passport{}
		for _, p := range strings.Split(line, " ") {
			pp := strings.SplitN(p, ":", 2)
			if len(pp) != 2 {
				continue
			}
			key, val := pp[0], pp[1]

			var err error
			switch key {
			case "byr":
				pass.BirthYear, err = strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}

			case "iyr":
				pass.IssueYear, err = strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}

			case "eyr":
				pass.ExpirationYear, err = strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}

			case "hgt":
				pass.Height = val
			case "hcl":
				pass.HairColor = val
			case "ecl":
				pass.EyeColor = val
			case "pid":
				pass.PassportID = val
			case "cid":
				pass.CountryID = val
			default:
				log.Fatal("invalid key " + key)
			}
		}

		if err := validate.Struct(pass); err != nil {
			//fmt.Printf("%v\n%++v\n", err, pass)
			continue
		}
		count++
		fmt.Println(line)
	}

	fmt.Println("Valid passports: ", count)

}
