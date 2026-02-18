(() => {

  console.log("DOCS UI JS LOADED");
  document.documentElement.setAttribute("data-docsui", "loaded");

  const $ = (id) => document.getElementById(id);

  const state = {
    spec: null,
    flatEndpoints: [], // for search
  };

  const fileState = {}; // endpointId -> File


  // const STORAGE = {
  //   token: "docsui_bearer_token",
  //   baseUrl: "docsui_base_url",
  // };
  const ENV = {
    dev: "http://localhost:8023",
    staging: "https://staging.jolojolo.com",
    prod: "https://api.jolojolo.com",
  };

  const STORAGE = {
    token: "docsui_bearer_token",
    baseUrl: "docsui_base_url",
    // theme: "docsui_theme",
    env: "docsui_env",
  };


  function setHash(hash) {
    if (!hash.startsWith("#")) hash = "#" + hash;
    window.location.hash = hash;
  }

  function toastSuccess(title, msg) {
  const root = document.getElementById("toastRoot");
  if (!root) return;

  const toast = document.createElement("div");
  toast.className = "toast";
  toast.innerHTML = `
    <div class="checkWrap" aria-hidden="true">
      <svg class="checkIcon" viewBox="0 0 24 24">
        <path d="M20 6L9 17l-5-5"></path>
      </svg>
    </div>
    <div style="display:flex;flex-direction:column">
      <div class="toastTitle">${escapeHtml(title || "Copied")}</div>
      <div class="toastMsg">${escapeHtml(msg || "")}</div>
    </div>
  `;

  root.appendChild(toast);

  setTimeout(() => {
    toast.classList.add("hide");
    setTimeout(() => toast.remove(), 200);
  }, 1400);
}

// function selectCodeInPre(preEl) {
//   try {
//     const range = document.createRange();
//     range.selectNodeContents(preEl);
//     const sel = window.getSelection();
//     sel.removeAllRanges();
//     sel.addRange(range);
//   } catch {}
// }

// Clipboard API with fallback for older/blocked contexts
async function copyTextReliable(text) {
  // Try modern clipboard API first
  if (navigator.clipboard && window.isSecureContext) {
    await navigator.clipboard.writeText(text);
    return true;
  }

  // Fallback: hidden textarea + execCommand
  const ta = document.createElement("textarea");
  ta.value = text;
  ta.setAttribute("readonly", "");
  ta.style.position = "fixed";
  ta.style.top = "-1000px";
  ta.style.left = "-1000px";
  document.body.appendChild(ta);
  ta.focus();
  ta.select();

  try {
    const ok = document.execCommand("copy");
    document.body.removeChild(ta);
    return ok;
  } catch {
    document.body.removeChild(ta);
    return false;
  }
}

function copyIconSvg() {
  // simple “copy” icon (two rectangles)
  return `
    <svg viewBox="0 0 24 24" aria-hidden="true">
      <rect x="9" y="9" width="10" height="10" rx="2"></rect>
      <path d="M7 15H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v1"></path>
    </svg>
  `;
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

  // function normalize() {
  //   const list = [];
  //   for (const g of state.spec.groups) {
  //     for (const s of g.sections) {
  //       for (const e of s.endpoints) {
  //         list.push({
  //           ...e,
  //           groupId: g.id,
  //           groupTitle: g.title,
  //           sectionId: s.id,
  //           sectionTitle: s.title,
  //           searchText: `${g.title} ${s.title} ${e.method} ${e.path} ${e.summary} ${e.description || ""}`.toLowerCase(),
  //         });
  //       }
  //     }
  //   }
  //   state.flatEndpoints = list;
  // }

  function normalize() {
    const list = [];
    function walkSection(group, section) {
      // endpoints in this section
      for (const e of (section.endpoints || [])) {
        list.push({
           ...e,
           groupId: group.id,
           groupTitle: group.title,
           sectionId: section.id,
           sectionTitle: section.title,
           searchText: `${group.title} ${section.title} ${e.method} ${e.path} ${e.summary} ${e.description || ""}`.toLowerCase(),
        });
      }

       // children folders
      for (const child of (section.children || [])) {
        walkSection(group, child);
      }
    }

    for (const g of state.spec.groups || []) {
      for (const s of g.sections || []) {
        walkSection(g, s);
      }
    }

    state.flatEndpoints = list;
  }

  // function renderSidebar(filtered = null) {
  //   const tree = $("sidebarTree");
  //   tree.innerHTML = "";

  //   const groups = state.spec.groups;

  //   for (const g of groups) {
  //     const wrap = document.createElement("div");
  //     wrap.className = "treeGroup";

  //     const head = document.createElement("div");
  //     head.className = "treeGroupHeader";
  //     head.innerHTML = `<strong>${escapeHtml(g.title)}</strong><span class="pill">${g.sections.length} sections</span>`;

  //     const body = document.createElement("div");
  //     body.className = "treeGroupBody open";

  //     head.addEventListener("click", () => body.classList.toggle("open"));

  //     for (const s of g.sections) {
  //       const sec = document.createElement("div");
  //       sec.className = "treeSection";
  //       sec.innerHTML = `<div class="treeSectionTitle">${escapeHtml(s.title)}</div>`;

  //       const endpoints = filtered
  //         ? filtered.filter(x => x.groupId === g.id && x.sectionId === s.id)
  //         : s.endpoints.map(e => ({...e, groupId: g.id, sectionId: s.id}));

  //       for (const e of endpoints) {
  //         const item = document.createElement("div");
  //         item.className = "treeItem";
  //         item.innerHTML = `
  //           <div style="display:flex;flex-direction:column;gap:3px;min-width:0">
  //             <div style="display:flex;gap:8px;align-items:center">
  //               <span class="method">${escapeHtml(e.method)}</span>
  //               <span style="font-size:12px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">${escapeHtml(e.summary || e.id)}</span>
  //             </div>
  //             <div class="path" style="white-space:nowrap;overflow:hidden;text-overflow:ellipsis">${escapeHtml(e.path)}</div>
  //           </div>
  //         `;
  //         item.addEventListener("click", () => {
  //           setHash(e.id);
  //           renderPage();
  //         });
  //         sec.appendChild(item);
  //       }

  //       body.appendChild(sec);
  //     }

  //     wrap.appendChild(head);
  //     wrap.appendChild(body);
  //     tree.appendChild(wrap);
  //   }
  // }

function renderSidebar(filtered = null) {
  const tree = $("sidebarTree");
  tree.innerHTML = "";

  function createEndpointItem(e) {
    const item = document.createElement("div");
    item.className = "treeItem";

    // Sidebar label = just "Login", "Forgot Password", etc.
    item.innerHTML = `
      <div style="display:flex;flex-direction:column;gap:3px;min-width:0">
        <div style="font-size:12px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">
          ${escapeHtml(e.summary || e.id)}
        </div>
      </div>
    `;

    item.addEventListener("click", () => {
      setHash(e.id);
      renderPage();
    });

    return item;
  }

  function renderSectionFolder(group, section, parentEl) {
    // collect endpoints for this folder
    const endpoints = filtered
      ? filtered.filter(x => x.groupId === group.id && x.sectionId === section.id)
      : (section.endpoints || []).map(e => ({ ...e, groupId: group.id, sectionId: section.id }));

    const children = section.children || [];

    // In search mode: hide folders with nothing inside
    if (filtered && endpoints.length === 0 && children.length === 0) return;

    const folder = document.createElement("div");
    folder.className = "treeGroup";
    folder.style.marginTop = "8px";

    const head = document.createElement("div");
    head.className = "treeGroupHeader";
    head.innerHTML = `<strong>${escapeHtml(section.title)}</strong><span class="pill">${endpoints.length}</span>`;

    const body = document.createElement("div");
    body.className = "treeGroupBody"; // collapsed by default

    // If searching, auto-open matching folders
    if (filtered) body.classList.add("open");

    head.addEventListener("click", () => body.classList.toggle("open"));

    // endpoints inside folder
    for (const e of endpoints) {
      body.appendChild(createEndpointItem(e));
    }

    // nested folders
    for (const child of children) {
      renderSectionFolder(group, child, body);
    }

    folder.appendChild(head);
    folder.appendChild(body);
    parentEl.appendChild(folder);
  }

  for (const g of (state.spec.groups || [])) {
    const wrap = document.createElement("div");
    wrap.className = "treeGroup";

    const head = document.createElement("div");
    head.className = "treeGroupHeader";
    head.innerHTML = `<strong>${escapeHtml(g.title)}</strong><span class="pill">folders</span>`;

    const body = document.createElement("div");
    body.className = "treeGroupBody";
    if (filtered) body.classList.add("open");

    head.addEventListener("click", () => body.classList.toggle("open"));

    for (const s of (g.sections || [])) {
      renderSectionFolder(g, s, body);
    }

    wrap.appendChild(head);
    wrap.appendChild(body);
    tree.appendChild(wrap);
  }
}

  
  // function quickStartHtml() {
  //   const q = state.spec.quickStart;
  //   const steps = (q.steps || []).map(s => `<li>${escapeHtml(s)}</li>`).join("");
  //   const overview = (q.overview && (q.overview.body || []).length) ? `
  //     <div class="card">
  //       <h2 class="h2">${escapeHtml(q.overview.title || "Overview")}</h2>
  //         ${(q.overview.body || []).map(p => `<p class="muted" style="margin:0 0 10px">${escapeHtml(p)}</p>`).join("")}
  //     </div>` : "";
  //   // const examples = (q.examples || []).map(ex => `
  //   //   <div class="card">
  //   //     <div class="h2">${escapeHtml(ex.title)}</div>
  //   //     <pre><code>${escapeHtml(ex.code)}</code></pre>
  //   //   </div>
  //   // `).join("");

  //   const examples = (q.examples || []).map(ex => `
  //     <div class="card">
  //       <div class="h2">${escapeHtml(ex.title)}</div>
  //       <div class="codeWrap">
  //         <button class="copyBtn" data-copy>Copy</button>
  //         <pre><code>${escapeHtml(ex.code)}</code></pre>
  //       </div>
  //     </div>
  //   `).join("");


  //   return `
  //     ${overview}
  //     <div class="card">
  //       <h2 class="h2">${escapeHtml(q.title || "Quick Start")}</h2>
  //       <div class="muted">Base URL: <span class="path">${escapeHtml(state.spec.baseUrl)}</span></div>
  //       <div style="height:10px"></div>
  //       <ol class="muted" style="margin:0;padding-left:18px">${steps}</ol>
  //     </div>
  //     ${examples}
  //     <div class="card">
  //       <div class="h2">Browse</div>
  //       <div class="muted">Use the sidebar to pick a group, then an endpoint. The sidebar stays open.</div>
  //     </div>
  //   `;
  // }
  
function quickStartHtml() {
  const q = state.spec.quickStart;

  const steps = (q.steps || []).map(s => `<li>${escapeHtml(s)}</li>`).join("");

  const overview = (q.overview && (q.overview.body || []).length) ? `
    <div class="card">
      <h2 class="h2">${escapeHtml(q.overview.title || "Overview")}</h2>
      ${(q.overview.body || []).map(p => `<p class="muted" style="margin:0 0 10px">${escapeHtml(p)}</p>`).join("")}
    </div>` : "";

  // const examples = (q.examples || []).map(ex => `
  //   <div class="card">
  //     <div class="h2">${escapeHtml(ex.title)}</div>
  //     <div class="codeWrap">
  //       <button class="copyIconBtn" data-copy type="button" aria-label="Copy code">
  //         ${copyIconSvg()}
  //       </button>
  //     </div>
  //   </div>
  // `).join("");

  const examples = (q.examples || []).map(ex => `
    <div class="card">
      <div class="h2">${escapeHtml(ex.title)}</div>
      <div class="codeWrap">
        <button class="copyIconBtn" data-copy type="button" aria-label="Copy code">
          ${copyIconSvg()}
        </button>
        <pre><code>${escapeHtml(ex.code)}</code></pre>
      </div>
    </div>
  `).join("");


  return `
    ${overview}
    <div class="card">
      <h2 class="h2">${escapeHtml(q.title || "Quick Start")}</h2>
      <div class="muted">Base URL: <span class="path">${escapeHtml(state.spec.baseUrl)}</span></div>
      <div style="height:10px"></div>
      <ol class="muted" style="margin:0;padding-left:18px">${steps}</ol>
    </div>
    ${examples}
    <div class="card">
      <div class="h2">Browse</div>
      <div class="muted">Use the sidebar to pick a group, then a folder, then an endpoint.</div>
    </div>
  `;
}


//   function endpointHtml(e) {

//     const usageBox = (e.usage && (e.usage.notes || []).length) ? `
//       <div class="card">
//         <div class="h2">${escapeHtml(e.usage.title || "Usage")}</div>
//         <ul class="muted" style="margin:0;padding-left:18px">
//           ${(e.usage.notes || []).map(n => `<li>${escapeHtml(n)}</li>`).join("")}
//         </ul>
//       </div>` : "";

//     // const req = e.request ? `
//     //   <div class="card">
//     //     <div class="h2">Request</div>
//     //     <div class="muted">Content-Type: <span class="path">${escapeHtml(e.request.contentType || "application/json")}</span></div>
//     //     <div style="height:10px"></div>
//     //     <pre><code>${escapeHtml(formatJson(e.request.example || {}))}</code></pre>
//     //   </div>
//     // ` : `
//     //   <div class="card">
//     //     <div class="h2">Request</div>
//     //     <div class="muted">No request body documented.</div>
//     //   </div>
//     // `;

//   const req = e.request ? `
//   <details open>
//     <summary><span class="h2">Request</span><span class="chev">toggle</span></summary>
//     <div class="detailsBody">
//       <div class="muted">Content-Type: <span class="path">${escapeHtml(e.request.contentType || "application/json")}</span></div>
//       <div style="height:10px"></div>
//       <div class="codeWrap">
//         <button class="copyBtn" data-copy>Copy</button>
//         <pre><code>${escapeHtml(formatJson(e.request.example || {}))}</code></pre>
//       </div>
//     </div>
//   </details>
// ` : `
//   <details open>
//     <summary><span class="h2">Request</span><span class="chev">toggle</span></summary>
//     <div class="detailsBody">
//       <div class="muted">No request body documented.</div>
//     </div>
//   </details>
// `;


//     // const responses = (e.responses || []).map(r => `
//     //   <div class="card">
//     //     <div class="h2">Response ${escapeHtml(r.status)}</div>
//     //     <div class="muted">${escapeHtml(r.description || "")}</div>
//     //     ${r.example ? `<div style="height:10px"></div><pre><code>${escapeHtml(formatJson(r.example))}</code></pre>` : ""}
//     //   </div>
//     // `).join("");

//   const responses = (e.responses || []).map(r => `
//   <details>
//     <summary><span class="h2">Response ${escapeHtml(r.status)}</span><span class="chev">toggle</span></summary>
//     <div class="detailsBody">
//       <div class="muted">${escapeHtml(r.description || "")}</div>
//       ${r.example ? `
//         <div style="height:10px"></div>
//         <div class="codeWrap">
//           <button class="copyBtn" data-copy>Copy</button>
//           <pre><code>${escapeHtml(formatJson(r.example))}</code></pre>
//         </div>` : ""}
//     </div>
//   </details>
// `).join("");


//     const tryItOut = `
//       <div class="card">
//         <div class="h2">Try it out</div>
//         <div class="muted">Uses Base URL + Path. Adds Authorization header if needed.</div>

//         <div class="tryRow">
//           <button class="btn" data-run="${escapeHtml(e.id)}">Run</button>
//           <span id="result-${escapeHtml(e.id)}" class="muted"></span>
//         </div>

//         <div style="height:10px"></div>
//         <div class="muted">Request body (JSON):</div>
//         <div style="height:6px"></div>
//         <textarea id="body-${escapeHtml(e.id)}" style="width:100%;min-height:140px;border-radius:12px;border:1px solid var(--border);background:rgba(2,6,23,0.6);color:var(--text);padding:10px;font-family:var(--mono);font-size:12px;outline:none">${escapeHtml(formatJson((e.request && e.request.example) ? e.request.example : {}))}</textarea>

//         <div style="height:10px"></div>
//         <div class="muted">Response:</div>
//         <div style="height:6px"></div>
//         <pre><code id="out-${escapeHtml(e.id)}">{}</code></pre>
//       </div>
//     `;

//     return `
//       <div class="card" id="${escapeHtml(e.id)}">
//         <div class="endpointHeader">
//           <div class="endpointTitle">
//             <div class="endpointTitleRow">
//               <span class="bigMethod">${escapeHtml(e.method)}</span>
//               <span class="bigPath">${escapeHtml(e.path)}</span>
//             </div>
//             <div class="muted">${escapeHtml(e.summary || "")}</div>
//           </div>
//           <span class="authTag">Auth: ${escapeHtml(e.auth || "none")}</span>
//         </div>
//         ${e.description ? `<div style="height:10px"></div><div class="muted">${escapeHtml(e.description)}</div>` : ""}
//       </div>

//       ${usageBox}
//       <div class="grid2">
//         ${req}
//         ${tryItOut}
//       </div>

//       ${responses || `
//         <div class="card">
//           <div class="h2">Responses</div>
//           <div class="muted">No responses documented.</div>
//         </div>
//       `}
//     `;
//   }

function endpointHtml(e) {
  const usageBox = (e.usage && (e.usage.notes || []).length) ? `
    <div class="card">
      <div class="h2">${escapeHtml(e.usage.title || "Usage")}</div>
      <ul class="muted" style="margin:0;padding-left:18px">
        ${(e.usage.notes || []).map(n => `<li>${escapeHtml(n)}</li>`).join("")}
      </ul>
    </div>` : "";

  const req = `
  <div class="card">
    <div class="sumRow">
      <span class="h2">Request</span>
      <button class="sumToggle" type="button" data-toggle="#req-${escapeHtml(e.id)}" aria-label="Toggle request"></button>
    </div>

    <div class="detailsBody">
      ${e.request ? `
        <div class="muted">Content-Type: <span class="path">${escapeHtml(e.request.contentType || "application/json")}</span></div>
        <div style="height:10px"></div>

        <div id="req-${escapeHtml(e.id)}" class="collapsibleBody open">
          <div class="codeWrap">
            <button class="copyIconBtn" data-copy type="button" aria-label="Copy request JSON">
              ${copyIconSvg()}
            </button>
            <pre><code>${escapeHtml(formatJson(e.request.example || {}))}</code></pre>
          </div>
        </div>
      ` : `
        <div class="muted">No request body documented.</div>
      `}
    </div>
  </div>
`;

  const responses = (e.responses || []).map(r => `
  <div class="card">
    <div class="sumRow">
      <span class="h2">Response ${escapeHtml(r.status)}</span>
      <button class="sumToggle" type="button" data-toggle="#res-${escapeHtml(e.id)}-${escapeHtml(String(r.status))}" aria-label="Toggle response"></button>
    </div>

    <div class="detailsBody">
      <div class="muted">${escapeHtml(r.description || "")}</div>

      ${r.example ? `
        <div style="height:10px"></div>

        <div id="res-${escapeHtml(e.id)}-${escapeHtml(String(r.status))}" class="collapsibleBody open">
          <div class="codeWrap">
            <button class="copyIconBtn" data-copy type="button" aria-label="Copy response JSON">
              ${copyIconSvg()}
            </button>
            <pre><code>${escapeHtml(formatJson(r.example))}</code></pre>
          </div>
        </div>
      ` : ""}
    </div>
  </div>
`).join("");


  // const tryItOut = `
  //   <div class="card">
  //     <div class="h2">Try it out</div>
  //     <div class="muted">Uses Base URL + Path. Adds Authorization header if needed.</div>

  //     <div class="tryRow">
  //       <button class="btn" data-run="${escapeHtml(e.id)}" type="button">Run</button>
  //       <span id="result-${escapeHtml(e.id)}" class="muted"></span>
  //     </div>

  //     <div style="height:10px"></div>
  //     <div class="muted">Request body (JSON):</div>
  //     <div style="height:6px"></div>
  //     <textarea
  //       id="body-${escapeHtml(e.id)}"
  //       style="width:100%;min-height:140px;border-radius:12px;border:1px solid var(--border);background:rgba(2,6,23,0.6);color:var(--text);padding:10px;font-family:var(--mono);font-size:12px;outline:none"
  //     >${escapeHtml(formatJson((e.request && e.request.example) ? e.request.example : {}))}</textarea>

  //     <div style="height:10px"></div>
  //     <div class="muted">Response:</div>
  //     <div style="height:6px"></div>
  //     <div class="codeWrap">
  //       <button class="copyIconBtn" data-copy type="button" aria-label="Copy response output">
  //         ${copyIconSvg()}
  //       </button>
  //       <pre><code id="out-${escapeHtml(e.id)}">{}</code></pre>
  //     </div>
  //   </div>
  // `;

  const tryItOut = `
  <div class="card tryCard">
    <div class="tryHead">
      <div>
        <div class="h2">Try it out</div>
        <div class="muted">Uses Base URL + Path. Adds Authorization header if needed.</div>
      </div>

      <div class="tryHeadRight">
        <button class="iconBtn" data-pickfile="${escapeHtml(e.id)}" type="button" title="Attach file" aria-label="Attach file">
          ${paperclipSvg()}
        </button>
        <input id="file-${escapeHtml(e.id)}" type="file" style="display:none"/>
      </div>
    </div>

    <div class="tryRow">
      <button class="btn" data-run="${escapeHtml(e.id)}" type="button">Run</button>
      <span id="result-${escapeHtml(e.id)}" class="muted"></span>
    </div>

    <div id="fileInfo-${escapeHtml(e.id)}" class="muted" style="margin-top:8px;display:none"></div>

    <div style="height:10px"></div>
    <div class="muted">Request body (JSON):</div>
    <div style="height:6px"></div>
    <textarea
      id="body-${escapeHtml(e.id)}"
      style="width:100%;min-height:140px;border-radius:12px;border:1px solid var(--border);background:rgba(2,6,23,0.6);color:var(--text);padding:10px;font-family:var(--mono);font-size:12px;outline:none"
    >${escapeHtml(formatJson((e.request && e.request.example) ? e.request.example : {}))}</textarea>

    <div style="height:10px"></div>
    <div class="muted">Response:</div>
    <div style="height:6px"></div>
    <div class="codeWrap">
      <button class="copyIconBtn" data-copy type="button" aria-label="Copy response output">
        ${copyIconSvg()}
      </button>
      <pre><code id="out-${escapeHtml(e.id)}">{}</code></pre>
    </div>
  </div>
`;

  return `
    <div class="card" id="${escapeHtml(e.id)}">
      <div class="endpointHeader">
        <div class="endpointTitle">
          <div class="endpointTitleRow">
            <span class="bigMethod">${escapeHtml(e.method)}</span>
            <code class="bigPath">${escapeHtml(e.path)}</code>

            <button
              class="copyIconBtn copyEndpointBtn"
              data-copy-endpoint="${escapeHtml(e.id)}"
              type="button"
              aria-label="Copy endpoint"
              title="Copy endpoint"
            >
              ${copyIconSvg()}
            </button>
          </div>
          <div class="muted">${escapeHtml(e.summary || "")}</div>
        </div>

        <span class="authTag">Auth: ${escapeHtml(e.auth || "none")}</span>
      </div>
      ${e.description ? `<div style="height:10px"></div><div class="muted">${escapeHtml(e.description)}</div>` : ""}
    </div>

    ${usageBox}
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

function wireMiniToggles(root = document) {
  root.querySelectorAll(".sumToggle[data-toggle]").forEach(btn => {
    if (btn.dataset.bound === "1") return;
    btn.dataset.bound = "1";

    btn.addEventListener("click", (e) => {
      e.preventDefault();
      e.stopPropagation();

      const sel = btn.getAttribute("data-toggle");
      if (!sel) return;

      const target = root.querySelector(sel) || document.querySelector(sel);
      if (!target) return;

      target.classList.toggle("open");
      btn.classList.toggle("isOn");
    });
  });
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
      wireCopyButtons(content);
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
    wireCopyButtons(content);
    wireMiniToggles(content);
    wireFilePickers(content)

    // attach run handler
    const runBtn = content.querySelector(`[data-run="${CSS.escape(e.id)}"]`);
    if (runBtn) runBtn.addEventListener("click", () => runEndpoint(e));

    // scroll into view
    const el = document.getElementById(e.id);
    if (el) el.scrollIntoView({ behavior: "smooth", block: "start" });
  }

  function paperclipSvg(){
  return `
    <svg viewBox="0 0 24 24" aria-hidden="true">
      <path d="M21 12.5l-8.5 8.5a6 6 0 0 1-8.5-8.5l10-10a4 4 0 0 1 5.5 5.5l-10 10a2.5 2.5 0 0 1-3.5-3.5l9-9"></path>
    </svg>
  `;
}

  // async function runEndpoint(e) {
  //   const token = $("bearerToken").value.trim();
  //   const baseUrl = $("baseUrl").value.trim() || "";
  //   const resultEl = document.getElementById(`result-${e.id}`);
  //   const outEl = document.getElementById(`out-${e.id}`);
  //   const bodyEl = document.getElementById(`body-${e.id}`);

  //   resultEl.textContent = "Running...";
  //   resultEl.className = "muted";
  //   outEl.textContent = "{}";

  //   const url = (baseUrl.replace(/\/+$/,"")) + e.path;

  //   const headers = {};
  //   if (e.request && e.request.contentType) headers["Content-Type"] = e.request.contentType;
  //   if ((e.auth || "").toLowerCase() === "bearer" && token) headers["Authorization"] = `Bearer ${token}`;

  //   let body = undefined;
  //   if (e.method !== "GET" && e.method !== "DELETE") {
  //     try {
  //       const parsed = JSON.parse(bodyEl.value || "{}");
  //       body = JSON.stringify(parsed);
  //     } catch {
  //       resultEl.textContent = "Invalid JSON body";
  //       resultEl.className = "resultErr";
  //       return;
  //     }
  //   }

  //   try {
  //     const res = await fetch(url, {
  //       method: e.method,
  //       headers,
  //       body,
  //     });

  //     const text = await res.text();
  //     let data = text;
  //     try { data = JSON.parse(text); } catch {}

  //     outEl.textContent = typeof data === "string" ? data : JSON.stringify(data, null, 2);

  //     if (res.ok) {
  //       resultEl.textContent = `OK (${res.status})`;
  //       resultEl.className = "resultOk";
  //     } else {
  //       resultEl.textContent = `ERROR (${res.status})`;
  //       resultEl.className = "resultErr";
  //     }
  //   } catch (err) {
  //     resultEl.textContent = "Network error";
  //     resultEl.className = "resultErr";
  //     outEl.textContent = String(err);
  //   }
  // }
  
// async function runEndpoint(e) {
//   const resultEl = document.getElementById(`result-${e.id}`);
//   const outEl = document.getElementById(`out-${e.id}`);
//   if (!resultEl || !outEl) return;

//   resultEl.textContent = "Running...";
//   outEl.textContent = "{}";

//   try {
//     const token = localStorage.getItem(STORAGE.token) || "";
//     const baseUrl = (localStorage.getItem(STORAGE.baseUrl) || state.spec.baseUrl || "").replace(/\/+$/, "");
//     const url = baseUrl + e.path;

//     const headers = {};
//     if (e.request?.contentType) headers["Content-Type"] = e.request.contentType;
//     else headers["Content-Type"] = "application/json";

//     if ((e.auth || "").toLowerCase() === "bearer" && token) {
//       headers["Authorization"] = `Bearer ${token}`;
//     }

//     let body = undefined;
//     const bodyBox = document.getElementById(`body-${e.id}`);
//     if (bodyBox && ["POST", "PUT", "PATCH"].includes((e.method || "").toUpperCase())) {
//       const raw = (bodyBox.value || "").trim();
//       if (raw) body = raw;
//     }

//     const res = await fetch(url, {
//       method: e.method,
//       headers,
//       body,
//     });

//     const ct = (res.headers.get("content-type") || "").toLowerCase();
//     const text = await res.text();

//     // Pretty print JSON if possible
//     if (ct.includes("application/json")) {
//       try {
//         const obj = JSON.parse(text || "{}");
//         outEl.textContent = JSON.stringify(obj, null, 2);
//       } catch {
//         outEl.textContent = text;
//       }
//     } else {
//       // maybe JSON even if header isn't set
//       try {
//         const obj = JSON.parse(text);
//         outEl.textContent = JSON.stringify(obj, null, 2);
//       } catch {
//         outEl.textContent = text;
//       }
//     }

//     resultEl.textContent = res.ok ? `OK (${res.status})` : `Error (${res.status})`;
//   } catch (err) {
//     resultEl.textContent = "Error";
//     outEl.textContent = String(err);
//   }
// }

async function runEndpoint(e) {
  const resultEl = document.getElementById(`result-${e.id}`);
  const outEl = document.getElementById(`out-${e.id}`);
  if (!resultEl || !outEl) return;

  resultEl.textContent = "Running...";
  outEl.textContent = "{}";

  try {
    const token = localStorage.getItem(STORAGE.token) || "";
    const baseUrl = (localStorage.getItem(STORAGE.baseUrl) || state.spec.baseUrl || "").replace(/\/+$/, "");
    const url = baseUrl + e.path;

    const method = (e.method || "GET").toUpperCase();

    const attached = fileState[e.id] || null;

    let headers = {};
    let body;

    // read textarea JSON
    const bodyBox = document.getElementById(`body-${e.id}`);
    const raw = bodyBox ? (bodyBox.value || "").trim() : "";
    let jsonObj = {};
    if (raw) {
      try { jsonObj = JSON.parse(raw); } catch { jsonObj = {}; }
    }

    if (attached) {
      // ✅ multipart upload
      const fd = new FormData();

      // put JSON fields into form data
      Object.entries(jsonObj || {}).forEach(([k, v]) => {
        if (v === undefined || v === null) return;
        if (typeof v === "object") fd.append(k, JSON.stringify(v));
        else fd.append(k, String(v));
      });

      const fieldName = e.request?.file?.fieldName || "file";
      fd.append(fieldName, attached);

      body = fd;
      
      if ((e.auth || "").toLowerCase() === "bearer" && token) {
        headers["Authorization"] = `Bearer ${token}`;
      }
    } else {
      // normal JSON request
      headers["Content-Type"] = "application/json";
      if ((e.auth || "").toLowerCase() === "bearer" && token) {
        headers["Authorization"] = `Bearer ${token}`;
      }
      if (["POST","PUT","PATCH","DELETE"].includes(method)) {
        body = raw ? raw : JSON.stringify(jsonObj || {});
      }
    }

    const res = await fetch(url, { method, headers, body });

    const text = await res.text();
    // pretty print JSON if possible
    try {
      const obj = JSON.parse(text || "{}");
      outEl.textContent = JSON.stringify(obj, null, 2);
    } catch {
      outEl.textContent = text;
    }

    resultEl.textContent = res.ok ? `OK (${res.status})` : `Error (${res.status})`;
  } catch (err) {
    resultEl.textContent = "Error";
    outEl.textContent = String(err);
  }
}


function wireFilePickers(root = document) {
  root.querySelectorAll("[data-pickfile]").forEach(btn => {
    if (btn.dataset.bound === "1") return;
    btn.dataset.bound = "1";

    btn.addEventListener("click", () => {
      const id = btn.getAttribute("data-pickfile");
      const input = document.getElementById(`file-${id}`);
      if (!input) return;

      // configure accept from spec if available
      const e = state.flatEndpoints.find(x => x.id === id);
      const accept = e?.request?.file?.accept?.join(",") || "image/*,video/*,application/pdf";
      input.accept = accept;
      input.multiple = !!e?.request?.file?.multiple;

      input.click();
    });
  });

  root.querySelectorAll(`input[type="file"][id^="file-"]`).forEach(inp => {
    if (inp.dataset.bound === "1") return;
    inp.dataset.bound = "1";

    inp.addEventListener("change", () => {
      const id = inp.id.replace("file-", "");
      const f = inp.files && inp.files[0] ? inp.files[0] : null;
      fileState[id] = f;

      const info = document.getElementById(`fileInfo-${id}`);
      if (info) {
        if (!f) {
          info.style.display = "none";
          info.textContent = "";
        } else {
          info.style.display = "block";
          info.textContent = `Attached: ${f.name} (${Math.round(f.size/1024)} KB)`;
        }
      }
    });
  });
}


  function wireTopControls() {
    const tokenInput = $("bearerToken");
    const baseUrlInput = $("baseUrl");

    const envSelect = $("envSelect");
    if (envSelect) {
      const savedEnv = localStorage.getItem(STORAGE.env);
        if (savedEnv && ENV[savedEnv]) envSelect.value = savedEnv;
          // If user never set baseUrl before, set it from env
          if (!localStorage.getItem(STORAGE.baseUrl) && ENV[envSelect.value]) {
              baseUrlInput.value = ENV[envSelect.value];
          }
          envSelect.addEventListener("change", () => {
          const v = envSelect.value;
          localStorage.setItem(STORAGE.env, v);
          if (ENV[v]) {
            baseUrlInput.value = ENV[v];
            localStorage.setItem(STORAGE.baseUrl, baseUrlInput.value);
          }
       });
    }

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

//   function wireDetailsToggles(root = document) {
//   root.querySelectorAll(".sumToggle").forEach(btn => {
//     if (btn.dataset.bound === "1") return;
//     btn.dataset.bound = "1";

//     btn.addEventListener("click", (ev) => {
//       ev.preventDefault();
//       ev.stopPropagation();

//       const details = btn.closest("details");
//       if (!details) return;

//       details.open = !details.open;
//     });
//   });
// }


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

//   function wireCopyButtons(root = document) {
//   root.querySelectorAll("[data-copy]").forEach(btn => {
//     btn.addEventListener("click", async () => {
//       const pre = btn.parentElement?.querySelector("pre");
//       const text = pre ? pre.innerText : "";
//       try {
//         await navigator.clipboard.writeText(text);
//         btn.textContent = "Copied";
//         setTimeout(() => (btn.textContent = "Copy"), 900);
//       } catch {
//         btn.textContent = "Failed";
//         setTimeout(() => (btn.textContent = "Copy"), 900);
//       }
//     });
//   });
// }

//   function wireThemeToggle() {
//   const btn = $("themeToggle");
//   if (!btn) return;

//   const saved = localStorage.getItem(STORAGE.theme);
//   if (saved) document.documentElement.setAttribute("data-theme", saved);

//   btn.addEventListener("click", () => {
//     const current = document.documentElement.getAttribute("data-theme") || "dark";
//     const next = current === "dark" ? "light" : "dark";
//     document.documentElement.setAttribute("data-theme", next);
//     localStorage.setItem(STORAGE.theme, next);
//   });
// }

// function wireCopyButtons(root = document) {
//   // Auto-select on hover for code blocks
//   root.querySelectorAll("pre[data-autoselect='1']").forEach(pre => {
//     pre.addEventListener("mouseenter", () => selectCodeInPre(pre));
//   });

//   root.querySelectorAll("[data-copy]").forEach(btn => {
//     btn.addEventListener("click", async () => {
//       const pre = btn.parentElement?.querySelector("pre");
//       if (!pre) return;

//       const text = pre.innerText || "";
//       const ok = await copyTextReliable(text);

//       if (ok) {
//         toastSuccess("Copied", "Code copied to clipboard");
//       } else {
//         toastSuccess("Copy failed", "Your browser blocked clipboard access");
//       }
//     });
//   });
// }

// function wireCopyButtons(root = document) {
//   // Auto-select on hover for code blocks
//   root.querySelectorAll("pre[data-autoselect='1']").forEach(pre => {
//     if (pre.dataset.bound === "1") return;
//     pre.dataset.bound = "1";
//     pre.addEventListener("mouseenter", () => selectCodeInPre(pre));
//   });

//   root.querySelectorAll("[data-copy]").forEach(btn => {
//     if (btn.dataset.bound === "1") return;
//     btn.dataset.bound = "1";

//     btn.addEventListener("click", async () => {
//       // We copy the nearest code block inside the same codeWrap
//       const wrap = btn.closest(".codeWrap") || btn.parentElement;
//       const code = wrap ? wrap.querySelector("pre code") : null;
//       const pre = wrap ? wrap.querySelector("pre") : null;

//       const text = (code?.innerText || pre?.innerText || "").trim();
//       if (!text) {
//         toastSuccess("Copy failed", "Nothing to copy");
//         return;
//       }

//       const ok = await copyTextReliable(text);
//       if (ok) {
//         toastSuccess("Copied", "Code copied to clipboard");
//       } else {
//         toastSuccess("Copy failed", "Your browser blocked clipboard access");
//       }
//     });
//   });
// }

function wireCopyButtons(root = document) {
  root.querySelectorAll("[data-copy]").forEach(btn => {
    if (btn.dataset.bound === "1") return;
    btn.dataset.bound = "1";

    btn.addEventListener("click", async () => {
      const wrap = btn.closest(".codeWrap") || btn.parentElement;
      const code = wrap ? wrap.querySelector("pre code") : null;
      const pre = wrap ? wrap.querySelector("pre") : null;

      const text = (code?.innerText || pre?.innerText || "").trim();
      if (!text) {
        toastSuccess("Copy failed", "Nothing to copy");
        return;
      }

      const ok = await copyTextReliable(text);
      if (ok) toastSuccess("Capied", "Code copied to clipboard");
      else toastSuccess("Copy failed", "Your browser blocked clipboard access");
    });
  });
}


  async function boot() {
    wireTopControls();
    // wireThemeToggle();


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
