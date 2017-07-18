package helloworld

import (
	"context"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func Get(ctx *routing.Context) error {
	return ctx.Write("Hello world!")
}

func GetAgain(ctx *routing.Context) error {
	return ctx.Write("Hello world again!")
}

func Ozzo(ctx context.Context, routeGroup *routing.RouteGroup) error {
	routeGroup.Use(content.TypeNegotiator(content.XML, content.JSON))
	routeGroup.Get("", Get)
	routeGroup.Get("/again", GetAgain)
	return nil
}
