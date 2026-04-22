package turbospa

import "syscall/js"

func Patch(parent js.Value, oldNode *VNode, newNode *VNode) {
    // 1. Si l'ancien nœud n'existe pas, on crée le nouveau
    if oldNode == nil {
        newNode.domElement = createDOMElement(*newNode)
        parent.Call("appendChild", newNode.domElement)
        return
    }

    // 2. Si le type de balise a changé, on remplace tout
    if oldNode.TagName != newNode.TagName {
        newEl := createDOMElement(*newNode)
        parent.Call("replaceChild", newEl, oldNode.domElement)
        newNode.domElement = newEl
        return
    }

    // 3. Si c'est le même type, on réutilise l'élément DOM existant
    newNode.domElement = oldNode.domElement

    // 4. Mise à jour du texte si besoin
    if oldNode.Text != newNode.Text {
        newNode.domElement.Set("innerText", newNode.Text)
    }

    // 5. Mise à jour des attributs (simplifiée)
    patchAttributes(newNode.domElement, oldNode.Attrs, newNode.Attrs)

    // 6. Réconciliation des enfants
    patchChildren(newNode.domElement, oldNode.Children, newNode.Children)
}
func patchChildren(parent js.Value, oldChildren []VNode, newChildren []VNode) {
    // On récupère les longueurs des deux tableaux
    oldLen := len(oldChildren)
    newLen := len(newChildren)
    
    // On détermine la longueur maximale pour boucler sur tous les changements possibles
    maxLen := oldLen
    if newLen > maxLen {
        maxLen = newLen
    }

    for i := 0; i < maxLen; i++ {
        if i >= oldLen {
            // Cas 1 : Le nouvel enfant n'existait pas avant (Ajout)
            // On appelle Patch avec nil pour l'ancien nœud
            Patch(parent, nil, &newChildren[i])
            
        } else if i >= newLen {
            // Cas 2 : L'ancien enfant n'existe plus dans le nouveau rendu (Suppression)
            // On retire l'élément du DOM réel
            parent.Call("removeChild", oldChildren[i].domElement)
            
        } else {
            // Cas 3 : L'enfant existe dans les deux versions (Mise à jour)
            // On appelle Patch récursivement pour synchroniser cet enfant
            Patch(parent, &oldChildren[i], &newChildren[i])
        }
    }
}

func patchAttributes(el js.Value, oldAttrs, newAttrs map[string]string) {
    // Supprimer les anciens qui ne sont plus là
    for k := range oldAttrs {
        if _, ok := newAttrs[k]; !ok {
            el.Call("removeAttribute", k)
        }
    }
    // Ajouter ou mettre à jour les nouveaux
    for k, v := range newAttrs {
        if oldAttrs[k] != v {
            el.Set(k, v)
        }
    }
}