package menus

import (
	"fmt"
	"github.com/faiface/pixel"
	"image/color"
	"math"
	"timsims1717/magicmissile/pkg/camera"
	"timsims1717/magicmissile/pkg/input"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/util"
)

const (
	MaxLines = 10
)

var (
	DefaultColor = color.RGBA{
		R: 109,
		G: 117,
		B: 141,
		A: 255,
	}
	HoverColor = color.RGBA{
		R: 223,
		G: 62,
		B: 35,
		A: 255,
	}
	DisabledColor = color.RGBA{
		R: 109,
		G: 117,
		B: 141,
		A: 255,
	}
	MenuSize = 0.45
	HoverSize = 0.5
)

type Menu struct {
	Key       string
	ItemMap   map[string]*Item
	Items     []*Item
	Hovered   int
	Top       int
	Title     bool
	Roll      bool
	TLines    int

	Tran *object.Object
	Cam  *camera.Camera

	backFn   func()
	openFn   func()
	closeFn  func()
	updateFn func(*input.Input)

	Opened bool
}

func New(key string, cam *camera.Camera) *Menu {
	tran := object.New()
	return &Menu{
		Key:     key,
		ItemMap: map[string]*Item{},
		Items:   []*Item{},
		Tran:    tran,
		Cam:     cam,
	}
}

func (m *Menu) AddItem(key, raw string, right bool) *Item {
	if _, ok := m.ItemMap[key]; ok {
		panic(fmt.Errorf("menu '%s' already has item '%s'", m.Key, key))
	}
	item := NewItem(key, raw, right)
	m.ItemMap[key] = item
	m.Items = append(m.Items, item)
	return item
}

func (m *Menu) InsertItem(key, raw, after string, right bool) *Item {
	if _, ok := m.ItemMap[key]; ok {
		panic(fmt.Errorf("menu '%s' already has item '%s'", m.Key, key))
	}
	item := NewItem(key, raw, right)
	m.ItemMap[key] = item
	i := 0
	for j, itemAfter := range m.Items {
		if itemAfter.Key == after {
			i = j+1
			break
		}
	}
	if i >= len(m.Items) {
		m.Items = append(m.Items, item)
	} else {
		m.Items = append(m.Items[:i], append([]*Item{item}, m.Items[i:]...)...)
	}
	return item
}

func (m *Menu) RemoveItem(key string) {
	index := -1
	for i, item := range m.Items {
		if item.Key == key {
			index = i
			break
		}
	}
	if index != -1 {
		if len(m.Items) > 1 {
			m.Items = append(m.Items[:index], m.Items[index+1:]...)
		} else {
			m.Items = []*Item{}
		}
	}
	delete(m.ItemMap, key)
}

func (m *Menu) Open() {
	m.setHover(-1)
	if m.openFn != nil {
		m.openFn()
	}
	m.Opened = true
}

func (m *Menu) IsOpen() bool {
	return m.Opened
}

func (m *Menu) IsClosed() bool {
	return !m.Opened
}

func (m *Menu) Close() {
	if m.closeFn != nil {
		m.closeFn()
	}
	m.Opened = false
}

func (m *Menu) CloseInstant() {
	if m.closeFn != nil {
		m.closeFn()
	}
}

func (m *Menu) Update(in *input.Input) {
	if m.Opened && in != nil {
		m.UpdateView(in)
	}
	m.UpdateSize()
	if m.Opened && in != nil {
		m.UpdateItems(in)
	}
}

func (m *Menu) UpdateView(in *input.Input) {
	if in.Get("scrollUp").JustPressed() {
		m.menuUp()
	}
	if in.Get("scrollDown").JustPressed() {
		m.menuDown()
	}
	dir := -1
	if in.Get("menuUp").JustPressed() || in.Get("menuUp").Repeated() {
		dir = 0
	} else if in.Get("menuDown").JustPressed() || in.Get("menuDown").Repeated() {
		dir = 1
	} else if in.Get("menuRight").JustPressed() || in.Get("menuRight").Repeated() {
		dir = 2
	} else if in.Get("menuLeft").JustPressed() || in.Get("menuLeft").Repeated() {
		dir = 3
	}
	if dir != -1 {
		m.GetNextHover(dir, m.Hovered, in)
	} else if in.MouseMoved {
		for i, item := range m.Items {
			if !item.Hovered && !item.Disabled && !item.NoHover && !item.noShowT {
				b := item.Text.Text.BoundsOf(item.Raw)
				point := in.World
				if item.Right {
					point.X += 15.
				} else {
					point.X -= b.W() * 0.5 * MenuSize
				}
				point.Y -= b.H() * 1.8 * MenuSize
				if util.PointInside(point, b, item.Text.Obj.Mat) {
					m.setHover(i)
				}
			}
		}
	}
}

func (m *Menu) UpdateSize() {
	minWidth := 8.
	minHeight := 8.
	sameLine := false
	lines := 0
	tLines := 0
	for i, item := range m.Items {
		if item.Ignore {
			item.noShowT = true
			continue
		}
		visible := (m.Title && i == 0) || (tLines >= m.Top && lines < MaxLines)
		//if (m.Title && i == 0) || (tLines >= m.Top && lines < MaxLines) {
		item.CurrLine = tLines
		bW := item.Text.Text.Bounds().W() * MenuSize
		sW := 0.
		if !item.Right && i+1 < len(m.Items) && m.Items[i+1].Right {
			next := m.Items[i+1]
			sW = (next.Text.Text.Bounds().W() + next.Text.Text.BoundsOf("   ").W()) * MenuSize
			sameLine = true
		}
		minWidth = math.Max(bW+sW, minWidth)
		if !sameLine {
			if visible {
				minHeight += item.Text.Text.LineHeight * MenuSize
				lines++
			}
			tLines++
		}
		sameLine = false
		item.noShowT = !visible
		//} else {
		//	item.CurrLine = tLines
		//	if !item.Right && i+1 < len(m.Items) && m.Items[i+1].Right {
		//		sameLine = true
		//	}
		//	if !sameLine {
		//		tLines++
		//	}
		//	sameLine = false
		//	item.noShowT = true
		//}
	}
	m.TLines = tLines
	minWidth += 15.
	line := 0
	for i, item := range m.Items {
		if !item.noShowT {
			if item.Right {
				item.Text.SetPos(pixel.V(minWidth*0.5 - 10., minHeight*0.5 - float64(line+1)*item.Text.Text.LineHeight * MenuSize))
			} else {
				nextY := minHeight*0.5 - float64(line+1)*item.Text.Text.LineHeight * MenuSize
				nextX := minWidth*-0.5 + 5.
				item.Text.SetPos(pixel.V(nextX, nextY))
			}
			if item.Right || i >= len(m.Items)-1 || !m.Items[i+1].Right {
				line++
			}
		}
	}
}

func (m *Menu) UpdateItems(in *input.Input) {
	if in.Get("menuBack").JustPressed() {
		m.Back()
		in.Get("menuBack").Consume()
	} else if in.Get("menuSelect").JustPressed() && m.Hovered != -1 {
		if m.Items[m.Hovered].clickFn != nil {
			m.Items[m.Hovered].clickFn()
		}
		in.Get("menuSelect").Consume()
	} else if in.Get("click").JustPressed() && m.Hovered != -1 {
		if m.Items[m.Hovered].clickFn != nil {
			m.Items[m.Hovered].clickFn()
		}
		in.Get("click").Consume()
	} else if in.Get("menuRight").JustPressed() && m.Hovered != -1 {
		if m.Items[m.Hovered].rightFn != nil {
			m.Items[m.Hovered].rightFn()
		}
		in.Get("menuRight").Consume()
	} else if in.Get("menuLeft").JustPressed() && m.Hovered != -1 {
		if m.Items[m.Hovered].leftFn != nil {
			m.Items[m.Hovered].leftFn()
		}
		in.Get("menuLeft").Consume()
	}
	for _, item := range m.Items {
		item.Update()
	}
}

func (m *Menu) setHover(nextI int) {
	if nextI == -1 {
		hover := false
		for i, item := range m.Items {
			if !hover && !item.disabled && !item.NoHover {
				m.setHover(i)
				hover = true
			} else {
				item.Hovered = false
			}
		}
	} else {
		if m.Hovered != -1 {
			prev := m.Items[m.Hovered]
			prev.Hovered = false
			if prev.unHoverFn != nil {
				prev.unHoverFn()
			}
		}
		next := m.Items[nextI]
		next.Hovered = true
		m.Hovered = nextI
		sfx.SoundPlayer.PlaySound("click", 0.0)
		if next.hoverFn != nil {
			next.hoverFn()
		}
		m.setTop(next.CurrLine)
	}
}

func (m *Menu) setTop(line int) {
	if line < m.Top {
		m.Top = line
	} else if m.Title && line >= m.Top+MaxLines-1 {
		m.Top = line - MaxLines + 2
	} else if line >= m.Top+MaxLines {
		m.Top = line - MaxLines + 1
	}
}

func (m *Menu) menuUp() {
	m.Top--
	if m.Top < 0 {
		m.Top = 0
	}
}

func (m *Menu) menuDown() {
	m.Top++
	if m.Top > m.TLines-MaxLines+1 {
		m.Top = m.TLines - MaxLines + 1
	}
}

func (m *Menu) UnhoverAll() {
	for _, item := range m.Items {
		if item.Hovered && item.unHoverFn != nil {
			item.unHoverFn()
		}
		item.Hovered = false
	}
	m.Hovered = -1
}

func (m *Menu) GetNextHover(dir, curr int, in *input.Input) {
	if curr == -1 {
		m.setHover(-1)
	}
	if dir == 0 || dir == 1 {
		r := false
		if curr != -1 {
			r = m.Items[curr].Right
		}
		m.GetNextHoverVert(dir, curr, r, in)
	} else {
		m.GetNextHoverHor(dir, curr, in)
	}
}

func (m *Menu) GetNextHoverHor(dir, curr int, in *input.Input) {
	this := m.Items[curr]
	nextI := -1
	if dir == 2 && !this.Right && curr < len(m.Items)-1 {
		nextI = curr + 1
	} else if dir == 3 && this.Right && curr > 0 {
		nextI = curr - 1
	}
	if nextI != -1 {
		next := m.Items[nextI]
		if next.Right != this.Right && !next.Disabled && !next.NoHover && !next.noShowT {
			m.setHover(nextI)
			if dir == 2 {
				in.Get("menuRight").Consume()
			} else {
				in.Get("menuLeft").Consume()
			}
		}
	}
}

func (m *Menu) GetNextHoverVert(dir, curr int, right bool, in *input.Input) {
	nextI := curr
	if dir == 0 {
		nextI--
	} else {
		nextI++
	}
	if !m.Roll && (nextI >= len(m.Items) || nextI < 0) {
		return
	}
	if nextI < 0 {
		nextI += len(m.Items)
	}
	nextI %= len(m.Items)
	next := m.Items[nextI]
	if next.Disabled || next.NoHover || next.Ignore || next.Right != right {
		m.GetNextHoverVert(dir, nextI, right, in)
	} else {
		m.setHover(nextI)
		if dir == 0 {
			in.Get("menuUp").Consume()
		} else {
			in.Get("menuDown").Consume()
		}
	}
}

func (m *Menu) Draw(target pixel.Target) {
	if m.Opened {
		for _, item := range m.Items {
			item.Draw(target)
		}
	}
}

func (m *Menu) Back() {
	if m.backFn != nil {
		m.backFn()
	} else {
		m.Close()
	}
}

func (m *Menu) SetBackFn(fn func()) {
	m.backFn = fn
}

func (m *Menu) SetOpenFn(fn func()) {
	m.openFn = fn
}

func (m *Menu) SetCloseFn(fn func()) {
	m.closeFn = fn
}
