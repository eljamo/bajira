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
- [ ] bajira update workspace --workspace_id
- [ ] bajira archive workspace --workspace_id
- [ ] bajira unarchive workspace --workspace_id

### Board Commands

- [ ] bajira create board --workspace_id
- [ ] bajira delete board --workspace_id --board_id
- [ ] bajira update board --workspace_id --board_id
- [ ] bajira archive board --workspace_id --board_id
- [ ] bajira unarchive board --workspace_id --board_id

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
- [ ] bajira list boards
  - [ ] bajira list boards --workspace_id
  - [ ] bajira list boards --workspace_id --all
  - [ ] bajira list boards --workspace_id --archived
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
