package commands

import (
	. "github.com/quarnster/util/text"
	. "lime/backend"
)

type (
	FindUnderExpandCommand struct {
		DefaultCommand
	}
	SingleSelectionCommand struct {
		DefaultCommand
	}
)

func (c *SingleSelectionCommand) Run(v *View, e *Edit, args Args) error {
	r := v.Sel().Get(0)
	v.Sel().Clear()
	v.Sel().Add(r)
	return nil
}

func (c *FindUnderExpandCommand) Run(v *View, e *Edit, args Args) error {
	sel := v.Sel()
	rs := sel.Regions()

	he := sel.HasEmpty()
	if he {
		for i, r := range rs {
			if r2 := v.Buffer().Word(r.A); r2.Size() > r.Size() {
				rs[i] = r2
			}
		}
		sel.Clear()
		sel.AddAll(rs)
	} else {
		last := rs[len(rs)-1]
		b := v.Buffer()
		data := b.SubstrR(last)
		next := last
		size := last.Size()
		next.A += size
		next.B += size
		buf := b.SubstrR(Region{next.A, next.B})
		for next.End() < b.Size() {
			buf[size-1] = b.Index(next.B - 1)
			found := true
			for j, r := range buf {
				if r != data[j] {
					found = false
					break
				}
			}
			if found {
				sel.Add(next)
				break
			}
			copy(buf, buf[1:])
			next.A += 1
			next.B += 1
		}
	}
	return nil
}

func init() {
	register([]cmd{
		{"find_under_expand", &FindUnderExpandCommand{}},
		{"single_selection", &SingleSelectionCommand{}},
	})
}
