package utils

import (
	"encoding/json"
	"fmt"

	"github.com/piaverous/pira/pira/types"
)

func InjectCustomFieldsFromJSON(customFields []types.JiraCustomField, body []byte, targetObject *types.JiraIssue) error {
	var result types.JiraIssueWithJSONFields
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	for _, cField := range customFields {
		value, ok := result.Fields[cField.Id]
		if !ok {
			continue
		}
		if value == nil {
			continue
		}

		customField := types.JiraCustomField{}
		customField.Alias = cField.Alias
		customField.Id = cField.Id
		customField.Value = fmt.Sprintf("%v", value)
		targetObject.Fields.CustomFields = append(targetObject.Fields.CustomFields, customField)
	}
	return nil
}

func InjectCustomFieldsFromJSONList(customFields []types.JiraCustomField, reference types.JiraIssueListWithJSONFields, targetObject *types.JiraIssueList) error {
	var result []types.JiraIssue
	for _, referenceIssue := range reference.Issues {
		for _, targetIssue := range targetObject.Issues {
			if referenceIssue.Key == targetIssue.Key {
				for _, cField := range customFields {
					value, ok := referenceIssue.Fields[cField.Id]
					if !ok {
						continue
					}
					if value == nil {
						continue
					}

					customField := types.JiraCustomField{}
					customField.Alias = cField.Alias
					customField.Id = cField.Id
					customField.Value = fmt.Sprintf("%v", value)
					targetIssue.Fields.CustomFields = append(targetIssue.Fields.CustomFields, customField)
					result = append(result, targetIssue)
				}
			}
		}
	}
	targetObject.Issues = result
	return nil
}
