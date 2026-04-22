//go:build js && wasm
package turbospa

import "syscall/js"

type VNode struct {
    TagName  string
    Attrs    map[string]string
    Children []VNode
    Text     string
    // Nouveau : pour stocker les événements
    Events   map[string]func(js.Value, []js.Value) interface{}
    domElement js.Value
}

// Fonction pour créer un nœud de manière fluide (Helper)
func El(tag string, attrs map[string]string, children ...VNode) VNode {
    return VNode{TagName: tag, Attrs: attrs, Children: children}
} 