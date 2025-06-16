package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/Rezab98/gophercon/automation"
)

func main() {
	// consistent timestamped logs
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	srv := server.NewMCPServer(
		"Desktop Automation MCP",
		"0.1.0",
		server.WithToolCapabilities(false),
	)

	log.Printf("[startup] Initialising Desktop-Automation MCP server – version %s", "0.1.0")

	// Simple click tool powered by robotgo (left-click).
	clickTool := mcp.NewTool(
		"click",
		mcp.WithDescription("Perform a left mouse-button click at the given screen coordinates"),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("Horizontal screen coordinate in pixels"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Vertical screen coordinate in pixels"),
		),
	)

	srv.AddTool(clickTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		xFloat, err := req.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		yFloat, err := req.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		x, y := int(xFloat), int(yFloat)

		if err := automation.Click(x, y); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Clicked at (%d,%d)", x, y)), nil
	})

	// ------------------ Move tool ------------------
	moveTool := mcp.NewTool(
		"move",
		mcp.WithDescription("Move the mouse to coordinates. If durationMs > 0 it animates the movement."),
		mcp.WithNumber("x", mcp.Required(), mcp.Description("Horizontal coordinate")),
		mcp.WithNumber("y", mcp.Required(), mcp.Description("Vertical coordinate")),
		mcp.WithNumber("durationMs", mcp.Description("Duration of animation in milliseconds (optional)")),
	)

	srv.AddTool(moveTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		xF, err := req.RequireFloat("x")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		yF, err := req.RequireFloat("y")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		x, y := int(xF), int(yF)

		argsMap := req.GetArguments()
		duration := 0.0
		if d, ok := argsMap["durationMs"].(float64); ok {
			duration = d
		}

		var actErr error
		if duration > 0 {
			actErr = automation.MoveSmooth(x, y, int(duration))
		} else {
			actErr = automation.Move(x, y)
		}
		if actErr != nil {
			return mcp.NewToolResultError(actErr.Error()), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Moved to (%d,%d)", x, y)), nil
	})

	// ------------------ Type tool ------------------
	typeTool := mcp.NewTool(
		"type",
		mcp.WithDescription("Type the provided text string with optional delay between characters."),
		mcp.WithString("text", mcp.Required(), mcp.Description("Text to type")),
		mcp.WithNumber("delayMs", mcp.Description("Delay per character in milliseconds (optional)")),
	)

	srv.AddTool(typeTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		txt, err := req.RequireString("text")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		argsMap := req.GetArguments()
		delay := 0.0
		if d, ok := argsMap["delayMs"].(float64); ok {
			delay = d
		}

		if err := automation.Type(txt, int(delay)); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("Typed successfully"), nil
	})

	log.Println("[startup] Waiting for MCP client on stdio … (Ctrl-C to quit)")

	if err := server.ServeStdio(srv); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Println("[shutdown] Context canceled – graceful shutdown")
			return
		}
		log.Fatalf("[fatal] MCP server error: %v", err)
	}

	log.Println("[shutdown] MCP server exited cleanly")
}
