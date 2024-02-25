package spring

type ResponseMetadata struct {
	BootVersion  BootVersion            `json:"bootVersion"`
	Dependencies RootObjectDependencies `json:"dependencies"`
	Type         Type                   `json:"type"`
	JavaVersion  BootVersion            `json:"javaVersion"`
}

type BootVersion struct {
	Default string             `json:"default"`
	Type    string             `json:"type"`
	Values  []BootVersionValue `json:"values"`
}

type BootVersionValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RootObjectDependencies struct {
	Type   string              `json:"type"`
	Values []DependenciesValue `json:"values"`
}

type DependenciesValue struct {
	Name   string       `json:"name"`
	Values []ValueValue `json:"values"`
}

type ValueValue struct {
	Description  string      `json:"description"`
	ID           string      `json:"id"`
	Links        *ValueLinks `json:"_links,omitempty"`
	Name         string      `json:"name"`
	VersionRange *string     `json:"versionRange,omitempty"`
}

type ValueLinks struct {
	Guide     *Guide          `json:"guide"`
	Home      *Home           `json:"home,omitempty"`
	Reference *ReferenceUnion `json:"reference"`
	Sample    *Home           `json:"sample,omitempty"`
}

type Home struct {
	Href  string  `json:"href"`
	Title *string `json:"title,omitempty"`
}

type ReferenceClass struct {
	Href      string  `json:"href"`
	Templated *bool   `json:"templated,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type GradleBuildClass struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type Type struct {
	Default string      `json:"default"`
	Values  []TypeValue `json:"values"`
}

type TypeValue struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Tags        Tags   `json:"tags"`
}

type Tags struct {
	Build   string  `json:"build"`
	Dialect *string `json:"dialect,omitempty"`
	Format  string  `json:"format"`
}

type Guide struct {
	Home      *Home
	HomeArray []Home
}

type ReferenceUnion struct {
	HomeArray      []Home
	ReferenceClass *ReferenceClass
}

type QueryString struct {
	Type         string
	Language     string
	BootVersion  string
	BasicDir     string
	GroupId      string
	ArtifactId   string
	Name         string
	Description  string
	PackageName  string
	Packaging    string
	JavaVersion  string
	Dependencies string
}
