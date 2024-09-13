package gotranslate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	Original   string
	Translated string
	From       string
	To         string
}

// javascript "encodeURI()"
// so we embed js to our golang program
func encodeURI(s string) string {
	return url.QueryEscape(s)
}

func Translate(source, from, to string) (Response, error) {
	var text []string
	var result []interface{}

	encodedSource := encodeURI(source)
	api := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" +
		from + "&tl=" + to + "&dt=t&q=" + encodedSource

	r, err := http.Get(api)
	if err != nil {
		return Response{}, fmt.Errorf("error getting translate.googleapis.com")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return Response{}, fmt.Errorf("error reading response body")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return Response{}, fmt.Errorf("error 400 (Bad Request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Response{}, fmt.Errorf("error unmarshalling data")
	}

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		cText := strings.Join(text, "")
		detectedLang := result[2].(string)

		return Response{
			Original:   source,
			Translated: cText,
			From:       detectedLang,
			To:         to,
		}, nil
	} else {
		return Response{}, fmt.Errorf("no translated data in response")
	}
}
