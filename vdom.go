//go:build js && wasm
package turbospa

import "syscall/js"

type VNode struct {
    TagName  string
    Attrs    map[string]string
    Children []VNode
    Text     string
    // Référence vers l'élément réel dans le navigateur
    domElement js.Value 
}

// Fonction pour créer un nœud de manière fluide (Helper)
func El(tag string, attrs map[string]string, children ...VNode) VNode {
    return VNode{TagName: tag, Attrs: attrs, Children: children}
} 