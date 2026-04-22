//go:build js && wasm

package turbospa

import (
	// "strings"
	"fmt"
	"syscall/js"
)

// func Mount(targetID string, comp Component) {
// 	document := js.Global().Get("document")
// 	root := document.Call("getElementById", targetID)

// 	// Premier rendu
// 	vnode := comp.Render()
// 	domEl := createDOMElement(vnode)
// 	root.Call("appendChild", domEl)

// 	// Bloquer le programme pour maintenir la SPA active
// 	select {}
// }
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

	for eventName, handler := range vnode.Events {
        // On crée une fonction JS à partir de la fonction Go
        jsFunc := js.FuncOf(handler)
        // On l'attache à l'élément (ex: "click")
        el.Call("addEventListener", eventName, jsFunc)
        fmt.Println("event", eventName,vnode.TagName)
        // Note pour plus tard : Dans une lib pro, il faudra stocker jsFunc 
        // pour appeler .Release() lors de la suppression du nœud.
    }

    if vnode.Text != "" {
        el.Set("textContent", vnode.Text) // textContent est plus sûr que innerText
    }

    for _, child := range vnode.Children {
        el.Call("appendChild", createDOMElement(child))
    }

    return el
}

// turbospa/engine.go



var Instance *Core

func Mount(targetID string, comp Component) {
    document := js.Global().Get("document")
    root := document.Call("getElementById", targetID)

    // Premier rendu
    vnode := comp.Render()
    domEl := createDOMElement(vnode)
    root.Call("appendChild", domEl)

    // On initialise l'instance globale pour pouvoir appeler Update() partout
    Instance = &Core{
        rootElement:  root,
        currentVNode: &vnode,
        app:          comp,
    }

    select {}
}
