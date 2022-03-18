package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node //根节点
}

type node struct {
	isLast   bool                //节点是否可以成为最终的路由规则。该节点能否成为一个独立的uri，是否自身就是一个终极节点
	segment  string              //uri中的字符串，代表这个节点表示的路由中某个段的字符串
	handlers []ControllerHandler // 代表这个节点中包含的控制器，用于最终加载调用
	childs   []*node             // 代表这个节点下的子节点
	parent   *node               //父节点 双向指针
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func newTree() *Tree {
	root := newNode()
	return &Tree{root}
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	//这里是传进来的segment是通配符
	// 如果segment是通配符，则所有下一层子节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, childNode := range n.childs {
		//这里是节点的segment是通配符
		if isWildSegment(childNode.segment) {
			// 如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, childNode)
		} else if segment == childNode.segment {
			// 如果下一层子节点没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, childNode)
		}
	}

	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符将uri切割为两个部分
	segments := strings.SplitN(uri, "/", 2)

	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	childNodes := n.filterChildNodes(segment)

	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if childNodes == nil || len(childNodes) == 0 {
		return nil
	}

	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点，判断这些cnode是否有isLast标志
		for _, currNode := range childNodes {
			if currNode.isLast {
				return currNode
			}
		}

		//都不是最后一个节点
		return nil
	}
	// 如果有大于1个segment, 递归每个子节点继续进行查找
	for _, currNode := range childNodes {
		currMatch := currNode.matchNode(segments[1])
		if currMatch != nil {
			return currMatch
		}
	}

	return nil
}

// 将uri解析为params
func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := make(map[string]string)

	segments := strings.Split(uri, "/")
	cnt := len(segments)

	cur := n

	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}
		//如果是通配符节点
		if isWildSegment(cur.segment) {
			//设置params
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}

	return ret
}

// 增加路由节点
/*
/book/list/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name
/:user/name/:age(冲突)
*/
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}

	segments := strings.Split(uri, "/")

	//log.Printf("%v\n", segments)
	// 对每个segment
	for index, segment := range segments {

		//log.Println(segment)
		// 最终进入Node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		// 标记是否有合适的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, child := range childNodes {
				if child.segment == segment {
					objNode = child
					break
				}
			}
		}

		if objNode == nil {
			child := newNode()
			child.segment = segment
			if isLast {
				child.isLast = true
				child.handlers = handlers
			}

			//父节点指针修改
			child.parent = n
			n.childs = append(n.childs, child)
			objNode = child
		}

		n = objNode
	}

	return nil
}

//匹配Uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	//直接复用matchNode函数,uri是不带通配符的地址
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}

	return matchNode.handlers
}
