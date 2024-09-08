package printers

import (
	"fmt"
	tdlib "github.com/zelenin/go-tdlib/client"
)

func PrintUserInfo(user *tdlib.User) {
	fmt.Println("========== User Information ==========")
	fmt.Printf("ID:                       %d\n", user.Id)
	fmt.Printf("First Name:               %s\n", user.FirstName)
	fmt.Printf("Last Name:                %s\n", user.LastName)
	if user.Usernames != nil {
		fmt.Printf("Usernames:                %v\n", user.Usernames)
	} else {
		fmt.Println("Usernames:                None")
	}
	fmt.Printf("Phone Number:             %s\n", user.PhoneNumber)
	fmt.Printf("Status:                   %s\n", user.Status) // Обробіть статус, як зазначено вище
	fmt.Printf("Is Contact:               %t\n", user.IsContact)
	fmt.Printf("Is Mutual Contact:        %t\n", user.IsMutualContact)
	fmt.Printf("Is Close Friend:          %t\n", user.IsCloseFriend)
	fmt.Printf("Is Verified:              %t\n", user.IsVerified)
	fmt.Printf("Is Premium:               %t\n", user.IsPremium)
	fmt.Printf("Is Support:               %t\n", user.IsSupport)
	fmt.Printf("Is Scam:                  %t\n", user.IsScam)
	fmt.Printf("Is Fake:                  %t\n", user.IsFake)
	fmt.Printf("Has Active Stories:       %t\n", user.HasActiveStories)
	fmt.Printf("Has Unread Active Stories:%t\n", user.HasUnreadActiveStories)
	fmt.Printf("Restricts New Chats:      %t\n", user.RestrictsNewChats)
	fmt.Printf("Have Access:              %t\n", user.HaveAccess)
	fmt.Printf("Language Code:            %s\n", user.LanguageCode)
	fmt.Printf("Added to Attachment Menu: %t\n", user.AddedToAttachmentMenu)
	if user.RestrictionReason != "" {
		fmt.Printf("Restriction Reason:       %s\n", user.RestrictionReason)
	} else {
		fmt.Println("Restriction Reason:       None")
	}
	fmt.Println("=====================================")
}

func PrintUserFullInfo(info *tdlib.UserFullInfo) {
	fmt.Println("========== User Full Information ==========")
	if info.PersonalPhoto != nil {
		fmt.Println("Personal Photo:     Available")
	} else {
		fmt.Println("Personal Photo:     Not Available")
	}

	if info.Photo != nil {
		fmt.Println("Profile Photo:      Available")
	} else {
		fmt.Println("Profile Photo:      Not Available")
	}

	if info.PublicPhoto != nil {
		fmt.Println("Public Photo:       Available")
	} else {
		fmt.Println("Public Photo:       Not Available")
	}

	if info.BlockList != nil {
		fmt.Println("Block List:         Available")
	} else {
		fmt.Println("Block List:         Not Available")
	}

	fmt.Printf("Can Be Called:                 %v\n", info.CanBeCalled)
	fmt.Printf("Supports Video Calls:          %v\n", info.SupportsVideoCalls)
	fmt.Printf("Has Private Calls:             %v\n", info.HasPrivateCalls)
	fmt.Printf("Has Private Forwards:          %v\n", info.HasPrivateForwards)
	fmt.Printf("Has Restricted Voice & Video Notes: %v\n", info.HasRestrictedVoiceAndVideoNoteMessages)
	fmt.Printf("Has Pinned Stories:            %v\n", info.HasPinnedStories)
	fmt.Printf("Needs Phone Number Privacy Exception: %v\n", info.NeedPhoneNumberPrivacyException)
	fmt.Printf("Set Chat Background:           %v\n", info.SetChatBackground)

	if info.Bio != nil {
		fmt.Printf("Bio:                          %s\n", info.Bio.Text)
	} else {
		fmt.Println("Bio:                          Not Available")
	}

	if info.Birthdate != nil {
		fmt.Printf("Birthdate:                     %02d-%02d-%d\n", info.Birthdate.Day, info.Birthdate.Month, info.Birthdate.Year)
	} else {
		fmt.Println("Birthdate:                     Not Available")
	}

	fmt.Printf("Personal Chat ID:              %d\n", info.PersonalChatId)

	if len(info.PremiumGiftOptions) > 0 {
		fmt.Println("Premium Gift Options:         Available")
	} else {
		fmt.Println("Premium Gift Options:         None")
	}

	fmt.Printf("Groups in Common Count:        %d\n", info.GroupInCommonCount)

	if info.BusinessInfo != nil {
		fmt.Println("Business Info:                 Available")
	} else {
		fmt.Println("Business Info:                 Not Available")
	}

	if info.BotInfo != nil {
		fmt.Println("Bot Info:                      Available")
	} else {
		fmt.Println("Bot Info:                      Not Available")
	}

	fmt.Println("=========================================")
}
