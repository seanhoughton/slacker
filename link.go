package slacker

import "net/url"

// LinkShareDefinition structure contains definition of the bot LinkShare
type LinkShareDefinition struct {
	Description string
	Example     string
	Handler     func(botCtx BotContext, request *url.URL, response ResponseWriter)
}

// NewBotLinkShare creates a new bot LinkShare object
func NewBotLinkShare(domain string, definition *LinkShareDefinition) BotLinkShare {
	return &botLinkShare{
		domain:     domain,
		definition: definition,
	}
}

// BotLinkShare interface
type BotLinkShare interface {
	Domain() string
	Definition() *LinkShareDefinition
	Execute(botCtx BotContext, request *url.URL, response ResponseWriter)
}

// botLinkShare structure contains the bot's LinkShare, description and handler
type botLinkShare struct {
	domain     string
	definition *LinkShareDefinition
}

// Description returns the LinkShare description
func (c *botLinkShare) Definition() *LinkShareDefinition {
	return c.definition
}

// Match determines whether the bot should respond based on the text received
func (c *botLinkShare) Domain() string {
	return c.domain
}

// Execute executes the handler logic
func (c *botLinkShare) Execute(botCtx BotContext, request *url.URL, response ResponseWriter) {
	if c.definition == nil || c.definition.Handler == nil {
		return
	}
	c.definition.Handler(botCtx, request, response)
}
