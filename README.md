# System udostępniania plików - projekt na przedmiot *Projektowanie i Integracja Systemów*

[![wakatime](https://wakatime.com/badge/github/maksnowak/FileSharingSystem.svg)](https://wakatime.com/badge/github/maksnowak/FileSharingSystem)

## Opis projektu
- System udostępniania plików umożliwiający użytkownikom udostępnianie i pobieranie plików, pozwalający wyszukiwać pliki (do których mamy dostęp) po nazwie, typie, rozmiarze oraz datach (modyfikacji, utworzenia).
- System będzie umożliwiał zabezpieczenie dostępu do folderów oraz pojedynczych plików hasłem. W komponencie bazy danych będą trzymane ścieżki do plików na komputerze, na którym osadzony jest program.
- System będzie można wyeksportować/zaimportować, aby umożliwić posadowienie programu na innym serwerze bez utraty zaszyfrowanych plików, kategorii i tagów.
- System będzie umożliwiał tworzenia kont użytkowników, gdzie każdy użytkownik będzie posiadał pliki prywatne, ogólnodostępne pliki publiczne oraz będzie mógł udostępnić swoje pliki innemu użytkownikowi. Będzie także prowadzona ewidencja dostępu do własnych plików użytkownika (ile było pobrań, kto i kiedy pobrał).

## Architektura:
Projekt zostanie zaimplementowany w następującej architekturze mikroserwisowej:
-	Client (frontend) + Serwer
-	Mikroserwis 1: Przesyłanie plików
-	Mikroserwis 2: Szyfrowanie plików
-	Mikroserwis 3: Tworzenie kont

W razie potrzeby można zmniejszyć/zwiększyć ilość mikroserwisów.

## Technologie:
-	Frontend: Vue.js
-	Backend: Go
-	Database: MongoDB  
### Środowisko Developerskie:
-	VCS + wrzucanie releasów: GitHub
-	CI/CD: Azure DevOps
-	Repozytorium mavenowe: nie dotyczy
-	Budowanie projektu: Makefile   
     -	Możliwość uruchomienia testów
     -	Możliwość *zabicia* aplikacji
-	Wspierane systemu: Ubuntu Linux
-	IDE: Goland (backend), VsCode (frontend)
-	Zarządzanie zadaniami : GitHub Issues
-	Pokrycie kodu testami: `go test –cover`, Vitest 

## Kontrybucje
<a href="https://github.com/maksnowak/FileSharingSystem/graphs/contributors">
<img src="https://contrib.rocks/image?repo=maksnowak/FileSharingSystem" />
</a>

Made with [contrib.rocks](https://contrib.rocks).