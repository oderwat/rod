// The .github/workflows/check-issues.yml will use it as an github action
// To test it locally, you can generate a personal github token: https://github.com/settings/tokens
// Then run this:
//   ROD_GITHUB_ROBOT=your_token go run ./lib/utils/check-issue

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

var githubToken = os.Getenv("ROD_GITHUB_ROBOT")

func main() {
	id, body := getIssue()

	log.Println("check issue", id)

	msg := check(body)

	if msg != "" {
		sendComment(id, msg)
	}
}

func check(body string) string {
	msg := []string{}

	for _, check := range []func(string) error{
		checkVersion, checkMarkdown, checkGoCode,
	} {
		err := check(body)
		if err != nil {
			msg = append(msg, err.Error())
		}
	}
	if len(msg) != 0 {
		return strings.Join(msg, "\n\n")
	}
	return ""
}

func getIssue() (int, string) {
	data, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	utils.E(err)

	issue := gson.New(data).Get("issue")

	for _, l := range issue.Get("labels").Arr() {
		name := l.Get("name").Str()
		if name != "question" && name != "bug" {
			log.Println("skip", name)
			os.Exit(0)
		}
	}

	id := issue.Get("number").Int()
	body := issue.Get("body").Str()

	return id, body
}

func sendComment(id int, msg string) {
	msg += fmt.Sprintf(
		"\n\n_<sub>generated by [check-issue](https://github.com/go-rod/rod/actions/runs/%s)</sub>_",
		os.Getenv("GITHUB_RUN_ID"),
	)

	q := req(fmt.Sprintf("/repos/go-rod/rod/issues/%d/comments", id))
	q.Method = http.MethodPost
	q.Body = ioutil.NopCloser(bytes.NewBuffer(utils.MustToJSONBytes(map[string]string{"body": msg})))
	res, err := http.DefaultClient.Do(q)
	utils.E(err)
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode >= 400 {
		str, err := ioutil.ReadAll(res.Body)
		utils.E(err)
		panic(string(str))
	}
}
