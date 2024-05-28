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
