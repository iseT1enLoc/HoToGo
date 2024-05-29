# HoToGo-LESSON2

# CREATE A MODULE

Outline some noticable points from this part:

1. **Create a single module**
2. **Call func from another module**
3. **Syntax notice in this part**

## 1. Create a single module

To allow dependency tracking, we type “**go mod init *module-path/module-name***”

In the above lesson, we use [example.com/greetings](http://example.com/greetings) for the corresponding part:

```go
go mod init example.com/greetings
```

In this part we learn about how to create a function in go lang, let take this function as an example to analyse:

```go
func PrintGreetingsMessage(msg string, emp_name string) (string, error) {
	if emp_name == "" || msg == "" {
		return "", errors.New("The msg or employee name might be empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, %v!.%v\n", emp_name, msg, GiveRandomWishes())
	return message, nil
}
```

- **Function parameters:** (parameter-name parameter-type), it is completely contrary to C or CPP which take the parameter type in the first position.
- **Return value:** Go lang allow multiple return values which make it very convenient. Once we want to return multiple value we must encapsulate them in parenthesis “**()**”.
- **Errors package:** this package allow returning message when errors do occurs in the code

## 2. Call function from another module

In Go, code executed as an application must be in a `main` package.

In this  part, we do the two command lines sequently,
1.`go mod edit -replace example.com/greetings=../greetings`

**go mod edit:** This is the main command and tells Go to edit the `go.mod` file, which manages dependencies in your project.
**-replace:** This flag specifies that you want to replace an existing dependency with a different source.

**example.com/greetings:** This is the path of the module you want to replace. It indicates a module named`greetings` located under the domain `example.com`

**..greetings:** This is the path to the new source that will replace the original dependency. The `../` signifies going up one directory level from the current working directory. In this case, it points to a local directory named `greetings` within your project workspace.

2.`go mod tidy`

`go mod tidy` is a valuable tool for keeping your Go project's dependencies organized, up-to-date, and free of unnecessary clutter.

In the end we can have a more understanding in go modules by reading [Go modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)

## 3.Syntax notice in this part

- make
    
    “make” is the built in function in go for initializing objects of type **map,slice,..**
    
    Example
    
    ```jsx
     **map_messages := make(map[string]string)**
    ```
    
- loop
    
    ```go
    //suppose we have a messages map
    //**messages = ["Dev1":"Hello dev1","dev2":"Hello dev2"]
    //we can print the value by doing so**
    for _, value := range messages {
    		fmt.Print(value)
    	}
    ```
    

## References:

[Create a go module: ](https://go.dev/doc/tutorial/create-module)