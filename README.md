# GitHub Action Automation: Go Calculator

## Purpose
This project demonstrates how to automate build and test workflows for a simple Go-based calculator using GitHub Actions. It serves as a template for CI/CD integration in Go projects and provides a basic example of modular Go code and unit testing.

## What This Project Does
- Implements a basic calculator in Go with functions for addition, subtraction, multiplication, and division.
- Provides unit tests for each calculator function to ensure correctness.
- Includes a GitHub Actions workflow (`.github/workflows/go-pipeline.yaml`) that automatically builds and tests the project on every push and pull request to the `main` branch.

## How to Use
1. **Clone the repository:**
   ```sh
   git clone https://github.com/itsmesuraj20/github-action-automation.git
   cd github-action-automation
   ```
2. **Run locally:**
   - Build the project:
     ```sh
     go build -v ./...
     ```
   - Run tests:
     ```sh
     go test -v ./...
     ```
3. **GitHub Actions:**
   - On every push or pull request to `main`, the workflow will automatically build and test your code.
   - Check the Actions tab in your GitHub repository for build and test results.

## What Was Implemented
- Calculator functions: `Add`, `Subtract`, `Multiply`, `Divide` (with error handling for division by zero).
- Unit tests for all calculator functions in `main_test.go`.
- Go module setup (`go.mod`) with correct module path and Go version.
- Automated CI pipeline using GitHub Actions for build and test.

## What Can Be Added in the Future
- **More Calculator Functions:** Support for advanced operations (power, square root, modulus, etc.).
- **API Layer:** Expose calculator functions via a REST API using Go's `net/http` or a framework like Gin.
- **CLI Tool:** Add a command-line interface for interactive use.
- **Docker Support:** Add a Dockerfile for containerized builds and deployments.
- **Code Coverage Reporting:** Integrate coverage tools and report results in CI.
- **Release Automation:** Add workflows for automated releases and tagging.
- **Multi-OS Testing:** Extend CI to test on Windows and macOS runners.
- **Static Analysis & Linting:** Integrate Go linters and static analysis tools in the workflow.

## Contributing
Feel free to fork the repository, open issues, or submit pull requests for improvements and new features.

---

**Author:** Suraj
**License:** MIT
