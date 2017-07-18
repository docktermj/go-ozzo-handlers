package streaming

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func Get(ctx *routing.Context) error {
	return ctx.Write("Hello world!")
}

func StreamDate(c *routing.Context) error {

	// Need to cast http.ResponseWriter as a "flusher".

	w := c.Response
	fmt.Printf("W:  %+v\n", w)
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	// Stream the response body.  Note: loop will not stop.

	ticker := time.NewTicker(time.Millisecond * 500)
	for aTime := range ticker.C {
		c.Write([]byte(aTime.String()))
		flusher.Flush()
	}

	return c.Write(">>> Shouldn't get here. <<<")
}

func Ozzo(ctx context.Context, routeGroup *routing.RouteGroup) error {
	routeGroup.Use(content.TypeNegotiator(content.XML, content.JSON))
	routeGroup.Get("", Get)
	routeGroup.Get("/date", StreamDate)
	return nil
}
