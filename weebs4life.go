package weebs4life

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Weeb struct, this is the struct for it, at the moment there's no use for that 'key', but in the future there it'll be needed, due to the owner of weebs4life stating that they're gonna be adding api-key only endpoints
type Weeb struct {
	KEY string
}

//NewWeeb create a new instance of WeEb
func NewWeeb() Weeb {
	return Weeb{}
}

//Hug get hugged owo
func (w Weeb) Hug() (*Aweeb, error) {
	var theWeeb Aweeb
	resp, err := http.Get(weebURL + "/hug")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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
	defer resp.Body.Close()
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
	defer resp.Body.Close()
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
	defer resp.Body.Close()
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
