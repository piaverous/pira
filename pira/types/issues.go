package types

// An object reprensenting an Issue in Jira.
type JiraIssue struct {
	Key    string
	Fields JiraIssueFields
}

// An object returned from a request to list Jira Issues.
type JiraIssueList struct {
	Total  int         // Number of Issues listed.
	Issues []JiraIssue // List of Issues.
}

// Fields of a Jira Issue.
type JiraIssueFields struct {
	Created      string // Creation date
	Summary      string
	Updated      string // Last update date
	Description  string
	Labels       []string
	CustomFields []JiraCustomField // List of custom fields found for this issue
	Assignee     JiraUser
	Resolution   JiraResolution
	Status       JiraStatus
}

// A user of the Jira instance
type JiraUser struct {
	Name         string
	EmailAddress string
	DisplayName  string
	Active       bool
	Timezone     string
}

// A representation of the status of a given Jira Issue.
type JiraStatus struct {
	Name           string
	Id             string
	Description    string
	StatusCategory JiraStatusCategory
}

// A status category in Jira. For example: "ongoing" or "to-do".
type JiraStatusCategory struct {
	Name string
	Id   int
}

// A representation of the resolution level of an Issue in Jira.
type JiraResolution struct {
	Name        string
	Id          string
	Description string
}

// Custom field representation
type JiraCustomField struct {
	Id    string
	Alias string
	Value string
	Type  string
}

// An object representing a partially parsed Issue.
//
// This is used in order to inject custom fields into a JiraIssue object.
// Custom fields have different IDs and significations depending on how
// a Jira instance is setup, so they prove difficult handle through static
// types.
type JiraIssueWithJSONFields struct {
	Key    string
	Fields map[string]interface{}
}

// An object returned from a request to list Jira Issues, in which issues were
// only partially parsed.
type JiraIssueListWithJSONFields struct {
	Total  int                       // Number of Issues listed.
	Issues []JiraIssueWithJSONFields // List of Issues.
}
