package parser

import (
	"galvarezl/mail-indexer/data"
	"regexp"
	"strings"
)

// receive a list of emails in string and for each email get the content in Email struct
func ParseEmails(emails []string) []data.Email {
	var emailList []data.Email
	for _, email := range emails {
		emailList = append(emailList, parseEmail(email))
	}
	return emailList
}

// parse email content from string to Email struct
func parseEmail(email string) data.Email {
	//get the email body
	body := ""
	matched := regexp.MustCompile(`X-FileName:.*(\r\n|\n)`).FindStringIndex(email)
	if matched == nil {
		panic("Is not possible to parse the email to get body content")
	} else {
		body = replaceSpacesAndBreakLines(email[matched[1]:], " ")
	}
	return data.Email{
		MessageID: getValue(regexp.MustCompile(`Message-ID:\s*([^\r\n]+)`).FindStringSubmatch(email)),
		Subject:   getValue(regexp.MustCompile(`Subject:\s*([^\r\n]+)`).FindStringSubmatch(email)),
		Date:      getValue(regexp.MustCompile(`Date:\s*([^\r\n]+)`).FindStringSubmatch(email)),
		From:      getValue(regexp.MustCompile(`From:\s*([^\r\n]+)`).FindStringSubmatch(email)),
		To:        stringToList(getValue(regexp.MustCompile(`To:\s(.+(\n\t.+)+)|To:*([^\r\n]+)`).FindStringSubmatch(email))),
		Cc:        stringToList(getValue(regexp.MustCompile(`Cc:\s(.+(\n\t.+)+)|Cc:*([^\r\n]+)`).FindStringSubmatch(email))),
		Bcc:       stringToList(getValue(regexp.MustCompile(`Bcc:\s(.+(\n\t.+)+)|Bcc:*([^\r\n]+)`).FindStringSubmatch(email))),
		Body:      string(body),
	}
}

func getValue(value []string) string {
	valueResult := ""
	if len(value) == 2 {
		valueResult = value[1]
	}
	if len(value) == 4 {
		if value[3] != "" {
			valueResult = value[3]
		} else {
			valueResult = value[1]
		}
	}
	return strings.TrimSpace(valueResult)
}

// receive a  string or nil value and return a list of strings separte by comma
func stringToList(value string) []string {
	if value == "" {
		return []string{}
	}
	value = replaceSpacesAndBreakLines(value, "")
	return strings.Split(value, ",")
}

func replaceSpacesAndBreakLines(value string, replaceFor string) string {
	value = strings.Replace(value, " ", replaceFor, -1)
	value = strings.Replace(value, "\t", replaceFor, -1)
	value = strings.Replace(value, "\n", replaceFor, -1)
	value = strings.Replace(value, "\r", replaceFor, -1)
	return value
}
