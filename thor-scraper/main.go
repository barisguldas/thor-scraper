package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
)


var reportLog *os.File

func main() {
	
	var err error
	reportLog, err = os.OpenFile("scan_report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer reportLog.Close()

	
	printBanner()

	// --- ANA MENÜ DÖNGÜSÜ ---
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(color.HiCyanString("\n┌──(THOR CTI MENU)"))
		fmt.Println(color.HiWhiteString("│"))
		fmt.Println(color.HiWhiteString("├── [1] ") + color.HiGreenString("TARAMAYI BAŞLAT") + color.HiBlackString(" (Start Scan)"))
		fmt.Println(color.HiWhiteString("├── [2] ") + color.HiYellowString("HEDEFLERİ GÖRÜNTÜLE") + color.HiBlackString(" (View Targets)"))
		fmt.Println(color.HiWhiteString("├── [3] ") + color.HiRedString("ÇIKIŞ") + color.HiBlackString(" (Exit)"))
		fmt.Println(color.HiWhiteString("│"))
		fmt.Print(color.HiCyanString("└──> Seçiminiz: "))

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			startScraping()
		case "2":
			showTargets()
		case "3":
			fmt.Println(color.HiRedString("\n[!] Sistemden Çıkış Yapılıyor..."))
			time.Sleep(1 * time.Second)
			os.Exit(0)
		default:
			fmt.Println(color.HiRedString("\n[HATALI GİRİŞ] Lütfen 1, 2 veya 3 seçeneğini kullanın."))
		}
	}
}

// --- FONKSİYON 1: TARAMA İŞLEMİ ---
func startScraping() {
	torProxy := "socks5://127.0.0.1:9150"

	fmt.Println()
	fmt.Println(color.HiYellowString("[*] Tor Bağlantısı Kuruluyor..."))

	// Chrome Ayarlarını Yapılandır
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer(torProxy),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// targets.yaml Dosyasını Oku
	file, err := os.Open("targets.yaml")
	if err != nil {
		fmt.Println(color.HiRedString("[!] targets.yaml dosyası bulunamadı!"))
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fmt.Println(color.HiGreenString("[+] Tarama Başlatıldı! (Lütfen Bekleyin)"))
	fmt.Println(color.HiWhiteString("------------------------------------------------"))

	// Her satır için döngü
	count := 0
	for fileScanner.Scan() {
		url := strings.TrimSpace(fileScanner.Text())
		if url == "" {
			continue
		}
		count++

		ctx, cancel := chromedp.NewContext(allocCtx)
		ctx, cancel = context.WithTimeout(ctx, 60*time.Second) // 60sn zaman aşımı

		scrapeOnionSite(ctx, url)
		cancel()
	}

	fmt.Println(color.HiWhiteString("------------------------------------------------"))
	fmt.Printf(color.HiGreenString("[OK] Tarama Tamamlandı. Toplam %d hedef denendi.\n"), count)
}


func scrapeOnionSite(ctx context.Context, url string) {
	fmt.Printf(color.HiCyanString(" -> Taranıyor: %s ... "), url)

	var htmlContent string
	var screenshot []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second), 
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.FullScreenshot(&screenshot, 90),
	)

	if err != nil {
		fmt.Println(color.HiRedString("BAŞARISIZ (Timeout)"))
		logToReport(url, "FAIL - "+err.Error())
		return
	}

	// Dosya Kaydetme İşlemleri
	safeName := strings.ReplaceAll(url, "https://", "")
	safeName = strings.ReplaceAll(safeName, "http://", "")
	safeName = strings.ReplaceAll(safeName, ".onion", "")
	safeName = strings.ReplaceAll(safeName, "/", "_")

	
	if len(safeName) > 50 {
		safeName = safeName[:50]
	}

	err1 := os.WriteFile("shot_"+safeName+".png", screenshot, 0644)
	err2 := os.WriteFile("html_"+safeName+".txt", []byte(htmlContent), 0644)

	if err1 == nil && err2 == nil {
		fmt.Println(color.HiGreenString("BAŞARILI"))
		logToReport(url, "SUCCESS - Veriler Kaydedildi")
	} else {
		fmt.Println(color.HiRedString("DOSYA YAZMA HATASI"))
	}
}

// --- FONKSİYON 3: HEDEFLERİ GÖSTERME ---
func showTargets() {
	fmt.Println(color.HiYellowString("\n--- HEDEF LİSTESİ (targets.yaml) ---"))

	file, err := os.Open("targets.yaml")
	if err != nil {
		fmt.Println(color.HiRedString("[!] targets.yaml dosyası bulunamadı!"))
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			fmt.Printf("%d. %s\n", i, line)
			i++
		}
	}
	fmt.Println(color.HiYellowString("------------------------------------"))
}

// --- LOGLAMA YARDIMCISI ---
func logToReport(url, status string) {
	logLine := fmt.Sprintf("[%s] URL: %s -> %s\n", time.Now().Format(time.RFC3339), url, status)
	if _, err := reportLog.WriteString(logLine); err != nil {
	}
}

func printBanner() {

	logoColor := color.New(color.FgHiWhite, color.Bold).SprintFunc()
	titleColor := color.New(color.FgHiRed, color.Bold).SprintFunc()
	subColor := color.New(color.FgWhite, color.Faint).SprintFunc()   
	appColor := color.New(color.FgHiYellow, color.Bold).SprintFunc() 


	fmt.Print("\033[H\033[2J")
	fmt.Println()

	
	lines := []string{
		"                                                            ",
		"                                                            ",
		"                                  ",
		"                      .:^!7JY55J~. ",
		"                  :!5#&@@@@@@G7:",
		"               :J#@@@@@@@@&?.                            ",
		"             ^G@@@@@@@@@@P.          .J~.  :7. ",
		"          .^@@@@@@@@@@@5              Y@&#@?  ",
		"         !GY@@@@@@@@@@#            .~P&@@@@B!. ",
		"        7@BY@@@@@@@@@@?                7@G^    ",
		"       ~@@@?@@@@@@@@@@7                 Y            ",
		"      .&@@@##@@@@@@@@@P                            ",
		"      !@@@@@&@@@@@@@@@@?!!77777!^:                ?! ",
		"      ~@@@@@@@@@@@@@@@@@@@@@@@@@@@5^:            J@Y ",
		"      !P@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#:        ^B@@Y ",
		"      7&B@@@@@@@@@@@@@@@@@@@@#J!^:   J      .7B@@@@! ",
		"      .&@&@@@@@@@@@@@@@@@@@&&&BY!^      ~?5#@@@@@@&. ",
		"       !@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@~ ",
		"        ?@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@7 ",
		"        !@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@&~ ",
		"          .P@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@P. ",
		"            G@@@@@@@@@@@@@@@@@@@@&&&&&&&&&#P   ",
		"              ?#@@@@@@@@@@@@@@@@@@@&&&#BY~   ",
		"                !5#&@@@&&&@@@@@@@@&#5!:   ",
		"                   ^!??7!~!??!^       ",
		"                                      ",
		"                                                            ",
		"                                                            ",
	}

	
	for i, line := range lines {

		fmt.Print(logoColor(line))

	
		if i == 12 {
			fmt.Print(titleColor("   SİBER VATAN"))
		}
		if i == 13 {
			fmt.Print(titleColor("   YILDIZ CTI"))
		}
		if i == 14 {
			fmt.Print(subColor("   ---------------"))
		}
		if i == 15 {
			fmt.Print(appColor("   THOR SCRAPER"))
		}


		fmt.Println()
		time.Sleep(5 * time.Millisecond)
	}

	fmt.Println()
	time.Sleep(1 * time.Second)
}


