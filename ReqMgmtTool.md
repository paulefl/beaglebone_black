
Requirements Management Setup for C / Rust / Go with StrictDoc

Overview

This document describes a practical setup for integrating requirements management into a mixed-language system using:

- C
- Rust
- Go
- StrictDoc (requirements management)
- Git + CI/CD

The goal is to achieve full traceability from requirements → implementation → tests.

---

1. Repository Structure

project-root/
│
├── requirements/
│   ├── system.sdoc
│   ├── safety.sdoc
│   └── software.sdoc
│
├── src/
│   ├── c_module/
│   │   ├── main.c
│   │   └── driver.c
│   │
│   ├── rust_module/
│   │   ├── Cargo.toml
│   │   └── src/lib.rs
│   │
│   └── go_service/
│       ├── go.mod
│       └── main.go
│
├── tests/
│   ├── c_tests/
│   ├── rust_tests/
│   └── go_tests/
│
├── docs/
│
├── strictdoc.toml
└── .gitlab-ci.yml / .github/workflows/

Key Principles

- Requirements live inside the repository
- No separation between code and requirements
- Git is the single source of truth

---

2. Requirements Definition (StrictDoc)

Example ("software.sdoc"):

[REQ-SW-001]
TITLE: System shall initialize hardware safely
STATEMENT: The system shall initialize all hardware interfaces before use.
RATIONALE: Prevent undefined behavior.

---

3. Traceability in Code

C Example

// REQ-SW-001
void hardware_init(void) {
    init_gpio();
    init_spi();
}

Rust Example

/// REQ-SW-001
pub fn hardware_init() {
    init_gpio();
    init_spi();
}

Go Example

// REQ-SW-001
func InitHardware() {
    initGPIO()
    initSPI()
}

Rule

- Always reference requirements via stable IDs
- Do not rename IDs after creation

---

4. Traceability in Tests

C Test

// TEST-REQ-SW-001
void test_hardware_init() {
    assert(init_successful());
}

Rust Test

/// TEST-REQ-SW-001
#[test]
fn test_hardware_init() {
    assert!(hardware_init_ok());
}

Go Test

// TEST-REQ-SW-001
func TestHardwareInit(t *testing.T) {
    if !InitHardwareOK() {
        t.Fail()
    }
}

---

5. Naming Conventions

Type| Prefix| Example
Requirement| REQ-| REQ-SW-001
Test| TEST-REQ-| TEST-REQ-SW-001
Design| DES-| DES-SW-001

---

6. Traceability Strategy

Each requirement must have:

- Implementation reference (code)
- Verification (test)

Mapping

REQ-SW-001
├── src/c_module/main.c
├── src/rust_module/lib.rs
├── src/go_service/main.go
└── tests/... (at least one test)

---

7. Automation

Option A: Simple (Recommended Start)

Search for references:

rg "REQ-SW-" src/ tests/

Option B: Scripted Validation

- Parse StrictDoc files
- Extract all REQ IDs
- Compare with:
  - code references
  - test references

Option C: Advanced

Use tools like:

- openfasttrace

to generate traceability matrices automatically

---

8. Web UI (StrictDoc)

Start server:

strictdoc server

Provides:

- Requirement browsing
- Navigation
- Traceability visualization (if maintained)

---

9. CI/CD Integration

Goals

- Detect missing traceability
- Ensure test coverage of requirements

Example Checks

- Every REQ appears in code or tests
- Every REQ has at least one TEST- reference
- No orphaned requirements

---

10. Architecture Guidelines

Core Principles

1. Requirements are the source of truth

- Defined in StrictDoc
- Versioned in Git

2. Code references requirements

- Never the other way around

3. Tests validate requirements

- Not just implementation

---

Recommended Workflow

1. Write requirement
2. Implement feature
3. Reference requirement in code
4. Write test referencing requirement
5. CI validates traceability

---

11. Optional Enhancements

- Coverage mapping (Requirement → Test Coverage)
- ReqIF export for external tools
- Integration with documentation (Sphinx)

---

12. Summary

This setup provides:

- Full traceability (REQ → Code → Test)
- Language-independent integration (C, Rust, Go)
- Lightweight but scalable RM approach
- Compatibility with modern DevOps workflows

It is especially suitable for:

- Embedded systems
- Safety-critical development
- Automotive (e.g., AUTOSAR, ISO 26262)

--- 
