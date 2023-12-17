# Todo-Gofr #
This is a simple Todo List website built with gofr for the server-side and HTML, CSS, and JavaScript for the frontend.

### Features ###  
* Create Todos: Add new tasks to your todo list by providing a title and status.  
* Read Todos: View all the todos in your list.  
* Update Todos: Modify existing todos by updating their title and status.  
* Delete Todos: Remove unwanted tasks from your list.    

### Technologies Used ###
* Server-side:
  * gofr: A lightweight Go web framework.
* Frontend:
  * HTML: Markup language for structuring the web page.
  * CSS: Stylesheet language for styling the web page.
  * JavaScript: Programming language for dynamic behavior and interactions.
* Database
  * MySQL: As part of the backend stack, MySQL serves as the relational database to store and manage data for the Todo List application.

### Setting Up a Local Database and Table with SQL ###
Write the following code on mySQL command line client:
* Create a database
  `CREATE DATABASE your_database_name;`
* Switch database
  `USE your_database_name;`
* Create a table
  `CREATE TABLE your_table_name (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255), status BOOL DEFAULT 0);`

Note: Make sure to change the table name and database name in main.go file according to your choice. By default your_database_name = "Todo" and your_table_name = "todos". And also edit username, password and host of mySQL server in main.go file.

### How to run ###

Run server  
* `cd Backend`
* `go run main.go`
* The server will be running at `http://localhost:8000`.

Run client
* Open the HTML file using Live Server
* You can view your website at a URL like `http://127.0.0.1:5500/index.html`

### (**All changes will appear after a quick refresh.**)

### Screenshots ###

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/48bc9e6a-e285-400a-bbe4-23ac4daceddd)
Type your task

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/024129ed-518b-4c11-b241-a48f7ce6bb4f)

Press the + sign to add the task in your todolist

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/67661e18-cd77-4c5b-b3bd-805a6dd8f507)

The task will be added to your todolist with the default value of status as "not done"

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/f2bd5e1e-0aa7-4475-afad-7bf6a5f0e6e8)

To delete the task click on the bin icon

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/0625365d-5575-414b-a6de-9754569a7849)

To update the status of task, click on the green tick icon

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/1ae7efaf-2ffd-416a-bf91-41314dae80b2)

![image](https://github.com/Shreya1201/Todo-Gofr/assets/93670796/c69f64a7-bd75-4b02-a998-1c405eacbb82)











