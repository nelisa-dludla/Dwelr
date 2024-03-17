// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ListingCard() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex h-72 w-full mt-8 max-w-6xl rounded shadow-xl hover:cursor-pointer hover:shadow-lime-800\"><div class=\"w-2/5\"><figure class=\"bg-black flex justify-center items-center h-full w-full\"><img class=\"h-full\" src=\"/views/images/Placeholder.svg\"></figure></div><div class=\"flex flex-col justify-center w-3/5 p-8\"><p class=\"font-bold text-2xl\">R 10 500</p><p class=\"my-4\">1 Bedroom Apartment</p><p class=\"font-bold\">Bluff, Durban</p><p class=\"my-4\">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi porta, ex in euismod iaculis, enim metus pharetra odio, at pellentesque purus nisi sed.</p><div class=\"flex font-bold my-4\"><p>1 Bedroom</p><p>1 Bathroom</p></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
