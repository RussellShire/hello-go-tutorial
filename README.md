# Hello Go Tutorial

Learning Go with the [Tutorial](https://go.dev/doc/tutorial/getting-started) from the docs.

A simple program that prints greetings to the terminal.

## Set up
Go must be installed

## Dev

```$ go run ./hello``` (from root)

```$ go run .``` (from /hello)

## Testing

```$ go test ./... -v```

## Deploy

### 1) Build
 
```$ go build```

---
### 2) Test (Optional)

#### On Mac/Linux:

```$ ./hello```


#### On Windows:

```$ hello.exe```

---

### 3) Find your Go Install Path:

```$ go list -f '{{.Target}}'```
(note: when copying end at 'bin', don't include '/hello')

---

### 4) Add to PATH

#### On Linux or Mac:

```$ export PATH=$PATH:/path/to/your/install/directory```

#### On Windows:

```$ set PATH=%PATH%;C:\path\to\your\install\directory```

--- 

### 5) Install
 
```$ go install```

---

### 6) Add to PATH permanently (Optional)

Export will only temporarily add the program to your PATH for one terminal window session, to make the change persist to need to update Zsh or Bash

#### On Zsh (Mac default):

Add PATH export from Step 4 to ```~/.zshrc``` (remember to end at bin)

#### On Bash (Often Linux and Windows default):git 

Add PATH export from Step 4 to ```~/.bashrc``` (remember to end at bin)

## Run

```$ hello```
