package parser

import (
	"fmt"
	"pdf_balance_parser/pkg/model/document"
	"regexp"
	"sort"
	"strings"

	"github.com/djimenez/iconv-go"
)

var regexes map[int]map[int]string
var regexesKeys []int

func LoadRegexMap() {
	regexes = make(map[int]map[int]string)
	regexes[0] = map[int]string{0: `(?i)Venda dis`, 1: `[0-9a-zA-ZÀ-ÿ ]*`, 2: `ócios`, 3: `(?i)ócios[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[1] = map[int]string{0: `(?i)IRRF`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().]*`, 2: `gar\)`, 3: `(?i)gar\)[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[3] = map[int]string{0: `(?i)\+Ou`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().]*`, 2: `ais`, 3: `(?i)ais[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[4] = map[int]string{0: `(?i) \Outros`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().#]*`, 2: `nota`, 3: `(?i)nota[0-9, |a-zA-Z]* [C|D]+ `}
	for key := range regexes {
		regexesKeys = append(regexesKeys, key)
	}
	sort.Ints(regexesKeys)
}

func replaceBoundariesRowValues(regex string, line *string) {
	regex = strings.Replace(regex, "\\)", ")", 10)
	*line = strings.Replace(strings.Replace(*line, " |", "", 3), regex, "", 1)
}

func adjustCurrencyFormat(line *string) {
	*line = strings.Replace(*line, ".", "", 10)
	*line = strings.Replace(*line, ",", ".", 10)
}

func ParseDocument(d document.Document) {
	output, _ := iconv.ConvertString(d.Content, "iso-8859-1", "iso-8859-1")
	fmt.Println(output)

	var headerRegex *regexp.Regexp
	var bodyRegex *regexp.Regexp

	for _, key := range regexesKeys {
		headerRegex, _ = regexp.Compile(regexes[key][0] + regexes[key][1] + regexes[key][2])
		for index, match := range headerRegex.FindStringSubmatch(d.Content) {
			fmt.Printf("[%d] %s\n", index, match)
		}
		bodyRegex, _ = regexp.Compile(regexes[key][3])
		for index, match := range bodyRegex.FindStringSubmatch(d.Content) {
			replaceBoundariesRowValues(regexes[key][2], &match)
			adjustCurrencyFormat(&match)
			fmt.Printf("[%d] %s\n", index, strings.Split(match, " "))
		}
	}

}
