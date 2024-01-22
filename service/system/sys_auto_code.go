package system

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
)

const (
	plugPath = "resource/plug_template"
)

var (
	injectionPaths []injectionMeta
	caser          = cases.Title(language.English)
)

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

type AutoCodeService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllTplFile
//@description: 获取 pathName 文件夹下所有 tpl 文件
//@param: pathName string, fileList []string
//@return: []string, error

func (autoCodeService *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = autoCodeService.GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

// injectionCode 封装代码注入
func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		if err := utils.AutoInjectionCode(meta.path, meta.funcName, code); err != nil {
			return err
		}
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}

type Visitor struct {
	ImportCode  string
	StructName  string
	PackageName string
	GroupName   string
}

func (vi *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.GenDecl:
		// 查找有没有import context包
		// Notice：没有考虑没有import任何包的情况
		if n.Tok == token.IMPORT && vi.ImportCode != "" {
			vi.addImport(n)
			// 不需要再遍历子树
			return nil
		}
		if n.Tok == token.TYPE && vi.StructName != "" && vi.PackageName != "" && vi.GroupName != "" {
			vi.addStruct(n)
			return nil
		}
	case *ast.FuncDecl:
		if n.Name.Name == "Routers" {
			vi.addFuncBodyVar(n)
			return nil
		}

	}
	return vi
}

func (vi *Visitor) addStruct(genDecl *ast.GenDecl) ast.Visitor {
	for i := range genDecl.Specs {
		switch n := genDecl.Specs[i].(type) {
		case *ast.TypeSpec:
			if strings.Index(n.Name.Name, "Group") > -1 {
				switch t := n.Type.(type) {
				case *ast.StructType:
					f := &ast.Field{
						Names: []*ast.Ident{
							{
								Name: vi.StructName,
								Obj: &ast.Object{
									Kind: ast.Var,
									Name: vi.StructName,
								},
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: vi.PackageName,
							},
							Sel: &ast.Ident{
								Name: vi.GroupName,
							},
						},
					}
					t.Fields.List = append(t.Fields.List, f)
				}
			}
		}
	}
	return vi
}

func (vi *Visitor) addImport(genDecl *ast.GenDecl) ast.Visitor {
	// 是否已经import
	hasImported := false
	for _, v := range genDecl.Specs {
		importSpec := v.(*ast.ImportSpec)
		// 如果已经包含
		if importSpec.Path.Value == strconv.Quote(vi.ImportCode) {
			hasImported = true
		}
	}
	if !hasImported {
		genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(vi.ImportCode),
			},
		})
	}
	return vi
}

func (vi *Visitor) addFuncBodyVar(funDecl *ast.FuncDecl) ast.Visitor {
	hasVar := false
	for _, v := range funDecl.Body.List {
		switch varSpec := v.(type) {
		case *ast.AssignStmt:
			for i := range varSpec.Lhs {
				switch nn := varSpec.Lhs[i].(type) {
				case *ast.Ident:
					if nn.Name == vi.PackageName+"Router" {
						hasVar = true
					}
				}
			}
		}
	}
	if !hasVar {
		assignStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: vi.PackageName + "Router",
					Obj: &ast.Object{
						Kind: ast.Var,
						Name: vi.PackageName + "Router",
					},
				},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "xhs_router",
						},
						Sel: &ast.Ident{
							Name: "RouterGroupApp",
						},
					},
					Sel: &ast.Ident{
						Name: caser.String(vi.PackageName),
					},
				},
			},
		}
		funDecl.Body.List = append(funDecl.Body.List, funDecl.Body.List[1])
		index := 1
		copy(funDecl.Body.List[index+1:], funDecl.Body.List[index:])
		funDecl.Body.List[index] = assignStmt
	}
	return vi
}

// CreatePlug 自动创建插件模板
func (autoCodeService *AutoCodeService) CreatePlug(plug system.AutoPlugReq) error {
	// 检查列表参数是否有效
	plug.CheckList()
	tplFileList, _ := autoCodeService.GetAllTplFile(plugPath, nil)
	for _, tpl := range tplFileList {
		temp, err := template.ParseFiles(tpl)
		if err != nil {
			zap.L().Error("parse err", zap.String("tpl", tpl), zap.Error(err))
			return err
		}
		pathArr := strings.SplitAfter(tpl, "/")
		if strings.Index(pathArr[2], "tpl") < 0 {
			dirPath := filepath.Join(global.BODO_CONFIG.AutoCode.Root, global.BODO_CONFIG.AutoCode.Server, fmt.Sprintf(global.BODO_CONFIG.AutoCode.SPlug, plug.Snake+"/"+pathArr[2]))
			os.MkdirAll(dirPath, 0755)
		}
		file := filepath.Join(global.BODO_CONFIG.AutoCode.Root, global.BODO_CONFIG.AutoCode.Server, fmt.Sprintf(global.BODO_CONFIG.AutoCode.SPlug, plug.Snake+"/"+tpl[len(plugPath):len(tpl)-4]))
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			zap.L().Error("open file", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
		defer f.Close()

		err = temp.Execute(f, plug)
		if err != nil {
			zap.L().Error("exec err", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
	}
	return nil
}
