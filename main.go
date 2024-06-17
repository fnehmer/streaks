/**
Plan:
leap year: <year> % 4 == 0

2024:
___________________________________________________________
Mo|  ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢ |
Tu| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢ |
We| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢ |
Th| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢  |
Fr| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢  |
Sa| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢  |
Su| ▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢▢  |
‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
▢: No log
▥: Maybe something in between
▩: Log

commands:
- streaks show          // Display everything
- streaks show <title>  // Display <title>
- streaks log <title>   // Logs current date in <title>
- streaks log <title>   // <date:YYYYMMDD> // Logs specified date in <title>
**/

package main

import "fmt"
import "time"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
    var Reset string = "\033[0m"
//    var Red string = "\033[31m" 
//    var Green string = "\033[32m" 
//    var Yellow string = "\033[33m" 
    var Blue string = "\033[34m" 
//    var Magenta string = "\033[35m" 
//    var Cyan string = "\033[36m" 
//    var Gray string = "\033[37m" 
//    var White string = "\033[97m"

    fmt.Println(Blue + "streaks v0" + Reset)

    var somedate time.Time
    somedate = somedate.AddDate(2023, 0, 0)
    weekday := time.Now().Weekday()

    fmt.Println(weekday)
    fmt.Println(somedate)
}

