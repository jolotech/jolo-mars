(() => {

  console.log("DOCS UI JS LOADED");
  document.documentElement.setAttribute("data-docsui", "loaded");

  const $ = (id) => document.getElementById(id);

  const state = {
    spec: null,
    flatEndpoints: [], // for search
  };

  const STORAGE = {
    token: "docsui_bearer_token",
    baseUrl: "docsui_base_url",
  };

  function setHash(hash) {
    if (!hash.startsWith("#")) hash = "#" + hash;
    window.location.hash = hash;
  }

  function escapeHtml(str) {
    return String(str)
      .replaceAll("&", "&amp;")
      .replaceAll("<", "&lt;")
      .replaceAll(">", "&gt;")
      .replaceAll('"', "&quot;")
      .replaceAll("'", "&#039;");
  }

  function formatJson(obj) {
    try { return JSON.stringify(obj, null, 2); } catch { return String(obj); }
  }

  function normalize() {
    const list = [];
    for (const g of state.spec.groups) {
      for (const s of g.sections) {
        for (const e of s.endpoints) {
          list.push({
            ...e,
            groupId: g.id,
            groupTitle: g.title,
            sectionId: s.id,
            sectionTitle: s.title,
            searchText: `${g.title} ${s.title} ${e.method} ${e.path} ${e.summary} ${e.description || ""}`.toLowerCase(),
          });
        }
      }
    }
    state.flatEndpoints = list;
  }

  function renderSidebar(filtered = null) {
    const tree = $("sidebarTree");
    tree.innerHTML = "";

    const groups = state.spec.groups;

    for (const g of groups) {
      const wrap = document.createElement("div");
      wrap.className = "treeGroup";

      const head = document.createElement("div");
      head.className = "treeGroupHeader";
      head.innerHTML = `<strong>${escapeHtml(g.title)}</strong><span class="pill">${g.sections.length} sections</span>`;

      const body = document.createElement("div");
      body.className = "treeGroupBody open";

      head.addEventListener("click", () => body.classList.toggle("open"));

      for (const s of g.sections) {
        const sec = document.createElement("div");
        sec.className = "treeSection";
        sec.innerHTML = `<div class="treeSectionTitle">${escapeHtml(s.title)}</div>`;

        const endpoints = filtered
          ? filtered.filter(x => x.groupId === g.id && x.sectionId === s.id)
          : s.endpoints.map(e => ({...e, groupId: g.id, sectionId: s.id}));

        for (const e of endpoints) {
          const item = document.createElement("div");
          item.className = "treeItem";
          item.innerHTML = `
            <div style="display:flex;flex-direction:column;gap:3px;min-width:0">
              <div style="display:flex;gap:8px;align-items:center">
                <span class="method">${escapeHtml(e.method)}</span>
                <span style="font-size:12px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">${escapeHtml(e.summary || e.id)}</span>
              </div>
              <div class="path" style="white-space:nowrap;overflow:hidden;text-overflow:ellipsis">${escapeHtml(e.path)}</div>
            </div>
          `;
          item.addEventListener("click", () => {
            setHash(e.id);
            renderPage();
          });
          sec.appendChild(item);
        }

        body.appendChild(sec);
      }

      wrap.appendChild(head);
      wrap.appendChild(body);
      tree.appendChild(wrap);
    }
  }

  function quickStartHtml() {
    const q = state.spec.quickStart;
    const steps = (q.steps || []).map(s => `<li>${escapeHtml(s)}</li>`).join("");
    const examples = (q.examples || []).map(ex => `
      <div class="card">
        <div class="h2">${escapeHtml(ex.title)}</div>
        <pre><code>${escapeHtml(ex.code)}</code></pre>
      </div>
    `).join("");

    return `
      <div class="card">
        <h2 class="h2">${escapeHtml(q.title || "Quick Start")}</h2>
        <div class="muted">Base URL: <span class="path">${escapeHtml(state.spec.baseUrl)}</span></div>
        <div style="height:10px"></div>
        <ol class="muted" style="margin:0;padding-left:18px">${steps}</ol>
      </div>
      ${examples}
      <div class="card">
        <div class="h2">Browse</div>
        <div class="muted">Use the sidebar to pick a group, then an endpoint. The sidebar stays open.</div>
      </div>
    `;
  }

  function endpointHtml(e) {
    const req = e.request ? `
      <div class="card">
        <div class="h2">Request</div>
        <div class="muted">Content-Type: <span class="path">${escapeHtml(e.request.contentType || "application/json")}</span></div>
        <div style="height:10px"></div>
        <pre><code>${escapeHtml(formatJson(e.request.example || {}))}</code></pre>
      </div>
    ` : `
      <div class="card">
        <div class="h2">Request</div>
        <div class="muted">No request body documented.</div>
      </div>
    `;

    const responses = (e.responses || []).map(r => `
      <div class="card">
        <div class="h2">Response ${escapeHtml(r.status)}</div>
        <div class="muted">${escapeHtml(r.description || "")}</div>
        ${r.example ? `<div style="height:10px"></div><pre><code>${escapeHtml(formatJson(r.example))}</code></pre>` : ""}
      </div>
    `).join("");

    const tryItOut = `
      <div class="card">
        <div class="h2">Try it out</div>
        <div class="muted">Uses Base URL + Path. Adds Authorization header if needed.</div>

        <div class="tryRow">
          <button class="btn" data-run="${escapeHtml(e.id)}">Run</button>
          <span id="result-${escapeHtml(e.id)}" class="muted"></span>
        </div>

        <div style="height:10px"></div>
        <div class="muted">Request body (JSON):</div>
        <div style="height:6px"></div>
        <textarea id="body-${escapeHtml(e.id)}" style="width:100%;min-height:140px;border-radius:12px;border:1px solid var(--border);background:rgba(2,6,23,0.6);color:var(--text);padding:10px;font-family:var(--mono);font-size:12px;outline:none">${escapeHtml(formatJson((e.request && e.request.example) ? e.request.example : {}))}</textarea>

        <div style="height:10px"></div>
        <div class="muted">Response:</div>
        <div style="height:6px"></div>
        <pre><code id="out-${escapeHtml(e.id)}">{}</code></pre>
      </div>
    `;

    return `
      <div class="card" id="${escapeHtml(e.id)}">
        <div class="endpointHeader">
          <div class="endpointTitle">
            <div class="endpointTitleRow">
              <span class="bigMethod">${escapeHtml(e.method)}</span>
              <span class="bigPath">${escapeHtml(e.path)}</span>
            </div>
            <div class="muted">${escapeHtml(e.summary || "")}</div>
          </div>
          <span class="authTag">Auth: ${escapeHtml(e.auth || "none")}</span>
        </div>

        ${e.description ? `<div style="height:10px"></div><div class="muted">${escapeHtml(e.description)}</div>` : ""}
      </div>

      <div class="grid2">
        ${req}
        ${tryItOut}
      </div>

      ${responses || `
        <div class="card">
          <div class="h2">Responses</div>
          <div class="muted">No responses documented.</div>
        </div>
      `}
    `;
  }

  function getSelectedEndpointId() {
    const hash = (window.location.hash || "").trim();
    if (!hash || hash === "#" || hash === "#quickstart") return "quickstart";
    return hash.replace("#", "");
  }

  function renderPage() {
    const content = $("pageContent");
    const id = getSelectedEndpointId();

    if (id === "quickstart") {
      content.innerHTML = quickStartHtml();
      return;
    }

    const e = state.flatEndpoints.find(x => x.id === id);
    if (!e) {
      content.innerHTML = `
        <div class="card">
          <div class="h2">Not found</div>
          <div class="muted">Endpoint "${escapeHtml(id)}" was not found. Go to Quick Start or pick one from sidebar.</div>
          <div style="height:10px"></div>
          <button class="btn" id="goBackQuick">Quick Start</button>
        </div>
      `;
      const btn = $("goBackQuick");
      if (btn) btn.onclick = () => { setHash("quickstart"); renderPage(); };
      return;
    }

    content.innerHTML = endpointHtml(e);

    // attach run handler
    const runBtn = content.querySelector(`[data-run="${CSS.escape(e.id)}"]`);
    if (runBtn) runBtn.addEventListener("click", () => runEndpoint(e));

    // scroll into view
    const el = document.getElementById(e.id);
    if (el) el.scrollIntoView({ behavior: "smooth", block: "start" });
  }

  async function runEndpoint(e) {
    const token = $("bearerToken").value.trim();
    const baseUrl = $("baseUrl").value.trim() || "";
    const resultEl = document.getElementById(`result-${e.id}`);
    const outEl = document.getElementById(`out-${e.id}`);
    const bodyEl = document.getElementById(`body-${e.id}`);

    resultEl.textContent = "Running...";
    resultEl.className = "muted";
    outEl.textContent = "{}";

    const url = (baseUrl.replace(/\/+$/,"")) + e.path;

    const headers = {};
    if (e.request && e.request.contentType) headers["Content-Type"] = e.request.contentType;
    if ((e.auth || "").toLowerCase() === "bearer" && token) headers["Authorization"] = `Bearer ${token}`;

    let body = undefined;
    if (e.method !== "GET" && e.method !== "DELETE") {
      try {
        const parsed = JSON.parse(bodyEl.value || "{}");
        body = JSON.stringify(parsed);
      } catch {
        resultEl.textContent = "Invalid JSON body";
        resultEl.className = "resultErr";
        return;
      }
    }

    try {
      const res = await fetch(url, {
        method: e.method,
        headers,
        body,
      });

      const text = await res.text();
      let data = text;
      try { data = JSON.parse(text); } catch {}

      outEl.textContent = typeof data === "string" ? data : JSON.stringify(data, null, 2);

      if (res.ok) {
        resultEl.textContent = `OK (${res.status})`;
        resultEl.className = "resultOk";
      } else {
        resultEl.textContent = `ERROR (${res.status})`;
        resultEl.className = "resultErr";
      }
    } catch (err) {
      resultEl.textContent = "Network error";
      resultEl.className = "resultErr";
      outEl.textContent = String(err);
    }
  }

  function wireTopControls() {
    const tokenInput = $("bearerToken");
    const baseUrlInput = $("baseUrl");

    // load saved values
    const savedToken = localStorage.getItem(STORAGE.token);
    const savedBase = localStorage.getItem(STORAGE.baseUrl);
    if (savedToken) tokenInput.value = savedToken;
    if (savedBase) baseUrlInput.value = savedBase;

    tokenInput.addEventListener("input", () => localStorage.setItem(STORAGE.token, tokenInput.value));
    baseUrlInput.addEventListener("input", () => localStorage.setItem(STORAGE.baseUrl, baseUrlInput.value));

    $("goQuickStart").addEventListener("click", () => {
      setHash("quickstart");
      renderPage();
    });
  }

  function wireSearch() {
    const input = $("search");
    input.addEventListener("input", () => {
      const q = input.value.trim().toLowerCase();
      if (!q) {
        renderSidebar(null);
        return;
      }
      const matches = state.flatEndpoints.filter(e => e.searchText.includes(q));
      renderSidebar(matches);
    });
  }

  async function boot() {
    wireTopControls();

    const res = await fetch("/docs/spec.json", { cache: "no-store" });
    state.spec = await res.json();
    normalize();

    renderSidebar(null);
    wireSearch();

    if (!window.location.hash) setHash("quickstart");
    renderPage();

    window.addEventListener("hashchange", renderPage);
  }

  boot().catch((e) => {
    $("pageContent").innerHTML = `
      <div class="card">
        <div class="h2">Failed to load docs</div>
        <pre><code>${escapeHtml(String(e))}</code></pre>
      </div>
    `;
  });
})();
