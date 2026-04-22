package turbospa

import "syscall/js"

type Core struct {
    rootElement js.Value
    currentVNode *VNode
    app         Component
}

func (c *Core) Update() {
    newVNode := c.app.Render()
    Patch(c.rootElement, c.currentVNode, &newVNode)
    c.currentVNode = &newVNode
}