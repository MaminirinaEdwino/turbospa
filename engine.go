//go:build js && wasm
package turbospa

import "syscall/js"

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
    doc := js.Global().Get("document")
    el := doc.Call("createElement", vnode.TagName)
    
    // Appliquer les attributs (id, class, etc.)
    for k, v := range vnode.Attrs {
        el.Set(k, v)
    }
    
    // Gérer le texte ou les enfants
    if vnode.Text != "" {
        el.Set("innerText", vnode.Text)
    }
    
    for _, child := range vnode.Children {
        el.Call("appendChild", createDOMElement(child))
    }
    
    return el
}