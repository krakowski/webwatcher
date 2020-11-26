package util

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"strings"
	"time"
)

type Watcher struct {
	client *resty.Client
	config Config
}

type PollResult int32

const (
	MATCH    PollResult = 0
	NO_MATCH PollResult = 1

	triggerUrlFormat = "https://maker.ifttt.com/trigger/%s/with/key/%s"
)

func NewWatcher(client *http.Client, config *Config) (*Watcher, error) {
	// Create a new client
	var restyClient *resty.Client
	if client == nil {
		restyClient = resty.New()
	} else {
		restyClient = resty.NewWithClient(client)
	}

	// Create watcher instance
	ret := &Watcher{
		client: restyClient,
		config: *config,
	}

	return ret, nil
}

func (watcher *Watcher) Watch() error  {

	duration, err := time.ParseDuration(watcher.config.Interval)
	if err != nil {
		return err
	}

	log.Printf("Starting to watch website...")
	for range time.Tick(duration) {
		result, err := watcher.poll()
		if err != nil {
			log.Printf("Polling website at '%s' failed. Reason: %s", watcher.config.Website, err.Error())
			continue
		}

		switch result {
			case MATCH:
				log.Printf("Found match. Sending out notification...")
				if err := watcher.notify(); err != nil {
					log.Printf("Notification failed. Reason: %s", err.Error())
				}
				return nil
			case NO_MATCH:
				log.Print("No match found.")
		}
	}

	return nil
}

func (watcher *Watcher) poll() (PollResult, error) {
	response, err := watcher.client.R().Get(watcher.config.Website)
	if err != nil {
		return NO_MATCH, err
	}

	if response.IsError() {
		return NO_MATCH, fmt.Errorf("website %s replied with unexpected status code %d", watcher.config.Website, response.StatusCode())
	}

	body := string(response.Body())
	if !strings.Contains(body, watcher.config.Check) {
		return NO_MATCH, fmt.Errorf("'%s' was not found within the website's body", watcher.config.Check)
	}

	for _, keyword := range(watcher.config.Keywords) {
		if strings.Contains(body, keyword) {
			return MATCH, nil
		}
	}

	return NO_MATCH, nil
}

func (watcher *Watcher) notify() error {
	triggerUrl := fmt.Sprintf(triggerUrlFormat, watcher.config.Trigger.Event, watcher.config.Trigger.Key)
	response, err := watcher.client.R().Post(triggerUrl)
	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("maker replied with unexpected status code %d", response.StatusCode())
	}

	return nil;
}



