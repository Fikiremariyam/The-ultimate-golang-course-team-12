package main

import (
	"fmt"
	"math/rand"
)

// Email struct holds email and skill level
type Email struct {
	Address    string
	SkillLevel string
}

// Function to shuffle emails and group them
func createMixedGroups(emails []Email, groupSize int) [][]Email {
	// Use a fixed seed to get consistent results
	r := rand.New(rand.NewSource(42)) // Replace 42 with any other fixed number if you prefer a different sequence
	r.Shuffle(len(emails), func(i, j int) {
		emails[i], emails[j] = emails[j], emails[i]
	})

	var groups [][]Email
	for i := 0; i < len(emails); i += groupSize {
		end := i + groupSize
		if end > len(emails) {
			end = len(emails)
		}
		groups = append(groups, emails[i:end])
	}
	return groups
}

// Function to calculate skill level percentages
func calculateSkillPercentages(emails []Email) map[string]float64 {
	skillCount := map[string]int{"beginner": 0, "intermediate": 0, "advanced": 0}
	totalEmails := len(emails)

	// Count each skill level
	for _, email := range emails {
		skillCount[email.SkillLevel]++
	}

	// Calculate percentages
	skillPercentages := make(map[string]float64)
	for skill, count := range skillCount {
		skillPercentages[skill] = (float64(count) / float64(totalEmails)) * 100
	}
	return skillPercentages
}

func main() {
	emails := []Email{
		{"bilalbinabdelkadir@gmail.com", "beginner"},
		{"gizeworkmelkamu70@gmail.com", "intermediate"},
		{"getinet.neu@gmail.com", "advanced"},
		{"getinetamare@gmail.com", "intermediate"},
		{"hawigirma926@gmail.com", "intermediate"},
		{"yordanoslegesse15@gmail.com", "intermediate"},
		{"teklumoges482@gmail.com", "intermediate"},
		{"kirubelkassahun9@gmail.com", "intermediate"},
		{"betekbebe@gmail.com", "advanced"},
		{"mebawubeshet@gmail.com", "advanced"},
		{"enkutatasheshetu96@gmail.com", "intermediate"},
		{"nathanyilmaasrat@gmail.com", "intermediate"},
		{"Kirubelkassahun9@gmail.com", "intermediate"},
		{"azmetefera07@gmail.com", "advanced"},
		{"fikirabrham3@gmail.com", "beginner"},
		{"yonbelay777@gmail.com", "intermediate"},
		{"liduhon3@gmail.com", "beginner"},
		{"captainkevin2008@gmail.com", "intermediate"},
		{"destaget@gmail.com", "intermediate"},
		{"melat.tesfaye.123456@gmail.com", "intermediate"},
		{"elishabu28@gmail.com", "beginner"},
		{"aaron.smith6722073@gmail.com", "intermediate"},
		{"shaanalam369@gmail.com", "advanced"},
		{"tshivam619@gmail.com", "beginner"},
		{"sasti.pk786@gmail.com", "intermediate"},
		{"henokbetiw2000@gmail.com", "intermediate"},
		{"ayelekanke46@gmail.com", "beginner"},
		{"getab2017@gmail.com", "beginner"},
		{"kidusshimelis3@gmail.com", "beginner"},
		{"tsegasolomon635@gmail.com", "intermediate"},
		{"Fikerbdu@gmail.com", "intermediate"},
		{"anteaddisue20@gmail.com", "intermediate"},
		{"eyobahacker@gmail.com", "intermediate"},
		{"fikerbdu@gmail.com", "advanced"},
		{"abenigech14@gmail.com", "beginner"},
		{"zekudk@gmail.com", "intermediate"},
		{"beimnetworku101@gmail.com", "advanced"},
		{"soltad65@gmail.com", "intermediate"},
		{"yonatanawoke@gmail.com", "intermediate"},
		{"rekiksolomon8@gmail.com", "advanced"},
		{"nuredinmaru5@gmail.com", "beginner"},
		{"takelemit25@gmail.com", "intermediate"},
		{"yohanssamuel2014@gmail.com", "advanced"},
		{"yonatanawoke@gmail.com ", "advanced"},
		{"devhaile2@gmail.com", "intermediate"},
		{"anteneh.mulu1921@gmail.com", "intermediate"},
		{"sarigtchw629@gmsarigtcgmail.com", "beginner"},
		{"eliaswakgari12@gmail.com", "beginner"},
		{"surafelwondimagegnhailu@gmail.com", "beginner"},
		{"lidethonelign43@gmail.com", "intermediate"},
		{"weldetsadik2535@gmail.com", "advanced"},
		{"Adonya.getachew@gmail.com", "intermediate"},
		{"anteneh.mulu1921@gmail.com", "intermediate"},
		{"gelilafassil659@gmail.com", "intermediate"},
		{"mearegteame99995555@gmail.com", "intermediate"},
		{"nuredinmaru5@gmail. com", "intermediate"},
		{"keroabdurehman@gmail.com", "beginner"},
		{"behailu.fekadu24@gmail.com", "beginner"},
	}

	groupSize := 5
	groups := createMixedGroups(emails, groupSize)
	skillPercentages := calculateSkillPercentages(emails)
	for i, group := range groups {
		fmt.Printf("Group %d:\n", i+1)
		for _, email := range group {
			fmt.Printf("  %s (%s)\n", email.Address, email.SkillLevel)
		}
		fmt.Println()
	}

	fmt.Println("Skill Level Percentages:")
	for skill, percentage := range skillPercentages {
		fmt.Printf("%s: %.2f%%\n", skill, percentage)
	}
}
