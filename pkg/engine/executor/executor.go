package executor

import "github.com/ishanmadhav/codefusion/internal/api"

// Checks which language it is, then passes te flow to the function for that language
func Execute(code *api.Code) error {
	switch code.Language {
	case "python":
		return executePython(code)
	case "c":
		return executeC(code)
	case "cpp":
		return executeCpp(code)
	case "java":
		return executeJava(code)
	default:
		return nil
	}

	return nil
}

// Executes python code
func executePython(code *api.Code) error {
	return nil
}

// Compiles and executes C code
func executeC(code *api.Code) error {
	return nil
}

// Compiles and executes C++ code
func executeCpp(code *api.Code) error {
	return nil
}

// Compiles and executes Java code
func executeJava(code *api.Code) error {
	return nil
}
