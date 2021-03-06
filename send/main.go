package send

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/absolutedevops/civostatsd/config"
	"github.com/absolutedevops/civostatsd/gather"
)

func ToAPI(configuration config.Config, s gather.Stats) {
	endpoint := configuration.Server + "/civostatsd"
	// Post a form containing config.token as token, s.CPU as cpu, s.Memory as memory and s.Disk as disk

	hc := http.Client{}
	form := url.Values{}
	form.Add("token", configuration.Token)
	form.Add("cpu", fmt.Sprintf("%f", s.CPU))
	form.Add("memory", fmt.Sprintf("%f", s.Memory))
	form.Add("disk", fmt.Sprintf("%f", s.Disk))
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
