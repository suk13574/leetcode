import (
	"sort"
	"strconv"
	"strings"
)

type UserStatus struct {
	isOnline bool
	endTime  int
}

func countMentions(numberOfUsers int, events [][]string) []int {
	// sort by timestamp
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] == events[j][1] {
			return events[i][0] == "OFFLINE"
		}
		ti, _ := strconv.Atoi(events[i][1])
		tj, _ := strconv.Atoi(events[j][1])
		return ti < tj
	})

	userStatus := make([]*UserStatus, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		userStatus[i] = &UserStatus{isOnline: true}
	}

	updateStatus := func(now int) {
		for _, us := range userStatus {
			if !us.isOnline && us.endTime <= now {
				us.isOnline = true
			}
		}
	}

	res := make([]int, numberOfUsers)
	base := 0
	for _, event := range events {
		now, _ := strconv.Atoi(event[1])
		users := strings.Split(event[2], " ")

		updateStatus(now)

		switch event[0] {
		case "MESSAGE":
			for _, u := range users {
				switch u {
				case "ALL":
					base++
				case "HERE":
					for i, us := range userStatus {
						if us.isOnline {
							res[i]++
						}
					}
				default:
					id, _ := strconv.Atoi(u[2:])
					res[id]++
				}
			}
		case "OFFLINE":
			id, _ := strconv.Atoi(users[0])
			userStatus[id].isOnline = false
			userStatus[id].endTime = now + 60
		}
	}

	if base > 0 {
		for i := 0; i < numberOfUsers; i++ {
			res[i] += base
		}
	}

	return res
}