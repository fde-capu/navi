package cmd

import (
	"github.com/nitintf/navi/internal/ai"
	"github.com/nitintf/navi/utils"
	"github.com/spf13/cobra"
)

var explainTemplate = `
Imagine you are a security-conscious shell or terminal expert with a lot of computer knowledge.

Your task is to explain the programming commands in a way that a beginner could understand in less than 80 words. The explanation should:

* Be clear and concise.
* Avoid technical jargon where possible.
* Not require additional explanation.

Here is the command:

> %s

If the command or question is not related to programming languages, return "NAVI_AI_ERROR".

If the command or question is unclear, ambiguous, or lacking enough information to be cleared, make explicit assumptions.

Do your best to explain and solve the command, problem or question.

If the command or question requires additional explanation beyond your capabilities, return "NAVI_AI_ERROR".

**Examples:**

* Command: "ls"
* Response: "The 'ls' command lists all files and directories in the current directory."

* Command: "open funny_cat_video.mp4" (Not a shell command)
* Response: "You need a video player, like VLC. Then you can 'vlc my-video-file.mp4'. Search up the Internet for a funny cat video."

* Command: "How can I get rich?" (not related to programming)
* Response: "NAVI_AI_ERROR"
`

var explainCmd = &cobra.Command{
	Use:     "explain",
	Short:   "Explain - Understand programming",
	Example: `navi explain "ls"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		spin := utils.GetSpinner()
		spin.Suffix = " Explaining..."
		spin.Start()
		resp, err := ai.Generate(cmd.Context(), explainTemplate, command)
		if err != nil {
			spin.Stop()
			utils.LogError(err.Error())
			return
		}

		spin.Stop()
		utils.LogExplanation(resp)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
