package utils

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

func getStringFromList(list []map[string]interface{}, key string) string {
	green := color.FgLightGreen.Render
	yellow := color.FgLightYellow.Render
	magenta := color.FgLightCyan.Render

	header := color.LightMagenta.Sprintf("\nMY %s:\n\n", strings.ToUpper(key))

	var b strings.Builder
	b.WriteString(header)
	for i, s := range list {
		switch k := key; k {
		case "experience":
			fmt.Fprintf(&b, "%d. At %s as %s for %s\n", i+1, green(s["where"]), yellow(s["what"]), magenta(s["howLong"]))
		case "skills":
			fmt.Fprintf(&b, "%s: %s\n", green(s["category"]), yellow(s["name"]))
		default:
			fmt.Fprintf(&b, "")
		}
	}
	return b.String()
}

// GetDisplayItem returns the info string to be displayed
func GetDisplayItem(data map[string]string, key string) string {
	green := color.FgLightGreen.Render
	yellow := color.FgLightYellow.Render
	magenta := color.FgLightCyan.Render

	header := color.LightMagenta.Sprintf("\nMY %s:\n\n", strings.ToUpper(key))
	dataHolder := map[string]string{
		"about": header + fmt.Sprintf("Name: %s\nAddress: %s\nContact: %s\n",
			yellow(data["name"]), green(data["address"]),
			magenta(data["contact"])),
		"contact": header + fmt.Sprintf("Email: %s\nSite: %s\nGithub: %s\nLinkedIn: %s\nTwitter: %s\n",
			yellow(data["email"]), green(data["site"]),
			yellow(data["github"]), green(data["linkedIn"]),
			magenta(data["twitter"])),
	}
	return dataHolder[key]
}

// GetDisplayList returns the string for the array of items to be displayed
func GetDisplayList(data []map[string]interface{}, key string) string {
	dataHolder := map[string]string{
		"experience": getStringFromList(data, "experience"),
		"skills":     getStringFromList(data, "skills"),
	}
	return dataHolder[key]
}
