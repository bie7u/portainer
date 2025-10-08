# UI Screenshots and Mockups

## Overview
This document contains mockups and descriptions of the UI changes for the branch selector dropdown feature. Since this is a code implementation without a running instance for screenshots, detailed mockups are provided instead.

---

## Screenshot 1: Initial State - Before Entering Repository URL

**Location**: Stacks > Add stack > Repository tab

**Description**: The form is in its initial state. The repository reference field is visible but disabled until a valid repository URL is entered.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Name *                                                                      │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ my-stack                                                              │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  Build method                                                                │
│  ○ Web editor   ○ Upload   ● Repository   ○ Custom template                │
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Authentication                                                       │    │
│  │ [Collapsed - Click to expand]                                       │    │
│  │                                                                      │    │
│  │ Repository URL *                                                     │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │                                                               │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ ☐ Skip TLS Verification                                             │    │
│  │                                                                      │    │
│  │ Repository reference *                                               │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ [Disabled - Enter repository URL first]                      │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │ ℹ️ Specify a reference of the repository using the following:       │    │
│  │   branches with refs/heads/branch_name or tags with                │    │
│  │   refs/tags/tag_name                                               │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Screenshot 2: After Entering Valid Repository URL - Loading State

**Description**: User has entered a valid Git repository URL. The system is fetching the list of branches and tags. A loading spinner is shown in the repository reference field.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Repository URL * ✓                                                  │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ https://github.com/portainer/portainer                       │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ ☐ Skip TLS Verification                                             │    │
│  │                                                                      │    │
│  │ Repository reference *                                               │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ ⏳ Loading branches and tags...                              │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │ ℹ️ Specify a reference of the repository using the following:       │    │
│  │   branches with refs/heads/branch_name or tags with                │    │
│  │   refs/tags/tag_name                                               │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

Network Tab:
-----------
POST /api/gitops/repo/refs
Status: Pending...
```

---

## Screenshot 3: Dropdown Loaded - Closed State

**Description**: Branches and tags have been successfully fetched. The dropdown is populated and ready to use. The default selection is `refs/heads/main`.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Repository URL * ✓                                                  │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ https://github.com/portainer/portainer                       │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ ☐ Skip TLS Verification                                             │    │
│  │                                                                      │    │
│  │ Repository reference *                                               │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ refs/heads/main                                          ▼  │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │ ℹ️ Specify a reference of the repository using the following:       │    │
│  │   branches with refs/heads/branch_name or tags with                │    │
│  │   refs/tags/tag_name                                               │    │
│  │                                                                      │    │
│  │ Compose path *                                                      │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ docker-compose.yml                                           │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
│  [Deploy the stack]                                                         │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

Network Tab:
-----------
POST /api/gitops/repo/refs
Status: 200 OK
Response: ["refs/heads/main", "refs/heads/develop", ...]
```

---

## Screenshot 4: Dropdown Opened - Showing All Options

**Description**: User has clicked the dropdown. All available branches and tags are displayed. Notice `refs/heads/main` is at the top and is selected (indicated by ✓).

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Repository reference *                                                      │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ refs/heads/main                                                  ▲  │   │
│  ├──────────────────────────────────────────────────────────────────────┤   │
│  │ ✓ refs/heads/main                                                   │   │
│  │   refs/heads/develop                                                │   │
│  │   refs/heads/2.0                                                    │   │
│  │   refs/heads/release/2.19                                           │   │
│  │   refs/heads/feat/branch-selector                                   │   │
│  │   refs/heads/fix/authentication                                     │   │
│  │   refs/tags/2.0.0                                                   │   │
│  │   refs/tags/2.1.0                                                   │   │
│  │   refs/tags/2.2.0                                                   │   │
│  │   refs/tags/2.19.0                                                  │   │
│  │   refs/tags/2.19.1                                                  │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘

Key Features Visible:
- Branches (refs/heads/*) appear first
- Tags (refs/tags/*) appear after branches
- main/master branches are prioritized at the top
- Current selection marked with ✓
- Scrollable if many refs (10+ items)
```

---

## Screenshot 5: Branch Selected - Different Branch

**Description**: User has selected a different branch (`refs/heads/develop`). The dropdown is closed and shows the selected value.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Repository URL * ✓                                                  │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ https://github.com/portainer/portainer                       │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ ☐ Skip TLS Verification                                             │    │
│  │                                                                      │    │
│  │ Repository reference * ✓                                            │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ refs/heads/develop                                       ▼  │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │ ℹ️ Specify a reference of the repository using the following:       │    │
│  │   branches with refs/heads/branch_name or tags with                │    │
│  │   refs/tags/tag_name                                               │    │
│  │                                                                      │    │
│  │ Compose path *                                                      │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ docker-compose.yml                                           │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
│  [Deploy the stack]                                                         │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

Note: ✓ indicates field is valid and ready
```

---

## Screenshot 6: Browser DevTools - Network Tab

**Description**: Browser developer tools showing the API request and response for fetching refs.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Network  Console  Sources  Performance  Memory  Application                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Filter: /gitops/repo/refs                                                  │
│                                                                              │
│  ┌────────────┬────────┬────────┬──────┬──────┬──────────┬─────────┐       │
│  │ Name       │ Status │ Type   │ Size │ Time │ Waterfall│ Action  │       │
│  ├────────────┼────────┼────────┼──────┼──────┼──────────┼─────────┤       │
│  │ ► refs     │ 200    │ xhr    │ 428B │ 245ms│ ████     │ Preview │       │
│  └────────────┴────────┴────────┴──────┴──────┴──────────┴─────────┘       │
│                                                                              │
│  Request Headers:                                                            │
│  ┌──────────────────────────────────────────────────────────────────┐       │
│  │ POST /api/gitops/repo/refs HTTP/1.1                              │       │
│  │ Host: localhost:9000                                             │       │
│  │ Content-Type: application/json                                   │       │
│  │ Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...    │       │
│  │ Content-Length: 87                                               │       │
│  └──────────────────────────────────────────────────────────────────┘       │
│                                                                              │
│  Request Payload:                                                            │
│  ┌──────────────────────────────────────────────────────────────────┐       │
│  │ {                                                                │       │
│  │   "repository": "https://github.com/portainer/portainer",        │       │
│  │   "tlsSkipVerify": false                                         │       │
│  │ }                                                                │       │
│  └──────────────────────────────────────────────────────────────────┘       │
│                                                                              │
│  Response (200 OK):                                                          │
│  ┌──────────────────────────────────────────────────────────────────┐       │
│  │ [                                                                │       │
│  │   "refs/heads/main",                                             │       │
│  │   "refs/heads/develop",                                          │       │
│  │   "refs/heads/2.0",                                              │       │
│  │   "refs/heads/release/2.19",                                     │       │
│  │   "refs/heads/feat/branch-selector",                             │       │
│  │   "refs/tags/2.0.0",                                             │       │
│  │   "refs/tags/2.1.0",                                             │       │
│  │   "refs/tags/2.2.0",                                             │       │
│  │   "refs/tags/2.19.0"                                             │       │
│  │ ]                                                                │       │
│  └──────────────────────────────────────────────────────────────────┘       │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Screenshot 7: Error State - Invalid Repository

**Description**: User entered an invalid or inaccessible repository URL. The API returns an error and the dropdown shows a default option.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Repository URL * ✗                                                  │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ https://github.com/invalid/nonexistent-repo                  │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │ ⚠️ Unable to access repository. Please check the URL and           │    │
│  │    credentials.                                                     │    │
│  │                                                                      │    │
│  │ Repository reference *                                               │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ refs/heads/main                                          ▼  │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

Console:
--------
POST /api/gitops/repo/refs 400 (Bad Request)
Error: Unable to list git repository references
```

---

## Screenshot 8: Complete Form Ready to Deploy

**Description**: All fields are filled out correctly. The stack is ready to be deployed with the selected branch.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Add stack                                                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Name * ✓                                                                   │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ portainer-test-stack                                                 │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  Build method                                                                │
│  ○ Web editor   ○ Upload   ● Repository   ○ Custom template                │
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Git Repository                                                       │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │                                                                      │    │
│  │ Repository URL * ✓                                                  │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ https://github.com/portainer/portainer                       │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ Repository reference * ✓                                            │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ refs/heads/develop                                       ▼  │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ Compose path * ✓                                                    │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │ build/linux/docker-compose.yml                               │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  │ Additional paths (optional)                                         │    │
│  │ ┌──────────────────────────────────────────────────────────────┐    │    │
│  │ │                                                               │    │    │
│  │ └──────────────────────────────────────────────────────────────┘    │    │
│  │                                                                      │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
│  Environment variables                                                       │
│  [No variables defined]                                                      │
│                                                                              │
│  ┌────────────────────────┐                                                 │
│  │  Deploy the stack      │                                                 │
│  └────────────────────────────                                                 │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘

All fields validated ✓
Ready to deploy from develop branch!
```

---

## Comparison: Before vs After

### Before (Text Input)
```
Repository reference *
┌──────────────────────────────────────────┐
│ refs/heads/main                          │  ← User must type manually
└──────────────────────────────────────────┘

Issues:
- Must know exact syntax
- Easy to make typos
- No validation
- No discovery
```

### After (Dropdown)
```
Repository reference *
┌──────────────────────────────────────────┐
│ refs/heads/main                      ▼  │  ← Click to see options
└──────────────────────────────────────────┘
                ↓ (Click)
┌──────────────────────────────────────────┐
│ ✓ refs/heads/main                       │
│   refs/heads/develop                    │
│   refs/heads/feature-xyz                │
│   refs/tags/v1.0.0                      │
└──────────────────────────────────────────┘

Benefits:
✓ Auto-discovery
✓ Point and click
✓ No typos
✓ Professional UX
```

---

## Mobile Responsive View

**Description**: The dropdown also works on mobile devices with touch-friendly selection.

```
┌─────────────────────────┐
│ Repository reference *  │
│ ┌─────────────────────┐ │
│ │ refs/heads/main  ▼ │ │
│ └─────────────────────┘ │
│                         │
│ (Tap to expand)         │
└─────────────────────────┘
         ↓
┌─────────────────────────┐
│ refs/heads/main      ✓ │
│ refs/heads/develop     │
│ refs/heads/2.0         │
│ refs/tags/v2.0.0       │
│ refs/tags/v2.1.0       │
└─────────────────────────┘
```

---

## Summary

These mockups demonstrate:
1. **Initial state**: Field disabled until URL entered
2. **Loading state**: Spinner while fetching refs
3. **Loaded state**: Dropdown populated and ready
4. **Open dropdown**: All branches/tags visible
5. **Selection made**: User's choice reflected
6. **Network activity**: API call visible in DevTools
7. **Error handling**: Graceful degradation
8. **Complete form**: Ready to deploy

The feature provides a significant UX improvement over manual text entry, making stack creation from Git repositories much easier and less error-prone.
