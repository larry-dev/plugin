package snowflake

import "testing"

func TestNode_Generate(t *testing.T) {
	node, err := NewNode(1)
	if err!=nil{
		t.Error(err)
	}
	id:=node.Generate()
	t.Log(id.String())
	t.Log(id.Time())
	t.Log(id.Node())
	t.Log(id.Step())
	t.Log("=============")
	id,err=ParseString(id.String())
	t.Log(id.String())
	t.Log(id.Time())
	t.Log(id.Node())
	t.Log(id.Step())
}