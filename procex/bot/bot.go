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
					Name:        "directory",
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
				{
					Name:        "personal_code",
					Description: "Code used for quick authentication only shown in DM",
					Type:        discordgo.ApplicationCommandOptionString,  // Added Type
				},
			},
		},
		{
			Name:        "stop",
			Description: "Stops a specific process based on project directory and stopping command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "process_name",  // Changed the name to follow the proper naming convention
					Description: "Name of the process to stop",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "auth_code",
					Description: "Uuid based code used for quick authentication only shown in DM",
					Type:        discordgo.ApplicationCommandOptionString,
					Required: true,  // Added Type
				},
				{
					Name:        "workspace_name",
					Description: "The workspace where the process exists(neccesary for auth)",
					Type:        discordgo.ApplicationCommandOptionString,  // Added Type
					Required:    true,
				},
			},
		},
		{
			Name:        "restart",
			Description: "Restarts a specific process",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "process_name",  // Changed the name to follow the proper naming convention
					Description: "Name of the process to restart",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "personal_code",
					Description: "Code used for quick authentication only shown in DM",
					Type:        discordgo.ApplicationCommandOptionString,  // Added Type
					Required:    true,
				},
				{
					Name:        "workspace_name",
					Description: "The workspace where the process exists(neccesary for auth)",
					Type:        discordgo.ApplicationCommandOptionString,  // Added Type
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
			Description: "Authenticates a user to access a workspace",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "workspace",  // Changed the name to lowercase
					Description: "A workspace to which the user gains access",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "user",  // Changed the name to lowercase
					Description: "The user to be authenticated",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
		{
			Name:        "workspaces",
			Description: "Current existing workspaces",
		},
		{
			Name:        "workspace_create",
			Description: "Creates a new workspace",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "name",
					Description: "Where your active directory would be",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "permitted_users",  // Fixed the typo in 'permitted'
					Description: "Users who have access to process management",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
				
				{
					Name:        "owner_token",
					Description: "Your personal GitHub access token used to fetch repositories and grant push/pull access",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "owner",  // Fixed the typo in 'permitted'
					Description: "The user who owns the workspace",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
				{
					Name:        "github_username",  // Fixed the typo in 'permitted'
					Description: "Used in order to authenticate requests to github",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:        "pull",
			Description: "Pull a repository and run the code from it",
		},
		{
			Name:        "project_create",
			Description: "Creates a new project",
		},
	}
	
	
	if err != nil {
        log.Fatal("Error creating Discord session:", err)
    }
	bot.AddHandler(CommandHandler)
	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = bot.Open()
	SlashCommandCreator(bot, commands)
	// deleteAllGlobalCommands(bot)
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
// func deleteAllGlobalCommands(s *discordgo.Session) error {
// 	// Fetch all global commands
// 	commands, err := s.ApplicationCommands(s.State.User.ID, "")
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch commands: %w", err)
// 	}

// 	// Loop through each command and delete it
// 	for _, command := range commands {
// 		err := s.ApplicationCommandDelete(s.State.User.ID, "",command.ID)
// 		if err != nil {
// 			return fmt.Errorf("failed to delete command ID %s: %w", command.ID, err)
// 		}
// 		fmt.Printf("Global command ID %s deleted successfully.\n", command.ID)
// 	}

// 	return nil
// }

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
	case "workspace_create":
		
		
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			
			Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Workspace inited successfully",
						Description: ProcessManager.WorkSpaceCreate(ProcessManager.WorkSpace{
							Name: m.ApplicationCommandData().Options[0].StringValue(),
							Owner:m.ApplicationCommandData().Options[3].UserValue(s).Username,
							Token:m.ApplicationCommandData().Options[2].StringValue(),
							PermitedUsers: []ProcessManager.PermitedUser{
								{
									UserName: m.ApplicationCommandData().Options[1].UserValue(s).Username,
									ID: m.ApplicationCommandData().Options[1].UserValue(s).ID,
								},
								{
									UserName: "",
									ID: "",

									

								},
							},
							GitHubUserName: m.ApplicationCommandData().Options[4].StringValue(),


						

						}),
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
						Title: "Process stopped successfully",
						Description: ProcessManager.Stop(ProcessManager.ProcessStop{
							ProcessName: m.ApplicationCommandData().Options[0].StringValue(),
							WorkSpace: m.ApplicationCommandData().Options[2].StringValue(),
							Uuid: m.ApplicationCommandData().Options[1].StringValue(),


						}),
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