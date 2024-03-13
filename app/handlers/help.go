package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// NewMessageHandler tạo một MessageHandler mới với các thiết lập mặc định.
func Help() *Config {
	return BotConfig()
}

// Setup thiết lập MessageHandler với một phiên Discord.
func (ms *Config) Setup(s *discordgo.Session) {
	// Đăng ký một event handler để xử lý lệnh !help
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Kiểm tra xem tin nhắn có bắt đầu bằng tiền tố của MessageHandler không
		if strings.HasPrefix(m.Content, ms.Prefix+"help") {
			// Tạo một danh sách các chức năng của bot
			commands := []string{
				"1. Lệnh 1",
				"2. Lệnh 2",
				"3. Lệnh 3",
				// Thêm các chức năng khác tại đây...
			}

			// Tạo một widget để hiển thị danh sách chức năng
			widget := &discordgo.MessageEmbed{
				Title:       "Danh sách chức năng của bot:",
				Description: "Dưới đây là danh sách các chức năng mà bot cung cấp:",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Chức năng",
						Value: strings.Join(commands, "\n"),
					},
				},
				Color: 0x00ff00,
			}

			// Gửi widget vào kênh mà tin nhắn được gửi từ
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, widget)
			if err != nil {
				fmt.Println("Error sending help message:", err)
			}
		}
	})
}
