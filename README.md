#  THOR SCRAPER - Onion Network Crawler

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)
![Tor Network](https://img.shields.io/badge/Network-Tor%20%2F%20Onion-7D4698?style=flat&logo=tor-browser)
![Platform](https://img.shields.io/badge/Platform-Windows%20%2F%20Linux-gray)
![License](https://img.shields.io/badge/License-MIT-green)

**Thor Scraper**, Siber Tehdit İstihbaratı (CTI) operasyonları için geliştirilmiş, **Tor Ağı (Dark Web)** üzerindeki hedefleri güvenli bir şekilde tarayan, ekran görüntüsü alan ve kaynak kodlarını arşivleyen gelişmiş bir web kazıma aracıdır.

**Siber Vatan** ve **Yıldız CTI** programı kapsamında geliştirilmiştir.

---

##  Özellikler

* ** Tam Gizlilik:** Tüm trafik `SOCKS5` protokolü üzerinden (127.0.0.1:9150) Tor ağına tünellenir. IP sızıntısı önlenmiştir.
* ** Screenshot Yeteneği:** `chromedp` (Headless Chrome) motoru sayesinde, JavaScript tabanlı modern onion sitelerinin tam sayfa ekran görüntüsünü alır.
* ** HTML Arşivleme:** Hedef sitelerin kaynak kodlarını (DOM) `.txt` formatında kaydeder.
* ** Otomatik Konfigürasyon:** `targets.yaml` dosyası yoksa otomatik oluşturur ve örnek verilerle doldurur.
* ** Detaylı Raporlama:** Taramaların durumunu, zaman damgaları ile birlikte `scan_report.log` dosyasına işler.
* ** İnteraktif CLI Arayüzü:** Kullanıcı dostu menü sistemi ve özel ASCII banner tasarımı.

---

## 🛠️ Gereksinimler

Bu aracı çalıştırmak için sisteminizde aşağıdakilerin yüklü olması gerekir:

1.  **Go (Golang):** [İndirmek için tıklayın](https://go.dev/dl/)
2.  **Tor Browser:** [İndirmek için tıklayın](https://www.torproject.org/download/)
    * *Not: Programın çalışması için Tor Browser'ın arka planda açık olması ve bağlantının kurulmuş olması şarttır.*

---

## 📥 Kurulum

Projeyi yerel makinenize klonlayın ve gerekli bağımlılıkları indirin:

```bash
# Projeyi klonla
git clone [https://github.com/KULLANICI_ADINIZ/thor-scraper.git](https://github.com/KULLANICI_ADINIZ/thor-scraper.git)

# Proje dizinine gir
cd thor-scraper

# Gerekli modülleri indir
go mod tidy
