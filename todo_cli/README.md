
# 📝 Tasks CLI


A simple command-line task manager built in Go, supporting add, list, complete, and delete operations. Tasks are stored in a local CSV file for persistence.


## 🚀 Features

- Add tasks with a description
- List tasks (pending only, or all with -a)
- Mark tasks as complete
- Delete tasks by ID
- Stores tasks in a CSV file for easy portability
- Friendly human-readable timestamps (e.g. "2 minutes ago")


## 🛠️ Built With

- Go – programming language
- Cobra – CLI framework
- timediff – friendly time differences
- CSV (encoding/csv) – lightweight task persistence
## Installation


1. Clone the Repository

```bash
git clone https://github.com/Shashivarunreddy/Go_projects.git
cd Go_projects/todo_cli
```

2. Install dependencies:

```bash
go mod tidy
```
3. Build and install to your Go bin: 


```bash
go install .

```
⚡ Ensure $HOME/go/bin is in your $PATH so you can run tasks globally.


## 🎯 Example Workflow


```bash
tasks add "Buy groceries" // add tasks
tasks add "Finish homework"

tasks list // list all the tasks

tasks complete 1 // Marks task with id 1 as completed
tasks list -a // lists all the incompleted tasks

tasks delete 2 // Delete task with id 2
tasks list -a

```



