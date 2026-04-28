package react

import (
	"cbp-gen/models"
	"embed"
	"math/rand"
	"os"
	"strings"
)

var templateFiles embed.FS

// generate a random name for a component string
func RandomComponentName() string {
	var name string
	var adj = []string{"apple", "bridge", "candle", "desert", "engine", "forest", "galaxy", "harbor", "island", "jungle", "kernel", "ladder", "magnet", "nebula", "oasis", "puzzle", "quartz", "rocket", "shadow", "tunnel", "umbrella", "valley", "whisper", "xenon", "yonder", "zephyr", "anchor", "blossom", "crystal", "drift", "ember", "fable", "glimmer", "horizon", "ignite", "jigsaw", "keystone", "lantern", "meadow", "nectar", "orbit", "paradox", "quiver", "ripple", "solstice", "thicket", "utopia", "voyage", "wander", "xylem", "yearn", "zenith", "azure", "breeze", "cascade", "dawn", "echo", "flare", "grove", "haven", "ivory", "jade", "kettle", "lunar", "marble", "nova", "opal", "prairie", "quest", "reef", "summit", "tide", "unity", "vortex", "willow", "xerox", "yacht", "zodiac", "arcade", "binary", "cobalt", "delta", "emberly", "fusion", "glyph", "helix", "ion", "jovial", "kinetic", "legend", "matrix", "nylon", "onyx", "plasma", "quantum", "radar", "signal", "vector"}
	
	id := rand.Intn(len(adj))
	name = adj[id]
	return name
}

func GenComponent(name string, ts bool, args []models.FuncArgs) {
	var output string
	if ts {
		output += name + ".tsx"
	} else {
		output += name + ".jsx"
	}

	if name == "main" {
		name = RandomComponentName()
	}

	data, err :=  templateFiles.ReadFile("component.tmpl")
	if err != nil {
		panic(err)
	}

	result := string(data)

	argsBlock := ""
	last := args[len(args) - 1]
	for _, arg := range args {
		if (arg.ArgName == last.ArgName) && (arg.ArgType == last.ArgType) {
			argsBlock += arg.ArgType + " " + arg.ArgName
		} else {
			argsBlock += arg.ArgType + " " + arg.ArgName + ", "
		}
	}

	result = strings.ReplaceAll(result, "{name}", name)
	result = strings.ReplaceAll(result, "{args}", argsBlock)

	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}