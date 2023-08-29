package harness

// Generated by https://quicktype.io

type Organizations []Organization

type Organization struct {
	Org            Org   `json:"org"`
	Created        int64 `json:"created"`
	Updated        int64 `json:"updated"`
	HarnessManaged bool  `json:"harness_managed"`
}

type Org struct {
	Identifier  string           `json:"identifier"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tags        OrganizationTags `json:"tags"`
}

type OrganizationTags struct {
}