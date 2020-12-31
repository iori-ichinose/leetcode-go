/*
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
 * 2020.12.20
 * Golang 12ms(92.13%), 6.7MB(90.02%)
 */
package binaryTree

import (
	"strconv"
	"strings"
)

type Codec struct {
	res        strings.Builder
	stringList []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) encoding(root *TreeNode) {
	if root == nil {
		this.res.WriteString("nil ")
	} else {
		this.res.WriteString(strconv.Itoa(root.Val))
		this.res.WriteRune(' ')
		this.encoding(root.Left)
		this.encoding(root.Right)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.encoding(root)
	return this.res.String()
}

func (this *Codec) decoding() *TreeNode {
	if this.stringList[0] == "nil" {
		this.stringList = this.stringList[1:]
		return nil
	}
	num, _ := strconv.Atoi(this.stringList[0])
	this.stringList = this.stringList[1:]

	root := &TreeNode{Val: num}
	root.Left = this.decoding()
	root.Right = this.decoding()

	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.stringList = strings.Split(data, " ")
	return this.decoding()
}
