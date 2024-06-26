// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Footer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer class=\"mt-auto h-72 w-screen px-10 pt-10 pb-5\"><div class=\"flex flex-row justify-between\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Logo().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex\"><ul class=\"mr-10\"><li class=\"font-bold\">Company</li><li class>Our Story</li></ul><ul class=\"mr-10\"><li class=\"font-bold\">Explore</li><li>Listings</li></ul><ul class=\"mr-10\"><li class=\"font-bold\">Support</li><li>Feedback</li><li>Resources</li></ul></div></div><div class=\"bg-black h-0.5 w-auto mt-28 mb-3\"></div><div class=\"flex justify-between items-end\"><div class=\"flex w-1/2 justify-between\"><ul><li>Copyright &copy 2024 Dwelr</li></ul><ul class=\"flex\"><li><a href=\"\" class=\"cursor-not-allowed\">Terms</a></li><li class=\"cursor-not-allowed ml-8\"><a href=\"\" class=\"cursor-not-allowed\">Privacy</a></li></ul></div><div class=\"flex w-1/2 justify-end\"><ul class=\"flex\"><li class=\"mr-10 text-2xl\"><a href=\"https://twitter.com/Nelisa_Dludla\"><i class=\"fa-brands fa-x-twitter\" style=\"color: #000000;\"></i></a></li><li class=\"mr-10 text-2xl\"><a href=\"www.linkedin.com/in/nelisa-dludla\"><i class=\"fa-brands fa-linkedin\" style=\"color: #000000;\"></i></a></li><li class=\"mr-10 text-2xl\"><a href=\"https://github.com/nelisa-dludla/Dwelr\"><i class=\"fa-brands fa-github\" style=\"color: #000000;\"></i></a></li></ul></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
