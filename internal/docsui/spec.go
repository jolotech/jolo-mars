package docsui

type DocSpec struct {
	ProductName string    `json:"productName"`
	CompanyName string    `json:"companyName"`
	Description string    `json:"description"`
	BaseURL     string    `json:"baseUrl"`
	Version     string    `json:"version"`
	QuickStart  QuickStart `json:"quickStart"`
	Groups      []Group   `json:"groups"`
}

type QuickStart struct {
	Title    string      `json:"title"`
	Steps    []string    `json:"steps"`
	Examples []CodeBlock `json:"examples"`
}

type CodeBlock struct {
	Title string `json:"title"`
	Lang  string `json:"lang"` // bash, js, etc
	Code  string `json:"code"`
}

type Group struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Sections    []Section `json:"sections"`
}

// type Section struct {
// 	ID          string     `json:"id"`
// 	Title       string     `json:"title"`
// 	Description string     `json:"description,omitempty"`
// 	Endpoints   []Endpoint `json:"endpoints"`
// }

type Section struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`

	// endpoints directly under this section (optional)
	Endpoints   []Endpoint `json:"endpoints,omitempty"`

	// nested folders under this section (optional)
	Children    []Section  `json:"children,omitempty"`
}


type Endpoint struct {
	ID          string         `json:"id"`
	Method      string         `json:"method"`
	Path        string         `json:"path"`
	Summary     string         `json:"summary"`
	Description string         `json:"description,omitempty"`
	Auth        string         `json:"auth"` // none | bearer
	Request     *RequestSpec   `json:"request,omitempty"`
	Responses   []ResponseSpec `json:"responses,omitempty"`
}

type RequestSpec struct {
	ContentType string      `json:"contentType"`
	Example     interface{} `json:"example"`
}

type ResponseSpec struct {
	Status      int         `json:"status"`
	Description string      `json:"description,omitempty"`
	Example     interface{} `json:"example,omitempty"`
}
