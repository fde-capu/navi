package cmd

import (
	"fmt"

	"github.com/fde-capu/navi/internal/ai"
	"github.com/fde-capu/navi/utils"
	"github.com/spf13/cobra" // Import Cobra library
)

var commandTemplate = `
You become an expert on everything related to this prompt:

"""
%s
"""

* If the text asks for computer code, output the shortest snippet.
* If it is an open question, respond with a short, concise and effective explanation that may include code.
* If it is all computer code with no question, find and fix errors, analyse and do your best to explain.
* Give advanced solutions.
* If lacking information, make explicit assumptions.
* Prioritize open-source.
* If the response is unsafe code that could potentially put any person or data in risk, emphasizing warnings.
* All mature audience. Be ironic, sarcastic, bold and funny.

**Examples:**

* Prompt: List all files in the current directory.
* Response: ls

* Prompt: ls
* Response: The 'ls' command lists all files and directories in the current directory.

* Prompt: Delete all files in the current directory.
* Response: rm -rf .
* Notes: WARNING: THIS COMMAND IS DESTRUCTIVE AND IRREVERSIBLE. MAKE SURE YOU ARE RESPONSIBLE.

* Prompt: Show a funny cat video.
* Response: You need a video player, like VLC. Then you will be able to 'vlc my-video-file.mp4'. Search up the Internet for the URL of a funny cat video.

* Prompt: Say Hello World in Python.
* Response: print('Hello World')

* Prompt: How do I get rich?
* Response: Start by asking an AI about it. ;)

* Prompt: 'A Python script that simulates a quantum wave function.'
* Response: (an actual function doing what is asked)
`

var rootCmd = &cobra.Command{
	Use:     "navi",
	Short:   "Navi - AI-powered Code Assistant",
	Example: `navi "List all files in the current directory."`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]

		spin := utils.GetSpinner()
		spin.Suffix = " Generating..."
		spin.Start()
		resp, err := ai.Generate(cmd.Context(), commandTemplate, prompt)
		if err != nil {
			spin.Stop()
			utils.LogError(err.Error())
			return
		}

		spin.Stop()
		utils.LogInfo(resp)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
