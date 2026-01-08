// Package apperrors defines custom error types for the sley application.
// These typed errors enable proper error handling with errors.Is and errors.As
// without coupling internal packages to the CLI framework.
//
// Error Handling Conventions:
//   - Always wrap errors from external packages with context
//   - Use sentinel errors for common, well-known conditions
//   - Use typed errors when callers need to extract structured information
//   - Include relevant context (file paths, values) in error messages
package apperrors

import (
	"errors"
	"fmt"
)

// Sentinel errors for common conditions.
// Use errors.Is() to check for these conditions.
var (
	// ErrNotFound indicates a resource was not found.
	ErrNotFound = errors.New("not found")

	// ErrInvalidInput indicates invalid user input.
	ErrInvalidInput = errors.New("invalid input")

	// ErrPermissionDenied indicates insufficient permissions.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrAlreadyExists indicates a resource already exists.
	ErrAlreadyExists = errors.New("already exists")

	// ErrTimeout indicates an operation timed out.
	ErrTimeout = errors.New("operation timed out")

	// ErrCanceled indicates an operation was canceled.
	ErrCanceled = errors.New("operation canceled")

	// ErrGitOperation indicates a git operation failed.
	ErrGitOperation = errors.New("git operation failed")

	// ErrExtension indicates an extension-related error.
	ErrExtension = errors.New("extension error")
)

// VersionFileNotFoundError indicates that the version file does not exist.
type VersionFileNotFoundError struct {
	Path string
}

func (e *VersionFileNotFoundError) Error() string {
	return fmt.Sprintf("version file not found at %s", e.Path)
}

// InvalidVersionError indicates that a version string does not conform to semver.
type InvalidVersionError struct {
	Version string
	Reason  string
}

func (e *InvalidVersionError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("invalid version format %q: %s", e.Version, e.Reason)
	}
	return fmt.Sprintf("invalid version format: %s", e.Version)
}

// InvalidBumpTypeError indicates an invalid bump type was specified.
type InvalidBumpTypeError struct {
	BumpType string
}

func (e *InvalidBumpTypeError) Error() string {
	return fmt.Sprintf("invalid bump type: %s (expected: patch, minor, or major)", e.BumpType)
}

// ConfigError indicates a configuration-related error.
type ConfigError struct {
	Operation string
	Err       error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("config %s failed: %v", e.Operation, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// CommandError indicates a command execution error.
type CommandError struct {
	Command string
	Err     error
	Timeout bool
}

func (e *CommandError) Error() string {
	if e.Timeout {
		return fmt.Sprintf("command %q timed out: %v", e.Command, e.Err)
	}
	return fmt.Sprintf("command %q failed: %v", e.Command, e.Err)
}

func (e *CommandError) Unwrap() error {
	return e.Err
}

// PathValidationError indicates an invalid or dangerous path was provided.
type PathValidationError struct {
	Path   string
	Reason string
}

func (e *PathValidationError) Error() string {
	return fmt.Sprintf("invalid path %q: %s", e.Path, e.Reason)
}

// HookError indicates an error during hook execution.
type HookError struct {
	HookName string
	Err      error
}

func (e *HookError) Error() string {
	return fmt.Sprintf("hook %q failed: %v", e.HookName, e.Err)
}

func (e *HookError) Unwrap() error {
	return e.Err
}

// GitError indicates a git operation error with command context.
type GitError struct {
	Operation string
	Err       error
}

func (e *GitError) Error() string {
	return fmt.Sprintf("git %s failed: %v", e.Operation, e.Err)
}

func (e *GitError) Unwrap() error {
	return e.Err
}

// Is allows errors.Is to match against ErrGitOperation.
func (e *GitError) Is(target error) bool {
	return target == ErrGitOperation
}

// FileError indicates a file operation error with path context.
type FileError struct {
	Op   string
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s %q: %v", e.Op, e.Path, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

// ExtensionError indicates an extension-related error.
type ExtensionError struct {
	Name string
	Op   string
	Err  error
}

func (e *ExtensionError) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("extension %q %s: %v", e.Name, e.Op, e.Err)
	}
	return fmt.Sprintf("extension %s: %v", e.Op, e.Err)
}

func (e *ExtensionError) Unwrap() error {
	return e.Err
}

// Is allows errors.Is to match against ErrExtension.
func (e *ExtensionError) Is(target error) bool {
	return target == ErrExtension
}

// WrapGit wraps an error as a git operation error.
func WrapGit(operation string, err error) error {
	if err == nil {
		return nil
	}
	return &GitError{Operation: operation, Err: err}
}

// WrapFile wraps an error as a file operation error.
func WrapFile(op, path string, err error) error {
	if err == nil {
		return nil
	}
	return &FileError{Op: op, Path: path, Err: err}
}

// WrapExtension wraps an error as an extension error.
func WrapExtension(name, op string, err error) error {
	if err == nil {
		return nil
	}
	return &ExtensionError{Name: name, Op: op, Err: err}
}
