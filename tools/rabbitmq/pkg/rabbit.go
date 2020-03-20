package pkg

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type ConsumerDetail struct {
	ConsumerTag string `json:"consumer_tag"`
	Queue       struct {
		Name  string
		Vhost string
	}
}

func FetchAllConsumer(uri, username, password string) ([]ConsumerDetail, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}

	var consumers []ConsumerDetail
	if err = json.Unmarshal(body, &consumers); err != nil {
		return nil, errors.Wrap(err, "fetchAllConsumer")
	}

	return consumers, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
