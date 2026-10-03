package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Content struct {
	Badge       string
	Title       string
	Lead        string
	Stat1Val    string
	Stat1Lbl    string
	Stat2Val    string
	Stat2Lbl    string
	Stat3Val    string
	Stat3Lbl    string
	AppleTag    string
	AppleTitle  string
	AppleText1  string
	AppleText2  string
	AppleBtn    string
	GoogleTag   string
	GoogleTitle string
	GoogleText1 string
	GoogleText2 string
	GoogleBtn   string
	MSTag       string
	MSTitle     string
	MSText1     string
	MSText2     string
	MSBtn       string
}

type PageData struct {
	Lang       string
	AltLangURL string
	AltLangText string
	MetaTitle  string
	MetaDesc   string
	Content    Content
}

var ruContent = Content{
	Badge:       "АНАЛИТИЧЕСКИЙ РАЗБОР ИМПЕРИЙ",
	Title:       "Корпорации Зла и Добра: Кто На Самом Деле Владеет Миром",
	Lead:        "Добро пожаловать в суровую реальность кремниевого капитализма. Здесь мы без соплей и корпоративной цензуры разбираем три главных монополии современности.",
	Stat1Val:    "$10T+", Stat1Lbl: "Суммарная капитализация",
	Stat2Val:    "3 млрд", Stat2Lbl: "Зомбированных юзеров",
	Stat3Val:    "0%", Stat3Lbl: "Совести у руководства",
	AppleTag:    "КУПЕРТИНОВСКАЯ СЕКТА",
	AppleTitle:  "Яблодрочеры: Религия Закрытого Кода",
	AppleText1:  "Яблочная контора построила гениальную бизнес-модель: продавать людям алюминий по цене тачки.",
	AppleText2:  "Закрытая экосистема работает как мышеловка: зайти легко, а выйти — хрен вам.",
	AppleBtn:    "Читать полное досье на яблодрочеров",
	GoogleTag:   "ИМПЕРИЯ АЛГОРИТМОВ",
	GoogleTitle: "Корпорация Добра: «Не Будь Злом» (Но Мы Уже)",
	GoogleText1: "Эти парни превратились в цифрового Левиафана, который знает о тебе больше, чем твоя мама.",
	GoogleText2: "Поисковая выдача забита рекламой, а Android высасывает телеметрию на раз-два.",
	GoogleBtn:   "Узнать всю правду про корпорацию добра",
	MSTag:       "РЕДМОНДСКИЙ ВЕТЕРАН",
	MSTitle:     "Мелкомягкие: Синдикат Синего Экрана Смерти",
	MSText1:     "Операционная система Windows — это памятник ленивому коду и костылям.",
	MSText2:     "Каждое обновление десятки — это русская рулетка для твоего ПК.",
	MSBtn:       "Изучить грехи мелкомягких",
}

var enContent = Content{
	Badge:       "EMPIRE ANALYTICS BREAKDOWN",
	Title:       "Tech Monopolies: Who Actually Rules The World",
	Lead:        "Welcome to the harsh reality of Silicon Valley capitalism. No fluff, no corporate censorship — just pure facts about the big three.",
	Stat1Val:    "$10T+", Stat1Lbl: "Total Market Cap",
	Stat2Val:    "3B", Stat2Lbl: "Zombified Users",
	Stat3Val:    "0%", Stat3Lbl: "Executive Conscience",
	AppleTag:    "CUPERTINO CULT",
	AppleTitle:  "Apple: The Religion of Closed Source",
	AppleText1:  "The fruit company built a genius business model: selling aluminum for the price of a used car.",
	AppleText2:  "Their closed ecosystem is a roach motel — easy to check in, impossible to leave.",
	AppleBtn:    "Read full Apple dossier",
	GoogleTag:   "ALGORITHM EMPIRE",
	GoogleTitle: "Google: Don't Be Evil (Too Late)",
	GoogleText1: "These guys turned into a digital Leviathan that knows you better than your own mother.",
	GoogleText2: "Search is a landfill of ads, and Android harvests your telemetry 24/7.",
	GoogleBtn:   "Uncover the truth about Google",
	MSTag:       "REDMOND VETERAN",
	MSTitle:     "Microsoft: Syndicate of the Blue Screen",
	MSText1:     "Windows is a monumental monument to lazy code and infinite workarounds.",
	MSText2:     "Every new update is Russian roulette for your hardware.",
	MSBtn:       "Explore Microsoft sins",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/en", handleHomeEN)

	log.Printf("Сервер запущен на порту %s... Погнали!", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplPath := filepath.Join("templates", tmpl)
	layoutPath := filepath.Join("templates", "layout.html")

	t, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Template error", 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.ExecuteTemplate(w, "layout", data)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := PageData{
		Lang:        "ru",
		AltLangURL:  "/en",
		AltLangText: "EN",
		MetaTitle:   "Корпорации Зла и Добра: Яблодрочеры, Google и Microsoft",
		MetaDesc:    "Полный аналитический разбор главных техно-монополий мира.",
		Content:     ruContent,
	}
	renderTemplate(w, "index.html", data)
}

func handleHomeEN(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Lang:        "en",
		AltLangURL:  "/",
		AltLangText: "РУ",
		MetaTitle:   "Tech Monopolies: Apple, Google and Microsoft Exposed",
		MetaDesc:    "Deep dive analysis into the biggest tech monopolies.",
		Content:     enContent,
	}
	renderTemplate(w, "index.html", data)
}