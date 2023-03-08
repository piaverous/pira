package pira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/piaverous/pira/pira/types"
	"github.com/piaverous/pira/pira/utils"
)

// Get the value of a given field in an object.
// Function taken from https://stackoverflow.com/a/66470232/10494684.
func getAttr(obj interface{}, fieldName string) (reflect.Value, error) {
	var curField reflect.Value
	pointToStruct := reflect.ValueOf(obj) // addressable
	if pointToStruct.Kind() != reflect.Interface || pointToStruct.Kind() != reflect.Pointer {
		return curField, fmt.Errorf("could not get attribute '%s' from object", fieldName)
	}
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		return curField, fmt.Errorf("could not get attribute '%s' from object : not a struct", fieldName)
	}
	curField = curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		return curField, fmt.Errorf("object does not have an attribute named '%s'", fieldName)
	}
	return curField, nil
}

func (app *App) GetJiraIssue(issueId string) (types.JiraIssue, error) {
	var cResp types.JiraIssue

	// 1. Build URL for Jira API call
	jiraUrl, err := url.JoinPath(app.Config.Jira.Url, "rest/api/latest", "issue", issueId)
	if err != nil {
		return cResp, err
	}

	// 2. Make authenticated API request
	req, err := http.NewRequest("GET", jiraUrl, nil)
	if err != nil {
		return cResp, err
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", app.Config.Jira.Token)},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cResp, err
	}

	// 3. Parse response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cResp, err
	}

	if err := json.Unmarshal(body, &cResp); err != nil {
		return cResp, err
	}

	var reference types.JiraIssueWithJSONFields
	if err := json.Unmarshal(body, &reference); err != nil {
		return cResp, err
	}

	// 4. Parse additional Custom Fields
	if err := utils.InjectCustomFieldsFromJSON(app.Config.Jira.CustomFields, reference, &cResp); err != nil {
		return cResp, err
	}

	return cResp, nil
}

func (app *App) ListJiraIssues(sprint string) (types.JiraIssueList, error) {
	var cResp types.JiraIssueList

	// 1. Build URL for Jira API call
	jiraUrl, err := url.JoinPath(app.Config.Jira.Url, "rest/api/latest", "search")
	if err != nil {
		return cResp, err
	}

	// 2. Make authenticated API request
	req, err := http.NewRequest("GET", jiraUrl, nil)
	if err != nil {
		return cResp, err
	}

	sprintFilterQuery := fmt.Sprintf("project=%s&sprint in (\"Sprint %s\")", app.Config.Jira.ProjectKey, sprint)
	if sprint == "" {
		sprintFilterQuery = fmt.Sprintf("project=%s&sprint in openSprints()&sprint not in futureSprints()", app.Config.Jira.ProjectKey)
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", app.Config.Jira.Token)},
	}
	queryParams := req.URL.Query()
	queryParams.Add("maxResults", app.Config.Jira.RequestMaxResults)
	queryParams.Add("jql", sprintFilterQuery)
	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cResp, err
	}

	// 3. Parse response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cResp, err
	}

	if err := json.Unmarshal(body, &cResp); err != nil {
		return cResp, err
	}

	var reference types.JiraIssueListWithJSONFields
	if err := json.Unmarshal(body, &reference); err != nil {
		return cResp, err
	}

	// 4. Parse additional Custom Fields
	if err := utils.InjectCustomFieldsFromJSONList(app.Config.Jira.CustomFields, reference, &cResp); err != nil {
		return cResp, err
	}

	return cResp, nil
}

func (app *App) StoryPointsFromIssue(issue types.JiraIssue) (int, error) {
	storyPointsFieldId := app.Config.Jira.SprintConfig.StoryPointFieldId
	if strings.Contains(storyPointsFieldId, "customfield") {
		for _, cField := range issue.Fields.CustomFields {
			if cField.Id == storyPointsFieldId {
				i, err := strconv.Atoi(cField.Value)
				if err != nil {
					return 0, err
				}
				return i, nil
			}
		}
		return 0, fmt.Errorf("unknown custom field : '%s'", storyPointsFieldId)
	} else {
		value, err := getAttr(issue, storyPointsFieldId)
		if err != nil {
			return 0, err
		}
		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(value.Int()), nil
		case reflect.String:
			return 0, fmt.Errorf("attribute '%s' is a string, expected a number", storyPointsFieldId)
		default:
			return 0, fmt.Errorf("attribute '%s' is not a number", storyPointsFieldId)
		}
	}
}
