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
- bajira delete task --workspace_key --board_key --task_number
- bajira update task --workspace_key --board_key --task_number
- bajira assign task
  - bajira assign task --workspace_key --board_key --task_number --to_self
  - bajira assign task --workspace_key --board_key --task_number --subtask_number --to_self
  - bajira assign task --workspace_key --board_key --task_number --assingee
  - bajira assign task --workspace_key --board_key --task_number --subtask_number --assignee
- bajira unassign task
  - bajira unassign task --workspace_key --board_key --task_number
  - bajira unassign task --workspace_key --board_key --task_number --subtask_number
- bajira move task
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --subtask_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --to_task_number
  - bajira move task --workspace_key --from_board_key --to_board_key --task_number --subtask_number --to_task_number
- bajira flag task
  - bajira flag task --workspace_key --from_board_key --to_board_key --task_number
  - bajira flag task --workspace_key --from_board_key --to_board_key --task_number --subtask_number

### List Commands

- bajira list workspaces
- bajira list boards --workspace_key
- bajira list tasks
  - bajira list tasks --workspace_key --board_key --assignee
  - bajira list tasks --workspace_key --board_key --self
  - bajira list tasks --workspace_key --assignee
  - bajira list tasks --workspace_key --self

### Kanban Commands

- bajira kanban
  - bajira kanban --workspace_key --board_key --assignee
  - bajira kanban --workspace_key --board_key --self
  - bajira kanban --workspace_key --assignee
  - bajira kanban --workspace_key --self

### Timer Commands

- bajira timer start task
  - bajira timer start task --workspace_key --board_key --task_number
  - bajira timer start task --workspace_key --board_key --task_number --subtask_number
  - bajira timer start task --workspace_key --board_key --task_number --pomodoro
  - bajira timer start task --workspace_key --board_key --task_number --5217
- bajira timer stop

### Config Commands

- bajira config self
  - bajira config self --global --name
  - bajira config self --workspace_key --name
