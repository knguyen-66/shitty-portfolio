// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package base

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Sidebar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-[15rem] flex-none flex h-screen flex-col justify-between border-r-2 border-primary\"><div class=\"px-4 py-6\"><span class=\"grid h-10 w-32 place-content-center rounded-lg bg-surface-1 hover:bg-surface-2 text-primary text-xs\">Logo</span><ul class=\"mt-6 space-y-1\"><li><a href=\"#\" class=\"block rounded-lg px-4 py-2 text-sm font-medium bg-surface-1 hover:bg-surface-2 text-primary\">Home</a></li><li><a href=\"#\" class=\"block rounded-lg px-4 py-2 text-sm font-medium bg-surface-1 hover:bg-surface-2 text-primary\">Blogs</a></li><li><a href=\"#\" class=\"block rounded-lg px-4 py-2 text-sm font-medium bg-surface-1 hover:bg-surface-2 text-primary\">Work Archives</a></li></ul></div><div class=\"sticky inset-x-0 bottom-0 border-t-2 border-secondary\"><a href=\"#\" class=\"flex items-center gap-2 bg-surface hover:bg-surface-2 p-4\"><img alt=\"\" src=\"https://cdn-icons-png.flaticon.com/512/9187/9187604.png\" class=\"size-10 rounded-full object-cover\"><div><p class=\"text-xs\"><strong class=\"block font-medium text-primary\">Khanh Nguyen</strong> <span class=\"text-secondary\">fake@email.com </span></p></div></a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
