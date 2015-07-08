package duke_nukem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"
)

//------ PACKAGE VARIABLES
var (
	BaseURI         string
	tenDukeUserName string
	tenDukePassword string
	endpointGraph   string
	profileID       string
	timeout         time.Duration
	templateLogin   *template.Template
)

//------ STRUCTS

// Client represents a 10duke API client and associated
// response data.
type Client struct {
	client *http.Client
	// Embedded field for method delegation.
	Result
}

// Result represents a data structure returned from 10duke. It's
// pretty convoluted so we capture all of it. We do NOT like string
// parsing for errors!
type Result struct {
	Code string `json:"resultCode"`
	User struct {
		ProfileId string `json:"id"`
	}
	ErrorMessage struct {
		MessageTemplate string
		Key             string
	}
	ProcessingResult struct {
		ErrorMessage struct {
			MessageTemplate string
			Key             string
		}
		ResultCode string
	}
	ParameterParserResult struct {
		ResultCode string
	}
	ErrorTechnical string
}

//------ PACKAGE SETUP

func init() {
	endpointGraph = "/graph/"
	timeout = time.Duration(10 * time.Second)
}

//------ EXPORTED FUNCTIONS

// NewClient returns a client struct with
// a preset timeout.
func NewClient() Client {
	BaseURI = os.Getenv("TENDUKE_URL")
	tenDukeUserName = os.Getenv("TENDUKE_USERNAME")
	tenDukePassword = os.Getenv("TENDUKE_PASSWORD")
	if BaseURI == "" || tenDukeUserName == "" || tenDukePassword == "" {
		log.Fatal("missing 10duke credentials, please check .env file")
	}

	transport := http.Transport{
		Dial: dialTimeout,
	}

	return Client{
		&http.Client{Transport: &transport},
		Result{},
	}
}

//------ CLIENT METHODS - GET

// GetRequest assembles a GET request to 10duke, processes the
// response and returns any error.
func (duke Client) GetRequest(urlstr string) (err error) {
	req, err := http.NewRequest("GET", urlstr, nil)
	resp, err := duke.client.Do(req)
	return duke.processResponse(resp, "")
}

//------ CLIENT METHODS - POST
//
//// Login2 logs a user into 10duke via POST.
func (duke Client) Login() (err error) {
	// Assemble the registration data.
	data := map[string]string{
		"userName": os.Getenv("TENDUKE_USERNAME"),
		"password": os.Getenv("TENDUKE_PASSWORD"),
	}
	URL, _ := url.ParseRequestURI(BaseURI)
	URL.Path = endpointGraph
	urlStr := fmt.Sprintf("%v", URL)
	// Prepare and send the HTTP request.
	return duke.PostRequest(urlStr, "Login", data)
}

// RegisterUser takes user data and sends it to
// the 10duke system, simulating a POST request via
// the registration form.
func (duke Client) RegisterUser(user User) (err error) {

	// Assemble the registration data. These URL parameters are long,
	// verbose and ugly so hid them away in templates.go.
	data := map[string]string{
		//userProfile:   "{randomUuid,profile}",
		userFirstName: user.FirstName,
		userLastName:  user.LastName,
		userEmail:     user.Email,
		userName:      user.Email,
		userPassword:  user.Password,
		userConfirm:   user.Password,
		//userAccountType: "{randomUuid,account}",
		userAccount: "personal",
		userTsAndCs: "true",
	}

	URL, _ := url.ParseRequestURI(BaseURI)
	URL.Path = endpointGraph
	urlStr := fmt.Sprintf("%v", URL)
	return duke.PostRequest(urlStr, "RegisterUser", data)
}

// PostRequest assembles a POST request to 10duke, processes the
// response and returns any error.
func (duke Client) PostRequest(urlstr, op string, postData map[string]string) (err error) {
	data := url.Values{}
	data.Set("operation", op)

	for name, value := range postData {
		data.Set(name, value)
	}

	encoded := data.Encode()

	req, err := http.NewRequest("POST", urlstr, bytes.NewBufferString(encoded))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))

	resp, err := duke.client.Do(req)
	return duke.processResponse(resp, op)
}

//------ ERROR HANDLING

func (duke Client) processResponse(resp *http.Response, op string) (err error) {
	if resp.Status != "200 OK" {
		return fmt.Errorf("HTTP error: %s\n", resp.Status)
	}
	// Unmarshall the response into the result struct.
	d, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(d), &duke.Result)

	if err != nil {
		return err
	}
	// No errors from the HTTP request or issues processing the
	// response body or marshalling it to JSON. Proceed to inspect
	// the response JSON from 10duke.
	return duke.Response(op)
}

// Response will parse the data structure returned from
// 10duke and return any appropriate error to the caller.
func (duke Client) Response(op string) (err error) {
	if op == "RegisterUser" && duke.Code == "Success" {
		profileID = duke.User.ProfileId
		log.Println("adding entitlement for profile", duke.User.ProfileId)
		return nil
	}
	errmsg := duke.ErrorMessage.MessageTemplate
	prsmsg := duke.ProcessingResult.ResultCode
	if errmsg != "" {
		return fmt.Errorf(
			"processing failed: %s\n",
			errmsg)
	}

	if prsmsg != "" {
		return fmt.Errorf(
			"processing failed: %s %s\n",
			prsmsg, duke.ErrorTechnical)
	}
	return nil
}

//------ UNEXPORTED UTILITY FUNCTIONS

// populateQuery takes a struct and compiles the large query string
// template contained in registerString.
func populateQuery(user interface{}, tmpl *template.Template) string {
	var doc bytes.Buffer
	err := tmpl.Execute(&doc, user)
	if err != nil {
		log.Fatal(err)
	}
	return doc.String()
}

// dialTimeout sets an HTTP timeout for use by the API client.
func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}
