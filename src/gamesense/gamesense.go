package gamesense

type Gamesense struct {
	AppId      string
	State      string
	Details    string
	LargeImage string
	LargeText  string
	SmallImage string
	SmallText  string

	ButtonName string
	ButtonUrl  string
}

func NewGamesense(appId string, state string, details string, largeImage string, largeText string, smallImage string, smallText string, buttonName string, buttonUrl string) *Gamesense {

	if appId == "" || appId == "YOUR_APP_ID" {
		panic("[ERROR] AppId is empty or YOUR_APP_ID")
	}

	return &Gamesense{
		AppId:      appId,
		State:      state,
		Details:    details,
		LargeImage: largeImage,
		LargeText:  largeText,
		SmallImage: smallImage,
		SmallText:  smallText,
		ButtonName: buttonName,
		ButtonUrl:  buttonUrl,
	}
}
