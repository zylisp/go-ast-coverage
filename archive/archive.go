package archive

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// SimpleASTBundle stores AST with source for perfect reconstruction
type SimpleASTBundle struct {
	// Store the formatted source code (guaranteed to round-trip perfectly)
	SourceCode string `gob:"source"`

	// Store the original filename and parse mode
	Filename  string      `gob:"filename"`
	ParseMode parser.Mode `gob:"parse_mode"`

	// Optional: store the cleaned AST for quick access to structure
	CleanedAST *ast.File `gob:"cleaned_ast,omitempty"`

	// Store any additional metadata
	Metadata map[string]interface{} `gob:"metadata,omitempty"`
}

// ASTArchive provides a convenient API for working with archived AST data.
// It wraps SimpleASTBundle with helper methods for common operations.
type ASTArchive struct {
	bundle *SimpleASTBundle
}

// GetSourceCode returns the source code stored in the archive.
func (a *ASTArchive) GetSourceCode() string {
	return a.bundle.SourceCode
}

// GetFilename returns the original filename.
func (a *ASTArchive) GetFilename() string {
	return a.bundle.Filename
}

// GetPackageName returns the package name from metadata.
func (a *ASTArchive) GetPackageName() string {
	if pkg, ok := a.bundle.Metadata["original_package"].(string); ok {
		return pkg
	}
	return ""
}

// GetAST reconstructs the complete AST with Scope/Object references by re-parsing.
// This gives you the full AST including all semantic information.
func (a *ASTArchive) GetAST() (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, a.bundle.Filename, a.bundle.SourceCode, a.bundle.ParseMode)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to reconstruct AST: %w", err)
	}
	return file, fset, nil
}

// GetCleanedAST returns the pre-cleaned AST without Scope/Object references.
// This is faster than GetAST() as it doesn't require re-parsing, but lacks semantic info.
func (a *ASTArchive) GetCleanedAST() *ast.File {
	return a.bundle.CleanedAST
}

// GetMetadata retrieves a metadata value by key.
func (a *ASTArchive) GetMetadata(key string) interface{} {
	return a.bundle.Metadata[key]
}

// NodeCount returns the total number of AST nodes by traversing the cleaned AST.
func (a *ASTArchive) NodeCount() int {
	if a.bundle.CleanedAST == nil {
		return 0
	}
	count := 0
	ast.Inspect(a.bundle.CleanedAST, func(n ast.Node) bool {
		if n != nil {
			count++
		}
		return true
	})
	return count
}

// DeclarationCount returns the number of top-level declarations.
func (a *ASTArchive) DeclarationCount() int {
	if count, ok := a.bundle.Metadata["num_declarations"].(int); ok {
		return count
	}
	return 0
}

// ImportCount returns the number of imports.
func (a *ASTArchive) ImportCount() int {
	if count, ok := a.bundle.Metadata["num_imports"].(int); ok {
		return count
	}
	return 0
}

// RegisterAllASTTypes registers all AST types with gob for serialization
func RegisterAllASTTypes() {
	// Core interfaces - these must be registered first
	gob.Register((*ast.Node)(nil))
	gob.Register((*ast.Expr)(nil))
	gob.Register((*ast.Stmt)(nil))
	gob.Register((*ast.Decl)(nil))
	gob.Register((*ast.Spec)(nil))

	// Expression types
	gob.Register(&ast.BadExpr{})
	gob.Register(&ast.Ident{})
	gob.Register(&ast.Ellipsis{})
	gob.Register(&ast.BasicLit{})
	gob.Register(&ast.FuncLit{})
	gob.Register(&ast.CompositeLit{})
	gob.Register(&ast.ParenExpr{})
	gob.Register(&ast.SelectorExpr{})
	gob.Register(&ast.IndexExpr{})
	gob.Register(&ast.IndexListExpr{}) // Go 1.18+
	gob.Register(&ast.SliceExpr{})
	gob.Register(&ast.TypeAssertExpr{})
	gob.Register(&ast.CallExpr{})
	gob.Register(&ast.StarExpr{})
	gob.Register(&ast.UnaryExpr{})
	gob.Register(&ast.BinaryExpr{})
	gob.Register(&ast.KeyValueExpr{})

	// Statement types
	gob.Register(&ast.BadStmt{})
	gob.Register(&ast.DeclStmt{})
	gob.Register(&ast.EmptyStmt{})
	gob.Register(&ast.LabeledStmt{})
	gob.Register(&ast.ExprStmt{})
	gob.Register(&ast.SendStmt{})
	gob.Register(&ast.IncDecStmt{})
	gob.Register(&ast.AssignStmt{})
	gob.Register(&ast.GoStmt{})
	gob.Register(&ast.DeferStmt{})
	gob.Register(&ast.ReturnStmt{})
	gob.Register(&ast.BranchStmt{})
	gob.Register(&ast.BlockStmt{})
	gob.Register(&ast.IfStmt{})
	gob.Register(&ast.CaseClause{})
	gob.Register(&ast.SwitchStmt{})
	gob.Register(&ast.TypeSwitchStmt{})
	gob.Register(&ast.CommClause{})
	gob.Register(&ast.SelectStmt{})
	gob.Register(&ast.ForStmt{})
	gob.Register(&ast.RangeStmt{})

	// Declaration types
	gob.Register(&ast.BadDecl{})
	gob.Register(&ast.GenDecl{})
	gob.Register(&ast.FuncDecl{})

	// Spec types
	gob.Register(&ast.ImportSpec{})
	gob.Register(&ast.ValueSpec{})
	gob.Register(&ast.TypeSpec{})

	// Type expression types
	gob.Register(&ast.ArrayType{})
	gob.Register(&ast.StructType{})
	gob.Register(&ast.FuncType{})
	gob.Register(&ast.InterfaceType{})
	gob.Register(&ast.MapType{})
	gob.Register(&ast.ChanType{})

	// Other important types
	gob.Register(&ast.Field{})
	gob.Register(&ast.FieldList{})
	gob.Register(&ast.File{})
	gob.Register(&ast.Package{})
	gob.Register(&ast.Comment{})
	gob.Register(&ast.CommentGroup{})

	// Slices of interfaces (these are crucial!)
	gob.Register([]ast.Expr{})
	gob.Register([]ast.Stmt{})
	gob.Register([]ast.Decl{})
	gob.Register([]ast.Spec{})
	gob.Register([]*ast.Ident{})
	gob.Register([]*ast.Field{})
	gob.Register([]*ast.Comment{})
	gob.Register([]*ast.CommentGroup{})
	gob.Register([]*ast.ImportSpec{})

	// Token types
	gob.Register(token.Token(0))
	gob.Register(token.Pos(0))
}

// SaveASTWithSourcePreservation saves AST by preserving source code
func SaveASTWithSourcePreservation(file *ast.File, fset *token.FileSet, filename, outputFile string) error {
	// Register all AST types for gob encoding
	RegisterAllASTTypes()

	// Convert AST back to source code
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		return fmt.Errorf("failed to format AST to source: %w", err)
	}

	sourceCode := buf.String()

	// Create a cleaned copy for structural analysis (optional)
	cleanedFile := deepCopyAndClean(file)

	bundle := SimpleASTBundle{
		SourceCode: sourceCode,
		Filename:   filename,
		ParseMode:  parser.ParseComments, // Preserve comments by default
		CleanedAST: cleanedFile,
		Metadata:   make(map[string]interface{}),
	}

	// Add useful metadata
	bundle.Metadata["original_package"] = file.Name.Name
	bundle.Metadata["num_declarations"] = len(file.Decls)
	bundle.Metadata["num_imports"] = len(file.Imports)

	// Serialize with gob
	var gobBuf bytes.Buffer
	encoder := gob.NewEncoder(&gobBuf)
	if err := encoder.Encode(&bundle); err != nil {
		return fmt.Errorf("failed to encode bundle: %w", err)
	}

	return os.WriteFile(outputFile, gobBuf.Bytes(), 0644)
}

// LoadASTWithSourceReconstruction loads AST and reconstructs all references
func LoadASTWithSourceReconstruction(filename string) (*ast.File, *token.FileSet, string, error) {
	// Register all AST types for gob decoding
	RegisterAllASTTypes()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to read file: %w", err)
	}

	var bundle SimpleASTBundle
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&bundle); err != nil {
		return nil, nil, "", fmt.Errorf("failed to decode bundle: %w", err)
	}

	// Re-parse the source code to get perfect AST with all references
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, bundle.Filename, bundle.SourceCode, bundle.ParseMode)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to re-parse source: %w", err)
	}

	return file, fset, bundle.SourceCode, nil
}

// Load loads a single AST archive and wraps it in the convenience API.
func Load(filename string) (*ASTArchive, error) {
	// Register all AST types for gob decoding
	RegisterAllASTTypes()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var bundle SimpleASTBundle
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&bundle); err != nil {
		return nil, fmt.Errorf("failed to decode bundle: %w", err)
	}

	return &ASTArchive{bundle: &bundle}, nil
}

// LoadAll loads all .asta files from a directory.
// Returns a slice of ASTArchive objects for easy iteration.
func LoadAll(dir string) ([]*ASTArchive, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var archives []*ASTArchive
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".asta") {
			continue
		}

		archivePath := filepath.Join(dir, entry.Name())
		archive, err := Load(archivePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", entry.Name(), err)
		}
		archives = append(archives, archive)
	}

	return archives, nil
}

// WalkArchives iterates over all .asta files in a directory, calling fn for each.
// If fn returns an error, iteration stops and that error is returned.
// This is useful for processing archives without loading them all into memory at once.
func WalkArchives(dir string, fn func(*ASTArchive) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".asta") {
			continue
		}

		archivePath := filepath.Join(dir, entry.Name())
		archive, err := Load(archivePath)
		if err != nil {
			return fmt.Errorf("failed to load %s: %w", entry.Name(), err)
		}

		if err := fn(archive); err != nil {
			return err
		}
	}

	return nil
}

// deepCopyAndClean creates a copy without circular references (for optional storage)
func deepCopyAndClean(file *ast.File) *ast.File {
	// Use go/format and re-parse to get a clean copy
	fset := token.NewFileSet()
	var buf bytes.Buffer
	format.Node(&buf, fset, file)

	// Parse without object resolution to avoid circular references
	cleanFile, _ := parser.ParseFile(fset, "", buf.String(), parser.SkipObjectResolution)
	return cleanFile
}

// VerifyPerfectFidelity ensures the loaded AST is identical to original
func VerifyPerfectFidelity(original, restored *ast.File, originalFset, restoredFset *token.FileSet) error {
	// Format both to source and compare
	var origBuf, restBuf bytes.Buffer

	if err := format.Node(&origBuf, originalFset, original); err != nil {
		return fmt.Errorf("failed to format original: %w", err)
	}

	if err := format.Node(&restBuf, restoredFset, restored); err != nil {
		return fmt.Errorf("failed to format restored: %w", err)
	}

	if origBuf.String() != restBuf.String() {
		return fmt.Errorf("source code does not match")
	}

	// Verify scope/object preservation
	if (original.Scope == nil) != (restored.Scope == nil) {
		return fmt.Errorf("scope preservation mismatch")
	}

	if original.Scope != nil && restored.Scope != nil {
		if len(original.Scope.Objects) != len(restored.Scope.Objects) {
			return fmt.Errorf("scope objects count mismatch: %d vs %d",
				len(original.Scope.Objects), len(restored.Scope.Objects))
		}
	}

	// Count identifier objects
	origIdents, restIdents := 0, 0
	ast.Inspect(original, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Obj != nil {
			origIdents++
		}
		return true
	})
	ast.Inspect(restored, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Obj != nil {
			restIdents++
		}
		return true
	})

	if origIdents != restIdents {
		return fmt.Errorf("identifier object count mismatch: %d vs %d", origIdents, restIdents)
	}

	return nil
}

// ExtractFunctions returns all function declarations from an archive.
// This reconstructs the AST to get access to function declarations.
func ExtractFunctions(archive *ASTArchive) ([]*ast.FuncDecl, error) {
	file, _, err := archive.GetAST()
	if err != nil {
		return nil, err
	}

	var funcs []*ast.FuncDecl
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			funcs = append(funcs, fn)
		}
	}
	return funcs, nil
}

// ExtractTypes returns all type declarations from an archive.
func ExtractTypes(archive *ASTArchive) ([]*ast.TypeSpec, error) {
	file, _, err := archive.GetAST()
	if err != nil {
		return nil, err
	}

	var types []*ast.TypeSpec
	for _, decl := range file.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gen.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					types = append(types, typeSpec)
				}
			}
		}
	}
	return types, nil
}

// GetImports returns all import paths from an archive.
func GetImports(archive *ASTArchive) []string {
	// Use the cleaned AST for fast access (no re-parsing needed)
	cleanAST := archive.GetCleanedAST()
	if cleanAST == nil {
		return nil
	}

	var imports []string
	for _, imp := range cleanAST.Imports {
		// Remove quotes from import path
		path := strings.Trim(imp.Path.Value, "\"")
		imports = append(imports, path)
	}
	return imports
}

// FindNodesByType finds all AST nodes of a specific type in the archive.
// nodeType should be a string like "*ast.FuncDecl", "*ast.IfStmt", etc.
// Returns nodes as []ast.Node - caller should type assert to specific types.
func FindNodesByType(archive *ASTArchive, nodeType string) ([]ast.Node, error) {
	file, _, err := archive.GetAST()
	if err != nil {
		return nil, err
	}

	var nodes []ast.Node
	ast.Inspect(file, func(n ast.Node) bool {
		if n != nil && fmt.Sprintf("%T", n) == nodeType {
			nodes = append(nodes, n)
		}
		return true
	})
	return nodes, nil
}

// GetFunctionNames returns the names of all functions in the archive.
func GetFunctionNames(archive *ASTArchive) ([]string, error) {
	funcs, err := ExtractFunctions(archive)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(funcs))
	for i, fn := range funcs {
		names[i] = fn.Name.Name
	}
	return names, nil
}

// GetTypeNames returns the names of all type declarations in the archive.
func GetTypeNames(archive *ASTArchive) ([]string, error) {
	types, err := ExtractTypes(archive)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(types))
	for i, t := range types {
		names[i] = t.Name.Name
	}
	return names, nil
}
