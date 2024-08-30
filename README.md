`go vet` 是用于检测 Go 语言的源代码中的可能的 bug、错误的用法和其他常见的代码问题。下面是 `go vet` 命令的一些常见用法：

1. **在当前目录下运行 vet**：如果在当前目录下运行 `go vet`（无参数）， `vet` 命令会检查当前目录中的 Go 语言文件。

    ```
    go vet
    ```

2. **在特定的.go文件中运行 vet**：可以指定`go vet` 检查的 Go 文件。

    ```
    go vet main.go
    ```

3. **在特定的包中运行 vet**：可以让 `go vet` 检查特定的 Go 语言包。例如，以下命令会检查 `mypackage` 中的文件：

    ```
    go vet mypackage
    ```

4. **在所有包中运行 vet**：可以用 `./...` 让 `go vet` 检查当前目录及其所有子目录中的所有 Go 语言文件：

    ```
    go vet ./...
    ```

5. **打开/关闭特定的检查器**：可以通过 `-vet` 标记来启用或禁用特定的检查器。例如，要禁用 `'composite'` 检查器（这个检查器检查使用结构字面量复合的方式是否正确），可以这样使用：

    ```
    go vet -vet=^composite ./...
    ```

6. **检查重新声明的变量**：Go 语言中，隐藏或者覆盖一个存在的变量通常是一种错误，可以用 `-shadow` 标记来检查这种错误：

    ```
    go tool vet -shadow .
    ```

    注意，`-shadow` 标记比较消耗性能，使用时需要小心。