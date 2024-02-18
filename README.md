# Command Dockerize App
make build

# Evidence Result in folder ./result

# HOW TO RUN THE APP
1. make sure docker is ready
2. make build
3. make up
4. make connect for mysql. Like this picture:![make connect mysql](result/makeconnect.PNG)
5. INSERT Query to mysql from database folder. Like this: ![insert query](result/INSERTQUERY.PNG)
6. API is ready. Like this picture: ![Alt text](result/API_READY1.PNG)
![Alt text](result/API_READY2.PNG)

# ERD
![Alt text](<result/ERD SPRINT ASIA.png>)
## API Endpoints

### GET /tasks
- Description: Get all tasks
- Response:
  - Status: 200 OK
  - Body: Array of task objects

### GET /tasks/history
- Description: Get all tasks history
- Response:
  - Status: 200 OK
  - Body: Array of task objects

### POST /tasks
- Description: Create a new task
- Request Body:
  - task: Task object
- Response:
  - Status: 201 Created
  - Body: Created task object

### PUT /tasks/{id}
- Description: Update a specific task by ID
- Parameters:
  - id: ID of the task
- Request Body:
  - task: Updated task object
- Response:
  - Status: 200 OK
  - Body: Updated task object

### DELETE /tasks/{id}
- Description: Delete a specific task by ID
- Parameters:
  - id: ID of the task
- Response:
  - Status: 200 Status ok

### PUT /tasks/{id}/checklist
- Description: checklist a specific task by ID
- Parameters:
  - id: ID of the task
- Response:
  - Status: 200 Status ok

### POST /tasks/{id}/subtasks/{subtaskid}
- Description: Create a new subtask
- Request Body:
  - task: Task object
- Response:
  - Status: 201 Created
  - Body: Created task object

### PUT /tasks/{id}/substasks/{subtaskid}
- Description: Update a specific subtask by subtaskID and taskid
- Parameters:
  - id: ID of the task
  - subtaskid: id subtask
- Request Body:
  - task: Updated task object
- Response:
  - Status: 200 OK
  - Body: Updated task object

### DELETE /tasks/{id}/{subtaskid}
- Description: Delete a specific task by subtaskID and taskid
- Parameters:
  - id: ID of the task
  - subtaskid: id of the subtask
- Response:
  - Status: 200 Status ok

### PUT /tasks/{id}/subtasks/{subtaskid}/checklist
- Description: checklist a specific task by subtask ID and checklist.
- Parameters:
  - id: ID of the task
  - subtaskid: id of the subtask
- Response:
  - Status: 200 Status ok


## POSTMAN COLLECTION
postman-collection-TASKLIST.JSON in the root directory of this project.