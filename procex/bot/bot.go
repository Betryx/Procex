package main

import (
	"fmt"
	"log"
	"os"

	"procex/ProcessManager"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load("config.env")
	token := os.Getenv("TOKEN")
	bot, err := discordgo.New("Bot " + token)
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "start",
            Description: "Starts a specific process based on project directory and starting command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name: "directory",
					Description: "Working directory from where you can start your project",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "command",
                    Description: "Command to start your project",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
				},
			
			},
		},
		{
			Name:        "stop",
            Description: "Stops a specific process based on project directory and stopping command",
            Options: []*discordgo.ApplicationCommandOption{
                {
                    Name: "pid",
                    Description: "id of the process to stop",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
                },
            
            },
		},
		{
			Name:        "list",
            Description: "Lists all running processes",
            
		},
		{
			Name:        "auth",
            Description: "Authenticates a user in order to asign a workspace",
           
		},
		{
			Name:        "workspaces",
            Description: "curent existing workspaces",
           
		},
		{
			Name:        "first_time_start",
            Description: "Configures the active directory for projects as well as permissions",
			Options: []*discordgo.ApplicationCommandOption{
                {
                    Name: "path",
                    Description: "where your active directory would be",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,

                },
				{
					Name:        "permited_users",
                    Description: "Users who have access to processual managment ",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
				},
				{
					Name:        "owner",
                    Description: "The workspace owner",
                    Type:        discordgo.ApplicationCommandOptionUser,
                    Required:    true,
				},
				{
					Name:        "token",
                    Description: "Your personal github access token used to fetch repositories as well as grant push pull access",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
				},
				
            
            
            },
		},
		{
			Name:        "pull",
            Description: "pull a repository and run the code from it",
           
		},
		{
			Name:        "project_create",
            Description: "Creates a new project ",
           
		},

}
	
	
	if err != nil {
        log.Fatal("Error creating Discord session:", err)
    }
	bot.AddHandler(CommandHandler)
	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = bot.Open()
	SlashCommandCreator(bot, commands)
	bot.UpdateWatchStatus(0, "the servers")
	if err != nil {
        log.Fatal("Error opening connection:", err)
    }
	defer bot.Close()
	fmt.Println("Bot is running...")
	select {}

}



func SlashCommandCreator(s *discordgo.Session, commands []*discordgo.ApplicationCommand) {
    for _, cmd := range commands {
        _, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		fmt.Printf("Created commands %q: %v\n", cmd.Name, err)
		
        if err != nil {
            fmt.Printf("Cannot create command %q: %v\n", cmd.Name, err)
        }
    }
}


func CommandHandler(s *discordgo.Session, m *discordgo.InteractionCreate){
	name := m.ApplicationCommandData().Name

	switch name {
	case "start":
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Process started",
						Description: "Process with XXXXXXXXXX ID launched",
					},
					

				},
            },
		})
	case "stop":
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Process started ",
						Description: "Process with XXXXXXXXXX ID launched",
					},
					

				},
            },
		})
	
	
	
	case "auth":
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Process started",
						Description: "Process with XXXXXXXXXX ID launched",
					},
					

				},
            },

		})
		


	case "list":
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: ProcessManager.List(),
            },
		})
	case "workspaces":
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Curent active workspaces",
						Description: ProcessManager.Fetch_workspaces(),
						Footer: &discordgo.MessageEmbedFooter{
							Text: "",
							IconURL: "https://flourishing-begonia-5c93da.netlify.app/WorkDir.png",
						},
					},
					

				},
            },
		})
	
	default:
		return
	}
	
		
	
}