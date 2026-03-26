# Publishing

## Prerequisites

1. **Install `vsce`** — the VS Code Extension CLI:
   ```bash
   npm install -g @vscode/vsce
   ```
2. **Visual Studio Marketplace account** — sign in at [marketplace.visualstudio.com/manage](https://marketplace.visualstudio.com/manage).
   The publisher ID for this extension is **`yuriity`**.

## One-time Setup: Personal Access Token

`vsce` authenticates via a Personal Access Token (PAT) scoped to the Marketplace.

1. Go to [dev.azure.com](https://dev.azure.com) → your organisation → **User Settings → Personal access tokens**.
2. Click **New Token** and set:
   - **Name** — something like `vsce-nimnus`
   - **Organisation** — `All accessible organizations`
   - **Scopes** → **Marketplace → Manage**
3. Copy the generated token — it is shown only once.
4. Log in locally:
   ```bash
   vsce login yuriity
   # Paste the PAT when prompted
   ```
   Credentials are stored in the system keychain; you only need to do this once per machine.

## Publishing a New Version

`vsce publish` bumps the version in `package.json` and uploads the packaged VSIX to the Marketplace in one step.

| Command | What it does |
|---|---|
| `vsce publish patch` | `0.2.0` → `0.2.1` — bug fixes |
| `vsce publish minor` | `0.2.0` → `0.3.0` — new features, backward-compatible |
| `vsce publish major` | `0.2.0` → `1.0.0` — breaking changes |
| `vsce publish` | Publishes the exact version in `package.json`, no bump |

Example — releasing a patch:
```bash
vsce publish patch
```

> **Tip:** Run `vsce package` first to inspect the VSIX contents before publishing.

## After Publishing

- The new version typically appears on the Marketplace within a few minutes.
- Check the live listing at:
  `https://marketplace.visualstudio.com/items?itemName=yuriity.nimnus`
- Commit and push the version bump that `vsce publish` wrote to `package.json`:
  ```bash
  git add package.json
  git commit -m "chore: bump version to $(node -p "require('./package.json').version")"
  git push
  ```
