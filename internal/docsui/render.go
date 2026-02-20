package docsui

import "html/template"


// pageTmpl is the HTML template for the documentation page. 
// It includes placeholders for product name, version, company name, description, and base URL. 
// The template also references static assets like styles.css and app.js, 
// which are served from the embedded filesystem.

var pageTmpl = template.Must(template.New("docs").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8"/>
  <meta name="viewport" content="width=device-width,initial-scale=1"/>
  <title>{{.ProductName}} Docs</title>

  <link rel="icon" href="/docs/assets/joloFav.png" type="image/png"/>
  <link rel="stylesheet" href="/docs/assets/styles.css"/>
</head>
<body>
  <div id="backdrop" class="backdrop" hidden></div>
  <header class="topbar">

    <div class="brand">
      <img class="logo" src="/docs/assets/joloImg.png" alt="logo"/>
      <div class="brandText">
        <div class="titleRow">
          <h1>{{.ProductName}}</h1>
          <span class="pill">v{{.Version}}</span>
        </div>
        <p class="subtitle">{{.CompanyName}} — {{.Description}}</p>
      </div>
    </div>

    <!-- Updated controls: token + env selector + base url + theme toggle -->
    <div class="topControls">
    <button id="mobileMenuBtn" class="btn small mobileOnly" type="button">Menu</button>

      <div class="tokenWrap">
        <label for="bearerToken">Bearer Token</label>
        <input id="bearerToken" placeholder="paste token here (optional)"/>
      </div>

      <div class="envWrap">
        <label for="envSelect">Environment</label>
        <select id="envSelect">
          <option value="dev">Dev</option>
          <option value="staging" selected>Staging</option>
          <option value="prod">Prod</option>
        </select>
      </div>

      <div class="envWrap">
        <label for="baseUrl">Base URL</label>
        <input id="baseUrl" value="{{.BaseURL}}" />
      </div>
      
    </div>
  </header>

  <div class="layout">
    <aside class="sidebar">
      <div class="sideTop">
        <input id="search" class="search" placeholder="Search endpoints..."/>
        <button id="goQuickStart" class="btn small" type="button">Quick Start</button>
      </div>
      <nav id="sidebarTree" class="tree"></nav>
    </aside>

    <main class="content">
      <div id="pageContent"></div>
    </main>
  </div>

  <div id="toastRoot" class="toastRoot" aria-live="polite" aria-atomic="true"></div>
  <script defer src="/docs/assets/app.js"></script>
</body>
</html>`))