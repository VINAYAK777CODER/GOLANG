📦 Go Packages — Quick Guide
✅ 1. Create a Go Module (Project)

Inside any folder, run:

go mod init myproject


This creates a go.mod file.
Now Go knows your project name → myproject.

✅ 2. Create a Folder for Your Custom Package

Example project structure:

myproject/
    go.mod
    main.go
    utils/
        math.go


Here, utils is your custom package.

✅ 3. Write Code Inside Your Package
utils/math.go
package utils

// Add is exported because it starts with capital A
func Add(a, b int) int {
    return a + b
}

// multiply is unexported (private)
func multiply(a, b int) int {
    return a * b
}

✅ 4. Use the Package in main.go
main.go
package main

import (
    "fmt"
    "myproject/utils"
)

func main() {
    result := utils.Add(10, 20)
    fmt.Println("Result:", result)
}


This works because:

Module name = myproject

Folder name = utils

Import path → "myproject/utils"

🔥 Important Rules About Packages
Rule 1: Folder name = package name

All .go files inside utils/ must begin with:

package utils

Rule 2: Export with Capital Letters

Capital letter → Public / Exported

Small letter → Private / Unexported

Rule 3: One package per folder

All .go files in the same folder must have the same package name.

⭐ 5. Run the Program

Run:

go run .


This will automatically detect main.go and execute it.

⭐ 6. Adding More Packages

You can add more folders:

myproject/
    utils/
    services/
    db/
    models/


Each folder = new package.

⭐ 7. You Can Also Create Sub-packages

Example:

utils/math/
utils/strings/


Import like:

import "myproject/utils/math"

⭐ 8. Creating Packages for GitHub Projects

If your module is a GitHub repo:

go mod init github.com/yourname/project


Imports look like:

import "github.com/yourname/project/utils"

⭐ 9. Shortcut: Custom Package Example with Output
utils/hello.go
package utils

func Hello() string {
    return "Hello from custom package!"
}

main.go
package main

import (
    "fmt"
    "myproject/utils"
)

func main() {
    fmt.Println(utils.Hello())
}

Output
Hello from custom package!

⭐ Final Summary
Concept	Meaning
package	Code group inside a folder
main package	Starting point of execution
library package	Reusable code package you create
import	To use functions from other packages
Capital letter	Exported (public)
small letter	Unexported (private)


// ------------------------------------------------------//
📦 Installing & Using Third-Party Packages in Go
✅ 1. Initialize a Go Module

Before installing any external package, create a Go module:

go mod init myproject


This generates a go.mod file which manages dependencies for your project.

✅ 2. Install a Third-Party Package

Use go get to download and add a package to your project.

Example (installing Google’s UUID package):

go get github.com/google/uuid


This will:

Download the package into Go’s module cache

Add the module under require in go.mod

Update go.sum with checksums

✅ 3. Import the Package in Code

Use the import path exactly as shown on GitHub:

package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
}

✅ 4. Run Your Program
go run .


Go will automatically load the installed third-party package.

⭐ Installing Specific Versions

Install a specific version:

go get github.com/gofiber/fiber/v2@v2.52.0


Upgrade to the latest version:

go get -u github.com/gofiber/fiber/v2

⭐ Multiple Packages

You can install multiple libraries:

go get github.com/gorilla/mux
go get github.com/spf13/viper

⭐ Where Packages Are Stored

Downloaded modules are stored in:

$GOPATH/pkg/mod


(Go manages this automatically.)

🎯 Summary
Step	Command / Action
Initialize module	go mod init myproject
Install package	go get github.com/user/repo
Import package	import "github.com/user/repo"
Run app	go run .

---------------------------------------------
go mod tidy is a Go command that cleans and fixes your module dependencies.

Think of it as an automatic cleaner + organizer for your go.mod and go.sum files.

✅ What go mod tidy Does
1️⃣ Removes unused dependencies

If your go.mod has packages you no longer import, go mod tidy deletes them.

2️⃣ Adds missing dependencies

If your code imports a package but it's not listed in go.mod, it adds it automatically.

3️⃣ Updates go.sum

It ensures checksums are correct and removes unused ones.

🔥 In short:
go mod tidy ensures your project has exactly the dependencies it needs — nothing more, nothing less.
📌 When to use go mod tidy?

Use it when:

✔ After adding new imports
✔ After deleting imports
✔ After moving files or refactoring
✔ Before pushing code to GitHub
✔ To fix broken go.mod / go.sum

📘 Example
go mod tidy


Output (example):

unused github.com/some/package removed
added github.com/new/dependency

🧠 Why it’s important?

Keeps your project clean

Removes bloat

Fixes dependency issues

Ensures reproducible builds

Prevents unnecessary packages in GitHub repos

🎯 Summary
Command	Purpose
go mod tidy	Add missing + remove unused dependencies
go mod download	Downloads all deps
go mod vendor	Copies deps to vendor/ folder