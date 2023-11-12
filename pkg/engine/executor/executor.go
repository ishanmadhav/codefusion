package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ishanmadhav/codefusion/internal/api"
)

// Checks which language it is, then passes te flow to the function for that language
func Execute(code *api.Code, uuid string) error {
	switch code.Language {
	case "python":
		return executePython(code, uuid)
	case "c":
		return executeC(code, uuid)
	case "cpp":
		return executeCpp(code, uuid)
	case "java":
		return executeJava(code)
	default:
		return nil
	}

	return nil
}

// Executes python code
func executePython(code *api.Code, uuid string) error {
	fileName := uuid + ".py"
	err := CreateAndWriteToFile(fileName, code.Code)
	if err != nil {
		return err
	}
	cmd := exec.Command("python3", fileName)
	cmd.Stdin = strings.NewReader(code.Input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	code.Output = string(out)
	_ = os.Remove(fileName)
	return nil
}

// Compiles and executes C code
func executeC(code *api.Code, uuid string) error {
	fileName := uuid + ".c"
	err := CreateAndWriteToFile(fileName, code.Code)
	if err != nil {
		return err
	}

	// Compile the C file
	outFile := uuid
	cmd := exec.Command("gcc", "-o", outFile, fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to compile: %w: %s", err, out)
	}

	// Execute the output file
	cmd = exec.Command("./" + outFile)
	cmd.Stdin = strings.NewReader(code.Input)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute: %w: %s", err, out)
	}
	code.Output = string(out)
	_ = os.Remove(fileName)
	_ = os.Remove(outFile)
	return nil
}

// Compiles and executes C++ code
func executeCpp(code *api.Code, uuid string) error {
	fileName := uuid + ".cpp"
	err := CreateAndWriteToFile(fileName, code.Code)
	if err != nil {
		return err
	}

	// Compile the C++ file
	outFile := uuid
	cmd := exec.Command("g++", "-o", outFile, fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to compile: %w: %s", err, out)
	}

	// Execute the output file
	cmd = exec.Command("./" + outFile)
	cmd.Stdin = strings.NewReader(code.Input)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute: %w: %s", err, out)
	}
	code.Output = string(out)
	_ = os.Remove(fileName)
	_ = os.Remove(outFile)
	return nil
}

// Compiles and executes Java code
func executeJava(code *api.Code) error {
	return nil
}

func CreateAndWriteToFile(fileName string, program string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(program)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}
