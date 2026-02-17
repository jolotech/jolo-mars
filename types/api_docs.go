package types

type DocSpec struct {
  ProductName string `json:"productName"`
  CompanyName string `json:"companyName"`
  Description string `json:"description"`
  BaseURL     string `json:"baseUrl"`
  Version     string `json:"version"`
  Groups      []Group `json:"groups"`
  QuickStart  QuickStart `json:"quickStart"`
}

type Group struct {
  ID          string     `json:"id"`          // "admin"
  Title       string     `json:"title"`       // "Admin"
  Description string     `json:"description"` // optional
  Sections    []Section  `json:"sections"`    // "Auth", "Management", ...
}

type Section struct {
  ID        string     `json:"id"`        // "admin-auth"
  Title     string     `json:"title"`     // "Auth"
  Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
  ID          string `json:"id"`          // "admin-login"
  Method      string `json:"method"`      // "POST"
  Path        string `json:"path"`        // "/api/v1/admin/login"
  Summary     string `json:"summary"`
  Description string `json:"description"`
  Auth        string `json:"auth"`        // "none" | "bearer"
  Request     *RequestSpec `json:"request,omitempty"`
  Responses   []ResponseSpec `json:"responses,omitempty"`
}

type RequestSpec struct {
  ContentType string `json:"contentType"`
  Example     any    `json:"example"`
}

type ResponseSpec struct {
  Status      int    `json:"status"`
  Description string `json:"description"`
  Example     any    `json:"example"`
}

type QuickStart struct {
  Title string   `json:"title"`
  Steps []string `json:"steps"`
  Examples []CodeBlock `json:"examples"`
}

type CodeBlock struct {
  Title string `json:"title"`
  Lang  string `json:"lang"`  // "bash", "js", etc
  Code  string `json:"code"`
}
