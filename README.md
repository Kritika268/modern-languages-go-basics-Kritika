🧠 Modern Programming Languages – Go (Golang) Basics 
====================================================

This repository contains my work for learning and implementing **Go (Golang)** through structured tasks, from setup to advanced topics like structs and pointers.Each task includes **source code**, **output screenshots**, and a **step-by-step explanation**.

## 📂 Repository Structure
---
```bash
MODERN-LANGUAGES-GO-BASICS/
│
├── screenshots/                  # All task output screenshots
│
├── src/
│   └── go/
│       ├── task1_docs/           # Concept documentation
│       │   └── concept_notebook.md
│       │
│       ├── task2_environment/    # Go installation & setup
│       │   ├── go setup/
│       │   ├── kotlin_setup/
│       │   └── environment_setup.md
│       │
│       ├── task3_datatypes/      # Data types and variables
│       │   ├── main.go
│       │   └── task3_output.png
│       │
│       ├── task4_control/        # Control structures
│       │   ├── main.go
│       │   └── task4_output.png
│       │
│       ├── task5_packages/       # Packages & modularization
│       │   ├── greetings/
│       │   │   └── greetings.go
│       │   ├── main.go
│       │   └── task5_output.png
│       │
│       ├── task6_collections/    # Arrays, slices, maps
│       │   ├── main.go
│       │   └── task6_output.png
│       │
│       ├── task7_structs/        # Structs and custom types
│       │   ├── main.go
│       │   └── task7_output.png
│       │
│       ├── task8_pointers/       # Pointers in Go
│       │   ├── main.go
│       │   └── task8_output.png
│       │
│       └── go.mod                # Go module file
│
├── .gitignore
└── README.md
```


🧩 Task Breakdown
-----------------

### **Task 1 – Documentation**

*   Created a concept\_notebook.md explaining key Go language fundamentals and syntax.
    
*   Compared Go with modern languages (F#, Clojure, Kotlin).
    

### **Task 2 – Environment Setup**

*   Installed Go using the official installer.
    
*   Configured environment variables (GOPATH, GOROOT).
    
*   go versiongo env
    
*   Created first Go program (Hello World).
    

### **Task 3 – Data Types**

*   Implemented variables and constants of various data types (int, float, string, bool).
    
*   Output screenshot: task3\_output.png
    

### **Task 4 – Control Structures**

*   Demonstrated if-else, for loops, and switch-case.
    
*   for i := 1; i <= 5; i++ { fmt.Println("Count:", i)}
    
*   Output screenshot: task4\_output.png
    

### **Task 5 – Packages**

*   Created a custom package greetings and imported it into main.go.
    
*   Showed how modular programming improves reusability.
    
*   Output screenshot: task5\_output.png
    

### **Task 6 – Collections**

*   Implemented examples using arrays, slices, and maps.
    
*   Demonstrated slice append and iteration.
    
*   Output screenshot: task6\_output.png
    

### **Task 7 – Structs**

*   Defined and accessed fields of structs.
    
*   Demonstrated composition and nested structs.
    
*   Output screenshot: task7\_output.png
    

### **Task 8 – Pointers**

*   Demonstrated how pointers work in Go with variable references.
    
*   var x int = 10var p \*int = &xfmt.Println(\*p)
    
*   Output screenshot: task8\_output.png
    

🚀 Task 9 – Final Submission
----------------------------

### ✅ What’s Included

*   All Go source codes (.go files).
    
*   Output screenshots under each task folder.
    
*   A detailed documentation notebook (concept\_notebook.md).
    
*   Proper folder structure with Go modules (go.mod).
    

🗂️ Steps to Run Locally
------------------------

1.  git clone https://github.com/Kritika268/modern-languages-go-basics-Kritika.git
    
2.  cd modern-languages-go-basics-Kritika/src/go
    
3.  go mod tidy
    
4.  cd task3\_datatypesgo run main.go
    
5.  **View Outputs**
    
    *   Open the corresponding taskX\_output.png under each task folder or in screenshots/

🖼️ Screenshots
---------------

All task outputs are available in the screenshots/ directory, for example:

TaskScreenshotTask 3 – Data TypesTask 4 – Control StructuresTask 5 – PackagesTask 6 – CollectionsTask 7 – StructsTask 8 – Pointers

🧾 Author
---------

**Name:** _\[Kritika 2301350028\]_ 
**Course:** _Modern Programming Languages (Go Basics)_ 
**Submission:** _Task 9 – Final Repository Submission_
