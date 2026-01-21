# Development Workflow

When working on issues:

1. **Pick an issue** from GitHub
   - List open issues with `gh issue list`
   - View issue details with `gh issue view {number}`

2. **Sync with upstream** before starting
   - `git checkout main`
   - `git pull origin main`

3. **Create a branch** for your work
   - `git checkout -b branch-name`
   - Use descriptive kebab-case names (e.g., `add-button-component`)

4. **Implement the changes**
   - Write code following existing patterns
   - Add tests
   - Run `go test ./...` to verify

5. **Commit your changes**
   - Use the `/git` skill for commit message conventions
   - Reference issues: "Fixes #123" or "See #123"

6. **Push to the fork**
   - `git push -u fork branch-name`
   - The fork remote is `maragubot/gomponents-shadcn-ui`

7. **Create a PR** to the main repo
   - `gh pr create --repo maragudk/gomponents-shadcn-ui`
   - Include a summary and test plan

8. **Address review feedback** (see Code Review Workflow below)

## Component Development

- **Style reference**: Use shadcn/ui **New York v4** styles from the local checkout at `ui/apps/v4/registry/new-york-v4/ui/`
- The shadcn/ui repository is checked out locally at `ui/` - use this instead of fetching from the web
- When adding a new component, update the demo page in `cmd/demo/main.go`
- When updating a component, check that any new style variables are added to `tailwind.css`

## Code Review Workflow

When asked to review code on a branch:

1. **Review the changes** using the `/code-review` skill
   - This dispatches two competing review agents for thorough coverage
   - They examine both architecture and implementation

2. **Submit the review on GitHub**
   - Post an overall summary using `gh api repos/{owner}/{repo}/pulls/{number}/reviews`
   - Add line-specific comments using `gh api repos/{owner}/{repo}/pulls/{number}/comments`

3. **Check for feedback** on your review comments
   - Fetch comments with `gh api repos/{owner}/{repo}/pulls/{number}/comments`
   - Look for replies (comments with `in_reply_to_id`)

4. **Address feedback**
   - Make code changes as requested
   - Create GitHub issues for items to track separately
   - Commit and push changes

5. **Reply to comments** confirming actions taken
   - Use `gh api repos/{owner}/{repo}/pulls/{number}/comments/{comment_id}/replies`

6. **Resolve conversations** when done
   - Get thread IDs via GraphQL query on `reviewThreads`
   - Resolve with `resolveReviewThread` mutation

7. **Merge the PR** (if you have permissions)
   - Use `gh pr merge {number} --squash`
