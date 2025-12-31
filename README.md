#  THOR SCRAPER - Onion Network Crawler

**Thor Scraper**; Tor Ağı üzerindeki hedefleri güvenli bir şekilde tarayan, ekran görüntüsü alan ve kaynak kodlarını arşivleyen gelişmiş bir web scraping aracıdır.

**Siber Vatan** ve **Yıldız CTI** programı kapsamında geliştirilmiştir.

---

##  Özellikler

*  Tam Gizlilik: Tüm trafik `SOCKS5` protokolü üzerinden (127.0.0.1:9150) Tor ağına tünellenir. 
*  Ekran Görüntüsü Alma: JavaScript tabanlı modern onion sitelerinin tam sayfa ekran görüntüsünü alır.
*  HTML Arşivleme: Hedef sitelerin kaynak kodlarını `.txt` formatında kaydeder.
*  Detaylı Raporlama: Taramaların durumunu, zaman damgaları ile birlikte `scan_report.log` dosyasına işler.


---

##  Gereksinimler

Bu aracı çalıştırmak için sisteminizde aşağıdakilerin yüklü olması gerekir:

1.  **Go (Golang):** [İndirmek için tıklayın](https://go.dev/dl/)
2.  **Tor Browser:** [İndirmek için tıklayın](https://www.torproject.org/download/)
    * *Not: Programın çalışması için Tor Browser'ın arka planda açık olması ve bağlantının kurulmuş olması şarttır.*

---

##  Kurulum

Projeyi yerel makinenize klonlayın ve gerekli bağımlılıkları indirin:

```bash


# Proje dizinine gir
cd thor-scraper

# Gerekli modülleri indir
go mod tidy
