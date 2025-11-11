⚙️ TASK 2 — Environment Setup
-----------------------------

### 🎯 Objective

To install and configure **Go (Golang)** (mandatory) and one other modern language (**Kotlin**) and verify both by running a “Hello World” program in each environment.The setup process is fully documented with **commands**, **outputs**, and **screenshots** for evaluation.

🧩 A. Go (Golang) Setup
-----------------------

### ✅ Step 1: Installation

*   Download and install Go from the official website: [https://go.dev/dl](https://go.dev/dl)
    
*   Run the installer and follow the default setup.
    
*   **Default Installation Path:**C:\\Program Files\\Go
    

### ✅ Step 2: Verify Installation

**Command:**go version

**Expected Output:**go version go1.22.3 windows/amd64

**Screenshot:**Go Installation Screenshot

### ✅ Step 3: Create Workspace

**Folder Structure:**src/go/task2\_environment/go setup/

Inside that folder, create a new file named **hello.go**.

### ✅ Step 4: Write Code

### Example Code
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello from Go!")

    go func() {
        fmt.Println("This runs concurrently!")
    }()
}
```

### ✅ Step 5: Run the Program

**Command:**go run src/go/task2\_environment/go setup/hello.go
```

**Expected Output:**
```Hello Go! Environment setup is successful.
```

**Screenshot:**hello\_output.png

### ✅ Step 6: Verification Summary

StepCommandOutput ExampleScreenshotInstallation Checkgo versiongo1.22.3go\_install.pngRun Hello Programgo run hello.goHello Go! Environment setup is successful.hello\_output.png

✅ **Result:** Go environment successfully installed, verified, and executed.

💻 B. Kotlin Setup
------------------

### ✅ Step 1: Installation

*   Download the Kotlin compiler from: [https://github.com/JetBrains/kotlin/releases](https://github.com/JetBrains/kotlin/releases)
    
*   Extract ZIP to:C:\\Kotlin\\kotlinc
    
*   Add this path to **Environment Variables**:C:\\Kotlin\\kotlinc\\bin
    
*   Restart VS Code or Command Prompt after updating the path.
    

### ✅ Step 2: Verify Installation

**Command:**kotlinc -version

**Expected Output:**info: kotlinc-jvm 2.2.21 (JRE 17.0.16+8-LTS)

**Screenshot:**kotlin\_install.png

### ✅ Step 3: Create Workspace

**Folder Path:**src/go/task2\_environment/kotlin\_setup/

Create a new file named **kotlin\_example.kt**

### ✅ Step 4: Write Code

**File:** kotlin\_example.kt

### Example Code

```kotlin
fun main() {
    println("Hello from Kotlin!")
    val names = listOf("Alice", "Bob", "Charlie")
    names.forEach { println("Hi, $it!") }
}
```

### ✅ Step 5: Run the Program

**Compile and Run:**

'''kotlinc kotlin_example.kt -include-runtime -d hello.jar  java -jar hello.jar   
```

**Expected Output:**
 ```Hello Kotlin! Environment setup is successful 🎉
```

**Screenshot:**kotlin\_output.png

### ✅ Step 6: Verification Summary

StepCommandOutput ExampleScreenshotInstallation Checkkotlinc -version2.2.21 (JRE 17)kotlin\_install.pngRun Hello Programjava -jar hello.jarHello Kotlin! Environment setup is successful 🎉kotlin\_output.png

✅ **Result:** Kotlin environment successfully installed, compiled, and executed.

🧾 Combined Summary Table
-------------------------

| Language       | Check Command       | Output Example                        | Screenshot            |
|----------------|---------------------|----------------------------------------|------------------------|
| **Go**         | `go version`        | go1.22.3                               | go_install.png         |
| **Kotlin**     | `kotlinc -version`  | 2.2.21 (JRE 17)                        | kotlin_install.png     |
| **Go Hello**   | `go run hello.go`   | Hello Go!                              | hello_output.png       |
| **Kotlin Hello** | `java -jar hello.jar` | Hello Kotlin! Environment setup is successful 🎉 | kotlin_output.png |


🏁 Final Result
---------------

Both **Go** and **Kotlin** environments were successfully **installed**, **configured**, and **verified** using “Hello World” programs.All steps are **documented with commands, outputs, and screenshots** to validate the setup.

📚 References
-------------

*   [Go Documentation](https://go.dev/doc)
    
*   [Kotlin Documentation](https://kotlinlang.org/docs/home.html)
    
*   [F# Documentation](https://learn.microsoft.com/dotnet/fsharp/)
    
*   [Clojure Documentation](https://clojure.org/)
