# ⚙️ TASK 2 — Environment Setup  

## 🎯 Objective  
To install and configure **Go (Golang)** (mandatory) and one other modern language (**Kotlin**) and verify both by running a “Hello World” program in each environment. The setup process is documented with commands, outputs, and screenshots for evaluation.

---

## 🧩 A. Go (Golang) Setup  

### ✅ Step 1: Installation  
Download and install **Go** from the official website:  
👉 [https://go.dev/dl](https://go.dev/dl)  

Run the installer and follow default setup.  
Default installation path:  
`C:\Program Files\Go`

---

### ✅ Step 2: Verify Installation  
Run the following command:  
`go version`

**Expected Output:**  
`go version go1.22.3 windows/amd64`

📸 **Screenshot:**  
![Go Installation Screenshot](../go%20setup/go_install.png)

---

### ✅ Step 3: Create Workspace  
Create the folder structure as follows:  
`src/go/task2_environment/go setup/`  

Inside that folder, create a new file named:  
`hello.go`

---

### ✅ Step 4: Write Code  
Paste this code in `hello.go`:

```go

package main
import "fmt"

func main() {
    fmt.Println("Hello Go! Environment setup is successful.")
}

### ✅ Step 5: Run the Program  
Run the file using:  
`go run src/go/task2_environment/go setup/hello.go`

**Expected Output:**  
`Hello Go! Environment setup is successful.`

📸 **Screenshot:**  
![Go Hello Output](../go%20setup/hello_output.png)

---

### ✅ Step 6: Verification Summary  

| Step | Command | Output Example | Screenshot |
|------|----------|----------------|-------------|
| Installation Check | `go version` | go1.22.3 | ![Go Install](../go%20setup/go_install.png) |
| Run Hello Program | `go run hello.go` | Hello Go! Environment setup is successful. | ![Go Output](../go%20setup/hello_output.png) |

✅ **Result:** Go environment successfully installed, verified, and executed.

---

## 💻 B. Kotlin Setup  

### ✅ Step 1: Installation  
Download the **Kotlin compiler** from:  
👉 [https://github.com/JetBrains/kotlin/releases](https://github.com/JetBrains/kotlin/releases)

Extract the ZIP file to:  
`C:\Kotlin\kotlinc`

Add this path to your system environment variables:  
`C:\Kotlin\kotlinc\bin`

Restart VS Code or Command Prompt after adding it.

---

### ✅ Step 2: Verify Installation  
Run:  
`kotlinc -version`

**Expected Output:**  
`info: kotlinc-jvm 2.2.21 (JRE 17.0.16+8-LTS)`

📸 **Screenshot:**  
![Kotlin Installation Screenshot](kotlin_install.png)

---

### ✅ Step 3: Create Workspace  
Folder path:  
`src/go/task2_environment/kotlin_setup/`

Create a new file:  
`kotlin_example.kt`

---

### ✅ Step 4: Write Code  
Paste this code in `kotlin_example.kt`:

```kotlin
// File: kotlin_example.kt
// Author: Kritika
// Description: Simple Kotlin Hello World program for ENSP415

fun main() {
    println("Hello Kotlin! Environment setup is successful 🎉")
}
### ✅ Step 5: Run the Program  
Run the file using:  
`go run src/go/task2_environment/go setup/hello.go`

**Expected Output:**  
`Hello Go! Environment setup is successful.`

📸 **Screenshot:**  
![Go Hello Output](../go%20setup/hello_output.png)

---

### ✅ Step 6: Verification Summary  

| Step | Command | Output Example | Screenshot |
|------|----------|----------------|-------------|
| Installation Check | `go version` | go1.22.3 | ![Go Install](../go%20setup/go_install.png) |
| Run Hello Program | `go run hello.go` | Hello Go! Environment setup is successful. | ![Go Output](../go%20setup/hello_output.png) |

✅ **Result:** Go environment successfully installed, verified, and executed.

---

## 💻 B. Kotlin Setup  

### ✅ Step 1: Installation  
Download the **Kotlin compiler** from:  
👉 [https://github.com/JetBrains/kotlin/releases](https://github.com/JetBrains/kotlin/releases)

Extract the ZIP file to:  
`C:\Kotlin\kotlinc`

Add this path to your system environment variables:  
`C:\Kotlin\kotlinc\bin`

Restart VS Code or Command Prompt after adding it.

---

### ✅ Step 2: Verify Installation  
Run:  
`kotlinc -version`

**Expected Output:**  
`info: kotlinc-jvm 2.2.21 (JRE 17.0.16+8-LTS)`

📸 **Screenshot:**  
![Kotlin Installation Screenshot](kotlin_install.png)

---

### ✅ Step 3: Create Workspace  
Folder path:  
`src/go/task2_environment/kotlin_setup/`

Create a new file:  
`kotlin_example.kt`

---

### ✅ Step 4: Write Code  
Paste this code in `kotlin_example.kt`:

```kotlin
// File: kotlin_example.kt
// Author: Kritika
// Description: Simple Kotlin Hello World program for ENSP415

fun main() {
    println("Hello Kotlin! Environment setup is successful 🎉")
}


**Expected Output:**  
`Hello Kotlin! Environment setup is successful 🎉`

📸 **Screenshot:**  
![Kotlin Output Screenshot](kotlin_output.png)

---

### ✅ Step 6: Verification Summary  

| Step | Command | Output Example | Screenshot |
|------|----------|----------------|-------------|
| Installation Check | `kotlinc -version` | 2.2.21 (JRE 17) | ![Kotlin Install](kotlin_install.png) |
| Run Hello Program | `java -jar hello.jar` | Hello Kotlin! Environment setup is successful 🎉 | ![Kotlin Output](kotlin_output.png) |

✅ **Result:** Kotlin environment successfully installed, compiled, and executed.

---

## 🧾 Combined Summary Table  

| Language | Check Command | Output Example | Screenshot |
|-----------|----------------|----------------|-------------|
| **Go** | `go version` | go1.22.3 | ![Go Install](../go%20setup/go_install.png) |
| **Kotlin** | `kotlinc -version` | 2.2.21 (JRE 17) | ![Kotlin Install](kotlin_install.png) |
| **Go Hello** | `go run hello.go` | Hello Go! | ![Go Hello Output](../go%20setup/hello_output.png) |
| **Kotlin Hello** | `java -jar hello.jar` | Hello Kotlin! | ![Kotlin Hello Output](kotlin_output.png) |

---

## 🏁 Result  
Both **Go** and **Kotlin** environments were successfully installed, configured, and verified using “Hello World” programs.  
All steps are documented with commands, expected outputs, and screenshots to validate the installation process.

---

## 📚 References  
- **Go Docs:** [https://go.dev/doc](https://go.dev/doc)  
- **Kotlin Docs:** [https://kotlinlang.org/docs/home.html](https://kotlinlang.org/docs/home.html)  
- **F# Docs:** [https://learn.microsoft.com/dotnet/fsharp/](https://learn.microsoft.com/dotnet/fsharp/)  
- **Clojure Docs:** [https://clojure.org/](https://clojure.org/)
