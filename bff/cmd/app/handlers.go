package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (app *application) getApiContent(url string) (*http.Response, error) {
	app.infoLog.Printf("Getting content from %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (app *application) postApiContent(url string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	app.infoLog.Printf("Posting content %v to %s\n", data, url)
	_, err = http.Post(url, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return err
	}

	return nil
}

func (app *application) putApiContent(url string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	app.infoLog.Printf("Putting content %v to %s\n", data, url)
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

func (app *application) deleteApiContent(url string) error {
	app.infoLog.Printf("Deleting content from %s\n", url)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}