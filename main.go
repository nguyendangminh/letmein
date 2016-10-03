package main 

import (
	"os"
	"fmt"
	"log"
	"regexp"
    "strings"
	"net/http"
	"github.com/nlopes/slack"
)

var port string
var slackToken string
var teamName string
var firstName string = "Super"
var lastName string = "Man"
var api *slack.Client
var inviteTmpl string = `
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="keyword" content="Slack, letmein, let me in">
    <meta name="description" content="Slack, let me in. A Slack inviter">
    <meta name="author" content="minhnd.com">
    <title>Slack, let me in!</title>
    <link href="http://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.6/css/materialize.min.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  </head>
  <body>
    <div class="container">
        <div class="section">
            <div class="row">
                <div class="col s12 center cyan-text darken-4-text">
                  <h3 class="header">Slack, let me in</h3>
                </div>
            </div>
            %s
            <div class="row">
                <div class="col s6 offset-s3">
                <form method="POST" action="/">
                    <div class="input-field">
                        <i class="material-icons prefix">email</i>
                        <input id="email" name="email" type="text">
                        <label for="email">Enter your email</label>
                    </div>
                    <button class="btn waves-effect waves-light red darken-2 right" type="submit" name="action">Next
                        <i class="material-icons right">send</i>
                    </button>
                </form>
                </div>
            </div>
        </div>
    </div>
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.6/js/materialize.min.js"></script>
    <script src='https://www.google.com/recaptcha/api.js'></script>
  </body>
</html>
`
var successMessage string = `
            <div class="row">
                <div class="col s6 offset-s3">
                    <div class="card light-blue darken-3">
                        <div class="card-content white-text">
                            <p>
                            Invitation already sent, check your email!
                            </p>
                        </div>
                    </div>
                </div>
            </div>
`
var errorMessage string = `
            <div class="row">
                <div class="col s6 offset-s3">
                    <div class="card red darken-3">
                        <div class="card-content white-text">
                            <p>
                            %s
                            </p>
                        </div>
                    </div>
                </div>
            </div>
`
var errorEmail string = "Please enter a valid email"
var errorServer string = "Something wrong happended, please try again later"
var errorInvited string = "Invited this email already. Please check your mailbox."

func inviteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, fmt.Sprintf(inviteTmpl, ""))
		return
	}

	r.ParseForm()
	email := r.Form.Get("email")

	var validEmail = regexp.MustCompile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
	if !validEmail.MatchString(email) {
		msg := fmt.Sprintf(errorMessage, errorEmail)
		fmt.Fprintf(w, fmt.Sprintf(inviteTmpl, msg))
		return
	}

	err := api.InviteToTeam(teamName, firstName, lastName, email)
    if err != nil {
            log.Println(err.Error())
            var msg string
            if isInvitedError(err) {
                    msg = fmt.Sprintf(errorMessage, errorInvited)
            } else {
                    msg = fmt.Sprintf(errorMessage, errorServer)
            }
            fmt.Fprintf(w, fmt.Sprintf(inviteTmpl, msg))
            return
    }

	fmt.Fprintf(w, fmt.Sprintf(inviteTmpl, successMessage))
}

func isInvitedError(err error) bool {
        return strings.Contains(err.Error(), "already_invited")
}

func init() {
	port = os.Getenv("PORT")
	slackToken = os.Getenv("SLACK_TOKEN")
	teamName = os.Getenv("TEAM_NAME")
	if (port=="") || (slackToken=="") || (teamName=="") {
		log.Fatal("You must set all environment variables: PORT, SLACK_TOKEN, TEAM_NAME")
	}
	if os.Getenv("FIRST_NAME") != "" {firstName = os.Getenv("FIRST_NAME")}
	if os.Getenv("LAST_NAME") != "" {lastName = os.Getenv("LAST_NAME")}
}

func main() {
	api = slack.New(slackToken)
	http.HandleFunc("/", inviteHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}