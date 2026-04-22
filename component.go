//go:build js && wasm
package turbospa

type Component interface {
    Render() VNode
}