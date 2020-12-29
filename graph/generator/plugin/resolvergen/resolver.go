package resolvergen

import (
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)
var ResolverPath string
func New(resolverPath string) plugin.Plugin {
	ResolverPath = resolverPath
	return &Plugin{}
}

type Plugin struct{}

var _ plugin.CodeGenerator = &Plugin{}

func (m *Plugin) Name() string {
	return "resolvergen"
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Resolver.IsDefined() {
		return nil
	}

	switch data.Config.Resolver.Layout {
	case config.LayoutSingleFile:
		return m.generateSingleFile(data)
	case config.LayoutFollowSchema:
		return m.generatePerSchema(data)
	}

	return nil
}

func (m *Plugin) generateSingleFile(data *codegen.Data) error {
	file := File{}

	if _, err := os.Stat(data.Config.Resolver.Filename); err == nil {
		// file already exists and we dont support updating resolvers with layout = single so just return
		return nil
	}

	for _, o := range data.Objects {
		if o.HasResolvers() {
			file.Objects = append(file.Objects, o)
		}
		for _, f := range o.Fields {
			if !f.IsResolver {
				continue
			}

			resolver := Resolver{o, f, `panic("not implemented")`}
			file.Resolvers = append(file.Resolvers, &resolver)
		}
	}

	resolverBuild := &ResolverBuild{
		File:         &file,
		PackageName:  data.Config.Resolver.Package,
		ResolverType: data.Config.Resolver.Type,
		HasRoot:      true,
	}

	return templates.Render(templates.Options{
		PackageName: data.Config.Resolver.Package,
		FileNotice:  `// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.`,
		Filename:    data.Config.Resolver.Filename,
		Data:        resolverBuild,
		Packages:    data.Config.Packages,
	})
}

func (m *Plugin) generatePerSchema(data *codegen.Data) error {
	//rewriter, err := rewrite.New(data.Config.Resolver.Dir())
	//if err != nil {
	//	return err
	//}

	files := map[string]*File{}

	for _, o := range data.Objects {
		if o.HasResolvers() {
			fn := gqlToResolverName(data.Config.Resolver.Dir(), o.Position.Src.Name, data.Config.Resolver.FilenameTemplate)

			if files[fn] == nil {
				files[fn] = &File{}
			}

		//	rewriter.Mafmt.Println(fn)rkStructCopied(templates.LcFirst(o.Name) + templates.UcFirst(data.Config.Resolver.Type))
		//	rewriter.GetMethodBody(data.Config.Resolver.Type, o.Name)
			files[fn].Objects = append(files[fn].Objects, o)
		}
		for _, f := range o.Fields {
			if !f.IsResolver {
				continue
			}


			implementation := `return resolvers.`+f.GoFieldName+`(ctx, sorts, filters)`

			if(f.Name == "paginate"){
				implementation = "return resolvers.Paginate(ctx, sorts, filters, subSorts, subFilters, page, size)"
			}
			if(len(f.Args) == 4 && f.Name != "paginate"){
				implementation = `return resolvers.`+f.GoFieldName+`(ctx, sorts, filters, subSorts, subFilters)`
			}

			if(strings.Contains(f.Description, "mutation-create")){
				sources := strings.Split(f.Description, ":")
				subscriptionAdd := ""
				if(len(sources) >= 2){
					if(strings.Contains(sources[1], "subscription")){
						subscription := strings.Split(sources[1], "-")

						subscriptionAdd = fmt.Sprintf(`r.mutex.Lock()
	for _, ch := range r.%sCreatedChannel {
		ch <- row
	}
	r.mutex.Unlock()`, subscription[1])
					}
					}

				temp :=  fmt.Sprintf(`row, err := resolvers.%s(ctx, input)
    %s
	return row, err`, f.GoFieldName, subscriptionAdd)
				implementation = temp
			}

			if(strings.Contains(f.Description, "mutation-update")){
				sources := strings.Split(f.Description, ":")
				subscriptionAdd := ""
				if(len(sources) >= 2){
					if(strings.Contains(sources[1], "subscription")){
						subscription := strings.Split(sources[1], "-")

						subscriptionAdd = fmt.Sprintf(`r.mutex.Lock()
	for _, ch := range r.%sUpdatedChannel {
		ch <- row
	}
	r.mutex.Unlock()`, subscription[1])
					}
				}

				temp :=  fmt.Sprintf(`row, err := resolvers.%s(ctx, id, input)
    %s
	return row, err`, f.GoFieldName, subscriptionAdd)
				implementation = temp
			}


			if(strings.Contains(f.Description, "mutation-delete")){
				sources := strings.Split(f.Description, ":")
				subscriptionAdd := ""
				if(len(sources) >= 2){
					if(strings.Contains(sources[1], "subscription")){
						subscription := strings.Split(sources[1], "-")

						subscriptionAdd = fmt.Sprintf(`r.mutex.Lock()
	for _, ch := range r.%sDeletedChannel {
		ch <- row
	}
	r.mutex.Unlock()`, subscription[1])
					}
				}

				temp :=  fmt.Sprintf(`row, err := resolvers.%s(ctx, id)
    %s
	return row, err`, f.GoFieldName, subscriptionAdd)
				implementation = temp
			}


			if(strings.Contains(f.Description, "subscription-created")){

				sources := strings.Split(f.Description, ":")

					implementation = fmt.Sprintf(`id := RandString(8)
	event := make(chan *models.%s, 1)
	r.mutex.Lock()
	r.%sCreatedChannel[id] = event
	r.mutex.Unlock()
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.%sCreatedChannel, id)
		r.mutex.Unlock()
	}()
	return event, nil`,
						sources[1],
						sources[1],
						sources[1],
					)

			}
			if(strings.Contains(f.Description, "subscription-updated")){

				sources := strings.Split(f.Description, ":")

				implementation = fmt.Sprintf(`id := RandString(8)
	event := make(chan *models.%s, 1)
	r.mutex.Lock()
	r.%sUpdatedChannel[id] = event
	r.mutex.Unlock()
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.%sUpdatedChannel, id)
		r.mutex.Unlock()
	}()
	return event, nil`,
					sources[1],
					sources[1],
					sources[1],
				)

			}
			if(strings.Contains(f.Description, "subscription-deleted")){

				sources := strings.Split(f.Description, ":")

				implementation = fmt.Sprintf(`id := RandString(8)
	event := make(chan *models.%s, 1)
	r.mutex.Lock()
	r.%sDeletedChannel[id] = event
	r.mutex.Unlock()
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.%sDeletedChannel, id)
		r.mutex.Unlock()
	}()
	return event, nil`,
					sources[1],
					sources[1],
					sources[1],
				)

			}




			resolver := Resolver{o, f, implementation}
			fn := gqlToResolverName(data.Config.Resolver.Dir(), f.Position.Src.Name, data.Config.Resolver.FilenameTemplate)
			if files[fn] == nil {
				files[fn] = &File{}
			}

			files[fn].Resolvers = append(files[fn].Resolvers, &resolver)
		}
	}

	//for filename, file := range files {
	//	//file.imports = rewriter.ExistingImports(filename)
	//	//file.RemainingSource = rewriter.RemainingSource(filename)
	//}

	for filename, file := range files {
		resolverBuild := &ResolverBuild{
			File:         file,
			PackageName:  data.Config.Resolver.Package,
			ResolverType: data.Config.Resolver.Type,
		}

		err := templates.Render(templates.Options{
			PackageName: data.Config.Resolver.Package,
			FileNotice: `
				// This file will be automatically regenerated based on the schema, any resolver implementations
				// will be copied through when generating and any unknown code will be moved to the end.`,
			Filename: filename,
			Data:     resolverBuild,
			Packages: data.Config.Packages,
		})
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(data.Config.Resolver.Filename); os.IsNotExist(errors.Cause(err)) {
		err := templates.Render(templates.Options{
			PackageName: data.Config.Resolver.Package,
			FileNotice: `
				// This file will not be regenerated automatically.
				//
				// It serves as dependency injection for your app, add any dependencies you require here.`,
			Template: `type {{.}} struct {}`,
			Filename: data.Config.Resolver.Filename,
			Data:     data.Config.Resolver.Type,
			Packages: data.Config.Packages,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type ResolverBuild struct {
	*File
	HasRoot      bool
	PackageName  string
	ResolverType string
}

type File struct {
	// These are separated because the type definition of the resolver object may live in a different file from the
	//resolver method implementations, for example when extending a type in a different graphql schema file
	Objects         []*codegen.Object
	Resolvers       []*Resolver

	RemainingSource string
}

func (f *File) Imports() string {


	return ResolverPath

}

type Resolver struct {
	Object         *codegen.Object
	Field          *codegen.Field
	Implementation string
}

func gqlToResolverName(base string, gqlname, filenameTmpl string) string {
	gqlname = filepath.Base(gqlname)
	ext := filepath.Ext(gqlname)
	if filenameTmpl == "" {
		filenameTmpl = "{name}.resolvers.go"
	}
	filename := strings.ReplaceAll(filenameTmpl, "{name}", strings.TrimSuffix(gqlname, ext))
	return filepath.Join(base, filename)
}
