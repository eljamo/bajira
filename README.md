# Bajira (バジラ)

A bug tracker, issue tracker, and project management tool written in Go. You have Workspaces, which have Boards, which have Tasks, and Tasks can have Subtasks. All data for Workspaces, Boards, Tasks and Subtasks will be stored in TOML files in some way. Tasks and Subtaks might use Markdown files. The plan is to make this distributed via `git`. If the workspace detects it is or within a `git` directory, support for `fetch`, `pull`, `commit`, and `push` will be enabled.

## Example of a Bajira folder

```
.
└── workspace/
    ├── WORK/
    │   ├── .git/
    │   ├── config.toml
    │   └── board/
    │       ├── config.toml
    │       ├── DEV/
    │       │   ├── config.toml
    │       │   └── task/
    │       │       ├── config.toml
    │       │       ├── 1/
    │       │       │   ├── config.toml
    │       │       │   ├── description.md
    │       │       │   └── subtask/
    │       │       │       └── 1/
    │       │       │           ├── config.toml
    │       │       │           └── description.md
    │       │       └── 2/
    │       │           ├── config.toml
    │       │           └── description.md
    │       └── PRODUCT/
    │           ├── config.toml
    │           └── task/
    │               ├── config.toml
    │               ├── 1/
    │               │   ├── config.toml
    │               │   └── description.md
    │               ├── 2/
    │               │   ├── config.toml
    │               │   └── description.md
    │               └── 3/
    │                   ├── config.toml
    │                   └── description.md
    └── PERSONAL/
        ├── .git/
        ├── config.toml
        └── board/
            ├── config.toml
            └── PROJECT1/
                └── ...
```

## Planned commands, flags are subject to change

### Workspace Commands

- bajira create workspace
- bajira delete workspace --workspace_key
- bajira update workspace --workspace_key
- bajira archive workspace --workspace_key
- bajira unarchive workspace --workspace_key

### Board Commands

- bajira create board --workspace_key
- bajira delete board --workspace_key --board_key
- bajira update board --workspace_key --board_key
- bajira archive board --workspace_key --board_key
- bajira unarchive board --workspace_key --board_key

### Task Commands

- bajira create task
  - bajira create task --workspace_key --board_key
  - bajira create task --workspace_key --board_key --task_number
- bajira delete task
  - bajira delete task --workspace_key --board_key --task_number
  - bajira delete task --workspace_key --board_key --task_number --subtask_number
- bajira update task
  - bajira update task --workspace_key --board_key --task_number
  - bajira update task --workspace_key --board_key --task_number --subtask_number
- bajira archive task
  - bajira archive task --workspace_key --board_key --task_number
  - bajira archive task --workspace_key --board_key --task_number --subtask_number
- bajira unarchive task
  - bajira unarchive task --workspace_key --board_key --task_number
  - bajira unarchive task --workspace_key --board_key --task_number --subtask_number
- bajira assign task
  - bajira assign task --workspace_key --board_key --task_number --to_self
  - bajira assign task --workspace_key --board_key --task_number --subtask_number --to_self
  - bajira assign task --workspace_key --board_key --task_number --assingee
  - bajira assign task --workspace_key --board_key --task_number --subtask_number --assignee
- bajira unassign task
  - bajira unassign task --workspace_key --board_key --task_number --self
  - bajira unassign task --workspace_key --board_key --task_number --subtask_number --self
  - bajira unassign task --workspace_key --board_key --task_number --assignee
  - bajira unassign task --workspace_key --board_key --task_number --subtask_number --assignee
- bajira move task
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --subtask_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --to_task_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --subtask_number --to_task_number
- bajira flag task
  - bajira flag task --workspace_key --from_board_key --to_board_key --task_number
  - bajira flag task --workspace_key --from_board_key --to_board_key --task_number --subtask_number

#### Other Task Commands

I'll probably add support for these kinds of attributes on tasks

- Blocking/Blocked by
- Depended on by/Depends on
- Tests/Tested by
- Relates to
- Watched by

### List Commands

- bajira list workspaces
  - bajira list workspaces --all
  - bajira list workspaces --archived
- bajira list boards
  - bajira list boards --workspace_key
  - bajira list boards --workspace_key --all
  - bajira list boards --workspace_key --archived
- bajira list tasks
  - bajira list tasks --workspace_key --board_key --assignee
  - bajira list tasks --workspace_key --board_key --self
  - bajira list tasks --workspace_key --board_key --flagged
  - bajira list tasks --workspace_key --board_key --archived
  - bajira list tasks --workspace_key --assignee
  - bajira list tasks --workspace_key --self
  - bajira list tasks --workspace_key --flagged
  - bajira list tasks --workspace_key --archived

### Kanban Commands

- bajira kanban
  - bajira kanban --workspace_key --board_key --assignee
  - bajira kanban --workspace_key --board_key --self
  - bajira kanban --workspace_key --board_key --flagged
  - bajira kanban --workspace_key --assignee
  - bajira kanban --workspace_key --self
  - bajira kanban --workspace_key --flagged

### Timer Commands

- bajira timer start task
  - bajira timer start task --workspace_key --board_key --task_number
  - bajira timer start task --workspace_key --board_key --task_number --subtask_number
  - bajira timer start task --workspace_key --board_key --task_number --pomodoro
  - bajira timer start task --workspace_key --board_key --task_number --5217
- bajira timer stop

### Config Commands

- bajira config
  - bajira config --data_directory
  - bajira config --config_directory
  - bajira config --cache_directory
  - bajira config --default_workspace_key
  - bajira config --locale

### Assignee Commands

- bajira assignee
  - bajira assignee --name
  - bajira assignee --global --name
  - bajira assignee --workspace_key --name

### Set Commands

This will allow you to set the current workspace so for other commands which have a --workspace_key flag like board, task, list, kanban, timer, and assignee. You can leave it out and it'll use the currently set workspace

- bajira set workspace --workspace_key
