package main
import ("fmt";"os";"strconv";"time")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: date-calc <days> [from_date]"); os.Exit(1) }
	days, err := strconv.Atoi(os.Args[1])
	if err != nil { fmt.Fprintln(os.Stderr,"Invalid number:", os.Args[1]); os.Exit(1) }
	from := time.Now()
	if len(os.Args) > 2 {
		from, err = time.Parse("2006-01-02", os.Args[2])
		if err != nil { fmt.Fprintln(os.Stderr,"Invalid date:", os.Args[2]); os.Exit(1) }
	}
	result := from.AddDate(0, 0, days)
	fmt.Println(result.Format("2006-01-02"))
}
