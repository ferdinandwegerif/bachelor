package main

// Code generated by vugu via vugugen. Please regenerate instead of editing or add additional code in a separate file. DO NOT EDIT.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

import (
	"github.com/vugu/vugu/vgform"
	"strings"
)

type Root struct {
	firstname	string
	lastname	string
	username	string
	password	string
	email		string
	gender		string
	errors		[]string
	success		bool
}

func (c *Root) HandleSubmit(event vugu.DOMEvent) {
	event.PreventDefault()
	c.errors = []string{}
	c.success = false

	if len(strings.TrimSpace(c.firstname)) == 0 {
		c.errors = append(c.errors, "Requires First Name.")
	}
	if len(strings.TrimSpace(c.lastname)) == 0 {
		c.errors = append(c.errors, "Requires Last Name.")
	}
	if len(strings.TrimSpace(c.username)) == 0 {
		c.errors = append(c.errors, "Requires Username.")
	}
	if len(strings.TrimSpace(c.password)) == 0 {
		c.errors = append(c.errors, "Requires Password.")
	}
	if len(strings.TrimSpace(c.email)) == 0 {
		c.errors = append(c.errors, "Requires E-Mail.")
	}
	if len(strings.TrimSpace(c.gender)) == 0 {
		c.errors = append(c.errors, "Require gender.")
	}
	if len(c.errors) == 0 {
		c.success = true
	}

	/*Code where we send information to website etc */
	/* Continuation from previous forms */
}

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "html", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "head", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "title", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Vugu Form"}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "link", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rel", Val: "stylesheet"}, vugu.VGAttribute{Namespace: "", Key: "href", Val: "https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.AppendCSS(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "body", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "container"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "main"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "h4", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "display-4"}}}
				vgparent.AppendChild(vgn)
				vgn.SetInnerHTML(vugu.HTML("Registration with Vugu"))
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "form", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"submit",
					Func:		func(event vugu.DOMEvent) { c.HandleSubmit(event) },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " First name and last name "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "row"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                            "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: "firstname"}}}
								vgparent.AppendChild(vgn)
								vgn.SetInnerHTML(vugu.HTML("First Name"))
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                            "}
								vgparent.AppendChild(vgn)
								{
									vgcompKey := vugu.MakeCompKey(0x2EE7D65B9CDEAFAF^vgin.CurrentPositionHash(), vgiterkey)
									// ask BuildEnv for prior instance of this specific component
									vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Input)
									if vgcomp == nil {
										// create new one if needed
										vgcomp = new(vgform.Input)
										vgin.BuildEnv.WireComponent(vgcomp)
									}
									vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
									vgcomp.Value = vgform.StringPtr{&c.firstname}
									vgcomp.AttrMap = make(map[string]interface{}, 8)
									vgcomp.AttrMap["type"] = "text"
									vgcomp.AttrMap["placeholder"] = "First Name"
									vgcomp.AttrMap["class"] = "form-control"
									vgcomp.AttrMap["id"] = "firstname"
									vgout.Components = append(vgout.Components, vgcomp)
									vgn = &vugu.VGNode{Component: vgcomp}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "col"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                            "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: "lastname"}}}
								vgparent.AppendChild(vgn)
								vgn.SetInnerHTML(vugu.HTML("Last Name"))
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                            "}
								vgparent.AppendChild(vgn)
								{
									vgcompKey := vugu.MakeCompKey(0xB0C7A97EB30A2991^vgin.CurrentPositionHash(), vgiterkey)
									// ask BuildEnv for prior instance of this specific component
									vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Input)
									if vgcomp == nil {
										// create new one if needed
										vgcomp = new(vgform.Input)
										vgin.BuildEnv.WireComponent(vgcomp)
									}
									vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
									vgcomp.Value = vgform.StringPtr{&c.lastname}
									vgcomp.AttrMap = make(map[string]interface{}, 8)
									vgcomp.AttrMap["type"] = "text"
									vgcomp.AttrMap["placeholder"] = "Last Name"
									vgcomp.AttrMap["class"] = "form-control"
									vgcomp.AttrMap["id"] = "lastname"
									vgout.Components = append(vgout.Components, vgcomp)
									vgn = &vugu.VGNode{Component: vgcomp}
									vgparent.AppendChild(vgn)
								}
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
								vgparent.AppendChild(vgn)
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Username "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: ""}}}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("Username"))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0x4F868533D305F468^vgin.CurrentPositionHash(), vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Input)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(vgform.Input)
								vgin.BuildEnv.WireComponent(vgcomp)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Value = vgform.StringPtr{&c.username}
							vgcomp.AttrMap = make(map[string]interface{}, 8)
							vgcomp.AttrMap["type"] = "text"
							vgcomp.AttrMap["placeholder"] = "Username"
							vgcomp.AttrMap["class"] = "form-control"
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Password "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: ""}}}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("Password"))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0x8B0C2FCAD3F8D9FE^vgin.CurrentPositionHash(), vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Input)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(vgform.Input)
								vgin.BuildEnv.WireComponent(vgcomp)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Value = vgform.StringPtr{&c.password}
							vgcomp.AttrMap = make(map[string]interface{}, 8)
							vgcomp.AttrMap["type"] = "password"
							vgcomp.AttrMap["placeholder"] = "Password"
							vgcomp.AttrMap["class"] = "form-control"
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Email "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: ""}}}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("E-mail"))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0xEE4A99A4E5AB0BF8^vgin.CurrentPositionHash(), vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Input)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(vgform.Input)
								vgin.BuildEnv.WireComponent(vgcomp)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Value = vgform.StringPtr{&c.email}
							vgcomp.AttrMap = make(map[string]interface{}, 8)
							vgcomp.AttrMap["type"] = "text"
							vgcomp.AttrMap["placeholder"] = "E-mail"
							vgcomp.AttrMap["class"] = "form-control"
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(4), Data: " Gender "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: ""}}}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("Gender"))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                    "}
						vgparent.AppendChild(vgn)
						{
							vgcompKey := vugu.MakeCompKey(0x5B2857874B96E67C^vgin.CurrentPositionHash(), vgiterkey)
							// ask BuildEnv for prior instance of this specific component
							vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*vgform.Select)
							if vgcomp == nil {
								// create new one if needed
								vgcomp = new(vgform.Select)
								vgin.BuildEnv.WireComponent(vgcomp)
							}
							vgin.BuildEnv.UseComponent(vgcompKey, vgcomp)	// ensure we can use this in the cache next time around
							vgcomp.Options = vgform.SliceOptions{"Male", "Female"}.Title()
							vgcomp.Value = vgform.StringPtrDefault(&c.gender, "Male")
							vgcomp.AttrMap = make(map[string]interface{}, 8)
							vgcomp.AttrMap["name"] = "gender"
							vgcomp.AttrMap["id"] = "gender"
							vgcomp.AttrMap["class"] = "form-control"
							vgout.Components = append(vgout.Components, vgcomp)
							vgn = &vugu.VGNode{Component: vgcomp}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "field form-group"}}}
					vgparent.AppendChild(vgn)
					vgn.SetInnerHTML(vugu.HTML("\n                    \x3Cbutton\x3ESubmit\x3C/button\x3E\n                "))
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n            "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n        "}
				vgparent.AppendChild(vgn)
				if c.success {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "values"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "h3", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("Current values"))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                First name: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.firstname)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                Last Name: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.lastname)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                Username: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.username)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                Password: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.password)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                E-mail: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.email)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                Gender: "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "span", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.gender)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n        "}
				vgparent.AppendChild(vgn)
				if len(c.errors) > 0 {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "errors"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "p", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(vugu.HTML("Please correct errors below."))
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						for vgiterkeyt, error := range c.errors {
							var vgiterkey interface{} = vgiterkeyt
							_ = vgiterkey
							error := error
							_ = error
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "ul", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
								vgparent.AppendChild(vgn)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "li", Attr: []vugu.VGAttribute(nil)}
								vgparent.AppendChild(vgn)
								vgn.SetInnerHTML(error)
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
								vgparent.AppendChild(vgn)
							}
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
		}
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
