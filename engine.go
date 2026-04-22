//go:build js && wasm

package turbospa

import (
	// "strings"
	"fmt"
	"syscall/js"
)

func Mount(targetID string, comp Component) {
	document := js.Global().Get("document")
	root := document.Call("getElementById", targetID)

	// Premier rendu
	vnode := comp.Render()
	domEl := createDOMElement(vnode)
	root.Call("appendChild", domEl)

	// Bloquer le programme pour maintenir la SPA active
	select {}
}
func createDOMElement(vnode VNode) js.Value {
	fmt.Printf("Création de l'élément : [%s]\n", vnode.TagName)
    doc := js.Global().Get("document")
    // Assure-toi que vnode.TagName n'est pas vide

	if vnode.TagName == "" {
        return doc.Call("createTextNode", vnode.Text)
    }
    el := doc.Call("createElement", vnode.TagName)

    // Remplace el.Set(k, v) par ceci :
    for k, v := range vnode.Attrs {
        el.Call("setAttribute", k, v)
    }

    if vnode.Text != "" {
        el.Set("textContent", vnode.Text) // textContent est plus sûr que innerText
    }

    for _, child := range vnode.Children {
        el.Call("appendChild", createDOMElement(child))
    }

    return el
}

// func createDOMElement(vnode VNode) js.Value {
// 	doc := js.Global().Get("document")
// 	tag := strings.TrimSpace(vnode.TagName)
// 	el := doc.Call("createElement", tag)
// 	// el := doc.Call("createElement", vnode.TagName)

// 	// Appliquer les attributs (id, class, etc.)
// 	for k, v := range vnode.Attrs {
// 		el.Set(k, v)
// 	}

// 	// Gérer le texte ou les enfants
// 	if vnode.Text != "" {
// 		el.Set("innerText", vnode.Text)
// 	}

// 	for _, child := range vnode.Children {
// 		el.Call("appendChild", createDOMElement(child))
// 	}
// 	for eventName, handler := range vnode.Events {
// 		// On transforme la fonction Go en fonction compatible JS
// 		jsCallback := js.FuncOf(handler)
// 		// On l'attache (ex: "click", "input")
// 		el.Call("addEventListener", eventName, jsCallback)
// 	}
// 	return el
// }
