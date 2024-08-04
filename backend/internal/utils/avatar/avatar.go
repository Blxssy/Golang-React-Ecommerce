package avatar

import "fmt"

var avatarBaseURL = "https://api.multiavatar.com/"

func GenerateRandomAvatar(username string) string {
	return fmt.Sprintf("%s%s.png", avatarBaseURL, username)
}
