package info

import (
	"encoding/json"
	"my/utils"
	"strings"
)

// About contains 'about me' information
type About struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact string `json:"contact"`
	// age int
}

// Contact contains links to the accounts for further contacts
type Contact struct {
	Email    string `json:"email"`
	Site     string `json:"site"`
	Github   string `json:"github"`
	LinkedIn string `json:"linkedIn"`
	Twitter  string `json:"twitter"`
}

// Experience contains the enumeration of various work experiences
type Experience struct {
	Where   string `json:"where"`
	What    string `json:"what"`
	HowLong string `json:"howLong"`
}

// ExperienceCollection contains the list of the Experience
type ExperienceCollection []Experience

// Skill contains the enumeration of various skills
type Skill struct {
	Category string   `json:"category"`
	Name     []string `json:"name"`
}

// SkillCollection contains list of Skills
type SkillCollection []Skill

// GetAbout returns 'about me' information
func GetAbout(aboutData []byte) (*About, error) {
	about := &About{}
	if aboutData != nil {
		err := json.Unmarshal(aboutData, &about)
		if err != nil {
			return about, err
		}
	}
	return about, nil
}

// DisplayAbout returns 'about me' information to be displayed
func DisplayAbout(about *About, err error) string {
	if err != nil {
		return ""
	}
	var a map[string]string
	aData, _ := json.Marshal(about)
	json.Unmarshal(aData, &a)
	abt := utils.GetDisplayItem(a, "about")
	return abt
}

// GetContact returns contact information
func GetContact(contactData []byte) (*Contact, error) {
	contact := &Contact{}
	if contactData != nil {
		err := json.Unmarshal(contactData, &contact)
		if err != nil {
			return contact, err
		}
	}
	return contact, nil
}

// DisplayContact returns contact information to be displayed
func DisplayContact(contact *Contact, err error) string {
	if err != nil {
		return ""
	}
	var c map[string]string
	cData, _ := json.Marshal(contact)
	json.Unmarshal(cData, &c)
	cnt := utils.GetDisplayItem(c, "contact")
	return cnt
}

// GetExperience returns list of experiences
func GetExperience(experienceData []byte) (ExperienceCollection, error) {
	// Need to range over the array items. (So don't use pointer to the struct)
	experience := ExperienceCollection{}
	if experienceData != nil {
		err := json.Unmarshal(experienceData, &experience)
		if err != nil {
			return experience, err
		}
	}
	return experience, nil
}

// DisplayExperience returns experiences to be displayed
func DisplayExperience(experience ExperienceCollection, err error) string {
	if err != nil {
		return ""
	}
	expMap := make([]map[string]interface{}, 0)
	for _, e := range experience {
		expMap = append(expMap, map[string]interface{}{"where": e.Where, "what": e.What, "howLong": e.HowLong})
	}
	exp := utils.GetDisplayList(expMap, "experience")
	return exp
}

// GetSkill returns list of skills
func GetSkill(skillData []byte) (SkillCollection, error) {
	// Need to range over the array items. (So don't use pointer to the struct)
	skill := SkillCollection{}
	if skillData != nil {
		err := json.Unmarshal(skillData, &skill)
		if err != nil {
			return skill, err
		}
	}
	return skill, nil
}

// DisplaySkill returns skills to be displayed
func DisplaySkill(skill SkillCollection, err error) string {
	if err != nil {
		return ""
	}
	skMap := make([]map[string]interface{}, 0)
	for _, s := range skill {
		skMap = append(skMap, map[string]interface{}{"category": s.Category, "name": strings.Join(s.Name, ", ")})
	}
	skl := utils.GetDisplayList(skMap, "skills")
	return skl
}
