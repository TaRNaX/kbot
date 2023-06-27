/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}
		kbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm kbot %s!", appVersion))
			case "joke":
				err = m.Send(fmt.Sprintf("%s", getRandomJoke()))
			case "fact":
				err = m.Send(fmt.Sprintf("%s", getRandomFact()))
			}

			return err

		})

		kbot.Start()

	},
}

func getRandomJoke() string {

	jokes := []string{
		"Why don't scientists trust atoms? Because they make up everything!",
		"I invented a new word! Plagiarism!",
		"Did you hear about the mathematician who’s afraid of negative numbers? He will stop at nothing to avoid them!",
		"I'm reading a book about anti-gravity. It's impossible to put down!",
		"What's the difference between a snowman and a snowwoman? Snowballs!",
		"Why don’t skeletons fight each other? They don’t have the guts!",
		"Why did the scarecrow win an award? Because he was outstanding in his field!",
		"Why did the bicycle fall over? Because it was two-tired!",
		"How do you organize a space party? You planet!",
		"Did you hear about the mathematician who’s afraid of negative numbers? He will stop at nothing to avoid them!",
	}

	// Get random index
	randomIndex := rand.Intn(len(jokes))

	// Retrieve and return the joke at the random index
	return jokes[randomIndex]
}

func getRandomFact() string {

	facts := []string{
		"The deepest place on Earth is the Mariana Trench in the Pacific Ocean. It’s 36,201 feet (11,034m) deep. That’s almost seven miles!",
		"The longest river in the world is the River Nile, clocking 6,853km in length. Its water resources are shared by 11 different countries, too.",
		"Lobsters are not 'biologically immortal', but they do produce an enzyme that repairs their cells and helps their DNA to replicate indefinitely. Tha's where the myth comes from.",
		"The deepest freshwater lake in the world is Lake Baikal, located in Siberia. It plunges to a whopping depth of 5,315 feet (1,620m). Woah!",
		"Pineapples take two years to grow.",
		"Acacia trees in Africa communicate with each other. They emit gasses to alert other trees to produce the toxin tannin, which protects them from hungry animals.",
		"Armadillos are bulletproof. (This is NOT an invitation to test the fact.)",
		"Niagara Falls never freezes.",
		"Each limestone/granite block that makes up the Great Pyramid of Giza weighs 2.5 tons. And there are 2.3 million of them. Yes, you read that correctly.",
		"It would take you approximately 18 months to walk all the way along The Great Wall of China. (It's over 5,000 miles long).",
		"Scientists say tears tell you the reason for someone crying. If the first drop comes from the right eye, it's tears of joy. Otherwise, it's because of pain.",
		"In the UK in 2019, renewable energy generated more electricity than fossil fuels for the first time ever. Also, did you know Norway gets 0 electricity from coal? And Germany has installed 1 kW of renewable capacity per person in the last decade?",
		"The longest reigning monarch ever was Louis XIV of France. He ruled for 72 years, 110 days. Exhausting.",
		"Marie Curie was the first person ever to win TWO Nobel Prizes - one for physics in 1903, the other for chemistry in 1911 for her work on radioactivity.",
		"King Henry VIII of England had servants called “Grooms of Stool”, who wiped him clean after he visited the toilet. Gross.",
		"0.5 percent of the male population are descended from Genghis Khan. (Scientists did a study in 2003 showing that about 16 million dudes share a Y chromosome with the famous Emperor.)",
		"Iceland, the Faroe Islands, and the Isle of Man all have claims to having the oldest parliament in history, all of which were founded in the 9th and 10th centuries.",
	}

	// Get random index
	randomIndex := rand.Intn(len(facts))

	// Retrieve and return the joke at the random index
	return facts[randomIndex]
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
