# Nimnus Color Themes

## Development Guide

### Inspecting Token Scopes

To see the scope of any code element in your editor, use the **Developer: Inspect Editor Tokens and Scopes** command. This will show you the token type and scope information for the element under your cursor.

### Theme Development Tools

1. Use the **Preferences: Color Theme** command to switch between themes and see how they affect your code
2. Generate a starting point from your current settings with the **Developer: Generate Color Theme From Current Settings** command

### Testing Your Theme

You can test color customizations immediately using the **workbench.colorCustomizations** and **editor.tokenColorCustomizations** settings.

### Color Theme Structure

Based on the VS Code API, your color theme should include:
- `tokenColors`: Array of scope-based color rules
- `colors`: Workbench UI colors
- `semanticHighlighting`: Enable semantic token support
- `semanticTokenColors`: Semantic-based color rules

Use the token inspector on your markdown file to see scopes like `markup.heading.markdown`, `markup.italic.markdown`, `string.other.link.markdown`, etc.
