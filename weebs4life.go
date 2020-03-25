package weebs4life

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Weeb struct, this is the constructor
type Weeb struct {
	KEY string
}

//Hug get hugged owo
func (w Weeb) Hug() (*Aweeb, error) {
	var theWeeb Aweeb
	resp, err := http.Get(weebURL + "/hug")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &theWeeb)
	if err != nil {
		return nil, err
	}
	return &theWeeb, nil
}

//Kiss get kissed uwu
func (w Weeb) Kiss() (*Aweeb, error) {
	var theWeeb Aweeb
	resp, err := http.Get(weebURL + "/kiss")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &theWeeb)
	if err != nil {
		return nil, err
	}
	return &theWeeb, nil
}

//Neko get a neko from weebs for life
func (w Weeb) Neko() (*Aweeb, error) {
	var theWeeb Aweeb
	resp, err := http.Get(weebURL + "/neko")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &theWeeb)
	if err != nil {
		return nil, err
	}
	return &theWeeb, nil
}

//Slap get slapped lmao
func (w Weeb) Slap() (*Aweeb, error) {
	var theWeeb Aweeb
	resp, err := http.Get(weebURL + "/slap")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &theWeeb)
	if err != nil {
		return nil, err
	}
	return &theWeeb, nil
}
