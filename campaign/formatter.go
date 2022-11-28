package campaign

import (
	"fmt"
	"strings"
)

//merubah keluar API dan mengubah besar kecilnya huruf depan
type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amout"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.Slug = campaign.Slug
	CampaignFormatter.UserId = campaign.UserID
	CampaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return CampaignFormatter
}

//Jika data where user id maka lempar ke formatCampaign
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	//pakai refactori {} default jika campaign kosong mengembalikan array kosong
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageUrl         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	BackerCount      int                       `json:"backer_count"`
	UserId           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	User             CampaignUserFormatter     `json:"user"`
	Perks            []string                  `json:"perks"`
	Image            []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	CampaignDetailFormatter := CampaignDetailFormatter{}
	CampaignDetailFormatter.ID = campaign.ID
	CampaignDetailFormatter.Name = campaign.Name
	CampaignDetailFormatter.ShortDescription = campaign.ShortDescription
	CampaignDetailFormatter.Description = campaign.Description
	CampaignDetailFormatter.GoalAmount = campaign.GoalAmount
	CampaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignDetailFormatter.BackerCount = campaign.BackerCount
	CampaignDetailFormatter.UserId = campaign.UserID
	CampaignDetailFormatter.Slug = campaign.Slug
	CampaignDetailFormatter.ImageUrl = ""
	fmt.Println("xxx :", campaign.CampaignImages[0].FileName)
	if len(campaign.CampaignImages) > 0 {
		CampaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		//trimSpace hapus spasi string
		perks = append(perks, strings.TrimSpace(perk))
	}
	CampaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	CampaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}
	for _, img := range campaign.CampaignImages {
		campaignImagesFormatter := CampaignImagesFormatter{}
		campaignImagesFormatter.ImageUrl = img.FileName

		isPrimary := false
		if img.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImagesFormatter.IsPrimary = isPrimary
		images = append(images, campaignImagesFormatter)

	}

	CampaignDetailFormatter.Image = images

	return CampaignDetailFormatter
}
