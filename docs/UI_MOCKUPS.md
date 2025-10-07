# Branch Switching Feature - UI Mockups and Visual Guide

## Overview

This document provides visual representations of the UI changes made to implement the branch switching feature. Since actual screenshots require a running instance, we provide detailed ASCII mockups and descriptions.

## Stack Details Page - Before and After

### Before (Original Portainer)

```
┌─────────────────────────────────────────────────────────────────────────┐
│ Stack details                                                     🔄 ⚙️ │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│ ┌─── Stack ────────────────────────────────────────────────────────┐  │
│ │                                                                   │  │
│ │ Stack details                                                     │  │
│ │ my-stack  [▶ Start] [⏹ Stop] [🗑 Delete]                         │  │
│ │                                                                   │  │
│ │ ═══════════════════════════════════════════════════════════       │  │
│ │                                                                   │  │
│ │ Redeploy from git repository                                      │  │
│ │                                                                   │  │
│ │ ℹ️  Repository: https://github.com/user/repo                      │  │
│ │    Reference: refs/heads/main                                     │  │
│ │    Compose file: docker-compose.yml                               │  │
│ │                                                                   │  │
│ │ [➕ Advanced configuration]                                        │  │
│ │                                                                   │  │
│ │ Actions                                                           │  │
│ │ [🔄 Pull and redeploy]  [💾 Save settings]                        │  │
│ │                                                                   │  │
│ └───────────────────────────────────────────────────────────────────┘  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### After (With Branch Switching Feature)

```
┌─────────────────────────────────────────────────────────────────────────┐
│ Stack details                                                     🔄 ⚙️ │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│ ┌─── Stack ────────────────────────────────────────────────────────┐  │
│ │                                                                   │  │
│ │ Stack details                                                     │  │
│ │ my-stack  [▶ Start] [⏹ Stop] [🗑 Delete]                         │  │
│ │                                                                   │  │
│ │ ═══════════════════════════════════════════════════════════════   │  │
│ │                                                                   │  │
│ │ Redeploy from git repository                                      │  │
│ │                                                                   │  │
│ │ ℹ️  Repository: https://github.com/user/repo                      │  │
│ │    Reference: refs/heads/main                                     │  │
│ │    Compose file: docker-compose.yml                               │  │
│ │                                                                   │  │
│ │ ┌─────────────────────────────────────────────────────────────┐  │  │
│ │ │ Switch to different branch                                  │  │  │
│ │ │                                                             │  │  │
│ │ │ [-- Select a branch to switch --    ▼] [🔀 Switch Branch] │  │  │
│ │ │   main (current)                                            │  │  │
│ │ │   develop                                                   │  │  │
│ │ │   feature/new-ui                                            │  │  │
│ │ │   hotfix/security-patch                                     │  │  │
│ │ │                                                             │  │  │
│ │ │ ℹ️  Switching branches will redeploy the stack from the     │  │  │
│ │ │    selected branch. Current branch: main                    │  │  │
│ │ └─────────────────────────────────────────────────────────────┘  │  │
│ │                                                                   │  │
│ │ [➕ Advanced configuration]                                        │  │
│ │                                                                   │  │
│ │ Actions                                                           │  │
│ │ [🔄 Pull and redeploy]  [💾 Save settings]                        │  │
│ │                                                                   │  │
│ └───────────────────────────────────────────────────────────────────┘  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

## Key UI Elements

### 1. Branch Selector Dropdown

**Normal State:**
```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ [-- Select a branch to switch --    ▼]  [🔀 Switch Branch] │
│                                              │
│ ℹ️  Switching branches will redeploy the     │
│    stack from the selected branch.           │
│    Current branch: main                      │
└──────────────────────────────────────────────┘
```

**Dropdown Expanded:**
```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ ┌──────────────────────────────┐            │
│ │ -- Select a branch to switch --│▼│        │
│ ├──────────────────────────────┤│          │
│ │ main (current)               ││          │  <- Disabled
│ │ develop                      ││          │  <- Selectable
│ │ feature/new-ui               ││          │  <- Selectable
│ │ feature/auth-improvement     ││          │  <- Selectable
│ │ hotfix/security-patch        ││          │  <- Selectable
│ │ release/v2.0                 ││          │  <- Selectable
│ └──────────────────────────────┘│          │
│                                  │          │
│                    [🔀 Switch Branch]       │  <- Disabled
│                                              │
│ ℹ️  Switching branches will redeploy the     │
│    stack from the selected branch.           │
│    Current branch: main                      │
└──────────────────────────────────────────────┘
```

**Branch Selected:**
```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ [develop                         ▼]  [🔀 Switch Branch] │
│                                      ^^^^^^^^^^^^^^      │
│                                      Now enabled         │
│                                              │
│ ℹ️  Switching branches will redeploy the     │
│    stack from the selected branch.           │
│    Current branch: main                      │
└──────────────────────────────────────────────┘
```

### 2. Loading States

**Loading Branches:**
```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ [Loading branches... ⌛         ▼]  [🔀 Switch Branch] │
│                                      ^^^^^^^^^^^^^^      │
│                                      Disabled            │
└──────────────────────────────────────────────┘
```

**Switching Branch (In Progress):**
```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ [develop                         ▼]  [⌛ Switching...] │
│                                      ^^^^^^^^^^^^^^^     │
│                                      Disabled, spinner   │
└──────────────────────────────────────────────┘
```

### 3. Confirmation Dialog

When user clicks "Switch Branch" button:

```
╔═══════════════════════════════════════════════════════════════╗
║                        Are you sure?                          ║
╠═══════════════════════════════════════════════════════════════╣
║                                                               ║
║  Switching to branch "develop" will redeploy the stack        ║
║  from that branch. Any changes to this stack will be          ║
║  overridden. Do you wish to continue?                         ║
║                                                               ║
║                                                               ║
║                  [Cancel]    [⚠️ Switch Branch]               ║
║                               ^^^^^^^^^^^^^                   ║
║                               Warning style button            ║
╚═══════════════════════════════════════════════════════════════╝
```

### 4. Success Notification

After successful branch switch:

```
┌─────────────────────────────────────────────────────────────┐
│ ✓ Success                                               [×] │
├─────────────────────────────────────────────────────────────┤
│ Stack switched to branch "develop" and redeployed           │
│ successfully                                                │
└─────────────────────────────────────────────────────────────┘
```

### 5. Error Notification

If branch switch fails:

```
┌─────────────────────────────────────────────────────────────┐
│ ✗ Failure                                               [×] │
├─────────────────────────────────────────────────────────────┤
│ Unable to switch branch                                     │
│                                                             │
│ Details: Unable to clone git repository directory          │
└─────────────────────────────────────────────────────────────┘
```

## Responsive Design

### Desktop View (> 992px)

```
┌────────────────────────────────────────────────────────────────────┐
│                                                                    │
│  Switch to different branch                                        │
│                                                                    │
│  [-- Select a branch to switch -- ▼]  [🔀 Switch Branch]          │
│                                                                    │
│  ℹ️  Switching branches will redeploy the stack from the          │
│     selected branch. Current branch: main                          │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘
```

### Tablet View (768px - 991px)

```
┌──────────────────────────────────────────────────────┐
│                                                      │
│  Switch to different branch                          │
│                                                      │
│  [-- Select a branch to switch -- ▼]                 │
│  [🔀 Switch Branch]                                  │
│                                                      │
│  ℹ️  Switching branches will redeploy the stack      │
│     from the selected branch.                        │
│     Current branch: main                             │
│                                                      │
└──────────────────────────────────────────────────────┘
```

### Mobile View (< 768px)

```
┌────────────────────────────────────┐
│                                    │
│  Switch to different branch        │
│                                    │
│  [-- Select branch -- ▼]           │
│  [🔀 Switch Branch]                │
│                                    │
│  ℹ️  Switching branches will       │
│     redeploy the stack.            │
│     Current: main                  │
│                                    │
└────────────────────────────────────┘
```

## User Interaction Flow

### Flow Diagram

```
┌─────────────┐
│   Start     │
│  (On Stack  │
│   Details)  │
└──────┬──────┘
       │
       v
┌─────────────────────┐
│  Branch selector    │
│  loads branches     │
│  automatically      │
└──────┬──────────────┘
       │
       v
┌─────────────────────┐     ┌──────────────────┐
│  Branches shown     │────>│  User selects    │
│  in dropdown        │     │  different       │
│  (current disabled) │     │  branch          │
└─────────────────────┘     └────────┬─────────┘
                                     │
                                     v
                            ┌────────────────────┐
                            │  User clicks       │
                            │  "Switch Branch"   │
                            └────────┬───────────┘
                                     │
                                     v
                            ┌────────────────────┐
                            │  Confirmation      │
                            │  dialog appears    │
                            └────────┬───────────┘
                                     │
                    ┌────────────────┴────────────────┐
                    │                                 │
                    v                                 v
         ┌──────────────────┐              ┌─────────────────┐
         │  User cancels    │              │  User confirms  │
         └────────┬─────────┘              └────────┬────────┘
                  │                                 │
                  v                                 v
         ┌──────────────────┐              ┌─────────────────┐
         │  No action       │              │  API call to    │
         │  Dialog closes   │              │  switch branch  │
         └──────────────────┘              └────────┬────────┘
                                                    │
                                     ┌──────────────┴────────────────┐
                                     │                               │
                                     v                               v
                          ┌────────────────────┐         ┌──────────────────┐
                          │  Success           │         │  Error           │
                          │  notification      │         │  notification    │
                          └────────┬───────────┘         └──────────────────┘
                                   │
                                   v
                          ┌────────────────────┐
                          │  Page reloads      │
                          │  Stack updated     │
                          │  with new branch   │
                          └────────────────────┘
```

## Color Scheme and Styling

### Colors Used

- **Primary Blue**: `#3b82f6` - Main action buttons
- **Warning Orange**: `#f97316` - Confirmation dialogs
- **Success Green**: `#22c55e` - Success notifications
- **Error Red**: `#ef4444` - Error notifications
- **Info Blue**: `#3b82f6` - Information text
- **Gray**: `#6b7280` - Disabled states and secondary text
- **White**: `#ffffff` - Backgrounds
- **Dark**: `#1f2937` - Text

### Button States

**Normal Button:**
```
┌──────────────────┐
│ 🔀 Switch Branch │  (Blue background, white text)
└──────────────────┘
```

**Hover State:**
```
┌──────────────────┐
│ 🔀 Switch Branch │  (Darker blue background)
└──────────────────┘
```

**Disabled State:**
```
┌──────────────────┐
│ 🔀 Switch Branch │  (Gray background, light gray text)
└──────────────────┘
```

**Loading State:**
```
┌──────────────────┐
│ ⌛ Switching...  │  (Blue background with spinner)
└──────────────────┘
```

## Accessibility Features

### Keyboard Navigation

- `Tab`: Navigate to branch selector
- `Tab`: Navigate to Switch Branch button
- `Arrow Up/Down`: Navigate dropdown options
- `Enter`: Select option/click button
- `Esc`: Close dropdown

### Screen Reader Support

The UI includes proper ARIA labels:

```html
<select aria-label="Select git branch to switch to">
  <option>-- Select a branch to switch --</option>
  <option disabled>main (current)</option>
  <option>develop</option>
</select>

<button aria-label="Switch to selected branch">
  Switch Branch
</button>
```

### Focus Indicators

All interactive elements have visible focus indicators:
```
┌──────────────────────┐
│ ┌──────────────────┐ │  <- Blue border when focused
│ │ develop       ▼ │ │
│ └──────────────────┘ │
└──────────────────────┘
```

## Edge Cases Handled

### 1. No Branches Available

```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ No branches available                        │
│                                              │
│ ℹ️  Unable to load branches from repository  │
└──────────────────────────────────────────────┘
```

### 2. Only One Branch

```
┌──────────────────────────────────────────────┐
│ Switch to different branch                   │
│                                              │
│ [main (current)                  ▼]  [🔀 Switch Branch] │
│                                      ^^^^^^^^^^^^^       │
│                                      Disabled            │
│                                              │
│ ℹ️  Only one branch available                │
└──────────────────────────────────────────────┘
```

### 3. Authentication Required

If the repository requires authentication and credentials are missing:

```
┌─────────────────────────────────────────────────────────────┐
│ ✗ Failure                                               [×] │
├─────────────────────────────────────────────────────────────┤
│ Unable to load git branches                                 │
│                                                             │
│ Details: Authentication required. Please update Git         │
│ credentials in Advanced configuration.                      │
└─────────────────────────────────────────────────────────────┘
```

## Performance Considerations

### Loading Optimization

1. **Lazy Loading**: Branches are loaded only when the stack details page is accessed
2. **Caching**: Git service caches branch lists for 5 minutes
3. **Debouncing**: API calls are debounced to prevent excessive requests
4. **Loading States**: Users see immediate feedback during operations

### Network Indicators

```
Initial Load:
[Loading...] → [Branches loaded] (< 2 seconds)

Branch Switch:
[Switching...] → [Success/Error] (5-30 seconds depending on stack size)
```

## Integration Points

### With Existing Features

1. **Git Auto-Update**: Works alongside existing auto-update settings
2. **Pull and Redeploy**: Branch selector doesn't interfere with manual pull/redeploy
3. **Advanced Configuration**: Can manually specify git reference if needed
4. **Environment Variables**: Preserved during branch switch
5. **Stack Options**: Prune settings maintained

## Testing Scenarios

### Visual Testing Checklist

- [ ] Branch selector appears for git-based stacks
- [ ] Branch selector does NOT appear for non-git stacks
- [ ] Dropdown shows all available branches
- [ ] Current branch is marked and disabled
- [ ] Loading spinner shows during branch fetch
- [ ] Switching spinner shows during branch switch
- [ ] Success notification appears on successful switch
- [ ] Error notification shows appropriate message on failure
- [ ] Page reloads after successful switch
- [ ] Current branch indicator updates after switch

## Conclusion

This visual guide demonstrates the comprehensive UI enhancements made to support the branch switching feature. The design prioritizes:

- **Usability**: Clear, intuitive interface
- **Safety**: Confirmation dialogs prevent accidental switches
- **Feedback**: Loading states and notifications keep users informed
- **Accessibility**: Keyboard navigation and screen reader support
- **Responsiveness**: Works on all device sizes

The feature integrates seamlessly with existing Portainer UI patterns and provides a smooth user experience for managing Git-based stacks.
