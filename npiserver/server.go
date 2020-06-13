package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	requestor "github.com/alexanderkarlis/npi/npirequest"
)

const port = "8080"
const welcomeString = `
Welcome to the NPPES NPI API (what a mouthful!) wrapper. Currently, there are two
supported verbs, GET (this page), and POST (which you can post a json of paramaters from
the NPPES website and get a JSON that NPPES would return). Have fun!`
const examplePost = `
After docker container is running locally...
<br /> 
> curl -d '{"number": "1689872012"}' -X POST 'http://localhost:8888/npi'`

var welcomeHTML = fmt.Sprintf(
	`<div>
			<h1>NPPES NPI Wrapper</h1>
			<p>%v</p>
			<h4>Example</h4>
			<p>%v</p>
	</div>
`, welcomeString, examplePost)

// Serve function serves the service to NPPES Registry.
func Serve() {
	http.HandleFunc("/npi", requestHandler)
	log.Println("** Service Started on Port " + port + " ***")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, welcomeHTML)
	case "POST":
		requestObj := new(requestor.NpiRequest)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestObj)

		if err != nil {
			panic(err)
		}
		npiData, err := makeNPPESRequest(requestObj)
		if err != nil {
			panic(err.Error())
		}

		w.Header().Set("Content-type", "application/json")
		fmt.Fprint(w, string(npiData))
	}
}

func makeNPPESRequest(requestObj *requestor.NpiRequest) ([]byte, error) {
	reqString := requestObj.ToRequestString()
	fmt.Println(reqString)

	resp, err := http.Get(reqString)
	if err != nil {
		return []byte{}, err
	}

	j, err := toJSON(resp)
	if err != nil {
		return []byte{}, err
	}

	return j, err
}

func toJSON(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		return []byte{}, err
	}
	j, err := json.MarshalIndent(m, "", "  ")
	return j, err
}
