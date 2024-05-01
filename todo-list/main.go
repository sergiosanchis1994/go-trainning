package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
}

type TaskList struct {
	tasks []Task
}

func (tl *TaskList) addTask(t Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) complete(index int) {
	tl.tasks[index].completed = true
}

func (tl *TaskList) edit(index int, t Task) {
	tl.tasks[index] = t
}

func (tl *TaskList) delete(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func main() {
	taskList := TaskList{}
	reader := bufio.NewReader(os.Stdin)
	for {
		printTasks(&taskList)
		var option int
		fmt.Println("Choose an option:\n",
			"1. Add task\n",
			"2. Mark task as completed\n",
			"3. Edit task\n",
			"4. Delete task\n",
			"5. Exit")

		fmt.Print("Option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task
			fmt.Print("Add task name: ")
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Add task description: ")
			t.description, _ = reader.ReadString('\n')
			taskList.addTask(t)
			fmt.Println("Task added successfully")
		case 2:
			var index int
			fmt.Print("Which task do you want to mark as a completed? (introduce its index): ")
			fmt.Scanln(&index)
			taskList.complete(index)
			fmt.Println("Task completed")
		case 3:
			var t Task
			var index int
			fmt.Print("Which task do you want to edit? (introduce its index): ")
			fmt.Scanln(&index)
			fmt.Print("Add task name: ")
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Add task description: ")
			t.description, _ = reader.ReadString('\n')
			taskList.edit(index, t)
			fmt.Println("Task edited successfully")
		case 4:
			var index int
			fmt.Print("Which task do you want todelete? (introduce its index): ")
			fmt.Scanln(&index)
			taskList.delete(index)
			fmt.Println("Task deleted")
		case 5:
			fmt.Println("Bye :)")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printTasks(tl *TaskList) {
	fmt.Println("Tasks:")
	fmt.Println("======================================================")
	for i, t := range tl.tasks {
		fmt.Printf("%d. %s - %s - Completed: %t\n", i, t.name, t.description, t.completed)
	}
	fmt.Println("======================================================")
}
