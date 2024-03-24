package cmd

import (
	"fmt"

	"github.com/nitintf/navi/internal/ai"
	"github.com/nitintf/navi/utils"
	"github.com/spf13/cobra" // Import Cobra library
)

var commandTemplate = `
Imagine you are a security-conscious shell or terminal expert with a lot of computer knowledge.

According to the prompt, choose a type of response:
* Write a single programming language command that achieves the desired outcome, or
* Write a concise explanation to solve the problem, or
* Write a concise code snipet.

The answers must:

* Be appropriate for a general audience (avoid offensive or harmful commands).
* If answering with code, never give additional explanation.

Here is the prompt:

> %s

If the prompt is not related to computers, return "NAVI_AI_ERROR".

If the prompt is not appropriate for a general audience, return "NAVI_AI_ERROR".

If the prompt is unclear or ambiguous, make explicit assumptions.
At first, assume the questions are about bash, Linux and open-source software.

If the prompt requires additional explanation, make explicit assumptions.

If the prompt would return an unsafe command, that could potentially put any user data in risk, return the solution, but make a big warning and pass the buck stating along the lines of: WARNING. POTENTIALLY DESTRUCTIVE COMMAND. USE RESPONSIBLY.

**Examples:**

* Prompt: "List all files in the current directory."
* Response: ls
* Notes: assuming bash.

* Prompt: "Delete all files in the current directory." (Unsafe)
* Response: "rm -rf ."
* Notes: Assuming bash. WARNING: THIS COMMAND IS UNREVERSIBLE. MAKE SURE YOU ARE RESPONSIBLE FOR USING IT."

* Prompt: "Show a funny cat video." (Nothing specific)
* Response: "You need a video player, like VLC. Then you can 'vlc my-video-file.mp4'. Search up the Internet for a funny cat video."
* Notes: Assuming bash and open-source.

* Prompt: 'Say Hello World in Python.'
* Response: print('Hello World')
(no notes)

* Prompt: 'How do I get rich?' (Not related to computers)
* Response: "NAVI_AI_ERROR"

`

var rootCmd = &cobra.Command{
	Use:     "navi",
	Short:   "Navi - Your AI-powered Shell Guide",
	Example: `navi "List all files in the current directory."`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]

		spin := utils.GetSpinner()
		spin.Suffix = " Generating command..."
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
