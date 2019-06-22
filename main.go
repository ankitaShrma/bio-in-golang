package main

import (
	"fmt"
	myInfo "my/info"
	"my/utils"
)

func setAbout() (*myInfo.About, error) {
	aboutHolder := []byte(`{
		"name": "your name here",
		"address":"your address here",
		"contact": "your contact here"
		}`)
	about, err := myInfo.GetAbout(aboutHolder)
	if err != nil {
		return about, err
	}
	return about, nil
}

func setContact() (*myInfo.Contact, error) {
	contactHolder := []byte(`{
			"email":    "your email here",
			"site":     "your website here",
			"github":   "your github link here",
			"linkedIn": "your linkedin link here",
			"twitter":  "your twitter link here"
		}`)
	contact, err := myInfo.GetContact(contactHolder)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

func setExperience() (myInfo.ExperienceCollection, error) {
	experienceHolder := []byte(`[{
		"where": "your work place here",
		"what":"your work here",
		"howLong": "your time period here"
	},
	{
		"where": "your work place here",
		"what":"your work here",
		"howLong": "your time period here"
	}]`)
	experience, err := myInfo.GetExperience(experienceHolder)
	if err != nil {
		return experience, err
	}
	return experience, nil
}

func setSkills() (myInfo.SkillCollection, error) {
	skillHolder := []byte(`[{
		"category": "your skill category here-1",
		"name":["<skill1>", "<skill2>", "<skill3>"]
	},
	{
		"category": "your skill category here-2",
		"name":["<skill1>", "<skill2>", "<skill3>"]
	}]`)
	skill, err := myInfo.GetSkill(skillHolder)
	if err != nil {
		return skill, err
	}
	return skill, nil
}

func main() {

	about, err := setAbout()
	utils.LogIfError(err)
	contact, err := setContact()
	utils.LogIfError(err)
	experience, err := setExperience()
	utils.LogIfError(err)
	skill, err := setSkills()
	utils.LogIfError(err)

	aboutValue := myInfo.DisplayAbout(about, err)
	contactValue := myInfo.DisplayContact(contact, err)
	experienceValue := myInfo.DisplayExperience(experience, err)
	skillValue := myInfo.DisplaySkill(skill, err)
	fmt.Println(aboutValue + contactValue + experienceValue + skillValue)

}
