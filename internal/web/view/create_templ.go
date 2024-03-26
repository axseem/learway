// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreatePage() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center w-full px-4 py-8\"><div class=\"w-full md:w-[736px] flex flex-col gap-8\"><div x-data=\"{ title: &#39;&#39;, cards: [] }\" class=\"flex flex-col w-full gap-4\"><div class=\"flex flex-col gap-1\"><label for=\"title\">Title of the deck</label> <input @change=\"(e)=&gt;title=e.target.value\" :value=\"title\" type=\"text\" class=\"px-4 py-2 border rounded-lg\"></div><ul class=\"flex flex-col gap-4\"><template x-for=\"(card, index) in cards\"><li class=\"flex flex-col overflow-hidden border divide-y rounded-2xl\"><input @change=\"(e)=&gt;cards[index][0]=e.target.value\" :value=\"card[0]\" :name=\"&#39;cards[&#39;+index+&#39;][0]&#39;\" class=\"w-full p-4\"> <input @change=\"(e)=&gt;cards[index][1]=e.target.value\" :value=\"card[1]\" :name=\"&#39;cards[&#39;+index+&#39;][1]&#39;\" class=\"w-full p-4 bg-neutral-50\"></li></template></ul><button @click=\"cards.push([])\" class=\"p-4 rounded-full bg-neutral-100\">+ Add card</button> <button @click=\"fetch(&#39;/create&#39;, {method: &#39;POST&#39;, body: JSON.stringify($data)}); window.location.replace(&#39;/&#39;)\" class=\"p-4 rounded-full bg-neutral-100\">Create</button></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
