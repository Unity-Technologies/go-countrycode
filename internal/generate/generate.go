package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	const sourceURL = "https://raw.githubusercontent.com/lukes/ISO-3166-Countries-with-Regional-Codes/master/all/all.csv"
	inputData, err := httpGet(sourceURL)
	if err != nil {
		panic(err)
	}

	inputReader := bytes.NewReader(inputData)
	records, err := csv.NewReader(inputReader).ReadAll()
	if err != nil {
		panic(err)
	}

	countryRecords := make(countryRecordList, 0, len(records)-1)
	for _, record := range records[1:] {
		countryRecords = append(countryRecords, countryRecord{
			Name:            record[0],
			Alpha3:          record[2],
			Alpha2:          record[1],
			Alpha2LowerCase: strings.ToLower(record[1]),
		})
	}
	sort.Stable(countryRecords)

	if err = writeASTFile("code.gen.go", countryRecords.GenerateAST()); err != nil {
		panic(err)
	}
}

func writeASTFile(filename string, astFile *ast.File) (returnedErr error) {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := fd.Close(); err != nil && returnedErr == nil {
			returnedErr = err
		}
	}()
	if err := format.Node(fd, token.NewFileSet(), astFile); err != nil {
		return err
	}
	return nil
}

func httpGet(urlStr string) ([]byte, error) {
	res, err := http.Get(urlStr) // nolint: gosec
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if closeErr := res.Body.Close(); err != nil {
		return nil, closeErr
	}
	return body, err
}

type countryRecord struct {
	Name            string
	Alpha2          string
	Alpha2LowerCase string
	Alpha3          string
}

type countryRecordList []countryRecord

func (l countryRecordList) Len() int {
	return len(l)
}

func (l countryRecordList) Less(i, j int) bool {
	return l[i].Alpha3 < l[j].Alpha3
}

func (l countryRecordList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l countryRecordList) GenerateAST() *ast.File {
	return &ast.File{
		Name: &ast.Ident{Name: pkgName},
		Decls: []ast.Decl{
			l.generateCountries(),
			l.generateCodes(),
		},
	}
}

func (l countryRecordList) generateCountries() ast.Decl {
	specs := make([]ast.Spec, 0, len(l))
	for i, cr := range l {
		specs = append(specs, &ast.ValueSpec{
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{Text: fmt.Sprintf("\n// %s is %s.", cr.Alpha3, cr.Name)},
				},
			},
			Names: []*ast.Ident{{Name: cr.Alpha3}},
			Values: []ast.Expr{
				&ast.CompositeLit{
					Type: &ast.Ident{Name: countryType},
					Elts: []ast.Expr{
						&ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", i+1)},
					},
				},
			},
		})
	}
	return &ast.GenDecl{
		Tok: token.VAR,
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "\n// Code generated. DO NOT EDIT."},
			},
		},
		Specs: specs,
	}
}

func (l countryRecordList) generateCodes() ast.Decl {
	elts := make([]ast.Expr, 0, len(l)+1)

	elts = append(elts, &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"---"`},
			&ast.BasicLit{Kind: token.STRING, Value: `"--"`},
			&ast.BasicLit{Kind: token.STRING, Value: `"--"`},
		},
	})

	for _, cr := range l {
		elts = append(elts, &ast.CompositeLit{
			Elts: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", cr.Alpha3)},
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", cr.Alpha2)},
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", cr.Alpha2LowerCase)},
			},
		})
	}

	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{{Name: codes}},
				Values: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{
							Len: &ast.Ellipsis{},
							Elt: &ast.ArrayType{
								Len: &ast.Ident{Name: formatsCount},
								Elt: &ast.Ident{Name: stringType},
							},
						},
						Elts: elts,
					},
				},
			},
		},
	}
}

const pkgName = "countrycode"
const countryType = "Country"
const codes = "codes"
const formatsCount = "formatsCount"
const stringType = "string"
