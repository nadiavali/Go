### 1. **Basic CRUD API**

**Challenge**: Build a simple REST API for managing a list of `tasks`. Each task should have an `ID`, `title`, `description`, and `completed` status.

- **Requirements**:
  - Implement basic CRUD operations (`Create`, `Read`, `Update`, `Delete`) for tasks.
  - Use Go’s `net/http` package without any third-party libraries.
  - Store data in memory (e.g., a map or slice).
  - Return appropriate HTTP status codes for each operation.

- **Example Endpoints**:
  - `POST /tasks`: Create a new task.
  - `GET /tasks`: Retrieve all tasks.
  - `GET /tasks/{id}`: Retrieve a specific task.
  - `PUT /tasks/{id}`: Update a specific task.
  - `DELETE /tasks/{id}`: Delete a specific task.

**Skills Assessed**: HTTP handling, JSON serialization, REST principles, error handling, working with maps and slices.

---

### 2. **Concurrency with Goroutines and Channels**

**Challenge**: Write a function that simulates a batch processing system for tasks. Each task takes a random amount of time (up to 1 second) to complete.

- **Requirements**:
  - Implement a `processTasks` function that receives a slice of tasks.
  - Each task should be processed concurrently.
  - Use a `channel` to signal when each task has completed.
  - Once all tasks are processed, print a message indicating the batch is complete.

- **Example**:
  ```go
  tasks := []string{"task1", "task2", "task3"}
  processTasks(tasks)
  ```

- **Expected Output**:
  ```
  task1 completed
  task2 completed
  task3 completed
  All tasks completed
  ```

**Skills Assessed**: Concurrency, goroutines, channels, synchronization.

---

### 3. **Implement a Rate Limiter**

**Challenge**: Create a basic rate limiter that restricts the number of API requests a user can make in a given timeframe.

- **Requirements**:
  - Limit requests to `n` requests per minute per user.
  - Use an in-memory map to store user request counts.
  - Reset the count every minute.
  - If a user exceeds the limit, return an HTTP `429 Too Many Requests` status code.

- **Bonus**: Use Go’s `time.AfterFunc` or `time.Ticker` to handle reset timing.

**Skills Assessed**: Time management, map handling, HTTP response handling, Go’s time library.

---

### 4. **String Manipulation - Alternating Case**

**Challenge**: Write a function `toAlternatingCase` that takes a string and returns a new string with each letter’s case reversed.

- **Requirements**:
  - Do not use any external packages.
  - The function should ignore non-alphabetic characters.
  
- **Example**:
  ```go
  toAlternatingCase("Hello World!") // Output: "hELLO wORLD!"
  ```

**Skills Assessed**: String manipulation, handling of ASCII/Unicode characters.

---

### 5. **Error Handling - Retry Mechanism**

**Challenge**: Implement a `retry` function that accepts a function as a parameter and retries it up to `n` times if it fails. Consider the function "failed" if it returns an error.

- **Requirements**:
  - If the function succeeds (returns no error), return the result immediately.
  - If it fails, retry up to `n` times with a small delay between attempts.
  - Return an error if all attempts fail.

- **Example**:
  ```go
  func exampleFunction() error {
      // Some logic that may fail
  }

  result, err := retry(exampleFunction, 3)
  ```

**Skills Assessed**: Error handling, function types, retry logic, Go’s time package.

---

### 6. **File Processing - Word Count**

**Challenge**: Write a program that reads a text file and counts the occurrences of each word, ignoring case and punctuation.

- **Requirements**:
  - Use Go’s standard library to read the file.
  - Return a map of words and their counts.
  - Handle errors appropriately, such as file not found.

- **Bonus**: Sort the output by word count before printing.

**Skills Assessed**: File handling, map usage, string manipulation, error handling.

---

### 7. **Data Processing with Structs - Sorting and Filtering**

**Challenge**: Create a list of `Product` structs, each with a `Name`, `Price`, and `Category`. Write functions to:

1. Sort products by price (ascending or descending).
2. Filter products by category.

- **Requirements**:
  - Use slices and sort the data in place.
  - Implement sorting and filtering functions separately.
  - Display the filtered or sorted list after each operation.

- **Example**:
  ```go
  type Product struct {
      Name     string
      Price    float64
      Category string
  }

  products := []Product{...}
  ```

**Skills Assessed**: Structs, sorting, filtering, slice manipulation.

---

### 8. **Simple Authentication System**

**Challenge**: Implement a basic login system where users can register and log in.

- **Requirements**:
  - Use an in-memory map to store user credentials (`username` and `password`).
  - Write endpoints for registering and logging in users.
  - Ensure passwords are stored in a hashed form using `bcrypt`.

- **Example Endpoints**:
  - `POST /register`: Registers a new user.
  - `POST /login`: Logs in a user and returns a simple authentication token.

**Skills Assessed**: Basic authentication, HTTP endpoints, bcrypt usage, map handling, password security.

---

### 9. **Palindrome Checker**

**Challenge**: Write a function that checks if a given string is a palindrome, ignoring spaces, punctuation, and case.

- **Requirements**:
  - The function should return `true` if the string is a palindrome and `false` otherwise.
  - Ignore non-alphabet characters.
  
- **Example**:
  ```go
  isPalindrome("A man, a plan, a canal, Panama") // Output: true
  ```

**Skills Assessed**: String manipulation, palindrome logic, ASCII/Unicode handling.

---

### 10. **Temperature Conversion API**

**Challenge**: Create a simple API with endpoints to convert temperatures between Celsius, Fahrenheit, and Kelvin.

- **Requirements**:
  - Implement endpoints like `/toCelsius`, `/toFahrenheit`, and `/toKelvin`.
  - Each endpoint should accept a temperature and return the converted result.
  - Handle basic error cases (e.g., invalid input).

- **Example Endpoints**:
  - `GET /toCelsius?temp=100&scale=F`: Convert 100°F to Celsius.

**Skills Assessed**: Basic math, HTTP endpoints, parameter handling, error handling.

---


You're correct to question if these challenges are ideal for a junior-level interview. For a junior position, some of these tasks might be a bit advanced, especially considering the 1.5-hour timeframe and the likelihood that you're relatively new to Go. Here’s a revised list of more beginner-friendly challenges tailored to test junior-level skills while still aligning with JustWatch's requirements:


---

### 11. **Basic CRUD API for Managing Streaming Data**

   - **Challenge**: Build a basic REST API that allows creating, reading, updating, and deleting items from a list of movies or shows.
   - **Skills Tested**: Basic REST API creation, JSON handling, struct and slice manipulation.
   - **Hints**: Focus on one or two endpoints, like `POST /movies` and `GET /movies`, to keep it manageable.

---

### 12. **Fetch and Process External JSON Data**

   - **Challenge**: Simulate fetching JSON data from a fake API (you can just define the data in a variable for simplicity) and convert it into a Go struct, then display it in a formatted response.
   - **Skills Tested**: JSON parsing and struct usage.
   - **Hints**: Keep the JSON data simple and focus on encoding/decoding.

---

### 13. **Rate Limiter for API Requests**

   - **Challenge**: Implement a basic rate limiter that limits API requests to a maximum of 5 requests per minute.
   - **Skills Tested**: Time-based control, basic concurrency.
   - **Hints**: Use a counter and reset it every 60 seconds.

---

### 4. **Filter and Sort Data from a List**

   - **Challenge**: Write a function that takes a list of movies (with fields like `Title`, `Genre`, and `Year`) and returns a filtered list based on genre and sorted by release year.
   - **Skills Tested**: Filtering and sorting, slice manipulation.
   - **Hints**: Use Go’s built-in `sort` package for sorting.

---

### 5. **Simulate Basic Concurrent Tasks**

   - **Challenge**: Simulate processing multiple tasks (e.g., "fetching data", "analyzing data") concurrently. Each task should take a random time between 100ms and 1 second to complete.
   - **Skills Tested**: Concurrency with goroutines and channels.
   - **Hints**: Focus on launching goroutines for each task and using channels to receive completion signals.

---

### 6. **Simple Middleware for Request Logging**

   - **Challenge**: Implement a middleware that logs each request (method and URL) before passing it to the handler.
   - **Skills Tested**: Middleware patterns, HTTP handling.
   - **Hints**: Use `http.HandlerFunc` and wrap it around existing handlers.

---

### 7. **Pagination for API Results**

   - **Challenge**: Implement pagination for a `GET /movies` endpoint that returns a large list of movies. Allow parameters like `page` and `limit` to control the number of results returned.
   - **Skills Tested**: Pagination, handling query parameters, slice manipulation.
   - **Hints**: Use slices to simulate paginated results.

---

### 8. **Basic Caching for API Responses**

   - **Challenge**: Add basic caching for API responses. If a request for the same endpoint with the same parameters is made within a short time, return the cached result.
   - **Skills Tested**: Simple caching logic, struct and map usage.
   - **Hints**: Use a map to store cached responses with a timestamp for expiration.

---

### 9. **Implement a Health Check Endpoint**

   - **Challenge**: Create an API endpoint (`/health`) that returns a JSON status message (e.g., `{ "status": "OK" }`) to indicate that the service is running.
   - **Skills Tested**: Basic HTTP handling, JSON response creation.
   - **Hints**: Focus on creating a simple handler and returning static JSON data.

---

### 10. **Calculate Click-Through Rate (CTR)**

   - **Challenge**: Write a function that calculates the click-through rate given a set of impressions and clicks for various items. CTR = (clicks / impressions) * 100.
   - **Skills Tested**: Basic arithmetic and data handling.
   - **Hints**: Keep it simple with a slice of structs representing impressions and clicks, and iterate over it to calculate the CTR.

---

### Summary

These challenges are simpler, focusing on core skills rather than complex concurrency or performance optimization. Here’s why they’re better suited for a junior interview:

- **Simplicity in Scope**: Each task is manageable in 15-20 minutes, leaving time for testing and questions.
- **Foundational Skills**: These challenges cover essential backend skills (basic concurrency, REST APIs, JSON handling, data manipulation).
- **Minimal Infrastructure Complexity**: They don’t require knowledge of advanced topics like Kubernetes, distributed systems, or heavy concurrency control.
  
This list should give you confidence as a junior developer while still allowing you to demonstrate your potential in Go and backend development fundamentals. Let me know if you’d like any specific examples or practice code for these!

- **Start with easier challenges** (e.g., CRUD API, string manipulation) to warm up.
- **Progress to slightly more complex challenges** (e.g., file processing, data sorting) to test data handling and struct manipulation.
- **End with concurrency-based challenges** (e.g., rate limiter, retry mechanism) to assess basic concurrency understanding.

These challenges cover key areas relevant to junior Go developers and provide a good gauge of their foundational knowledge and readiness for real-world tasks in Go.