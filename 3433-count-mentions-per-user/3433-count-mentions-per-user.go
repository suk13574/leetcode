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
			return events[i][0] > events[j][0]
		}
		ie, _ := strconv.Atoi(events[i][1])
		je, _ := strconv.Atoi(events[j][1])
		return ie < je
	})

	userStatus := make([]*UserStatus, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		userStatus[i] = &UserStatus{isOnline: true}
	}

	res := make([]int, numberOfUsers)
	base := 0
	for _, event := range events {
		now, _ := strconv.Atoi(event[1])
		users := strings.Split(event[2], " ")

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
						} else {
							if us.endTime <= now {
								res[i]++
								us.isOnline = true
							}
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
