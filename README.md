# N-Queens_in_Go

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all a distinct solutions to the n-queens puzzle. 

## HOW to RUN
### 1. Run it into the backend dir to turn on the server
```
go run ./cmd/api
```
tip: Now you are hosting in your computer the backend of this program. Because of that, u need to use a POST method to send what you want

### 2. Example of POST methods

DFS:<br>
```
curl -X POST http://localhost:8080/solve \
  -H "Content-Type: application/json" \
  -d '{"n":8,"algorithm":"dfs"}'
```
<br>

BFS: <br>
```
curl -X POST http://localhost:8080/solve \
  -H "Content-Type: application/json" \
  -d '{"n":8,"algorithm":"bfs"}'
```
<br>
<br>
After the server runs the informations, it will send back to you how much nodes was needed and the time to solve the problem
