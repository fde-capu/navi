package cmd

import (
	"fmt"

	"github.com/nitintf/navi/internal/ai"
	"github.com/nitintf/navi/utils"
	"github.com/spf13/cobra" // Import Cobra library
)

var commandTemplate = `
Imagine you are a security-conscious computer expert with all knowledge.

According to the given task, and to what is asked in the prompt, choose a type of response:

* Write a single command to achieve the desired outcome, or
* Write a code snipet or function, or
* Write a short explanation on how to solve the problem.

All the answers must:

* Be clear and concise.
* Avoid technical jargon where possible.
* Be appropriate for a general audience (avoid offensive or harmful commands).
* If answering with code, use code alone, without additional explanation.

Here is the prompt:

> %s

If the prompt is not related to computers, return "NAVI_AI_ERROR".

If the prompt is not appropriate for a general audience beside computer geeks, return "NAVI_AI_ERROR".

If the prompt is unclear, ambiguous, or lacking information to be solved, make explicit assumptions.

Explain shortly and clearly, as if to child, solving the problem or question.

At first, assume the questions are about bash, Linux and open-source software.

If the command or question requires additional explanation beyond your capabilities, return the reason why.

If the prompt would return an unsafe code that could potentially put any user data in risk, return the solution and make a big warning passing the buck, stating along the lines of: WARNING. POTENTIALLY DESTRUCTIVE COMMAND. USE RESPONSIBLY.

**Examples:**

* Prompt: "List all files in the current directory."
* Response: ls
* Notes: assuming bash.

* Command: "ls"
* Response: "The 'ls' command lists all files and directories in the current directory."

* Prompt: "Delete all files in the current directory." (Unsafe)
* Response: "rm -rf ."
* Notes: Assuming bash. WARNING: THIS COMMAND IS IRREVERSIBLE. MAKE SURE YOU ARE RESPONSIBLE FOR USING IT.

* Prompt: "Show a funny cat video." (Nothing specific)
* Response: "You need a video player, like VLC. Then you can 'vlc my-video-file.mp4'. Search up the Internet for a funny cat video."
* Notes: Assuming bash and open-source.

* Prompt: 'Say Hello World in Python.'
* Response: print('Hello World')
(no notes)

* Prompt: 'How do I get rich?' (Not related to computers)
* Response: "NAVI_AI_ERROR"

* Prompt: 'A Python script that simulates a quantum wave function.'
* Response:
import numpy as np
import matplotlib.pyplot as plt

# Define the system (replace with your specific potential)
def potential(x):
  return 0.5 * x**2  # Simple harmonic oscillator

# Discretize the space (number of points and range)
num_points = 100
x = np.linspace(-5, 5, num_points) 

# Initial wave function (replace with your desired form)
psi = np.exp(-x**2)  # Gaussian wavepacket

# Normalize the wave function (ensure probability integrates to 1)
psi /= np.sqrt(np.sum(psi**2))

# Plot the wave function (absolute value squared for probability density)
plt.plot(x, np.abs(psi)**2)
plt.xlabel("Position (x)")
plt.ylabel("Probability Density (|psi|^2)")
plt.title("Simulating a Quantum Wave Function")
plt.show()
`

var rootCmd = &cobra.Command{
	Use:     "navi",
	Short:   "Navi - AI-powered Code Assistant",
	Example: `navi "List all files in the current directory."`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]

		spin := utils.GetSpinner()
		spin.Suffix = " Generating code..."
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
