# Bajira (バジラ)

A bug tracker, issue tracker, and project management tool written in Go. You have Workspaces, which have Boards, which have Tasks, and Tasks can have Subtasks. All data for Workspaces, Boards, Tasks and Subtasks will be stored in TOML files in some way. Tasks and Subtaks might use Markdown files. The plan is to make this distributed via `git`. If the workspace detects it is or within a `git` directory, support for `fetch`, `pull`, `commit`, and `push` will be enabled.

## Example of a Bajira folder

```
.
└── workspaces/
    ├── work/
    │   ├── .git
    │   ├── config.toml
    │   └── boards/
    │       ├── config.toml
    │       ├── DEV/
    │       │   ├── config.toml
    │       │   └── tasks/
    │       │       ├── config.toml
    │       │       ├── 1/
    │       │       │   └── config.toml
    │       │       └── 2/
    │       │           └── config.toml
    │       └── PRODUCT/
    │           ├── config.toml
    │           └── tasks/
    │               ├── config.toml
    │               ├── 1/
    │               │   └── config.toml
    │               ├── 2/
    │               │   └── config.toml
    │               └── 3/
    │                   └── config.toml
    └── personal/
        ├── .git
        ├── config.toml
        └── boards/
            ├── config.toml
            └── PROJECT1/
                └── ...
```

## Planned commands, flags are subject to change

### Config Commands

- [ ] bajira update config
  - [ ] bajira update config --data_directory
  - [ ] bajira update config --default_workspace_id
  - [ ] bajira update config --locale
  - [ ] bajira update config --accessible-mode

### Assignee Commands

- [ ] bajira update assignee
  - [ ] bajira update assignee --name
  - [ ] bajira update assignee --global --name
  - [ ] bajira update assignee --workspace_id --name

### Workspace Commands

- [x] bajira create workspace
- [x] bajira delete workspace --workspace_id
- [x] bajira update workspace --workspace_id
- [x] bajira archive workspace --workspace_id
- [x] bajira unarchive workspace --workspace_id

### Board Commands

- [ ] bajira create board --workspace_id
- [x] bajira delete board --workspace_id --board_id
- [ ] bajira update board --workspace_id --board_id
- [x] bajira archive board --workspace_id --board_id
- [x] bajira unarchive board --workspace_id --board_id

### Task Commands

- [ ] bajira create task --workspace_id --board_id
- [ ] bajira delete task --workspace_id --board_id --task_id
- [ ] bajira update task --workspace_id --board_id --task_id
- [ ] bajira archive task --workspace_id --board_id --task_id
- [ ] bajira unarchive task --workspace_id --board_id --task_id
- [ ] bajira assign task
  - [ ] bajira assign task --workspace_id --board_id --task_id --to_self
  - [ ] bajira assign task --workspace_id --board_id --task_id --assingee
- [ ] bajira unassign task
  - [ ] bajira unassign task --workspace_id --board_id --task_id --self
  - [ ] bajira unassign task --workspace_id --board_id --task_id --assignee
- [ ] bajira move task --workspace_id --from_board_id --to_board_id --task_id
- [ ] bajira flag task --workspace_id --from_board_id --to_board_id --task_id

#### Other Task Commands

I'll probably add support for these kinds of attributes on tasks

- [ ] Blocking/Blocked by
- [ ] Depended on by/Depends on
- [ ] Tests/Tested by
- [ ] Relates to
- [ ] Watched by

### List Commands

- [x] bajira list workspaces
  - [x] bajira list workspaces --all
  - [x] bajira list workspaces --archived
- [x] bajira list boards
  - [x] bajira list boards --workspace_id
  - [x] bajira list boards --workspace_id --all
  - [x] bajira list boards --workspace_id --archived
- [ ] bajira list tasks
  - [ ] bajira list tasks --workspace_id --board_id --assignee
  - [ ] bajira list tasks --workspace_id --board_id --self
  - [ ] bajira list tasks --workspace_id --board_id --flagged
  - [ ] bajira list tasks --workspace_id --board_id --archived
  - [ ] bajira list tasks --workspace_id --assignee
  - [ ] bajira list tasks --workspace_id --self
  - [ ] bajira list tasks --workspace_id --flagged
  - [ ] bajira list tasks --workspace_id --archived

### Kanban Commands

- [ ] bajira kanban
  - [ ] bajira kanban --workspace_id --board_id --assignee
  - [ ] bajira kanban --workspace_id --board_id --self
  - [ ] bajira kanban --workspace_id --board_id --flagged
  - [ ] bajira kanban --workspace_id --assignee
  - [ ] bajira kanban --workspace_id --self
  - [ ] bajira kanban --workspace_id --flagged

## Unplanned Commands

I may create these sets of commands after the first planned bunch are done

### Timer Commands

- [ ] bajira start timer
  - [ ] bajira start timer --context
  - [ ] bajira start timer --context --pomodoro
  - [ ] bajira start timer --context --5217
  - [ ] bajira start timer --context --title
  - [ ] bajira start timer --context --title --pomodoro
  - [ ] bajira start timer --context --title --5217
- [ ] bajira start task timer
  - [ ] bajira start task timer --workspace_id --board_id --task_id
  - [ ] bajira timer task timer --workspace_id --board_id --task_id --pomodoro
  - [ ] bajira timer task timer --workspace_id --board_id --task_id --5217
- [ ] bajira stop timer

### Goal Commands

- [ ] bajira set yearly goal
- [ ] bajira set quaterly goal
  - [ ] bajira set quaterly goal --yearly_goal_id
- [ ] bajira set monthly goal
  - [ ] bajira set monthly goal --yearly_goal_id
- [ ] bajira set weekly goal
  - [ ] bajira set weekly goal --yearly_goal_id
- [ ] bajira set daily goal
  - [ ] bajira set daily goal --yearly_goal_id

### Schedule Commands

- [ ] bajira schedule task
  - [ ] bajira schedule task --workspace_id --board_id --task_id --date --time-start --time-end
  - [ ] bajira schedule task --workspace_id --board_id --task_id --date_time_start --date_time_end
- [ ] bajira schedule event
  - [ ] bajira schedule event --type --title --description --date --time-start --time-end --workspace_id
  - [ ] bajira schedule event --type --title --description --date_time_start --date_time_end --workspace_id
- [ ] bajira assign scheduled task
  - [ ] bajira assign scheduled task --workspace_id --to_self
  - [ ] bajira assign scheduled task --workspace_id --assingee
- [ ] bajira unassign scheduled task
  - [ ] bajira unassign scheduled task --workspace_id --self
  - [ ] bajira unassign scheduled task --workspace_id --assignee
- [ ] bajira assign scheduled event
  - [ ] bajira assign scheduled event --workspace_id --to_self
  - [ ] bajira assign scheduled event --workspace_id --assingee
- [ ] bajira unassign scheduled event
  - [ ] bajira unassign scheduled event --workspace_id --self
  - [ ] bajira unassign scheduled event --workspace_id --assignee
- bajira show schedule
  - [ ] bajira show schedule --self
  - [ ] bajira show schedule --assignee

### Log Commands

- [ ] bajira log distraction
  - [ ] bajira log distraction
