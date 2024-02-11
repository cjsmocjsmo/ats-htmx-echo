package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

type Template struct {
	templates *template.Template
}


func init() {
	godotenv.Load()
	dbpath := os.Getenv("ATS_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create comments table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		acctid TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		date TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt2 := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		acctid TEXT NOT NULL,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		rating TEXT NOT NULL,
		comment TEXT NOT NULL,
		date TEXT NOT NULL,
		media TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt2)
	if err != nil {
		panic(err)
	}
	sqlStmt3 := `
	CREATE TABLE IF NOT EXISTS estimates (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		acctid TEXT NOT NULL,
		estid TEXT NOT NULL,
		name TEXT NOT NULL,
		address TEXT NOT NULL,
		city TEXT NOT NULL,
		phone TEXT NOT NULL,
		email TEXT NOT NULL,
		servdate TEXT NOT NULL,
		comment TEXT NOT NULL,
		media TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt3)
	if err != nil {
		panic(err)
	}

}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	t := &Template{
		templates: template.Must(template.ParseGlob("AtsTemplates/*")),
	}
	e.Renderer = t

	e.GET("/", ats_index)
	e.GET("/about", ats_about)
	e.GET("/comments", ats_comments)
	e.GET("/estimates", ats_estimates)
	e.GET("/images", ats_images)
	e.GET("/services", ats_services)
	e.GET("/videos", ats_videos)
	e.GET("/port1", ats_port1)
	e.GET("/port2", ats_port2)
	e.GET("/port3", ats_port3)
	e.GET("/port4", ats_port4)
	e.GET("/port5", ats_port5)
	e.GET("/port6", ats_port6)
	e.GET("/port7", ats_port7)
	e.GET("/port8", ats_port8)
	e.GET("/port9", ats_port9)
	e.GET("/port10", ats_port10)
	e.POST("/comupload", com_upload)
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":8181"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ats_index(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_index", "WORKED")
}

func ats_about(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_about", "WORKED")
}

func ats_comments(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_comments", "WORKED")
}

func ats_estimates(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_estimates", "WORKED")
}

func ats_images(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_images", "WORKED")
}

func ats_services(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_services", "WORKED")
}

func ats_videos(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_videos", "WORKED")
}

func ats_port1(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port1", "WORKED")
}

func ats_port2(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port2", "WORKED")
}

func ats_port3(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port3", "WORKED")
}

func ats_port4(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port4", "WORKED")
}

func ats_port5(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port5", "WORKED")
}

func ats_port6(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port6", "WORKED")
}

func ats_port7(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port7", "WORKED")
}

func ats_port8(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port8", "WORKED")
}

func ats_port9(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port9", "WORKED")
}

func ats_port10(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port10", "WORKED")
}

func com_upload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	rating := c.FormValue("rating")
	comment := c.FormValue("comment")
	areInputsValid := checkInputs(name, email, rating, comment)
	if !areInputsValid {
		return c.Render(http.StatusOK, "ats_rejected", "WORKED")
	}

	hasAccount := accountCheck(email)
	var acctid string
	if !hasAccount {
		acctid = createAccount(email)
	} else {
		acctid = acountInfoByEmail(email)
	}

	file, err := c.FormFile("filepicker")
	if err != nil {
		println("filepicker error: ")
	}

	comid := atsUUID()

	media, err := save_file(comid, file)
	if err != nil {
		println("save_file error: ")
	}

	today := todaysDate()

	dbPath := os.Getenv("ATS_DB_PATH")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO comments (acctid, comid, name, email, rating, comment, date, media) VALUES (?, ?, ?, ?, ?, ?, datetime('now'), ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(acctid, comid, name, email, rating, comment, today, media)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "ats_thanks", "WORKED")
}

func atsUUID() string {
	comid, err := uuid.NewRandom()
	if err != nil {
		println("Unbable to create UUID")
	}
	return comid.String()
}

func todaysDate() string {
	t := time.Now()
	return t.Format("12-12-2006")
}

func nameCheck(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func emailCheck(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ratingCheck(rating string) bool {
	i, err := strconv.Atoi(rating)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return i >= 0 || i <= 5
}

func commentCheck(comment string) bool {
	bad_word_list := badwords()
	regex := regexp.MustCompile(`\w+`)
	words := regex.FindAllString(comment, -1)
	for _, word := range words {
		for _, bad_word := range bad_word_list {
			if word == bad_word {
				return true
			}
		}
	}
	return false
}

func checkInputs(name string, email string, rating string, comment string) bool {
	isValidName := nameCheck(name)
	isValidEmail := emailCheck(email)
	hasAccount := accountCheck(email)
	if !hasAccount {
		createAccount(email)
	}
	isValidRating := ratingCheck(rating)
	isValidComment := commentCheck(comment)
	if isValidName && isValidEmail && isValidRating && isValidComment {
		return true
	} else {
		return false
	}
}

func save_file(comid string, file *multipart.FileHeader) (string, error) {
	out_dir := os.Getenv("ATS_UPLOADS_PATH")
	out_path := out_dir + "/" + comid + "_" + file.Filename
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	dst, err := os.Create(out_path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return out_path, nil
}

func accountCheck(email string) bool {
	dbPath := os.Getenv("ATS_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT email FROM accounts WHERE email = ?", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

type AccountInfo struct {
	Acctid string
	Email  string
	Date   string
}

func acountInfoByEmail(email string) string {
	dbPath := os.Getenv("ATS_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ?", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//place the values into AccountInfo and return it
	var ai AccountInfo
	for rows.Next() {
		
		err = rows.Scan(&ai.Acctid, &ai.Email, &ai.Date)
		if err != nil {
			panic(err)
		}
		
	}
	return ai.Acctid
}

// func accountInfoByID(acctid string) AccountInfo {
// 	dbPath := os.Getenv("ATS_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM accounts WHERE acctid = ?", acctid)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var ai AccountInfo
// 	for rows.Next() {
// 		err = rows.Scan(&ai.Acctid, &ai.Email, &ai.Date)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return ai
// }

func createAccount(email string) string {
	acctid := atsUUID()
	date := todaysDate()
	dbPath := os.Getenv("ATS_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO accounts (acctid, email, date) VALUES (?, ?, ?)", acctid, email, date)
	if err != nil {
		panic(err)
	}
	return acctid
}

func badwords() []string {
	var badwords []string
	badwords = append(badwords, "abbo",
		"alla",
		"allah",
		"alligatorbait",
		"anal",
		"analannie",
		"analsex",
		"angry",
		"anus",
		"areola",
		"argie",
		"aroused",
		"arse",
		"arsehole",
		"ass",
		"assassin",
		"assassinate",
		"assassination",
		"assault",
		"assbagger",
		"assblaster",
		"assclown",
		"asscowboy",
		"asses",
		"assfuck",
		"assfucker",
		"asshat",
		"asshole",
		"assholes",
		"asshore",
		"assjockey",
		"asskiss",
		"asskisser",
		"assklown",
		"asslick",
		"asslicker",
		"asslover",
		"assman",
		"assmonkey",
		"assmunch",
		"assmuncher",
		"asspacker",
		"asspirate",
		"asspuppies",
		"assranger",
		"asswhore",
		"asswipe",
		"attack",
		"backdoor",
		"backdoorman",
		"backseat",
		"badfuck",
		"balllicker",
		"balls",
		"ballsack",
		"banging",
		"barelylegal",
		"bastard ",
		"bazongas",
		"bazooms",
		"beaner",
		"beastality",
		"beastial",
		"beastiality",
		"beatoff",
		"beat-off",
		"beatyourmeat",
		"beaver",
		"bestial",
		"bestiality",
		"bi",
		"biatch",
		"bicurious",
		"bigass",
		"bigbastard",
		"bigbutt",
		"bisexual",
		"bi-sexual",
		"bitch",
		"bitcher",
		"bitches",
		"bitchez",
		"bitchin",
		"bitching",
		"bitchslap",
		"bitchy",
		"biteme",
		"blackman",
		"blackout",
		"blacks",
		"blow",
		"blowjob",
		"boang",
		"bogan",
		"bohunk",
		"bollick",
		"bollock",
		"bomb",
		"bombers",
		"bombing",
		"bombs",
		"bomd",
		"bondage",
		"boner",
		"boob",
		"boobies",
		"boobs",
		"booby",
		"boody",
		"boom",
		"boong",
		"boonga",
		"boonie",
		"booty",
		"bootycall",
		"bountybar",
		"bra",
		"brea5t",
		"breast",
		"breastjob",
		"breastlover",
		"breastman",
		"brothel",
		"bullcrap",
		"bulldike",
		"bulldyke",
		"bullshit",
		"bumblefuck",
		"bumfuck",
		"bunga",
		"bunghole",
		"butchbabes",
		"butchdike",
		"butchdyke",
		"butt",
		"buttbang",
		"butt-bang",
		"buttface",
		"buttfuck",
		"butt-fuck",
		"buttfucker",
		"butt-fucker",
		"buttfuckers",
		"butt-fuckers",
		"butthead",
		"buttman",
		"buttmunch",
		"buttmuncher",
		"buttpirate",
		"buttplug",
		"buttstain",
		"byatch",
		"cameljockey",
		"cameltoe",
		"carpetmuncher",
		"chav",
		"cherrypopper",
		"chickslick",
		"chink",
		"chinky",
		"clamdigger",
		"clamdiver",
		"clit",
		"clitoris",
		"clogwog",
		"cock",
		"cockblock",
		"cockblocker",
		"cockcowboy",
		"cockfight",
		"cockhead",
		"cockknob",
		"cocklicker",
		"cocklover",
		"cocknob",
		"cockqueen",
		"cockrider",
		"cocksman",
		"cocksmith",
		"cocksmoker",
		"cocksucer",
		"cocksuck ",
		"cocksucked ",
		"cocksucker",
		"cocksucking",
		"cocktail",
		"cocktease",
		"cocky",
		"cohee",
		"coitus",
		"commie",
		"communist",
		"condom",
		"conspiracy",
		"coolie",
		"cooly",
		"coon",
		"coondog",
		"copulate",
		"cornhole",
		"crabs",
		"crack",
		"crackpipe",
		"crackwhore",
		"crack-whore",
		"crap",
		"crapola",
		"crapper",
		"crotch",
		"crotchjockey",
		"crotchmonkey",
		"crotchrot",
		"cum",
		"cumbubble",
		"cumfest",
		"cumjockey",
		"cumm",
		"cummer",
		"cumming",
		"cumquat",
		"cumqueen",
		"cumshot",
		"cunilingus",
		"cunillingus",
		"cunn",
		"cunnilingus",
		"cunntt",
		"cunt",
		"cunteyed",
		"cuntfuck",
		"cuntfucker",
		"cuntlick ",
		"cuntlicker ",
		"cuntlicking ",
		"cuntsucker",
		"cybersex",
		"dahmer",
		"dammit",
		"damn",
		"damnation",
		"damnit",
		"darkie",
		"darky",
		"datnigga",
		"dead",
		"deapthroat",
		"deepthroat",
		"defecate",
		"demon",
		"deth",
		"devil",
		"devilworshippe",
		"dick",
		"dickbrain",
		"dickforbrains",
		"dickhead",
		"dickless",
		"dicklick",
		"dicklicker",
		"dickman",
		"dickwad",
		"dickweed",
		"diddle",
		"dike",
		"dildo",
		"dingleberry",
		"dipshit",
		"dipstick",
		"dix",
		"dixiedike",
		"dixiedyke",
		"doggiestyle",
		"doggystyle",
		"dong",
		"doodoo",
		"doo-doo",
		"doom",
		"dragqueen",
		"dragqween",
		"dripdick",
		"dumbass",
		"dumbbitch",
		"dumbfuck",
		"dyke",
		"easyslut",
		"eatballs",
		"eatme",
		"eatpussy",
		"ejaculate",
		"ejaculated",
		"ejaculating ",
		"ejaculation",
		"enema",
		"enemy",
		"erection",
		"excrement",
		"executioner",
		"explosion",
		"facefucker",
		"faeces",
		"fag",
		"fagging",
		"faggot",
		"fagot",
		"fannyfucker",
		"fart",
		"farted ",
		"farting ",
		"farty ",
		"fastfuck",
		"fat",
		"fatah",
		"fatass",
		"fatfuck",
		"fatfucker",
		"fatso",
		"fckcum",
		"feces",
		"felatio ",
		"felch",
		"felcher",
		"felching",
		"fellatio",
		"feltch",
		"feltcher",
		"feltching",
		"fetish",
		"fingerfood",
		"fingerfuck ",
		"fingerfucked ",
		"fingerfucker ",
		"fingerfuckers",
		"fingerfucking ",
		"fister",
		"fistfuck",
		"fistfucked ",
		"fistfucker ",
		"fistfucking ",
		"fisting",
		"flange",
		"flasher",
		"flatulence",
		"fondle",
		"footaction",
		"footfuck",
		"footfucker",
		"footlicker",
		"footstar",
		"foreskin",
		"forni",
		"fornicate",
		"foursome",
		"fourtwenty",
		"freakfuck",
		"freakyfucker",
		"freefuck",
		"fuc",
		"fucck",
		"fuck",
		"fucka",
		"fuckable",
		"fuckbag",
		"fuckbuddy",
		"fucked",
		"fuckedup",
		"fucker",
		"fuckers",
		"fuckface",
		"fuckfest",
		"fuckfreak",
		"fuckfriend",
		"fuckhead",
		"fuckher",
		"fuckin",
		"fuckina",
		"fucking",
		"fuckingbitch",
		"fuckinnuts",
		"fuckinright",
		"fuckit",
		"fuckknob",
		"fuckme ",
		"fuckmehard",
		"fuckmonkey",
		"fuckoff",
		"fuckpig",
		"fucks",
		"fucktard",
		"fuckwhore",
		"fuckyou",
		"fudgepacker",
		"fugly",
		"fuk",
		"fuks",
		"funfuck",
		"fuuck",
		"gangbang",
		"gangbanged ",
		"gangbanger",
		"gangsta",
		"gatorbait",
		"gay",
		"gaymuthafuckin",
		"gaysex ",
		"genital",
		"german",
		"getiton",
		"ginzo",
		"gipp",
		"girls",
		"givehead",
		"glazeddonut",
		"gob",
		"god",
		"godammit",
		"goddamit",
		"goddammit",
		"goddamn",
		"goddamned",
		"goddamnes",
		"goddamnit",
		"goddamnmuthafu",
		"goldenshower",
		"gonorrehea",
		"gonzagas",
		"gook",
		"gotohell",
		"greaseball",
		"gringo",
		"grostulation",
		"gun",
		"gyp",
		"gypo",
		"gypp",
		"gyppie",
		"gyppo",
		"gyppy",
		"hamas",
		"handjob",
		"hapa",
		"harder",
		"hardon",
		"harem",
		"headfuck",
		"headlights",
		"hebe",
		"heeb",
		"hell",
		"henhouse",
		"heroin",
		"herpes",
		"hijack",
		"hijacker",
		"hijacking",
		"hillbillies",
		"hiscock",
		"hitler",
		"hitlerism",
		"hitlerist",
		"hiv",
		"ho",
		"hobo",
		"hodgie",
		"hoes",
		"hole",
		"holestuffer",
		"homo",
		"homobangers",
		"homosexual",
		"honger",
		"honk",
		"honkers",
		"honkey",
		"honky",
		"hook",
		"hooker",
		"hookers",
		"hooters",
		"hore",
		"hork",
		"horn",
		"horney",
		"horniest",
		"horny",
		"horseshit",
		"hosejob",
		"hoser",
		"hostage",
		"hotdamn",
		"hotpussy",
		"hottotrot",
		"hummer",
		"hussy",
		"hustler",
		"hymen",
		"hymie",
		"iblowu",
		"idiot",
		"ikey",
		"incest",
		"insest",
		"intercourse",
		"interracial",
		"intheass",
		"inthebuff",
		"itch",
		"jackass",
		"jackoff",
		"jackshit",
		"jacktheripper",
		"jade",
		"jap",
		"japcrap",
		"jebus",
		"jeez",
		"jerkoff",
		"jesus",
		"jesuschrist",
		"jiga",
		"jigaboo",
		"jigg",
		"jigga",
		"jiggabo",
		"jigger ",
		"jiggy",
		"jihad",
		"jijjiboo",
		"jimfish",
		"jism",
		"jiz ",
		"jizim",
		"jizjuice",
		"jizm ",
		"jizz",
		"jizzim",
		"jizzum",
		"juggalo",
		"jugs",
		"junglebunny",
		"kaffer",
		"kaffir",
		"kaffre",
		"kafir",
		"kanake",
		"kid",
		"kigger",
		"kike",
		"kink",
		"kinky",
		"kissass",
		"kkk",
		"knife",
		"knockers",
		"kock",
		"kondum",
		"koon",
		"kotex",
		"krap",
		"krappy",
		"kraut",
		"kum",
		"kumbubble",
		"kumbullbe",
		"kummer",
		"kumming",
		"kumquat",
		"kums",
		"kunilingus",
		"kunnilingus",
		"kunt",
		"ky",
		"kyke",
		"lactate",
		"laid",
		"lapdance",
		"latin",
		"lesbain",
		"lesbayn",
		"lesbian",
		"lesbin",
		"lesbo",
		"lez",
		"lezbe",
		"lezbefriends",
		"lezbo",
		"lezz",
		"lezzo",
		"libido",
		"licker",
		"lickme",
		"lies",
		"limey",
		"limpdick",
		"limy",
		"lingerie",
		"liquor",
		"livesex",
		"loadedgun",
		"lolita",
		"lovebone",
		"lovegoo",
		"lovegun",
		"lovejuice",
		"lovemuscle",
		"lovepistol",
		"loverocket",
		"lowlife",
		"lubejob",
		"lucifer",
		"luckycammeltoe",
		"magicwand",
		"mams",
		"manhater",
		"manpaste",
		"marijuana",
		"mastabate",
		"mastabater",
		"masterbate",
		"masterblaster",
		"mastrabator",
		"masturbate",
		"masturbating",
		"mattressprince",
		"meatbeatter",
		"meatrack",
		"mickeyfinn",
		"mideast",
		"milf",
		"minority",
		"mockey",
		"mockie",
		"mocky",
		"mofo",
		"moky",
		"moles",
		"molest",
		"molestation",
		"molester",
		"molestor",
		"moneyshot",
		"mooncricket",
		"moron",
		"moslem",
		"mosshead",
		"mothafuck",
		"mothafucka",
		"mothafuckaz",
		"mothafucked ",
		"mothafucker",
		"mothafuckin",
		"mothafucking ",
		"mothafuckings",
		"motherfuck",
		"motherfucked",
		"motherfucker",
		"motherfuckin",
		"motherfucking",
		"motherfuckings",
		"motherlovebone",
		"muff",
		"muffdive",
		"muffdiver",
		"muffindiver",
		"mufflikcer",
		"naked",
		"nastybitch",
		"nastyho",
		"nastyslut",
		"nastywhore",
		"nazi",
		"necro",
		"negro",
		"negroes",
		"negroid",
		"negro's",
		"nig",
		"niger",
		"nigerian",
		"nigerians",
		"nigg",
		"nigga",
		"niggah",
		"niggaracci",
		"niggard",
		"niggarded",
		"niggarding",
		"niggardliness",
		"niggardliness'",
		"niggardly",
		"niggards",
		"niggard's",
		"niggaz",
		"nigger",
		"niggerhead",
		"niggerhole",
		"niggers",
		"nigger's",
		"niggle",
		"niggled",
		"niggles",
		"niggling",
		"nigglings",
		"niggor",
		"niggur",
		"niglet",
		"nignog",
		"nigr",
		"nigra",
		"nigre",
		"nip",
		"nipple",
		"nipplering",
		"nittit",
		"nlgger",
		"nlggor",
		"nofuckingway",
		"nook",
		"nookey",
		"nookie",
		"noonan",
		"nooner",
		"nude",
		"nudger",
		"nuke",
		"nutfucker",
		"nymph",
		"ontherag",
		"oral",
		"orga",
		"orgasim ",
		"orgasm",
		"orgies",
		"orgy",
		"panti",
		"panties",
		"payo",
		"pearlnecklace",
		"peck",
		"pecker",
		"peckerwood",
		"pee",
		"peehole",
		"pee-pee",
		"peepshow",
		"peepshpw",
		"pendy",
		"penetration",
		"peni5",
		"penile",
		"penis",
		"penises",
		"penthouse",
		"period",
		"perv",
		"phonesex",
		"phuk",
		"phuked",
		"phuking",
		"phukked",
		"phukking",
		"phungky",
		"phuq",
		"pi55",
		"picaninny",
		"piccaninny",
		"pickaninny",
		"piker",
		"pikey",
		"piky",
		"pimp",
		"pimped",
		"pimper",
		"pimpjuic",
		"pimpjuice",
		"pimpsimp",
		"pindick",
		"piss",
		"pissed",
		"pisser",
		"pisses ",
		"pisshead",
		"pissin ",
		"pissing",
		"pissoff ",
		"pistol",
		"pixie",
		"pixy",
		"playboy",
		"playgirl",
		"pocha",
		"pocho",
		"pocketpool",
		"pohm",
		"polack",
		"pom",
		"pommie",
		"pommy",
		"poo",
		"poon",
		"poontang",
		"poop",
		"pooper",
		"pooperscooper",
		"pooping",
		"poorwhitetrash",
		"popimp",
		"porchmonkey",
		"porn",
		"pornflick",
		"pornking",
		"porno",
		"pornography",
		"pornprincess",
		"pot",
		"poverty",
		"premature",
		"pric",
		"prick",
		"prickhead",
		"primetime",
		"propaganda",
		"pros",
		"prostitute",
		"protestant",
		"pu55i",
		"pu55y",
		"pube",
		"pubic",
		"pubiclice",
		"pud",
		"pudboy",
		"pudd",
		"puddboy",
		"puke",
		"puntang",
		"purinapricness",
		"puss",
		"pussie",
		"pussies",
		"pussy",
		"pussycat",
		"pussyeater",
		"pussyfucker",
		"pussylicker",
		"pussylips",
		"pussylover",
		"pussypounder",
		"pusy",
		"quashie",
		"queef",
		"queer",
		"quickie",
		"quim",
		"racial",
		"racist",
		"radical",
		"radicals",
		"raghead",
		"randy",
		"rape",
		"raped",
		"raper",
		"rapist",
		"rearend",
		"rearentry",
		"rectum",
		"redlight",
		"redneck",
		"reefer",
		"reestie",
		"refugee",
		"reject",
		"remains",
		"rentafuck",
		"republican",
		"rere",
		"retard",
		"retarded",
		"ribbed",
		"rigger",
		"rimjob",
		"rimming",
		"roach",
		"robber",
		"roundeye",
		"rump",
		"russki",
		"russkie",
		"sadis",
		"sadom",
		"samckdaddy",
		"sandm",
		"sandnigger",
		"satan",
		"scag",
		"scallywag",
		"scat",
		"schlong",
		"screw",
		"screwyou",
		"scrotum",
		"scum",
		"semen",
		"seppo",
		"servant",
		"sex",
		"sexed",
		"sexfarm",
		"sexhound",
		"sexhouse",
		"sexing",
		"sexkitten",
		"sexpot",
		"sexslave",
		"sextogo",
		"sextoy",
		"sextoys",
		"sexual",
		"sexually",
		"sexwhore",
		"sexy",
		"sexymoma",
		"sexy-slim",
		"shag",
		"shaggin",
		"shagging",
		"shat",
		"shav",
		"shawtypimp",
		"sheeney",
		"shhit",
		"shinola",
		"shit",
		"shitcan",
		"shitdick",
		"shite",
		"shiteater",
		"shited",
		"shitface",
		"shitfaced",
		"shitfit",
		"shitforbrains",
		"shitfuck",
		"shitfucker",
		"shitfull",
		"shithapens",
		"shithappens",
		"shithead",
		"shithouse",
		"shiting",
		"shitlist",
		"shitola",
		"shitoutofluck",
		"shits",
		"shitstain",
		"shitted",
		"shitter",
		"shitting",
		"shitty ",
		"shoot",
		"shooting",
		"shortfuck",
		"showtime",
		"sick",
		"sissy",
		"sixsixsix",
		"sixtynine",
		"sixtyniner",
		"skank",
		"skankbitch",
		"skankfuck",
		"skankwhore",
		"skanky",
		"skankybitch",
		"skankywhore",
		"skinflute",
		"skum",
		"skumbag",
		"slant",
		"slanteye",
		"slapper",
		"slaughter",
		"slav",
		"slave",
		"slavedriver",
		"sleezebag",
		"sleezeball",
		"slideitin",
		"slime",
		"slimeball",
		"slimebucket",
		"slopehead",
		"slopey",
		"slopy",
		"slut",
		"sluts",
		"slutt",
		"slutting",
		"slutty",
		"slutwear",
		"slutwhore",
		"smack",
		"smackthemonkey",
		"smut",
		"snatch",
		"snatchpatch",
		"snigger",
		"sniggered",
		"sniggering",
		"sniggers",
		"snigger's",
		"sniper",
		"snot",
		"snowback",
		"snownigger",
		"sob",
		"sodom",
		"sodomise",
		"sodomite",
		"sodomize",
		"sodomy",
		"sonofabitch",
		"sonofbitch",
		"sooty",
		"sos",
		"soviet",
		"spaghettibende",
		"spaghettinigge",
		"spank",
		"spankthemonkey",
		"sperm",
		"spermacide",
		"spermbag",
		"spermhearder",
		"spermherder",
		"spic",
		"spick",
		"spig",
		"spigotty",
		"spik",
		"spit",
		"spitter",
		"splittail",
		"spooge",
		"spreadeagle",
		"spunk",
		"spunky",
		"squaw",
		"stagg",
		"stiffy",
		"strapon",
		"stringer",
		"stripclub",
		"stroke",
		"stroking",
		"stupid",
		"stupidfuck",
		"stupidfucker",
		"suck",
		"suckdick",
		"sucker",
		"suckme",
		"suckmyass",
		"suckmydick",
		"suckmytit",
		"suckoff",
		"suicide",
		"swallow",
		"swallower",
		"swalow",
		"swastika",
		"sweetness",
		"syphilis",
		"taboo",
		"taff",
		"tampon",
		"tang",
		"tantra",
		"tarbaby",
		"tard",
		"teste",
		"testicle",
		"testicles",
		"thicklips",
		"thirdeye",
		"thirdleg",
		"threesome",
		"threeway",
		"timbernigger",
		"tinkle",
		"tit",
		"titbitnipply",
		"titfuck",
		"titfucker",
		"titfuckin",
		"titjob",
		"titlicker",
		"titlover",
		"tits",
		"tittie",
		"titties",
		"titty",
		"tnt",
		"toilet",
		"tongethruster",
		"tongue",
		"tonguethrust",
		"tonguetramp",
		"tortur",
		"torture",
		"tosser",
		"towelhead",
		"trailertrash",
		"tramp",
		"trannie",
		"tranny",
		"transexual",
		"transsexual",
		"transvestite",
		"triplex",
		"trisexual",
		"trojan",
		"trots",
		"tuckahoe",
		"tunneloflove",
		"turd",
		"turnon",
		"twat",
		"twink",
		"twinkie",
		"twobitwhore",
		"uck",
		"uk",
		"unfuckable",
		"upskirt",
		"uptheass",
		"upthebutt",
		"urinary",
		"urinate",
		"urine",
		"usama",
		"uterus",
		"vagina",
		"vaginal",
		"vatican",
		"vibr",
		"vibrater",
		"vibrator",
		"vietcong",
		"violence",
		"virgin",
		"virginbreaker",
		"vomit",
		"vulva",
		"wab",
		"wank",
		"wanker",
		"wanking",
		"waysted",
		"weapon",
		"weenie",
		"weewee",
		"welcher",
		"welfare",
		"wetb",
		"wetback",
		"wetspot",
		"whacker",
		"whash",
		"whigger",
		"whiskey",
		"whiskeydick",
		"whiskydick",
		"whit",
		"whitenigger",
		"whites",
		"whitetrash",
		"whitey",
		"whiz",
		"whop",
		"whore",
		"whorefucker",
		"whorehouse",
		"wigger",
		"willie",
		"williewanker",
		"willy",
		"wn",
		"wog",
		"women's",
		"wop",
		"wtf",
		"wuss",
		"wuzzie",
		"xtc",
		"xxx",
		"yankee",
		"yellowman",
		"zigabo",
		"zipperhead",
	)
	return badwords
}
