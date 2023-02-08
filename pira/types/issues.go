package types

type JiraIssue struct {
	Key    string
	Fields JiraIssueFields
}

type JiraIssueList struct {
	StartAt int
	Total   int
	Issues  []JiraIssue
}

type JiraIssueFields struct {
	Created      string
	Summary      string
	Updated      string
	Description  string
	Labels       []string
	CustomFields []JiraCustomField
	Assignee     JiraAssignee
	Resolution   JiraResolution
	Status       JiraStatus
}

type JiraAssignee struct {
	Name         string
	EmailAddress string
	DisplayName  string
	Active       bool
	Timezone     string
}

type JiraStatus struct {
	Name           string
	Id             string
	Description    string
	StatusCategory JiraStatusCategory
}
type JiraStatusCategory struct {
	Name string
	Id   int
}

type JiraResolution struct {
	Name        string
	Id          string
	Description string
}

type JiraCustomField struct {
	Id    string
	Alias string
	Value string
	Type  string
}

type JiraIssueWithUnknownFields struct {
	Key    string
	Fields map[string]interface{}
}
type JiraIssueListWithUnknownFields struct {
	StartAt int
	Total   int
	Issues  []JiraIssueWithUnknownFields
}
