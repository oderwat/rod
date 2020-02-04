package rod_test

import (
	"fmt"
	"time"

	"github.com/ysmood/kit"
)

func (s *S) TestClosePage() {
	page := s.browser.Page(s.htmlFile("fixtures/click.html"))
	defer page.Close()
	page.Element("button")
}

func (s *S) TestPageElements() {
	s.page.Navigate(s.htmlFile("fixtures/input.html"))
	list := s.page.Elements("input")
	s.Equal("submit", list[2].Eval("() => this.value").String())
}

func (s *S) TestPages() {
	page := s.browser.Page(s.htmlFile("fixtures/click.html"))
	defer page.Close()

	page.Element("button")
	pages := s.browser.Pages()

	s.Len(pages, 3)
}

func (s *S) TestUntilPage() {
	page := s.page.Timeout(3 * time.Second).Navigate(s.htmlFile("fixtures/open-page.html"))
	defer page.CancelTimeout()

	wait := kit.All(func() {
		page.Element("a").Click()
	})

	newPage := page.WaitPage()

	s.Equal("click me", newPage.Element("button").Text())

	wait()
}

func (s *S) TestAlert() {
	page := s.page.Navigate(s.htmlFile("fixtures/alert.html"))

	wait := kit.All(func() {
		page.Element("button").Click()
	})

	page.HandleDialog(true, "")

	wait()
}

func (s *S) TestDownloadFile() {
	srv := kit.MustServer(":0")
	host := srv.Listener.Addr().String()
	content := "test content"

	srv.Engine.GET("/d", func(ctx kit.GinContext) {
		ctx.Writer.WriteHeader(200)
		kit.E(ctx.Writer.Write([]byte(content)))
	})
	srv.Engine.GET("/", func(ctx kit.GinContext) {
		ctx.Header("Content-Type", "text/html;")
		data := []byte(fmt.Sprintf(`<html><a href="//%s/d" download>click</a></html>`, host))
		kit.E(ctx.Writer.Write(data))
	})

	go srv.MustDo()

	page := s.page.Navigate("http://" + host)

	wait := kit.All(func() {
		page.Element("a").Click()
	})

	_, data := page.GetDownloadFile("*")

	s.Equal(content, string(data))
	wait()
}

func (s *S) TestMouse() {
	page := s.page.Navigate(s.htmlFile("fixtures/click.html"))
	mouse := page.Mouse

	mouse.Move(140, 160)
	mouse.Down("left")
	mouse.Up("left")

	s.True(page.Has("[a=ok]"))
}
func (s *S) TestMouseClick() {
	s.browser.Slowmotion = 1
	defer func() { s.browser.Slowmotion = 0 }()

	page := s.page.Navigate(s.htmlFile("fixtures/click.html"))
	mouse := page.Mouse
	mouse.Click("left")
	s.True(page.Has("[a=ok]"))
}

func (s *S) TestDrag() {
	s.T().Skip("not able to use mouse event to simulate it for now")

	page := s.page.Navigate(s.htmlFile("fixtures/drag.html"))
	mouse := page.Mouse

	mouse.Move(60, 30)
	mouse.Down("left")
	kit.E(mouse.MoveE(60, 80, 5))
	mouse.Up("left")

	page.Element(".dropzone:nth-child(2) #draggable")
}

func (s *S) TestPageOthers() {
	p := s.page.Navigate(s.htmlFile("fixtures/input.html"))

	s.Equal("body", p.ElementByJS(`() => document.body`).Describe().Get("node.localName").String())
	s.Len(p.ElementsByJS(`() => document.querySelectorAll('input')`), 3)
	s.EqualValues(1, p.Eval(`() => 1`).Int())
}
