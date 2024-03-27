# Development

### Naming conventions

#### Git branches and Commits
We try to follow the following conventions when naming branches and messaging in the commit message. When creating a new branch, prefix it with the issue ID of the Story which has subtasks, and try to make a commit entry for each task where possible. Reference the Task ID in the commit message and prefix the message with the appropriate commit type such as;

"feat(TSDK-33): user can log out of the application"
"fix(TSDK-44): error message no longer appears after dismissal"


feat - for new feature
fix - for bug fixes


In  the case where the task or a story is a stand alone, still prefix the branch with the issue ID and reference it in the commit message also.


#### Code Review
Since the branch is limited to a single story or single task than the PR shouldn't be large, which is purposeful. A PR should not be left open for too long without code review, therefore take the opportunity, where it is possible, to review the code of others so that any issue may be quickly resolved or the branch may be merged.
